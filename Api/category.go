package Api

import (
	"ginblog/Models"
	"ginblog/Utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddCategory 添加分类
func AddCategory(c *gin.Context) {
	var data Models.Category
	_ = c.ShouldBindJSON(&data)
	code := Models.CheckCategory(data.Name)
	if code == errmsg.SUCCSE {
		Models.CreateCate(&data)
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// GetCateInfo 查询分类信息
func GetCateInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := Models.GetCateInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

// GetCate 查询分类列表
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := Models.GetCate(pageSize, pageNum)
	code := errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 查询单个分类
//func GetCateInfo(c *gin.Context)  {
//	id, _ := strconv.Atoi(c.Param("id"))
//
//	data,code := Models.GetCateInfo(id)
//
//	c.JSON(http.StatusOK, gin.H{
//		"status":  code,
//		"data":    data,
//		"message": errmsg.GetErrMsg(code),
//	})
//}

// EditCate 编辑分类名
func EditCate(c *gin.Context) {
	var data Models.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := Models.CheckCategory(data.Name)
	if code == errmsg.SUCCSE {
		Models.EditCate(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		c.Abort()
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DeleteCate 删除用户
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := Models.DeleteCate(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
