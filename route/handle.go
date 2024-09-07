package route

import (
	"codefast_2024/app"
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerHandlers(engine *gin.Engine, app *app.App) {
	registerStaticFiles(engine, app)
}

func registerStaticFiles(engine *gin.Engine, app *app.App) {
	group := engine.Group("")
	staticFile(group, app.PageFS, "/labor", "labor_check.html", "text/html")
}

func staticFile(r *gin.RouterGroup, fs embed.FS, path, filePath, contentType string) {
	r.GET(path, func(ctx *gin.Context) {
		file, err := fs.ReadFile(filePath)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.Header("Cache-Control", "private, max-age=3600")
		ctx.Data(http.StatusOK, contentType, file)
	})
}
