package ses

import (
	"context"
	"encoding/csv"
	"fmt"

	awsxConfig "libs/pkg/awsx/config"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

type (
	EmailTemplate struct {
		Product      string
		Sender       string
		Recipient    string
		Subject      string
		TemplateName string

		// An object that defines the values to use for message variables in the template.
		// This object is a set of key-value pairs. Each key defines a message variable in
		// the template. The corresponding value defines the value to use for that
		// variable. eg: `{"code":"test_code"}`
		TemplateData         string
		ConfigurationSetName string
	}

	Client struct {
		Credential awsxConfig.Credential
		Client     *sesv2.Client
	}
)

func NewClient(credential awsxConfig.Credential) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(credential.Key, credential.Secret, "")),
		config.WithRegion(credential.Region),
	)
	if err != nil {
		return nil, err
	}
	client := sesv2.NewFromConfig(cfg)
	return &Client{Credential: credential, Client: client}, nil
}

func SendEmailByTemplate(client *sesv2.Client, emailTemplate EmailTemplate) (*sesv2.SendEmailOutput, error) {
	input := &sesv2.SendEmailInput{
		Content: &types.EmailContent{
			Raw: nil,
			Template: &types.Template{
				TemplateArn:  nil,
				TemplateData: aws.String(emailTemplate.TemplateData),
				TemplateName: aws.String(emailTemplate.TemplateName),
			},
		},
		Destination: &types.Destination{
			BccAddresses: []string{},
			CcAddresses:  []string{},
			ToAddresses:  []string{emailTemplate.Recipient},
		},
		EmailTags: []types.MessageTag{{
			Name:  aws.String("product"),
			Value: aws.String(emailTemplate.Product),
		}},
		FeedbackForwardingEmailAddress:            nil,
		FeedbackForwardingEmailAddressIdentityArn: nil,
		FromEmailAddress:                          aws.String(emailTemplate.Sender),
		FromEmailAddressIdentityArn:               nil,
		ListManagementOptions:                     nil,
		ReplyToAddresses:                          nil,
	}

	if emailTemplate.ConfigurationSetName != "" {
		input.ConfigurationSetName = aws.String(emailTemplate.ConfigurationSetName)
	}

	return client.SendEmail(context.Background(), input)
}

func GetSuppressedDestination(client *sesv2.Client, email string) (*sesv2.GetSuppressedDestinationOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return client.GetSuppressedDestination(ctx, &sesv2.GetSuppressedDestinationInput{
		EmailAddress: aws.String(email),
	})
}

func PutSuppressedDestination(client *sesv2.Client, email string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	_, err := client.PutSuppressedDestination(ctx, &sesv2.PutSuppressedDestinationInput{
		EmailAddress: aws.String(email),
		Reason:       types.SuppressionListReasonBounce,
	})
	return err
}

func ListSuppressedEmail(client *sesv2.Client, startTime, endTime time.Time, nextToken string) (*sesv2.ListSuppressedDestinationsOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	var nextTokenVal *string
	if nextToken != "" {
		nextTokenVal = aws.String(nextToken)
	}
	return client.ListSuppressedDestinations(ctx, &sesv2.ListSuppressedDestinationsInput{
		EndDate:   &endTime,
		NextToken: nextTokenVal,
		PageSize:  aws.Int32(100),
		Reasons:   []types.SuppressionListReason{},
		StartDate: &startTime,
	})
}

func ExportSuppressedEmailsToCsv(client *sesv2.Client, name string, startTime, endTime time.Time, nextToken string) (string, error) {
	csvName := fmt.Sprintf("%s_suppressed_emails_%s_%s.csv", name, startTime.Format("20060102150405"), endTime.Format("20060102150405"))
	file, err := os.Create(csvName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for {
		output, err := ListSuppressedEmail(client, startTime, endTime, nextToken)
		if err != nil {
			return "", err
		}

		for _, email := range output.SuppressedDestinationSummaries {
			writer.Write([]string{*email.EmailAddress, string(email.Reason)})
		}

		if output.NextToken == nil {
			break
		}

		nextToken = *output.NextToken
		fmt.Println("Next token: ", nextToken)
		time.Sleep(time.Second * 10)
	}
	return nextToken, nil
}

func ImportSuppressedEmailFromCsv(client *sesv2.Client, filename string) error {
	// 打开CSV文件
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Couldn't open the csv file:%v", err)
	}

	// 解析CSV文件
	reader := csv.NewReader(file)
	isStart := false
	// 逐行读取CSV文件内容
	for {
		record, err := reader.Read()
		if err != nil {
			return err
		}

		if len(record) == 0 {
			break
		}
		if record[0] == "parhamarhami8@gmail.com" {
			isStart = true
			continue
		}
		if !isStart {
			continue
		}
		err = PutSuppressedDestination(client, record[0])
		if err != nil {
			return err
		}
		fmt.Println(record[0])
		time.Sleep(time.Second)
	}

	if err := file.Close(); err != nil {
		return err
	}
	return nil
}
