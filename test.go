package main

func main() {
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopack/internal/archiver"
)

func main() {
	inputPath := flag.String("src", "", "")
	outputPath := flag.String("out", "output.tar.gz", "")
	decompress := flag.Bool("x", false, "")

	flag.Parse()

	if *inputPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *decompress {
		err := archiver.Unarchive(*inputPath, *outputPath)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := archiver.Archive(*inputPath, *outputPath)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Done")
}	
}
