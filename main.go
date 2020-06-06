package main

import (
	"fmt"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	wkhtmltopdf.SetPath("/usr/local/bin/wkhtmltopdf")

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		panic(err)
	}

	page := wkhtmltopdf.NewPage("https://css-tricks.com")
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./simplesample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}
