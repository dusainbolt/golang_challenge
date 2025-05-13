package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("== 1. Overview ==")
	log.Println("Đây là log mặc định.")

	fmt.Println("\n== 2. Types ==")
	// Các loại log thường dùng
	log.Print("Print log (Print)")
	log.Println("Println log (Println)")
	log.Printf("Printf log: %s", "Ghi log có định dạng")

	fmt.Println("\n== 3. Logging to a file ==")
	// Tạo file để ghi log
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Không thể mở file log:", err)
	}
	defer file.Close()

	// Gán output log mặc định vào file
	log.SetOutput(file)
	log.Println("Log này được ghi vào file app.log")

	fmt.Println("\n== 4. Creating custom loggers ==")
	// Tạo logger riêng
	errorLogger := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger := log.New(os.Stdout, "[INFO] ", log.LstdFlags)

	infoLogger.Println("Thông tin chạy chương trình.")
	errorLogger.Println("Có lỗi xảy ra.")

	fmt.Println("\n== 5. Log flags ==")
	flags := []struct {
		name string
		flag int
	}{
		{"Ldate", log.Ldate},
		{"Ltime", log.Ltime},
		{"Lmicroseconds", log.Lmicroseconds},
		{"Llongfile", log.Llongfile},
		{"Lshortfile", log.Lshortfile},
		{"LUTC", log.LUTC},
		{"LstdFlags (Ldate|Ltime)", log.LstdFlags},
	}

	for _, f := range flags {
		logger := log.New(os.Stdout, "["+f.name+"] ", f.flag)
		logger.Println("Test log flag.")
	}
}
