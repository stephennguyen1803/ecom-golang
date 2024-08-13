package user

import "github.com/gin-gonic/gin"

type ProductRouter struct {
}

func (pr *ProductRouter) InitProductRouter(routers *gin.RouterGroup) {
	//public route - external router
	productRouterPublic := routers.Group("/product")
	{
		//call middleware
		productRouterPublic.GET("/search")
		productRouterPublic.GET("/detail/:id")
	}
	//private route - internal router
}
