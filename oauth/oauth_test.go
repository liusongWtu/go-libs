package oauth

//import (
//	"fmt"
//	"testing"
//)
//
//func TestManager_Authorize(t *testing.T) {
//	conf := make(map[string]string)
//	conf["appKey"] = ""
//	conf["appSecret"] = ""
//	m, _ := NewManager(WECHAT, conf)
//	fmt.Println(m.Authorize("0210r0wV1HFv6T0f6qwV154KvV10r0wJ"))
//}
//
//func TestQQOAuth_GetAccesstoken(t *testing.T) {
//	conf := make(map[string]string)
//	conf["appKey"] = ""
//	conf["appSecret"] = ""
//	conf["redirectUrl"] = ""
//	m, _ := NewManager(QQ, conf)
//	m.oauth.GetAccessToken("7B6DEDC59F858A0B63F4D2C42F6D0E71")
//}
//
//func TestQQOAuth_Authorize(t *testing.T) {
//	conf := make(map[string]string)
//	conf["appKey"] = ""
//	conf["appSecret"] = ""
//	conf["redirectUrl"] = ""
//	m, _ := NewManager(QQ, conf)
//	m.oauth.Authorize("2EB4B7C981FAFCC1BAD97A2B2466EEDD")
//}
//
//func TestWeiboOAuth_Authorize(t *testing.T) {
//	conf := make(map[string]string)
//	conf["appKey"] = ""
//	conf["appSecret"] = ""
//	conf["redirectUrl"] = ""
//	m, _ := NewManager(WEIBO, conf)
//	fmt.Println(m.oauth.Authorize("63033d599df176a23cac0363db0787cf"))
//}
