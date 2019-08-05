package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type jsonHandlerInterface interface {
	ResponseJson(c *gin.Context)
	EditJson(c *gin.Context)
	SaveJson(c *gin.Context)
}

func NewJsonHandler() jsonHandlerInterface {
	return &jsonHandler{}
}

type jsonHandler struct {
}

// 指定された名前のjsonを返すメソッド
func (h *jsonHandler) ResponseJson(c *gin.Context) {
	jsonFileName := c.Param("jsonfilename")

	jsonMap, err := ioutil.ReadFile("./resources/" + jsonFileName + ".json")
	if err != nil {
		c.JSON(http.StatusOK, err)
	} else {
		c.String(http.StatusOK, string(jsonMap))
	}
}

func (h *jsonHandler) EditJson(c *gin.Context) {
	jsonFileName := c.Param("jsonfilename")

	jsonStr, err := ioutil.ReadFile("./resources/" + jsonFileName + ".json")
	if err == nil {
		c.HTML(http.StatusOK, "edit_json.html", gin.H{
			"json": string(jsonStr),
			"path": string("../save/" + jsonFileName),
		})
	} else {
		c.String(http.StatusOK, string("404 Not Found"))
	}
}

func (h *jsonHandler) SaveJson(c *gin.Context) {
	jsonFileName := c.Param("jsonfilename")
	c.Request.ParseForm()
	fmt.Println(c.Request.Form["json"])
	fmt.Println(jsonFileName)
	json := c.Request.Form["json"]
	data := []byte(strings.Join(json, ""))
	err := ioutil.WriteFile("./resources/"+jsonFileName+".json", data, 0755)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, string("ファイルを保存しました"))
}
