package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// Cấu hình kết nối
	connStr := "postgresql://postgres:postgres@localhost:5432/checker_test"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal("Lỗi mở kết nối:", err)
	}
	defer db.Close()

	// Kiểm tra kết nối
	err = db.Ping()
	if err != nil {
		log.Fatal("Không thể kết nối tới database:", err)
	}

	fmt.Println("✅ Kết nối PostgreSQL thành công!")

	// Ví dụ truy vấn
	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		log.Fatal("Lỗi truy vấn:", err)
	}
	fmt.Println("PostgreSQL version:", version)
}
