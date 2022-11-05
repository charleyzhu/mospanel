/*
 * @Author: Charley
 * @Date: 2022-11-05 16:22:50
 * @LastEditors: Charley
 * @LastEditTime: 2022-11-05 16:56:05
 * @FilePath: /mospanel/web/controllers/indexController.go
 * @Description: manager index
 */
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
