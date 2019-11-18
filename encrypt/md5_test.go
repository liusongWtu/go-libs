package encrypt

import "testing"

func BenchmarkMd5(b *testing.B) {
	val :="test string 123"
	for i := 0; i < b.N; i++ {
		Md5(val)
	}
}

//func BenchmarkMd5New(b *testing.B) {
//	val :="test string 123"
//	for i := 0; i < b.N; i++ {
//		Md5New(val)
//	}
//}