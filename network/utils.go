package network

import "github.com/gin-gonic/gin"

func (n *Network) registerGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.GET(path, handler...)
}

func (n *Network) registerPOST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.POST(path, handler...)
}

func (n *Network) registerUPDATE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.PUT(path, handler...)
}

func (n *Network) registerDELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.DELETE(path, handler...)
}

//response 형태 맞추기위한 util함수
func (n *Network) okResponse(c *gin.Context, result interface{}) {
	c.JSON(200, result)
}

func (n *Network) failedResponse(c *gin.Context, result interface{}) {
	c.JSON(400, result)
}
