package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/alumni_tracing")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Tes koneksi ke database
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("DB ", db)
	fmt.Println("Koneksi berhasil!")

	rows, err := db.Query("SELECT nomer, tanggal FROM yudisium")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	//fmt.Println("rows :", rows)

	// Iterasi hasil query
	for rows.Next() {
		var nomer int
		var tanggal string
		err = rows.Scan(&nomer, &tanggal)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("nomer yudisium\t\t: ", nomer)
		fmt.Println("tanggal yudisium\t: ", tanggal)
	}
	// Periksa kesalahan iterasi
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}

}
