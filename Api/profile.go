package Api

import (
	"ginblog/Models"
	"ginblog/Utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProfile(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := Models.GetProfile(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func UpdateProfile(c *gin.Context) {
	var data Models.Profile
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := Models.UpdateProfile(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}