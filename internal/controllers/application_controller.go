package controllers

import "github.com/gin-gonic/gin"

type applicationController struct {
	engine *gin.Engine
	home   *homeController
}

// Bind an engine to controllers
func Bind(engine *gin.Engine) {
	app := &applicationController{engine: engine}

	app.home = newHomeController(app)

	app.routes()
}

func (app *applicationController) routes() {
	app.home.routes()
}
