package cryptography

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func TestAesCtrCrypt(t *testing.T) {
	type args struct {
		plainText []byte
		key       []byte
		iv        []byte
	}

	key, _ := hex.DecodeString("BD3A6DE43FF080A512A00AE7B036134092A0600D1FAF0BF641183F95D4EEC001")
	iv, _ := hex.DecodeString("8DA8F20226552F19F907767B3529C001")
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				// plainText: []byte(`{"app_code":"chatgpt","user_id":"","google_id":"google_id_test","android_id":"android_id_test","timezone_offset":28800,"hour":11,"country_code":"CN","adb":0,"mac":"mac","network":"","model":"","brand":"","os_rom":"","de_width":0,"de_height":0,"density":0,"wifi_proxy":0,"de_version":"","vc":0,"vn":"","request_time":0,"is_vpn":0,"sim":0,"boot_time":0,"language":"","wifi_mac":"","wifi_name":"","oaid":""}`),
				plainText: []byte(`RxruQYor8EjHfNXQJIXgzAh6mv94fPcgKPLm82_UcJ-alU29yjOjML1WqsdH27JvsmAS6vweWzEv5CzCDvKaxAYrzAqk8tdeREyNzVJvLdd_9YkhE6wSXaaOU12rJMYCf4qCRmRJ9ibTDQZnJYTpOAHUhbbyfvU9OU4t-lyw9oIA3F3E0DChnxJr-9XHdhPhEaicn59jFydPyeAfxHuY2j8RVyjMREJVaGTyGgUTlu3ipUv9tMd_Kqq77p5VY6A`),
				key:       key,
				iv:        iv,
			},
			want:    []byte{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbuf := make([]byte, base64.StdEncoding.DecodedLen(len(tt.args.plainText)))
			n, err := base64.StdEncoding.Decode(dbuf, tt.args.plainText)
			decodeString := dbuf[:n]
			ciphertext, err := AesCtrCrypt(decodeString, tt.args.key, tt.args.iv)
			buf := make([]byte, base64.StdEncoding.EncodedLen(len(ciphertext)))
			base64.StdEncoding.Encode(buf, ciphertext)
			val := string(buf)
			_ = val
			if (err != nil) != tt.wantErr {
				t.Errorf("AesCtrCrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}