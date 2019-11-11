package oauth

import (
	"fmt"
)

const (
	WECHAT = "wechat"
	QQ     = "qq"
	WEIBO  = "weibo"
)

var oauthes = make(map[string]OAuth)

type OAuth interface {
	Init(conf map[string]string)
	GetAccessToken(code string) (map[string]interface{}, error)
	GetUserInfo(accessToken string, openid string) (map[string]interface{}, error)
	Authorize(code string) (AuthorizeResult, error)
	Code2Session(code string) (map[string]interface{}, error)
}

func RegisterPlatform(name string, oauth OAuth) {
	if oauth == nil {
		panic("Register simpleoauth instance is nil")
	}
	_, dup := oauthes[name]
	if dup {
		panic("The platform has registered already")
	}
	oauthes[name] = oauth
}

type Manager struct {
	oauth OAuth
}

func GetOauth(platformName string, conf map[string]string) (OAuth, error) {
	oauth, ok := oauthes[platformName]
	if !ok {
		return nil, fmt.Errorf("unknown platform %q", platformName)
	}
	oauth.Init(conf)
	return oauth, nil
}

type AuthorizeResult struct {
	Result   bool
	UserInfo map[string]interface{}
}
