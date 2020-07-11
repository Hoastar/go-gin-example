package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/hoastar/go-gin-example/middleware/protobuf"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/hoastar/go-gin-example/docs"
	"github.com/hoastar/go-gin-example/pkg/setting"
	"github.com/hoastar/go-gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	/*r.Use(gin.BasicAuth(gin.Accounts{
		"admin": "q1w2e3r4",
	}))
	 */
	gin.SetMode(setting.ServerSetting.RunMode)
	//auth验证
	//r.GET("/auth", api.GetAuth)
	// gin.BasicAuth
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Home")
	})

	r.GET("/protobuf", func(c *gin.Context) {
		data := &protobuf.User {
			Name: "张三",
			Age: 23,
		}

		c.ProtoBuf(http.StatusOK, data)
	})
	adminGroup := r.Group("/admin")
	adminGroup.Use(gin.BasicAuth(gin.Accounts{
		"admin": "q1w2e3r4",
	}))
	// adminGroup.Use(costTime.CostTime())

	adminGroup.GET("/index", func(c *gin.Context) {
		c.JSON(200, "后台首页")
	})



	apiv1 := r.Group("/api/v1")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//文章相关
		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取单个文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)

	}
	return r
}
