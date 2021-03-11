package services_test

import (
	"testing"

	"github.com/GuilhermeFirmiano/grok"
	"github.com/GuilhermeFirmiano/street-market/pkg/errors"
	"github.com/GuilhermeFirmiano/street-market/pkg/services"
	"github.com/GuilhermeFirmiano/street-market/pkg/settings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FileServiceTestSuite struct {
	suite.Suite
	assert   *assert.Assertions
	settings *settings.Settings
	service  *services.FileService
}

func TestFileServiceTestSuite(t *testing.T) {
	suite.Run(t, new(FileServiceTestSuite))
}

func (s *FileServiceTestSuite) SetupSuite() {
	s.assert = assert.New(s.T())

	s.settings = new(settings.Settings)
	err := grok.FromYAML("../../config.test.yaml", s.settings)
	s.assert.NoError(err)

	s.service = services.NewFileService()
}

func (s *FileServiceTestSuite) TestReadFileCSVErrorFileExtension() {
	r, err := s.service.ReadFileCSV("../../test-files/test.txt")

	s.assert.EqualError(err, errors.ErrFileExt.Error())
	s.assert.Nil(r)
}

func (s *FileServiceTestSuite) TestReadFileCSVErrorOpen() {
	r, err := s.service.ReadFileCSV("test.csv")

	s.assert.Error(err)
	s.assert.Nil(r)
}

func (s *FileServiceTestSuite) TestReadFileCSV() {
	r, err := s.service.ReadFileCSV("../../test-files/test.csv")

	s.assert.NoError(err)
	s.assert.NotNil(r)
	s.assert.Equal(r[1][0], "1")
}
