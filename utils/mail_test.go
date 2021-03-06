package utils

import "testing"

func TestMail(t *testing.T) {
	config := `{"username":"user001@gmail.com","password":"user001","host":"smtp.gmail.com","port":587}`
	mail := NewEMail(config)
	if mail.Username != "user001@gmail.com" {
		t.Fatal("email parse get username error")
	}
	if mail.Password != "user001" {
		t.Fatal("email parse get password error")
	}
	if mail.Host != "smtp.gmail.com" {
		t.Fatal("email parse get host error")
	}
	if mail.Port != 587 {
		t.Fatal("email parse get port error")
	}
	mail.To = []string{"xiemengjun@gmail.com"}
	mail.From = "user001@gmail.com"
	mail.Subject = "hi, just from beego!"
	mail.Text = "Text Body is, of course, supported!"
	mail.HTML = "<h1>Fancy Html is supported, too!</h1>"
	mail.AttachFile("/Users/user001/11.txt")
	mail.Send()
}
