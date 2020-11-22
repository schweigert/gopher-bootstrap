package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/gopher-bootstrap/internal/controllers/forms"
	"github.com/schweigert/gopher-bootstrap/internal/services"
)

type anagramController struct {
	parent  *applicationController
	service *services.Anagram
}

func newAnagramController(app *applicationController) *anagramController {
	return &anagramController{parent: app, service: services.NewAnagram()}
}

func (controller *anagramController) routes() {
	controller.parent.engine.GET("/api/anagram", controller.GET)
}

func (controller *anagramController) GET(c *gin.Context) {
	form := &forms.TwoStrings{}
	if c.Bind(form) != nil {
		c.JSON(http.StatusOK, "value not found")
		return
	}

	c.JSON(http.StatusOK, controller.service.Check(form.A, form.B))
}
