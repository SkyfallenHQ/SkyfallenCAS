package logging

import "log"

func Fatal(err error) {

	if err != nil {

		log.Println("There was a critical error in a function.")
		log.Println("SkyfallenCAS could not proceed without successfully executing this action.")
		log.Println(err)
		log.Fatal("SkyfallenCAS has crashed.")

	}

}