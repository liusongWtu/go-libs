package oauth

import (
	"fmt"
	"testing"
)

/*

{"errMsg":"getUserInfo:ok","rawData":"{\"nickName\":\"Dreamer\",\"gender\":1,\"language\":\"zh_CN\",\"city\":\"Chaoyang\",\"province\":\"Beijing\",\"country\":\"China\",\"avatarUrl\":\"https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTKhUo24y1c3vHjnaz4TGtulEsTptWy99j2XzhHRRMSeBdfJpbXKTmTABI48TYFEFibPkJ3b0cU5tOQ/132\"}","userInfo":{"nickName":"Dreamer","gender":1,"language":"zh_CN","city":"Chaoyang","province":"Beijing","country":"China","avatarUrl":"https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTKhUo24y1c3vHjnaz4TGtulEsTptWy99j2XzhHRRMSeBdfJpbXKTmTABI48TYFEFibPkJ3b0cU5tOQ/132"},"signature":"5200b79c80107cdf8b2e0da04403b85021483554","encryptedData":"jlFW8a9Isq0dXsiKPqfllAnxVUiKmqAH71WZlgK+n/ee8tx6Qme1+evu5RTpHLXN9wvw8akO4E/mNWp7MdEHX+VjS7icuNoC95EUYe4pMZbEsIVLF1cRcd/tloewdKOIo3jCT7lLPNcUhquh5O5astaUnAPAeWtCsKniFr0K57+4fWF3EfgsxQBCMrLbdsdRr+mLHQchCDODLZxML7t4MyKxfCXed65oMPoVAFn9JmBKHZk7Q/NmyWqh0PZJM2pb6+vOSjWnzRefe0Ku+qmx/Ri2p164H5M5qwd/GtRZCSeRFUAh+7jkcSuwtFN2Mbbxkd3JsKUOuJbQ2BrR397ZZOS5P3XV7iVwtXHdcFi3iSBl9Fvfgk4lQwprCOE+7ef1UemnXN/Oj/Yrlw7tfv/qWUz66q51tY/Fx0rWD4qK4+0dulasQqcjHiLLn7U0B1CkiMfv6BixiEP4sgkfVPWerIC5n8wNoXk3aDX3qoPuSNc=","iv":"u95DtEMQhaGgi8lWgy3OeQ=="}

*/

func TestWXDecryptUserInfo(t *testing.T) {
	sessionKey := "CfnXiF/Y67eaC2DKw+IQMA=="	//变化的
	rawData := "{\"nickName\":\"Dreamer\",\"gender\":1,\"language\":\"zh_CN\",\"city\":\"Chaoyang\",\"province\":\"Beijing\",\"country\":\"China\",\"avatarUrl\":\"https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTKhUo24y1c3vHjnaz4TGtulEsTptWy99j2XzhHRRMSeBdfJpbXKTmTABI48TYFEFibPkJ3b0cU5tOQ/132\"}"
	encryptedData := "jLszyhNCxdgxGSN1+A6VLW3c7qubQZEoZncL+OBqTiaBUnIK3DUZ7Ziv7OvW8tOxDXiY6N2Ybu7bPJL6eSDai2giK7S9KNZZD9oQZ/+/LkvuAj/BzYXlch/Ajn5xorJ0L6YdThy47b1fCnyQPGNEU3Lbdqqo9nxDajX5ZguT+86LuaHJAsDRj4Bn5wcZAK/O8qfVpoHcDZPruVLIh0pUSPgutFZ0JUbZdXNKLDj4FpbRkkqB+lVLxz0NjIVHOb9IcZdsNeDAFfmgCTCvw1rDtboVdNsELcNM78aH2SfWTbR3N68AYK75J6g3s/37m4DK4KKOXwxpXNzOADq3tWMGJlPDmilhyO8rhjEaHvQNnAAbGAxddTmIgEHJXtU9FF9k96wb4+8F8Cz7EAzx89NJAVjtCVOLWY4GY+Mn4RZsujPJO7PXfamkBRBqj3ni/ETvG0Wd+8wexPWMRRcJ27bHudhLE244lkSz3rfAxIxrTYQ="
	signature := "1a18de586774c571f15e80f8742afbfa80be9b3c"
	iv := "ycL/utQw3fTMLracA2dE8A=="
	userInfo, err := WXDecryptUserInfo(sessionKey, rawData, encryptedData, signature, iv)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(userInfo)
}

func TestWXDecryptUserInfo2(t *testing.T) {
	//appId := "wx4f4bc4dec97d474b"
	sessionKey := "tiihtNczf5v6AKRyjwEUhQ=="	//变化的
	rawData := "{\"nickName\":\"Dreamer\",\"gender\":1,\"language\":\"zh_CN\",\"city\":\"Chaoyang\",\"province\":\"Beijing\",\"country\":\"China\",\"avatarUrl\":\"https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTKhUo24y1c3vHjnaz4TGtulEsTptWy99j2XzhHRRMSeBdfJpbXKTmTABI48TYFEFibPkJ3b0cU5tOQ/132\"}"
	encryptedData := "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
	signature := "8b3594108004d160e324b401b67d47b397c4932f"
	iv := "r7BXXKkLb8qrSNn05n0qiA=="
	userInfo, err := WXDecryptUserInfo(sessionKey, rawData, encryptedData, signature, iv)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(userInfo)
}