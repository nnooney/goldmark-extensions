package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"

	extensions "github.com/nnooney/goldmark-extensions"
	"github.com/nnooney/goldmark-extensions/renderer/ast"
	"github.com/yuin/goldmark"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Declare Command Line Flags
	inputPtr := flag.String("input", "", "input markdown file")
	outputPtr := flag.String("output", "", "output file")
	formatPtr := flag.String("format", "html", "output format (html or ast)")
	flag.Parse()

	// Read the input (markdown) file
	input, err := ioutil.ReadFile(*inputPtr)
	check(err)

	// Choose the renderer based upon the format flag
	renderer := goldmark.DefaultRenderer()
	switch *formatPtr {
	case "ast":
		renderer = ast.NewRenderer()
	case "html":
		renderer = goldmark.DefaultRenderer()
	default:
		panic("Invalid --format")
	}

	// Convert the markdown into an output format
	var buf bytes.Buffer
	md := goldmark.New(
		goldmark.WithRenderer(
			renderer,
		),
		goldmark.WithExtensions(
			extensions.Tufte,
		),
	)
	err = md.Convert(input, &buf)
	check(err)

	// Write the output (HTML) to the file or stdout
	if *outputPtr != "" {
		output, err := os.Create(*outputPtr)
		check(err)
		defer output.Close()

		_, err = output.Write(buf.Bytes())
		check(err)
		output.Sync()
	} else {
		os.Stdout.Write(buf.Bytes())
		os.Stdout.Sync()
	}
}
