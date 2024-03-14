package network

import (
	"github.com/gin-gonic/gin"
	"github.com/jjang-go/golang-standard/service"
)

// api,route

type Network struct {
	engine *gin.Engine

	service *service.Service
}

func NewNetwork(service *service.Service) *Network {
	r := &Network{
		engine: gin.New(),
	}

	newUserRouter(r)

	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engine.Run(port)
}
