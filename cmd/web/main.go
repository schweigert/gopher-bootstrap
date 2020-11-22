package main

import (
	"github.com/gin-gonic/gin"
	"github.com/schweigert/gopher-bootstrap/internal/controllers"
)

func main() {
	engine := gin.Default()

	controllers.Bind(engine)

	engine.Run()
}
