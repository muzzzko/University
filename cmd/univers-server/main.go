package main

import (
	"log"
	"univers/pkg/config"
	"univers/pkg/restapi"
	"univers/pkg/restapi/operations"

	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
)

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewUniversAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Univers"
	parser.LongDescription = swaggerSpec.Spec().Info.Description

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	server.SetHandler(configureAPI(api))
	conf := config.GetConfig()
	server.Host = conf.AppHost
	server.Port = conf.AppPort
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
