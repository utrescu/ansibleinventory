package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/utrescu/listIP"
	"gopkg.in/yaml.v2"
)

type configuration struct {
	Name    string
	Address []string
}

type configurationList struct {
	Networks []configuration
}

var (
	portNumber int
	timeout    string
	filename   string
	debug      bool
)

func init() {
	flag.IntVar(&portNumber, "port", 22, "Port to scan")
	flag.IntVar(&portNumber, "p", 22, "Port to scan")
	flag.StringVar(&timeout, "timeout", "1000ms", "Network timeout")
	flag.StringVar(&timeout, "t", "1000ms", "Network timeout")
	flag.StringVar(&filename, "input", "conf.yaml", "Name of configuration time")
	flag.StringVar(&filename, "i", "conf.yaml", "Name of configuration time")
	flag.BoolVar(&debug, "debug", false, "verbose output")
}

func outputFormat(label string, resultats []string) {
	if len(resultats) > 0 {
		fmt.Println("[" + label + "]")

		for i := range resultats {
			fmt.Println(resultats[i])
		}
		fmt.Println()
	}
}

func main() {

	flag.Parse()

	_, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatal("Incorrect Duration\n", err)
	}

	// Processar el fitxer yaml
	var configs configurationList

	source, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("File '" + filename + "' not found")
	}

	err = yaml.Unmarshal(source, &configs)
	if err != nil {
		fmt.Println("Incorrect format")
		panic(err)
	}

	// startTime := time.Now()

	for tria := range configs.Networks {
		if debug == true {
			fmt.Println("... Trying ", configs.Networks[tria].Address)
		}
		rangs := configs.Networks[tria].Address
		resultats, _ := listIP.Check(rangs, portNumber, timeout)
		outputFormat(configs.Networks[tria].Name, resultats)
	}

	// scanDuration := time.Since(startTime)
	// fmt.Println(scanDuration)
}
