/*
 * @Author: Charley
 * @Date: 2022-11-05 16:38:57
 * @LastEditors: Charley
 * @LastEditTime: 2022-11-09 09:30:22
 * @FilePath: /mospanel/web/controllers/loginController.go
 * @Description: login controller
 */
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"title": "登陆",
	})
}
