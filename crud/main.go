// main
package main

import (
	"crud/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//ID查询
	router.POST("/show", userController.ShowByKey)

	//查询所有
	router.GET("/show", userController.ShowAll)

	//添加
	router.POST("/create", userController.Create)

	//删除
	router.POST("/delete", userController.DeleteUser)

	//更新
	router.POST("/update", userController.UpdateUser)

	router.Run()
}
