package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type homeController struct {
	parent *applicationController
}

func newHomeController(app *applicationController) *homeController {
	return &homeController{parent: app}
}

func (controller *homeController) routes() {
	controller.parent.engine.StaticFile("/", "static/dist/index.html")
	controller.parent.engine.StaticFile("/favicon.icon", "static/dist/favicon.ico")
	controller.parent.engine.Static("/img", "static/dist/img")
	controller.parent.engine.Static("/css", "static/dist/css")
	controller.parent.engine.Static("/js", "static/dist/js")

	controller.parent.engine.LoadHTMLFiles("static/dist/index.html")
	controller.parent.engine.NoRoute(controller.error404)
}

func (controller *homeController) error404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "index.html", gin.H{})
}
