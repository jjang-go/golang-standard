package cmd

import (
	"github.com/jjang-go/golang-standard/config"
	"github.com/jjang-go/golang-standard/network"
	"github.com/jjang-go/golang-standard/repository"
	"github.com/jjang-go/golang-standard/service"
)

// 모든 항목들을 Cmd에서 구성 후 main에서 사용

type Cmd struct {
	config *config.Config

	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	c.repository = repository.Newrepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	c.network.ServerStart(c.config.Server.Port)

	return c
}
