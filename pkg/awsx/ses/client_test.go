package ses

import (
	awsxConfig "libs/pkg/awsx/config"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/stretchr/testify/assert"
)

func getConfig() awsxConfig.Credential {
	return awsxConfig.Credential{
		Key:    "AKIAT37CEV2ZYVJ2367W",
		Secret: "",
		Region: "ap-northeast-1",
	}
}

func getJumpConfig() awsxConfig.Credential {
	return awsxConfig.Credential{
		Key:    "AKIAT37CEV2ZTHJZMFUT",
		Secret: "",
		Region: "eu-central-1",
	}
}

func getBiu2SESConfig() awsxConfig.Credential {
	return awsxConfig.Credential{
		Key:    "AKIA5TZWO4VXR2HM2ZFG",
		Secret: "",
		Region: "eu-central-1",
	}
}

func TestGetSuppressedDestination(t *testing.T) {
	// credential := getJumpConfig()
	credential := getConfig()
	client, err := NewClient(credential)
	assert.NoError(t, err)

	resp, err := GetSuppressedDestination(client.Client, "sajnadimi1@gmail.com")
	assert.NoError(t, err)
	t.Log(resp)

}

func TestPutSuppressedDestination(t *testing.T) {
	credential := getJumpConfig()
	client, err := NewClient(credential)
	assert.NoError(t, err)

	PutSuppressedDestination(client.Client, "sajnadimi1@gmail.com")
}

func TestExportSuppressedEmailsToCsv(t *testing.T) {
	credential := getJumpConfig()
	jumpClient, err := NewClient(credential)
	assert.NoError(t, err)

	credential = getConfig()
	biuClient, err := NewClient(credential)
	assert.NoError(t, err)

	type args struct {
		client    *sesv2.Client
		name      string
		startTime time.Time
		endTime   time.Time
		nextToken string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "jumpjump",
			args: args{
				client:    jumpClient.Client,
				name:      "jumpjump",
				startTime: time.Now().AddDate(-1, 0, 0),
				endTime:   time.Now(),
				nextToken: "",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "biubiu",
			args: args{
				client:    biuClient.Client,
				name:      "biubiu",
				startTime: time.Now().AddDate(-2, 0, 0),
				endTime:   time.Now(),
				nextToken: "",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExportSuppressedEmailsToCsv(tt.args.client, tt.args.name, tt.args.startTime, tt.args.endTime, tt.args.nextToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExportSuppressedEmailsToCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Log(got)
		})
	}
}

func TestImportSuppressedEmailFromCsv(t *testing.T) {

	credential := getBiu2SESConfig()
	client, err := NewClient(credential)
	assert.NoError(t, err)
	type args struct {
		client   *sesv2.Client
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// {
		// 	name: "biubiu",
		// 	args: args{
		// 		client:   client.Client,
		// 		filename: "biubiu_suppressed_emails_20211116195143_20231116195143.csv",
		// 	},
		// 	wantErr: false,
		// },
		{
			name: "jump",
			args: args{
				client:   client.Client,
				filename: "jumpjump_suppressed_emails_20221116195143_20231116195143.csv",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ImportSuppressedEmailFromCsv(tt.args.client, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("ImportSuppressedEmailFromCsv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
