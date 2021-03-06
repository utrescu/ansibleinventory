package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/utrescu/listIP"
	"gopkg.in/yaml.v2"
)

type configuration struct {
	Name     string
	Networks []string
}

type configurationList struct {
	Groups []configuration
}

var (
	portNumber int
	timeout    string
	parallel   int
	filename   string
	outputFile string
	debug      bool
	outFile    *os.File
)

func init() {
	flag.IntVar(&portNumber, "port", 22, "Port to scan")
	flag.IntVar(&portNumber, "p", 22, "Port to scan")
	flag.IntVar(&parallel, "parallel", 32, "Number of simultanius connections")
	flag.StringVar(&timeout, "timeout", "1000ms", "Network timeout")
	flag.StringVar(&timeout, "t", "1000ms", "Network timeout")
	flag.StringVar(&filename, "input", "conf.yaml", "Name of configuration time")
	flag.StringVar(&filename, "i", "conf.yaml", "Name of configuration time")
	flag.StringVar(&outputFile, "o", "", "Output file")
	flag.StringVar(&outputFile, "output", "", "Output file")
	flag.BoolVar(&debug, "debug", false, "verbose output")
}

func outputFormat(outFile *os.File, label string, resultats []string) {

	if outFile == nil {
		outFile = os.Stdout
	}
	if len(resultats) > 0 {

		fmt.Fprintln(outFile, "["+label+"]")

		for i := range resultats {
			fmt.Fprintln(outFile, resultats[i])
		}
		fmt.Fprintln(outFile, "")
	}
}

func main() {

	flag.Parse()

	_, err := time.ParseDuration(timeout)
	if err != nil {
		fmt.Println("Incorrect Duration\n", err)
		os.Exit(-1)
	}

	// Processar el fitxer yaml
	var configs configurationList

	source, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File '" + filename + "' not found")
		os.Exit(-1)
	}

	if outputFile != "" {
		outFile, err = os.OpenFile(outputFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
		defer outFile.Close()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
	}

	err = yaml.UnmarshalStrict(source, &configs)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	// startTime := time.Now()

	for tria := range configs.Groups {
		if debug == true {
			fmt.Println("... Trying", configs.Groups[tria].Name, ":", configs.Groups[tria].Networks)
		}
		rangs := configs.Groups[tria].Networks
		resultats, _, err := listIP.Check(rangs, portNumber, parallel, timeout)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		outputFormat(outFile, configs.Groups[tria].Name, resultats)
	}

	// scanDuration := time.Since(startTime)
	// fmt.Println(scanDuration)
}
