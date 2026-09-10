package main

import (
	"Assignment-1/builder"
	"fmt"
	"log"
)

func main() {
	director := builder.NewEMailDirector()

	objBuilder := builder.NewConcreteEmailBuilder()
	director.MakeWelcomeEmail(objBuilder)
	emailObj, err := objBuilder.GetResult()
	if err != nil {
		log.Fatalf("Failed to build emailobject: %v", err)
	}

	fmt.Println(emailObj)
	fmt.Println()

	previewBuilder := builder.NewEmailPreviewBuilder()
	director.MakePasswordResetEmail(previewBuilder)

	previewString, err := previewBuilder.GetResult()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(previewString)
}
