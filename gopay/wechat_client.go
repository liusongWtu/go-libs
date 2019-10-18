package gopay

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
)

type weChatClient struct {
	AppId     string
	MchId     string
	secretKey string
	signType  string
	isProd    bool
}

//初始化微信客户端 ok
//    appId：应用ID
//    mchID：商户ID
//    secretKey：Key值
//    isProd：是否是正式环境
func NewWeChatClient(appId, mchId, secretKey, signType string, isProd bool) *weChatClient {
	client := new(weChatClient)
	client.AppId = appId
	client.MchId = mchId
	client.secretKey = secretKey
	client.signType = signType
	client.isProd = isProd
	return client
}

//提交付款码支付 ok
//文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_10&index=1
func (this *weChatClient) Micropay(body BodyMap) (wxRsp *WeChatMicropayResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		tlsConfig := new(tls.Config)
		tlsConfig.InsecureSkipVerify = true
		bytes, err = this.doWeChat(body, wxURL_Micropay, tlsConfig)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_Micropay)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatMicropayResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//统一下单 ok
//文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
func (this *weChatClient) UnifiedOrder(body BodyMap) (wxRsp *WeChatUnifiedOrderResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		tlsConfig := new(tls.Config)
		tlsConfig.InsecureSkipVerify = true
		bytes, err = this.doWeChat(body, wxURL_UnifiedOrder, tlsConfig)
		if err != nil {
			return nil, err
		}
	} else {
		body.Set("total_fee", 101)
		bytes, err = this.doWeChat(body, wxURL_SanBox_UnifiedOrder)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatUnifiedOrderResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//查询订单 ok
//文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2
func (this *weChatClient) QueryOrder(body BodyMap) (wxRsp *WeChatQueryOrderResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		tlsConfig := new(tls.Config)
		tlsConfig.InsecureSkipVerify = true
		bytes, err = this.doWeChat(body, wxURL_OrderQuery, tlsConfig)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_OrderQuery)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatQueryOrderResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//关闭订单 ok
//文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3
func (this *weChatClient) CloseOrder(body BodyMap) (wxRsp *WeChatCloseOrderResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		tlsConfig := new(tls.Config)
		tlsConfig.InsecureSkipVerify = true
		bytes, err = this.doWeChat(body, wxURL_CloseOrder, tlsConfig)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_CloseOrder)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatCloseOrderResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//撤销订单 ok
//文档地址：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3
func (this *weChatClient) Reverse(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *WeChatReverseResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		pkcsPool := x509.NewCertPool()
		pkcs, err := ioutil.ReadFile(pkcs12FilePath)
		if err != nil {
			return nil, err
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		certificate, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
		if err != nil {
			return nil, err
		}
		tlsConfig := new(tls.Config)
		tlsConfig.Certificates = []tls.Certificate{certificate}
		tlsConfig.RootCAs = pkcsPool
		tlsConfig.InsecureSkipVerify = true

		bytes, err = this.doWeChat(body, wxURL_Reverse, tlsConfig)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_Reverse)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatReverseResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//申请退款 ok
//文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
func (this *weChatClient) Refund(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp *WeChatRefundResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		pkcsPool := x509.NewCertPool()
		pkcs, err := ioutil.ReadFile(pkcs12FilePath)
		if err != nil {
			return nil, err
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		certificate, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
		if err != nil {
			return nil, err
		}
		tlsConfig := new(tls.Config)
		tlsConfig.Certificates = []tls.Certificate{certificate}
		tlsConfig.RootCAs = pkcsPool
		tlsConfig.InsecureSkipVerify = true

		bytes, err = this.doWeChat(body, wxURL_Refund, tlsConfig)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_Refund)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatRefundResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//查询退款 ok
//文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5
func (this *weChatClient) QueryRefund(body BodyMap) (wxRsp *WeChatQueryRefundResponse, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		tlsConfig := new(tls.Config)
		tlsConfig.InsecureSkipVerify = true
		bytes, err = this.doWeChat(body, wxURL_RefundQuery, tlsConfig)
		if err != nil {
			return nil, err
		}
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_RefundQuery)
		if err != nil {
			return nil, err
		}
	}

	wxRsp = new(WeChatQueryRefundResponse)
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//下载对账单 ok
//文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6
func (this *weChatClient) DownloadBill(body BodyMap) (wxRsp string, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		tlsConfig := new(tls.Config)
		tlsConfig.InsecureSkipVerify = true
		bytes, err = this.doWeChat(body, wxURL_DownloadBill, tlsConfig)
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_DownloadBill)
	}
	wxRsp = string(bytes)
	if err != nil {
		return wxRsp, err
	}
	return wxRsp, nil
}

