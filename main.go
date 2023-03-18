package main

import (
	"os"

	"github.com/Rbsn-joses/microsservicos/authentication"
	"github.com/Rbsn-joses/microsservicos/cursos"
)

func main() {
	ms := os.Getenv("MICROSSERVICE_TYPE")
	if ms == "auth" {
		authentication.MicroserviceAuthentication()
	} else {
		cursos.MicroserviceCursos()
	}
	//

}
