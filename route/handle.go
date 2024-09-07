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
}

func registerStaticFiles(engine *gin.Engine, app *app.App) {
	group := engine.Group("")
	staticFile(group, app.PageFS, "/labor", "labor.html", "text/html")
	staticFile(group, app.PageFS, "/hospital", "hospital.html", "text/html")
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
