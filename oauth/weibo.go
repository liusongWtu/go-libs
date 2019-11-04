package oauth

import "errors"

const weibo_getaccesstoken_url = "https://api.weibo.com/oauth2/access_token"
const weibo_getuserinfo_url = "https://api.weibo.com/2/users/show.json"

var weiboOAuth = &WeiboOAuth{}

type WeiboOAuth struct {
	appKey      string
	appSecret   string
	redirectUrl string
}

func (oauth *WeiboOAuth) Init(conf map[string]string) {
	oauth.appKey = conf["appKey"]
	oauth.appSecret = conf["appSecret"]
	oauth.redirectUrl = conf["redirectUrl"]
}

func (oauth *WeiboOAuth) GetAccessToken(code string) (map[string]interface{}, error) {
	request := Post(weibo_getaccesstoken_url)
	request.Param("client_id", oauth.appKey)
	request.Param("client_secret", oauth.appSecret)
	request.Param("grant_type", "authorization_code")
	request.Param("code", code)
	request.Param("redirect_uri", oauth.redirectUrl)
	var response map[string]interface{}
	err := request.ToJSON(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (oauth *WeiboOAuth) GetUserInfo(accessToken string, openid string) (map[string]interface{}, error) {
	request := Get(weibo_getuserinfo_url)
	request.Param("access_token", accessToken)
	request.Param("uid", openid)
	var response map[string]interface{}
	err := request.ToJSON(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (oauth *WeiboOAuth) Authorize(code string) (AuthorizeResult, error) {
	accessTokenResponse, err := oauth.GetAccessToken(code)
	if err != nil {
		return AuthorizeResult{false, nil}, err
	}
	if accessTokenResponse == nil {
		return AuthorizeResult{false, nil}, nil
	}
	_, ok := accessTokenResponse["error_code"] //获取accesstoken接口返回错误码
	if ok {
		return AuthorizeResult{false, accessTokenResponse}, errors.New("accesstoken接口返回错误码")
	}
	openid := accessTokenResponse["uid"].(string)
	accessToken := accessTokenResponse["access_token"].(string)
	expire := accessTokenResponse["expires_in"].(float64)
	userInfo, err := oauth.GetUserInfo(accessToken, openid)
	if err != nil {
		return AuthorizeResult{false, nil}, err
	}
	if userInfo == nil {
		return AuthorizeResult{false, nil}, nil
	}
	_, ok = userInfo["error_code"] //获取用户信息接口返回错误码
	if ok {
		return AuthorizeResult{false, userInfo}, errors.New("用户信息接口返回错误")
	}
	var sex int
	if userInfo["gender"].(string) == "m" {
		sex = 1
	} else if userInfo["gender"].(string) == "f" {
		sex = 2
	} else if userInfo["gender"].(string) == "n" {
		sex = 0
	}
	return AuthorizeResult{true, map[string]interface{}{
		"openid":       openid,
		"unionid":      "",
		"nickname":     userInfo["screen_name"].(string),
		"sex":          sex,
		"avatar_url":   userInfo["profile_image_url"].(string),
		"access_token": accessToken,
		"expire":       expire,
		"platform":     "weibo",
	}}, nil
}

func init() {
	RegisterPlatform("weibo", weiboOAuth)
}
