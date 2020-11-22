package controllers

import "github.com/gin-gonic/gin"

type applicationController struct {
	engine  *gin.Engine
	home    *homeController
	fibo    *fiboController
	anagram *anagramController
}

// Bind an engine to controllers
func Bind(engine *gin.Engine) {
	app := &applicationController{engine: engine}

	app.home = newHomeController(app)
	app.fibo = newFiboController(app)
	app.anagram = newAnagramController(app)

	app.routes()
}

func (app *applicationController) routes() {
	app.home.routes()
	app.fibo.routes()
	app.anagram.routes()
}
