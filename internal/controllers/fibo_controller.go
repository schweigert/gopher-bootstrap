package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/gopher-bootstrap/internal/controllers/forms"
	"github.com/schweigert/gopher-bootstrap/internal/services"
)

type fiboController struct {
	parent  *applicationController
	service *services.Fibo
}

func newFiboController(app *applicationController) *fiboController {
	return &fiboController{parent: app, service: services.NewFibo()}
}

func (controller *fiboController) routes() {
	controller.parent.engine.GET("/api/fibo", controller.GET)
}

func (controller *fiboController) GET(c *gin.Context) {
	form := &forms.PositiveNumber{}
	if c.Bind(form) != nil {
		c.JSON(http.StatusOK, "value not found")
		return
	}

	c.JSON(http.StatusOK, controller.service.Fibo(form.Value))
}
