package config

import "github.com/jdcd/account_balance/internal/infrastructure/http/server"

// GetRouterDependencies configures dependencies injection
func GetRouterDependencies(config *AppConfiguration) *server.RouterDependencies {
	return &server.RouterDependencies{
		CheckController: &server.PingController{},
	}
}
