package gopay

import (
	"bytes"
	"encoding/xml"
	"net/http"
)

//解析支付完成后的回调信息
func ParseNotifyResult(req *http.Request) (notifyRsp *WeChatNotifyRequest, err error) {
	notifyRsp = new(WeChatNotifyRequest)
	defer req.Body.Close()
	err = xml.NewDecoder(req.Body).Decode(notifyRsp)
	if err != nil {
		return nil, err
	}
	return
}

type WeChatNotifyResponse struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
}

func (wn *WeChatNotifyResponse) ToXmlString() (xmlStr string) {
	buffer := new(bytes.Buffer)
	buffer.WriteString("<xml><return_code><![CDATA[")
	buffer.WriteString(wn.ReturnCode)
	buffer.WriteString("]]></return_code>")

	buffer.WriteString("<return_msg><![CDATA[")
	buffer.WriteString(wn.ReturnMsg)
	buffer.WriteString("]]></return_msg></xml>")
	xmlStr = buffer.String()
	return
}
