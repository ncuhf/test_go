// userController
package userController

import (
	"crud/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var db = user.Connection()

//创建用户
func Create(tx *gin.Context) {

	name := tx.PostForm("Name")
	number := tx.PostForm("Number")
	passWord := tx.PostForm("Password")

	db.Create(&user.User{Name: name, Number: number, PassWord: passWord})

	tx.Redirect(http.StatusMovedPermanently, "/show")
}

//ID查询
func ShowByKey(tx *gin.Context) {
	u := new(user.User)
	id, _ := strconv.Atoi(tx.PostForm("id"))
	db.Where("id = ?", id).First(&u)

	tx.JSON(200, gin.H{"user": u})
}

//查询所有
func ShowAll(tx *gin.Context) {
	users := make([]user.User, 0)
	db.Find(&users)
	for k, v := range users {
		key := "user" + strconv.Itoa(k+1)
		tx.JSON(200, gin.H{key: v})
	}

}

//删除
func DeleteUser(tx *gin.Context) {
	u := new(user.User)
	id, _ := strconv.Atoi(tx.PostForm("id"))
	db.Where("id = ?", id).First(&u)
	db.Delete(u)
	tx.Redirect(http.StatusMovedPermanently, "/show")
	tx.JSON()
}

//更新
func UpdateUser(tx *gin.Context) {
	u := new(user.User)
	id, _ := strconv.Atoi(tx.PostForm("id"))
	name := tx.PostForm("Name")
	number := tx.PostForm("Number")

	db.Where("id = ?", id).First(&u)
	db.Model(u).Updates(user.User{Name: name, Number: number})
	tx.Redirect(http.StatusMovedPermanently, "/show")
}
