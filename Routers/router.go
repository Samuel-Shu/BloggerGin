package Routers

import (
	"ginblog/Api"
	"ginblog/Middleware"
	"ginblog/Utils"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(Utils.AppMode)
	r := gin.New()
	r.HTMLRender = createMyRender()
	r.Use(Middleware.Log())
	r.Use(gin.Recovery())
	r.Use(Middleware.Cors())
	
	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})
	
	/*
	后台管理路由接口
	 */
	auth := r.Group("Api/v1")
	auth.Use(Middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", Api.GetUsers)
		auth.PUT("user/:id", Api.EditUser)
		auth.DELETE("user/:id", Api.DeleteUser)
		    //修改密码
		auth.PUT("admin/changepw/:id", Api.ChangeUserPassword)
		// 分类模块的路由接口
		auth.GET("admin/category", Api.GetCate)
		auth.POST("category/add", Api.AddCategory)
		auth.PUT("category/:id", Api.EditCate)
		auth.DELETE("category/:id", Api.DeleteCate)
		// 文章模块的路由接口
		auth.GET("admin/article/info/:id", Api.GetArtInfo)
		auth.GET("admin/article", Api.GetArt)
		auth.POST("article/add", Api.AddArticle)
		auth.PUT("article/:id", Api.EditArt)
		auth.DELETE("article/:id", Api.DeleteArt)
		// 上传文件
		auth.POST("Upload-photo", Api.UpLoad)
		// 更新个人设置
		auth.GET("admin/profile/:id", Api.GetProfile)
		auth.PUT("profile/:id", Api.UpdateProfile)
		// 评论模块
		auth.GET("comment/list", Api.GetCommentList)
		auth.DELETE("delcomment/:id", Api.DeleteComment)
		auth.PUT("checkcomment/:id", Api.CheckComment)
		auth.PUT("uncheckcomment/:id", Api.UncheckComment)
	}
	
	/*
	前端展示页面接口
	 */
	router := r.Group("Api/v1")
	{
		// 用户信息模块
		router.POST("user/add", Api.AddUser)
		router.GET("user/:id", Api.GetUserInfo)
		router.GET("users", Api.GetUsers)

		// 文章分类信息模块
		router.GET("category", Api.GetCate)
		router.GET("category/:id", Api.GetCateInfo)

		// 文章模块
		router.GET("article", Api.GetArt)
		router.GET("article/list/:id", Api.GetCateArt)
		router.GET("article/info/:id", Api.GetArtInfo)

		// 登录控制模块
		router.POST("login", Api.Login)
		router.POST("loginfront", Api.LoginFront)

		// 获取个人设置信息
		router.GET("profile/:id", Api.GetProfile)

		// 评论模块
		router.POST("addcomment", Api.AddComment)
		router.GET("comment/info/:id", Api.GetComment)
		router.GET("commentfront/:id", Api.GetCommentListFront)
		router.GET("commentcount/:id", Api.GetCommentCount)
	}

	_ = r.Run(Utils.HttpPort)

}
