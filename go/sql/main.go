package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Customers struct {
	customersID        int
	customersFirstname string
	customersLastname  interface{}
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(10.0.5.71:3306)/glamoursales_20190913")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select customers_id, customers_firstname, customers_lastname from customers limit 100")
	if err != nil {
		// ログメッセージを出力した後、os.Exit(1) を呼び出してプログラムを終了
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var c Customers
		err := rows.Scan(&c.customersID, &c.customersFirstname, &c.customersLastname)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(c.customersID, c.customersFirstname)

		getTypea(c.customersLastname)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
}

func getTypea(x interface{}) {
	switch x.(type) {
	case bool:
		fmt.Println("bool")
	case int, uint:
		fmt.Println("int or uint")
	case string:
		fmt.Println("string")
	case byte:
		fmt.Println("byte")
	case rune:
		fmt.Println("rune")
	case interface{}:
		fmt.Println("interface")
	default:
		fmt.Println("don't know")
	}
}
