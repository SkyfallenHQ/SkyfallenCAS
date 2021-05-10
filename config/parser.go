package config

import (
	"github.com/joho/godotenv"
	"log"
)

func Parse() map[string]string {

	log.Printf("Parsing SkyfallenCAS configuration file.")

	var SkyfallenCASConf map[string]string

	SkyfallenCASConf, err := godotenv.Read("CAS.skyfallenconfig")

	if err != nil {
		log.Println("An error occurred in during SkyfallenCAS boot.")
		log.Println("Please check if the config file exists and is valid.")
		log.Println(err)
		log.Fatal("SkyfallenCAS has crashed.")
		return nil
	}

	return SkyfallenCASConf

}
