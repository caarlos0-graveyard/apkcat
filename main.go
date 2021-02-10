package main

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"text/tabwriter"
)

func main() {
	f, err := os.Open(os.Args[1])

	gz, err := gzip.NewReader(f)
	if err != nil {
		log.Fatalln(err)
	}
	defer gz.Close()
	tr := tar.NewReader(gz)
	var w = new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	for {
		hdr, err := tr.Next()
		if errors.Is(err, io.EOF) {
			break // End of archive
		}
		if err != nil {
			log.Fatalln(err)
		}

		for k, v := range hdr.PAXRecords {
			fmt.Fprintf(w, "%s\t%s\t%s\n", hdr.Name, k, v)
		}
	}
}
