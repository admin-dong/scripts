package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // 根据你的数据库类型选择适当的驱动引入
)

func main() {
	// 连接数据库
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 获取所有数据库
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var dbName string
		err := rows.Scan(&dbName)
		if err != nil {
			log.Fatal(err)
		}

		// 进入数据库
		_, err = db.Exec(fmt.Sprintf("USE %s", dbName))
		if err != nil {
			log.Fatal(err)
		}

		// 获取当前数据库的所有表
		tables, err := db.Query("SHOW TABLES")
		if err != nil {
			log.Fatal(err)
		}
		defer tables.Close()

		for tables.Next() {
			var tableName string
			err := tables.Scan(&tableName)
			if err != nil {
				log.Fatal(err)
			}

			// 获取当前表的所有字段
			columns, err := db.Query(fmt.Sprintf("SHOW COLUMNS FROM %s", tableName))
			if err != nil {
				log.Fatal(err)
			}
			defer columns.Close()

			for columns.Next() {
				var columnName string
				err := columns.Scan(&columnName)
				if err != nil {
					log.Fatal(err)
				}
				// 打印字段名
				fmt.Printf("数据库: %s, 表: %s, 字段: %s\n", dbName, tableName, columnName)
			}
		}
	}
}
