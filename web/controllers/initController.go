/*
 * @Author: Charley
 * @Date: 2022-11-04 17:49:47
 * @LastEditors: Charley
 * @LastEditTime: 2022-11-05 16:46:28
 * @FilePath: /mospanel/web/controllers/initController.go
 * @Description: controller router init
 */
package controllers

import "github.com/gin-gonic/gin"

func InitControllerRouter(r *gin.Engine) {
	r.GET("/", IndexHome)
	r.GET("/login", LoginHome)
}
