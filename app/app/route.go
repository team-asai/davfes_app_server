package app

import (
	"davfes_app/handlers"

	"github.com/gin-gonic/gin"
)

func RouteV1(app *gin.Engine) {
	jsonHandler := handlers.NewJsonHandler()
	imageHandler := handlers.NewImageHandler()
	logHandler := handlers.NewLogHandler()

	apiGroup := app.Group("api")
	{
		apiGroup.GET("/resources/:jsonfilename", jsonHandler.ResponseJson)
		apiGroup.GET("/json/edit/:jsonfilename", jsonHandler.EditJson)
		apiGroup.POST("/json/save/:jsonfilename", jsonHandler.SaveJson)
		apiGroup.GET("/images/news/icon/id=:id", imageHandler.ResNewsIcon)
		apiGroup.GET("/images/news/picture/id=:id", imageHandler.ResNewsPicture)
		apiGroup.GET("/images/seisaku/id=:id", imageHandler.ResSeisakuPicture)
		apiGroup.GET("/images/map/:place", imageHandler.ResMapImage)
		apiGroup.GET("/images/number/:num", imageHandler.ResNumberImage)
		apiGroup.GET("/images/place/:place", imageHandler.ResPlaceImage)
		apiGroup.GET("/images/others/:filename", imageHandler.ResOtherImage)
		apiGroup.GET("/images/upload", imageHandler.UploadImage)
		apiGroup.POST("/images/save", imageHandler.SaveImage)
		apiGroup.POST("/post/beacon", logHandler.SaveBeaconLog)
		apiGroup.POST("/post/favorite", logHandler.SaveFavoriteLog)
	}
}
