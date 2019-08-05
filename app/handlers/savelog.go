package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type logHandlerInterface interface {
	SaveBeaconLog(c *gin.Context)
	SaveFavoriteLog(c *gin.Context)
}

func NewLogHandler() logHandlerInterface {
	return &logHandler{}
}

type logHandler struct {
}

type FavoriteLog struct {
	Os          string `json:"os" binding:"required"`
	DeviceToken string `json:"deviceToken" binding:"required"`
	FavoIdList  string `json:"favoIdList" binding:"required"`
}

type BeaconLog struct {
	Os          string `json:"os" binding:"required"`
	DeviceToken string `json:"deviceToken" binding:"required"`
	Uuid        string `json:"uuid" binding:"required"`
	Major       string `json:"major" binding:"required"`
	Minor       string `json:"minor" binding:"required"`
}

func (h *logHandler) SaveBeaconLog(c *gin.Context) {
	// インスタンス作成
	var logData BeaconLog
	// 型に合わせてPOSTされたJSONをパースして変数spotに格納してくれる超便利なメソッド
	if c.BindJSON(&logData) == nil {
		log.Println("Binding success.....................")
		log.Println(logData)
	} else {
		log.Println("Binding failed......................")
	}
	// DBへの接続
	db := ConnectDB()
	// この関数が終わったら自動的に接続を閉じるように
	defer db.Close()
	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	// マイグレーション
	db.AutoMigrate(&BeaconLog{})
	// もらった情報をINSERT
	db.Create(&logData)
	// 一応正常に受け取れた返事
	c.JSON(http.StatusOK, &logData)
}

func (h *logHandler) SaveFavoriteLog(c *gin.Context) {
	// インスタンス作成
	var logData FavoriteLog
	// 型に合わせてPOSTされたJSONをパースして変数spotに格納してくれる超便利なメソッド
	if c.BindJSON(&logData) == nil {
		log.Println("Binding success.....................")
		log.Println(logData)
	} else {
		log.Println("Binding failed......................")
	}
	// DBへの接続
	db := ConnectDB()
	// この関数が終わったら自動的に接続を閉じるように
	defer db.Close()
	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	// マイグレーション
	db.AutoMigrate(&FavoriteLog{})
	// もらった情報をINSERT
	db.Create(&logData)
	// 一応正常に受け取れた返事
	c.JSON(http.StatusOK, &logData)
}
