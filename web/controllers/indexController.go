/*
 * @Author: Charley
 * @Date: 2022-11-05 16:22:50
 * @LastEditors: Charley
 * @LastEditTime: 2022-11-25 14:28:31
 * @FilePath: /mospanel/web/controllers/indexController.go
 * @Description: manager index
 */
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHome(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	ctx.HTML(http.StatusOK, "test.html", gin.H{})
}
