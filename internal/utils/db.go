/*
 * @Author: sucy suchunyu1998@gmail.com
 * @Date: 2023-11-17 20:03:25
 * @LastEditTime: 2023-11-17 21:29:25
 * @LastEditors: Suchunyu
 * @Description: 
 * @FilePath: /GoBackendOfPcsSystem/internal/utils/db.go
 * Copyright (c) 2023 by Suchunyu, All Rights Reserved. 
 */
package utils

import(
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
)

func init(){
	Db,err = sql.Open("mysql","root:123456@tcp(localhost:3306)/test")
	if err!=nil{
		panic(err.Error())
	}
}