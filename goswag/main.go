package main

import (
	"github.com/diegoclair/go_boilerplate/application/service"
	"github.com/diegoclair/go_boilerplate/infra/config"
	"github.com/diegoclair/go_boilerplate/transport/rest"
)

// the swagger documentation will be generated by https://github.com/diegoclair/goswag
func main() {
	//	@title			Go Boilerplate API
	//	@version		1.0
	//	@description	This is an API documentation for go boilerplate

	//	@contact.name	Email
	//	@contact.email	diego93rodrigues@gmail.com

	//	@host			localhost:5000
	//	@schemes		http
	//	@servers.url http://localhost:5000

	cfg := &config.Config{
		Log: config.LogConfig{},
	}

	server := rest.NewRestServer(&service.Services{}, nil, cfg)
	server.Router.GenerateSwagger()
}
