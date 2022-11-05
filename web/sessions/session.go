/*
 * @Author: Charley
 * @Date: 2022-11-04 17:52:35
 * @LastEditors: Charley
 * @LastEditTime: 2022-11-04 17:52:46
 * @FilePath: /mospanel/web/sessions/session.go
 * @Description: session store
 */
package sessions

import (
	"github.com/gin-contrib/sessions/cookie"
)

// SessionSecret init store
func NewSessionStore(secret string) cookie.Store {
	return cookie.NewStore([]byte(secret))
}
