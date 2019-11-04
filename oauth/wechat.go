package oauth

import "errors"

const wechat_getaccesstoken_url = "https://api.weixin.qq.com/sns/oauth2/access_token"
const wechat_getuserinfo_url = "https://api.weixin.qq.com/sns/userinfo"

var weChatOAuth = &WeChatOAuth{}

type WeChatOAuth struct {
	appKey    string
	appSecret string
}

func (oauth *WeChatOAuth) Init(conf map[string]string) {
	oauth.appKey = conf["appKey"]
	oauth.appSecret = conf["appSecret"]
}

func (oauth *WeChatOAuth) GetAccessToken(code string) (map[string]interface{}, error) {
	request := Get(wechat_getaccesstoken_url)
	request.Param("appid", oauth.appKey)
	request.Param("secret", oauth.appSecret)
	request.Param("code", code)
	request.Param("grant_type", "authorization_code")
	var response map[string]interface{}
	err := request.ToJSON(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (oauth *WeChatOAuth) GetUserInfo(accessToken string, openid string) (map[string]interface{}, error) {
	request := Get(wechat_getuserinfo_url)
	request.Param("access_token", accessToken)
	request.Param("openid", openid)
	var response map[string]interface{}
	err := request.ToJSON(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (oauth *WeChatOAuth) Authorize(code string) (AuthorizeResult, error) {
	accessTokenResponse, err := oauth.GetAccessToken(code)
	if err != nil {
		return AuthorizeResult{false, nil}, err
	}
	if accessTokenResponse == nil {
		return AuthorizeResult{false, nil}, nil
	}
	_, ok := accessTokenResponse["errcode"] //获取accesstoken接口返回错误码
	if ok {
		return AuthorizeResult{false, accessTokenResponse}, errors.New("accesstoken接口返回错误")
	}
	openid := accessTokenResponse["openid"].(string)
	accessToken := accessTokenResponse["access_token"].(string)
	expire := accessTokenResponse["expires_in"].(float64)
	userInfo, err := oauth.GetUserInfo(accessToken, openid)
	if err != nil {
		return AuthorizeResult{false, nil}, err
	}
	if userInfo == nil {
		return AuthorizeResult{false, nil}, nil
	}
	_, ok = userInfo["errcode"] //获取用户信息接口返回错误码
	if ok {
		return AuthorizeResult{false, userInfo}, errors.New("用户信息接口返回错误")
	}

	return AuthorizeResult{true, map[string]interface{}{
		"openid":       userInfo["openid"].(string),
		"unionid":      userInfo["unionid"].(string),
		"nickname":     userInfo["nickname"].(string),
		"sex":          userInfo["sex"].(float64),
		"avatar_url":   userInfo["headimgurl"].(string),
		"access_token": accessToken,
		"expire":       expire,
		"platform":     "wechat",
	}}, nil
}

func init() {
	RegisterPlatform("wechat", weChatOAuth)
}
