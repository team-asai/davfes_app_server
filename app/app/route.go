package app

import (
	"github.com/gin-gonic/gin"
	"davfes_app/handlers"
)

func RouteV1(app *gin.Engine) {
	jsonHandler := handlers.NewJsonHandler()
	imageHandler := handlers.NewImageHandler()

	apiGroup := app.Group("api")
	{
		apiGroup.GET("/resources/:jsonfilename", jsonHandler.ResponseJson)
		apiGroup.GET("/images/news/icon/id=:id", imageHandler.ResNewsIcon)
		apiGroup.GET("/images/news/picture/id=:id", imageHandler.ResNewsPicture)
		apiGroup.GET("/images/seisaku/id=:id", imageHandler.ResSeisakuPicture)

	}

	editGroup := app.Group("edit")
	{
		editGroup.GET("/json/:jsonfilename", jsonHandler.EditJson)
	}	
}