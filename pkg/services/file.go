package services

import (
	"encoding/csv"
	"os"
	"path/filepath"

	"github.com/GuilhermeFirmiano/street-market/pkg/errors"
)

//FileService ...
type FileService struct{}

//NewFileService ...
func NewFileService() *FileService {
	return &FileService{}
}

//ReadFileCSV ...
func (f *FileService) ReadFileCSV(filePath string) ([][]string, error) {
	fileExt := filepath.Ext(filePath)

	if fileExt != ".csv" {
		return nil, errors.ErrFileExt
	}

	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	return csv.NewReader(csvFile).ReadAll()
}
