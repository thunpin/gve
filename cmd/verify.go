package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/thunpin/gve/pkg/business/network"
	"github.com/thunpin/gve/pkg/business/parser"
	"github.com/thunpin/gve/pkg/business/reader"
)

var filePath = flag.String("filePath", "", "istio yaml file path")

func main() {
	flag.Parse()

	log.Println("Using " + *filePath)

	content, err := reader.FromFile(*filePath)
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}

	element, err := parser.Yaml(content)
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}

	err = network.VerifyEndpoint(element)
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}

	fmt.Println("SUCCESS")
}
