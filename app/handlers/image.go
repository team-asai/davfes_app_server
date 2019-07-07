package handlers

import (
    "github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type imageHandlerInterface interface {
	ResNewsIcon(c *gin.Context)
	ResNewsPicture(c *gin.Context)
	ResSeisakuPicture(c *gin.Context)
}

func NewImageHandler() imageHandlerInterface {
    return &imageHandler{}
}

type imageHandler struct {
}

// 指定されたidのニュースアイコンを返すメソッド
func (h *imageHandler) ResNewsIcon(c *gin.Context) {
	id := c.Param("id")
	path := "./images/News/Icon/i" + string(id) + ".jpg"
	
	file, err := os.Open(path)
    defer file.Close()
	if err != nil {
		c.JSON(http.StatusOK, err)
	} else {
		c.File(path)
	}
}

// 指定されたidのニュース画像を返すメソッド
func (h *imageHandler) ResNewsPicture(c *gin.Context) {
	id := c.Param("id")
	path := "./images/News/Picture/p" + string(id) + ".jpg"
	
	file, err := os.Open(path)
    defer file.Close()
	if err != nil {
		c.JSON(http.StatusOK, err)
	} else {
		c.File(path)
	}
}

// 指定されたidの制作教室の画像を返すメソッド
func (h *imageHandler) ResSeisakuPicture(c *gin.Context) {
	id := c.Param("id")
	path := "./images/SeisakuKyoushitsu/k" + string(id) + ".jpg"
	
	file, err := os.Open(path)
    defer file.Close()
	if err != nil {
		c.JSON(http.StatusOK, err)
	} else {
		c.File(path)
	}
}