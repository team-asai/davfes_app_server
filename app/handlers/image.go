package handlers

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
)

type imageHandlerInterface interface {
	ResNewsIcon(c *gin.Context)
	ResNewsPicture(c *gin.Context)
	ResSeisakuPicture(c *gin.Context)
	ResMapImage(c *gin.Context)
	ResNumberImage(c *gin.Context)
	ResPlaceImage(c *gin.Context)
	ResOtherImage(c *gin.Context)
	UploadImage(c *gin.Context)
	SaveImage(c *gin.Context)
}

func NewImageHandler() imageHandlerInterface {
	return &imageHandler{}
}

type imageHandler struct {
}

// 指定されたidのニュースアイコンを返すメソッド
func (h *imageHandler) ResNewsIcon(c *gin.Context) {
	id := c.Param("id")
	path := "./images/news/icon/i" + string(id) + ".jpg"

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
	path := "./images/news/picture/p" + string(id) + ".jpg"

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
	path := "./images/seisakuKyoushitsu/k" + string(id) + ".jpg"

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		c.JSON(http.StatusOK, err)
	} else {
		c.File(path)
	}
}

// 指定されたidの制作教室の画像を返すメソッド
func (h *imageHandler) ResMapImage(c *gin.Context) {
	place := c.Param("place")
	path := "./images/map/" + string(place) + ".png"

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		c.JSON(http.StatusOK, err)
	} else {
		c.File(path)
	}
}

// 指定されたidの制作教室の画像を返すメソッド
func (h *imageHandler) ResNumberImage(c *gin.Context) {
	num := c.Param("num")
	path := "./images/number/number_" + string(num) + ".png"

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		c.JSON(http.StatusOK, err)
	} else {
		c.File(path)
	}
}

// 指定されたidの制作教室の画像を返すメソッド
func (h *imageHandler) ResPlaceImage(c *gin.Context) {
	place := c.Param("place")
	path := "./images/place/" + string(place) + ".png"

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		c.JSON(http.StatusOK, err)
	} else {
		c.File(path)
	}
}

func (h *imageHandler) ResOtherImage(c *gin.Context) {
	filename := c.Param("filename")
	path := "./images/others/" + string(filename)

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		c.JSON(http.StatusOK, err)
	} else {
		c.File(path)
	}
}

func (h *imageHandler) UploadImage(c *gin.Context) {
	c.HTML(http.StatusOK, "upload_image.html", gin.H{})
}

func (h *imageHandler) SaveImage(c *gin.Context) {
	c.Request.ParseMultipartForm(1000)
	t := c.Request.PostForm["type"]
	id := c.Request.PostForm["id"]
	log.Println(t)
	log.Println(id)
	img, err := imageupload.Process(c.Request, "upload_file")

	if err != nil {
		panic(err)
	}

	thumb, err := imageupload.ThumbnailPNG(img, 300, 300)

	if err != nil {
		panic(err)
	}

	img.Save("images/news/picture/p" + strings.Join(id, "") + ".jpg")
	thumb.Save("images/news/icon/i" + strings.Join(id, "") + ".jpg")

	c.String(http.StatusOK, string("ファイルを保存しました"))
}
