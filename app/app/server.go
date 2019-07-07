package app

import (
    "github.com/gin-gonic/gin"  
)

func Start() error {
    app := setup()
    return app.Run(":80")
}

func setup() *gin.Engine {
    app := gin.New()
    app.LoadHTMLGlob("./html/*")
    RouteV1(app)
    return app
}