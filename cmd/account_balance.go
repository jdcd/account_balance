package main

import (
	"fmt"
	"os"

	"github.com/jdcd/account_balance/config"
	"github.com/jdcd/account_balance/internal/infrastructure/http/server"
	"github.com/jdcd/account_balance/pkg"
)

func main() {
	appConfiguration := config.GetConfigurations()
	err := appConfiguration.CheckData()
	if err != nil {
		pkg.ErrorLogger().Fatalf("error reading configuration: %s", err)
		return
	}
	router := server.SetupRouter(config.GetRouterDependencies(appConfiguration))

	port := os.Getenv("PORT")

	if err := router.Run(); err != nil {
		errorDetail := fmt.Sprintf("unable to start app on the port: %s , %s", port, err.Error())
		pkg.ErrorLogger().Fatal(errorDetail)
	}
}
