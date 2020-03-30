/*
 * PDF to text: Extract all text for each page of a pdf file.
 *
 * Run as: go run pdf_extract_text.go input.pdf
 */

package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unipdf/v3/extractor"
	pdf "github.com/unidoc/unipdf/v3/model"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: go run pdf_extract_text.go input.pdf\n")
		os.Exit(1)
	}

	// For debugging.
	// common.SetLogger(common.NewConsoleLogger(common.LogLevelDebug))

	inputPath := os.Args[1]

	err := outputPdfText(inputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

// outputPdfText prints out contents of PDF file to stdout.
func outputPdfText(inputPath string) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	fmt.Printf("--------------------\n")
	fmt.Printf("PDF to text extraction:\n")
	fmt.Printf("--------------------\n")
	pageNum := numPages

	page, err := pdfReader.GetPage(pageNum)
	if err != nil {
		return err
	}

	ex, err := extractor.New(page)
	if err != nil {
		return err
	}

	text, err := ex.ExtractText()
	if err != nil {
		return err
	}

	fmt.Println("------------------------------")
	fmt.Printf("Page %d:\n", pageNum)
	fmt.Printf("\"%s\"\n", text)
	fmt.Println("------------------------------")

	page, err = pdfReader.GetPage(pageNum-1)
	if err != nil {
		return err
	}

	ex, err = extractor.New(page)
	if err != nil {
		return err
	}

	text, err = ex.ExtractText()
	if err != nil {
		return err
	}

	fmt.Println("------------------------------")
	fmt.Printf("Page %d:\n", pageNum)
	fmt.Printf("\"%s\"\n", text)
	fmt.Println("------------------------------")

	return nil
}
