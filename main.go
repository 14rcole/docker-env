package main

import (
	"fmt"
	"strings"
	"github.com/docker-env/operatingsys/operatingsys"
	"github.com/docker-env/errors/errors"
)

func main() {
	var os string
	//User picks operating system
	for true {
		fmt.Println("Input operating system: ")
		_, err = fmt.Scanf("%s", &os)
		if err != nil {
			errors.InputError("Operating System input")
		} else {
			break
		}
	}

	operatingSystem := operatingsys.GetOS(os);

	//User picks software to be installed through package manager
}
