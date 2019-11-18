package oauth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
)

const pkcs7blocksize = 32

type watermark struct {
	AppID     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}

// UserInfo 解密后的用户信息
type UserInfo struct {
	OpenID    string    `json:"openId"`
	Nickname  string    `json:"nickName"`
	Gender    int       `json:"gender"`
	Province  string    `json:"province"`
	Language  string    `json:"language"`
	Country   string    `json:"country"`
	City      string    `json:"city"`
	Avatar    string    `json:"avatarUrl"`
	UnionID   string    `json:"unionId"`
	Watermark watermark `json:"watermark"`
}

// DecryptUserInfo 解密用户信息
//
// sessionKey 微信 session_key
// rawData 不包括敏感信息的原始数据字符串，用于计算签名。
// encryptedData 包括敏感数据在内的完整用户信息的加密数据
// signature 使用 sha1( rawData + session_key ) 得到字符串，用于校验用户信息
// iv 加密算法的初始向量
func WXDecryptUserInfo(sessionKey, rawData, encryptedData, signature, iv string) (*UserInfo, error) {

	if ok := validateUserInfo(signature, rawData, sessionKey); !ok {
		return nil, errors.New("failed to validate signature")
	}

	raw, err := decryptUserData("", sessionKey, encryptedData, iv)
	if err != nil {
		return nil, err
	}

	info := new(UserInfo)
	if err := json.Unmarshal(raw, info); err != nil {
		return nil, err
	}
	return info, err
}

// 校验用户数据数据
func validateUserInfo(signature, rawData, ssk string) bool {
	raw := sha1.Sum([]byte(rawData + ssk))
	return signature == hex.EncodeToString(raw[:])
}

func decryptUserData(appId, sessionKey, encryptedData, iv string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return nil, err
	}

	cipher, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}

	rawIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	return cbcDecrypt(key, cipher, rawIV)
}

// CBC解密数据
func cbcDecrypt(key, ciphertext, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	size := aes.BlockSize
	iv = iv[:size]
	// ciphertext = ciphertext[size:] TODO: really useless?

	if len(ciphertext) < size {
		return nil, errors.New("ciphertext too short")
	}

	if len(ciphertext)%size != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return pkcs7decode(ciphertext), nil
}

// pkcs7decode 对解密后的明文进行补位删除
// plaintext 解密后的明文
// 返回删除填充补位后的明文和
func pkcs7decode(plaintext []byte) []byte {
	ln := len(plaintext)

	// 获取最后一个字符的 ASCII
	pad := int(plaintext[ln-1])
	if pad < 1 || pad > pkcs7blocksize {
		pad = 0
	}

	return plaintext[:(ln - pad)]
}
