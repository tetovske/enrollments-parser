package Parsers

import (
	"bytes"
	"github.com/ledongthuc/pdf"
	logger "github.com/sirupsen/logrus"
)

type Parser interface {
	Parse() interface{}
}

type PDFParser struct {
	fileToParse string
}

func NewPDFParser(fileToParse string) *PDFParser {
	return &PDFParser{fileToParse: fileToParse}
}

func (p *PDFParser) Parse() (error) {
	pdfData, err := readFile(p.fileToParse)
	if err != nil {
		return err
	}

	logger.Info(pdfData)
	return err
}

func readFile(path string) (string, error) {
	var paths string = "/home/tetovske/go/src/github.com/tetovske/enrollments-parser/assets/docs/2017/paid/20170808_p.pdf"
	//f, r, err := pdf.Open("/home/tetovske/go/src/github.com/tetovske/enrollments-parser/assets/docs/2016/budget/03082016.pdf")
	//defer f.Close()
	//if err != nil {
	//	return "", err
	//}
	//
	//b, err := r.GetPlainText()
	//
	//var buffer bytes.Buffer
	//_, err = buffer.ReadFrom(b)
	//if err != nil {
	//	return "", err
	//}
	f, r, err := pdf.Open(paths)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	str := buf.String()
	return str, nil
}
