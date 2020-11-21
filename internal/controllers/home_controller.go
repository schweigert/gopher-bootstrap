package controllers

type homeController struct {
	parent *applicationController
}

func newHomeController(app *applicationController) *homeController {
	return &homeController{parent: app}
}

func (controller *homeController) routes() {
	controller.parent.engine.Static("/", "static/dist")
}
