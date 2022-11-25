/*
 * @Author: Charley
 * @Date: 2022-11-04 16:34:47
 * @LastEditors: Charley
 * @LastEditTime: 2022-11-08 17:46:41
 * @FilePath: /mospanel/web/web.go
 * @Description: In User Settings Edit
 */
package web

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"mime"
	"net/http"
	"time"

	"github.com/charleyzhu/mospanel/configs"
	"github.com/charleyzhu/mospanel/web/controllers"
	"github.com/charleyzhu/mospanel/web/sessions"
	gsession "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//go:embed assets/*
var assetsFS embed.FS

//go:embed html/*
var htmlFS embed.FS

var startTime = time.Now()

type wrapAssetsFS struct {
	embed.FS
}

func (f *wrapAssetsFS) Open(name string) (fs.File, error) {
	file, err := f.FS.Open("assets/" + name)
	if err != nil {
		return nil, err
	}
	return &wrapAssetsFile{
		File: file,
	}, nil
}

type wrapAssetsFile struct {
	fs.File
}

func (f *wrapAssetsFile) Stat() (fs.FileInfo, error) {
	info, err := f.File.Stat()
	if err != nil {
		return nil, err
	}
	return &wrapAssetsFileInfo{
		FileInfo: info,
	}, nil
}

type wrapAssetsFileInfo struct {
	fs.FileInfo
}

func (f *wrapAssetsFileInfo) ModTime() time.Time {
	return startTime
}

type WebServer struct {
	cfg *configs.PanelConfig
}

func NewWebServer(cfg *configs.PanelConfig) *WebServer {
	return &WebServer{
		cfg: cfg,
	}
}

func (w *WebServer) Run() error {
	r := w.initRouter()

	return r.Run(fmt.Sprintf("%s:%d", w.cfg.Bind, w.cfg.Port))
}

func (w *WebServer) initRouter() *gin.Engine {
	mime.AddExtensionType(".map", "text/plain")
	r := gin.Default()
	r.SetHTMLTemplate(template.Must(template.New("").ParseFS(htmlFS, "html/**/*")))
	r.StaticFS("assets", http.FS(&wrapAssetsFS{FS: assetsFS}))

	r.Use(gsession.Sessions("session", sessions.NewSessionStore(w.cfg.SessionSecret)))
	controllers.InitControllerRouter(r)
	return r
}
