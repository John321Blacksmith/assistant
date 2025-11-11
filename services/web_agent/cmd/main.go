// Package main is an entrypoint
// for initialization of all business engine.
package main

import (
	"log"
	"os"
	"osint_agent/services/web_agent/cmd/engine"
	"osint_agent/services/web_agent/internal/adapter/config"
)

func main() {
	// input := `Little/ did they/ know the ice cream was invented in Russia.
	// 		This kind/ of food was liked by almost everyone/ in the world.
	// 		  because of it's taste/ and unique, nutrient/ features/.`

	// load configs
	cfgPath := os.Getenv("CONFIG_PATH")
	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatalf("Cannot load settings and start the service \"web_agent\": %v", err)
	}

	// implement engine
	engine, err := engine.NewEngine(cfg)
	if err != nil {
		log.Fatalf("Cannot start the engine because: %v", err)
	}

	// start the engine
	errChan := engine.Run()

	// error check
	potentialError := <-errChan
	if potentialError != nil {
		log.Fatalf("Error occurred during the engine startup: %v", err)
	}

}
