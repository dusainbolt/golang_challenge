package main

import (
	"encoding/base64"
	"fmt"
)

// DecodeSecret giải mã một chuỗi bí mật được mã hóa bằng Base64,
// sau đó dịch lùi mỗi ký tự trong chuỗi đã giải mã đi 1 đơn vị (ví dụ: 'b' -> 'a')
func DecodeSecret(message string) string {
	// Giải mã chuỗi base64 -> nhận được chuỗi gốc (dạng []byte)
	data, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		fmt.Println(err)
	}

	// In ra chuỗi sau khi giải mã base64 (dùng để kiểm tra)
	fmt.Println(string(data))

	var secret []rune

	// Duyệt qua từng ký tự trong chuỗi đã giải mã
	for _, char := range data {
		// In ký tự được dịch tiến +1 để kiểm tra (không ảnh hưởng đến logic chính)
		fmt.Println(rune(char + 1))

		// Dịch lùi mỗi ký tự -1 (ví dụ: 'b' -> 'a') và thêm vào mảng kết quả
		secret = append(secret, rune(char-1))
	}

	// In ra mảng các ký tự sau khi đã dịch lùi (dạng []rune)
	fmt.Println(secret)

	// Trả về chuỗi kết quả sau khi chuyển []rune về string
	return string(secret)
}

func main() {
	fmt.Println("Decode the Secret")

	message := "VEZEU0ZVVFVTSk9I"
	result := DecodeSecret(message)
	fmt.Println(result)

}
