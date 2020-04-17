package main

import (
	"github.com/AlexeyRyabichev/ShowItGate"
	"log"
	"net/http"

	"github.com/AlexeyRyabichev/ShowItUsers/internal"
)

var cfgFile = "cfg.json"

func main() {
	nodeCfg, err := ShowItGate.ReadCfgFromJSON(cfgFile)
	if err != nil {
		log.Fatal(err)
	}

	if nodeCfg.Token == "" {
		if err := nodeCfg.RegisterNode(); err != nil {
			log.Fatalf("Cannot register node: %v", err)
		}

		if err := nodeCfg.SaveCfgToJSON(cfgFile); err != nil {
			log.Fatal(err)
		}
	}

	router := internal.NewRouter(nodeCfg)

	log.Printf("Server started")
	log.Fatal(http.ListenAndServe(":7054", router.Router))
}
