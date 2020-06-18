package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type student struct {
	Id    int    `json:”id”`
	Nama  string `json:”nama”`
	Umur  int    `json:”umur”`
	Kelas int    `json:”kelas”`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	dbHost := os.Getenv("MYSQL_HOST")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	mysqlDSN := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	fmt.Println("MysqlDSN: ", mysqlDSN)
	db, err := gorm.Open("mysql", mysqlDSN)

	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Failed to connect DB , %v", err))
	}

	defer db.Close()
	db.AutoMigrate(&student{})
	students := []student{}

	db.Find(&students)

	json.NewEncoder(w).Encode(students)
}
