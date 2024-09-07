package route

import (
	"codefast_2024/app"
	"codefast_2024/controller"
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerHandlers(engine *gin.Engine, app *app.App) {
	registerStaticFiles(engine, app)
	registerApi(engine, app)
	engine.GET("/lunar", controller.Lunar(app))
}

func registerStaticFiles(engine *gin.Engine, app *app.App) {
	group := engine.Group("")
	staticFile(group, app.PageFS, "/labor", "labor.html", "text/html")
	staticFile(group, app.PageFS, "/hospital", "hospital.html", "text/html")
	staticFile(group, app.PageFS, "/lottery", "lottery.html", "text/html")
	staticFile(group, app.PageFS, "/lottery.css", "lottery.css", "text/css")
	staticFile(group, app.PageFS, "/lottery.js", "lottery.js", "text/javascript")
	staticFile(group, app.PageFS, "/right_arrow.png", "right_arrow.png", "image/png")
}

func staticFile(r *gin.RouterGroup, fs embed.FS, path, filePath, contentType string) {
	r.GET(path, func(ctx *gin.Context) {
		file, err := fs.ReadFile(filePath)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		ctx.Header("Cache-Control", "private, max-age=0, no-cache")
		ctx.Data(http.StatusOK, contentType, file)
	})
}

func registerApi(engine *gin.Engine, app *app.App) {
	apiGroup := engine.Group("/api")
	questionGroup := apiGroup.Group("/question")

	staticFile(questionGroup, app.PageFS, "/labor", "labor.json", "application/json")

	answerGroup := apiGroup.Group("/answer")
	answerGroup.POST("/labor", controller.AnswerLabor)
}
