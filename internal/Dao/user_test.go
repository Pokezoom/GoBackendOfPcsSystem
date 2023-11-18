/*
 * @Author: sucy suchunyu1998@gmail.com
 * @Date: 2023-11-17 19:28:06
 * @LastEditTime: 2023-11-17 21:21:02
 * @LastEditors: Suchunyu
 * @Description: 
 * @FilePath: /GoBackendOfPcsSystem/internal/Dao/user_test.go
 * Copyright (c) 2023 by Suchunyu, All Rights Reserved. 
 */
 package Dao

 import (
	 "fmt"
	 "testing"
 )
 

func TestUser(t *testing.T){
	fmt.Println("测试user中的函数")
	// t.Run("插入用户: ",testCreate)
	t.Run("验证Login: ",testLogin)
	// t.Run("验证注册: ",testRegist)


}

func testLogin(t *testing.T){
	user,_ :=CheckUserNameAndPassword("admin","123456")
	fmt.Println("获取用户信息室：",user)

}

func testRegist(t *testing.T){
	user,_ :=CheckUserName("admin")
	fmt.Println("获取用户信息室：",user)
}

func testCreate(t *testing.T){
	CreateUser("admin3","123456","admin@gamil.com")
}