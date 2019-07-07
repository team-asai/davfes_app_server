package handlers

import (
    "github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
)

type jsonHandlerInterface interface {
	ResponseJson(c *gin.Context)
	EditJson(x *gin.Context)
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

// 指定された名前のjsonを編集するメソッド(今は適当にHTML返してるだけ。後日作成予定)
func (h *jsonHandler) EditJson(c *gin.Context) {
	// jsonFileName := c.Param("jsonfilename")
	c.HTML(http.StatusOK, "edit_json.html", gin.H{
		"content": "ハゲハゲのハゲ",
	})
}