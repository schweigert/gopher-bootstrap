package controllers

import "github.com/gin-gonic/gin"

type applicationController struct {
	engine *gin.Engine
	home   *homeController
	fibo   *fiboController
}

// Bind an engine to controllers
func Bind(engine *gin.Engine) {
	app := &applicationController{engine: engine}

	app.home = newHomeController(app)
	app.fibo = newFiboController(app)

	app.routes()
}

func (app *applicationController) routes() {
	app.fibo.routes()
	app.home.routes()
}
