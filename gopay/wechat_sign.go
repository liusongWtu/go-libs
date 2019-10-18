package gopay

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/url"
	"sort"
	"strings"
)

//JSAPI支付，支付参数后，再次计算出小程序用的paySign
func getMiniPaySign(appId, nonceStr, prepayId, signType, timeStamp, secretKey string) (paySign string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("appId=")
	buffer.WriteString(appId)

	buffer.WriteString("&nonceStr=")
	buffer.WriteString(nonceStr)

	buffer.WriteString("&package=")
	buffer.WriteString(prepayId)

	buffer.WriteString("&signType=")
	buffer.WriteString(signType)

	buffer.WriteString("&timeStamp=")
	buffer.WriteString(timeStamp)

	buffer.WriteString("&key=")
	buffer.WriteString(secretKey)

	signStr := buffer.String()

	var hashSign []byte
	if signType == SignType_MD5 {
		hash := md5.New()
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	} else {
		hash := hmac.New(sha256.New, []byte(secretKey))
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	}
	paySign = strings.ToUpper(hex.EncodeToString(hashSign))
	return
}

//本地通过支付参数计算Sign值
func getLocalSign(secretKey string, signType string, body BodyMap) (sign string) {
	signStr := sortSignParams(secretKey, body)
	var hashSign []byte
	if signType == SignType_MD5 {
		hash := md5.New()
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	} else {
		hash := hmac.New(sha256.New, []byte(secretKey))
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	}
	sign = strings.ToUpper(hex.EncodeToString(hashSign))
	return
}

func sortSignParams(secretKey string, body BodyMap) string {
	keyList := make([]string, 0)
	for k, _ := range body {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	buffer := new(bytes.Buffer)
	for _, k := range keyList {
		buffer.WriteString(k)
		buffer.WriteString("=")

		valueStr := convert2String(body[k])
		buffer.WriteString(valueStr)

		buffer.WriteString("&")
	}
	buffer.WriteString("key=")
	buffer.WriteString(secretKey)
	str := buffer.String()
	return str
}

//企业支付参数计算Sign值
func getEnterpriseSign(secretKey string, signType string, body BodyMap) (sign string) {
	signStr := sortEnterpriseSignParams(secretKey, body)
	var hashSign []byte
	if signType == SignType_MD5 {
		hash := md5.New()
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	} else {
		hash := hmac.New(sha256.New, []byte(secretKey))
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	}
	sign = strings.ToUpper(hex.EncodeToString(hashSign))
	return
}

//获取企业支付根据Key排序后的请求参数字符串
func sortEnterpriseSignParams(secretKey string, body BodyMap) string {
	keyList := make([]string, 0)
	for k, v := range body {
		if fmt.Sprintf("%v", v) == "" {
			continue
		}
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	str := ""
	for _, k := range keyList {
		str += k + "=" + convert2String(body[k]) + "&"
	}
	str = strings.TrimRight(str, "&")
	r, _ := url.QueryUnescape(str)
	return r + "&key=" + secretKey
}

//从微信提供的接口获取：SandboxSignKey
func getSanBoxSign(mchId, nonceStr, secretKey, signType string) (key string, err error) {
	body := make(BodyMap)
	body.Set("mch_id", mchId)
	body.Set("nonce_str", nonceStr)

	//计算沙箱参数Sign
	sanboxSign := getLocalSign(secretKey, signType, body)
	//沙箱环境：获取key后，重新计算Sign
	key, err = getSanBoxSignKey(mchId, nonceStr, sanboxSign)
	if err != nil {
		return "", err
	}
	return
}

//从微信提供的接口获取：SandboxSignkey
func getSanBoxSignKey(mchId, nonceStr, sign string) (key string, err error) {
	reqs := make(BodyMap)
	reqs.Set("mch_id", mchId)
	reqs.Set("nonce_str", nonceStr)
	reqs.Set("sign", sign)

	reqXml := generateXml(reqs)
	//fmt.Println("req:::", reqXml)
	_, byteList, errorList := gorequest.New().
		Post(wxURL_SanBox_GetSignKey).
		Type("xml").
		SendString(reqXml).EndBytes()
	if len(errorList) > 0 {
		return "", errorList[0]
	}
	keyResponse := new(getSignKeyResponse)
	err = xml.Unmarshal(byteList, keyResponse)
	if err != nil {
		return "", err
	}
	if keyResponse.ReturnCode == "FAIL" {
		return "", errors.New(keyResponse.Retmsg)
	}
	return keyResponse.SandboxSignkey, nil
}

//初始化ca证书
func getTLSConfig(certFilePath, keyFilePath string) (*tls.Config, error) {
	// load cert
	cert, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		return nil, err
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	return tlsConfig, nil
}

//生成请求XML的Body体
func generateXml(bm BodyMap) (reqXml string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("<xml>")

	for k, v := range bm {
		buffer.WriteString("<")
		buffer.WriteString(k)
		buffer.WriteString("><![CDATA[")

		valueStr := convert2String(v)
		buffer.WriteString(valueStr)

		buffer.WriteString("]]></")
		buffer.WriteString(k)
		buffer.WriteString(">")
	}
	buffer.WriteString("</xml>")
	reqXml = buffer.String()
	return
}
