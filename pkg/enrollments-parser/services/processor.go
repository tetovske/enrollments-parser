package services

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tetovske/enrollments-parser/pkg/enrollments-parser/models"
	"github.com/tetovske/enrollments-parser/pkg/enrollments-parser/services/Parsers"
)

type regexModel struct {
	Regex string 			`mapstructure:"regex"`
	Model []string 			`mapstructure:"model"`
}

type docConfig struct {
	Year uint  				`mapstructure:"year"`
	FormOfStudy string 		`mapstructure:"form_of_study"`
	FileName string 		`mapstructure:"file_name"`
	Regexes []regexModel 	`mapstructure:"regexes"`
}

type docConfigurator struct {
	DocsData []docConfig 	`mapstructure:"docs_data"`
	DocsPath string 		`mapstructure:"docs_path"`
}

type ParsedDocument struct {
	Year uint
	FormOfStudy string
	Regex string
	Students []models.Student
}

type Processor struct {
	config docConfigurator
	db interface{} // TODO: Add MongoDB support
}

func NewProcessor() (*Processor, error) {
	conf := docConfigurator{}
	err := viper.Unmarshal(&conf)
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	return &Processor{
		config: conf,
		db: nil,
	}, err
}

func (p *Processor) ProcessDocuments() error {
	for _, doc := range p.config.DocsData {
		documentPath := fmt.Sprintf(
			"%s/%d/%s/%s",
			p.config.DocsPath,
			doc.Year,
			doc.FormOfStudy,
			doc.FileName,
			)
		var pdfParser = Parsers.NewPDFParser(documentPath)
		var err = pdfParser.Parse()
		if err != nil {
			return err
		}
	}

	return nil
}

//
//func NewDocument(params map[string]interface{}) *Document {
//	return Document{
//		Year:        0,
//		FormOfStudy: "",
//		Regex:       "",
//		Students:    nil,
//	}
//}

//func ParseStudents(documents []Document) (Document, error) {
//
//}