//下载资金账单 ok
//文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_18&index=7
//好像不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
func (this *weChatClient) DownloadFundFlow(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp string, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		pkcsPool := x509.NewCertPool()
		pkcs, err := ioutil.ReadFile(pkcs12FilePath)
		if err != nil {
			return "", err
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		certificate, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
		if err != nil {
			return "", err
		}
		tlsConfig := new(tls.Config)
		tlsConfig.Certificates = []tls.Certificate{certificate}
		tlsConfig.RootCAs = pkcsPool
		tlsConfig.InsecureSkipVerify = true

		bytes, err = this.doWeChat(body, wxURL_DownloadFundFlow, tlsConfig)
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_DownloadFundFlow)
	}

	if err != nil {
		return "", err
	}
	wxRsp = string(bytes)
	return wxRsp, nil
}

//拉取订单评价数据
//文档地址：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_17&index=11
//好像不支持沙箱环境，因为沙箱环境默认需要用MD5签名，但是此接口仅支持HMAC-SHA256签名
func (this *weChatClient) BatchQueryComment(body BodyMap, certFilePath, keyFilePath, pkcs12FilePath string) (wxRsp string, err error) {
	var bytes []byte
	if this.isProd {
		//正式环境
		body.Set("sign_type", SignType_HMAC_SHA256)

		pkcsPool := x509.NewCertPool()
		pkcs, err := ioutil.ReadFile(pkcs12FilePath)
		if err != nil {
			return "", err
		}
		pkcsPool.AppendCertsFromPEM(pkcs)
		certificate, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
		if err != nil {
			return "", err
		}
		tlsConfig := new(tls.Config)
		tlsConfig.Certificates = []tls.Certificate{certificate}
		tlsConfig.RootCAs = pkcsPool
		tlsConfig.InsecureSkipVerify = true

		bytes, err = this.doWeChat(body, wxURL_BatchQueryComment, tlsConfig)
	} else {
		bytes, err = this.doWeChat(body, wxURL_SanBox_BatchQueryComment)
	}

	if err != nil {
		return "", err
	}

	wxRsp = string(bytes)
	return wxRsp, nil
}

//企业付款到零钱
//文档地址：https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
func (this *weChatClient) MchTransfer(body BodyMap, certFilePath, keyFilePath string) (wxRsp *WeChatMchTransferResponse, err error) {
	//发送请求到微信支付接口
	tlsConfig, _ := getTLSConfig(certFilePath, keyFilePath)
	bytes, err := this.doEnterpriseWeChat(body, wxURL_Mch_Transfer, tlsConfig)
	if err != nil {
		return nil, err
	}
	//微信信息返回
	wxRsp = new(WeChatMchTransferResponse)
	fmt.Println(string(bytes))
	err = xml.Unmarshal(bytes, wxRsp)
	if err != nil {
		return nil, err
	}
	return wxRsp, nil
}

//向微信发送请求 ok
func (this *weChatClient) doWeChat(body BodyMap, url string, tlsConfig ...*tls.Config) (bytes []byte, err error) {
	//===============生成参数===================
	body.Set("appid", this.AppId)
	body.Set("mch_id", this.MchId)

	var sign string
	if !this.isProd {
		//沙箱环境
		body.Set("sign_type", SignType_MD5)
		//从微信接口获取SanBoxSignKey
		key, err := getSanBoxSign(this.MchId, body.Get("nonce_str"), this.secretKey, body.Get("sign_type"))
		if err != nil {
			return nil, err
		}
		sign = getLocalSign(key, body.Get("sign_type"), body)
	} else {
		//正式环境
		//本地计算Sign
		sign = getLocalSign(this.secretKey, body.Get("sign_type"), body)
	}
	body.Set("sign", sign)
	reqXML := generateXml(body)
	//===============发起请求===================
	agent := gorequest.New()
	if this.isProd && tlsConfig != nil {
		agent.TLSClientConfig(tlsConfig[0])
	}
	agent.Post(url)
	agent.Type("xml")
	agent.SendString(reqXML)
	_, bytes, errs := agent.EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return bytes, nil
}

//发送微信企业支付请求
func (this *weChatClient) doEnterpriseWeChat(body BodyMap, url string, tlsConfig ...*tls.Config) (bytes []byte, err error) {
	//===============生成参数===================
	body.Set("mch_appid", this.AppId)
	body.Set("mchid", this.MchId)
	sign := getEnterpriseSign(this.secretKey, this.signType, body)
	body.Set("sign", sign)
	reqXML := generateXml(body)
	//===============发起请求===================
	agent := gorequest.New()
	if this.isProd && tlsConfig != nil {
		agent.TLSClientConfig(tlsConfig[0])
	}
	agent.Post(url)
	agent.Type("xml")
	agent.SendString(reqXML)
	_, bytes, errs := agent.EndBytes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return bytes, nil
}
