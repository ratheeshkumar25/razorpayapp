package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ratheesh-g-kumar/controller"
	"github.com/ratheesh-g-kumar/initializer"
)

func init(){
	initializer.LoadEnv()
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.GET("/app",controller.App)
	r.GET("/payment-success",controller.PaymentStatus)
	r.Run(":3000")

}