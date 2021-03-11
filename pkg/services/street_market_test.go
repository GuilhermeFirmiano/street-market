package services_test

import (
	"context"
	"testing"

	"github.com/GuilhermeFirmiano/grok"
	"github.com/GuilhermeFirmiano/street-market/pkg/models"
	"github.com/GuilhermeFirmiano/street-market/pkg/repositories"
	"github.com/GuilhermeFirmiano/street-market/pkg/services"
	"github.com/GuilhermeFirmiano/street-market/pkg/settings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StreetMarketServiceTestSuite struct {
	suite.Suite
	assert     *assert.Assertions
	settings   *settings.Settings
	service    *services.StreetMarketService
	repository *repositories.StreetMarketRepository
	ctx        context.Context
}

func TestStreetMarketServiceTestSuite(t *testing.T) {
	suite.Run(t, new(StreetMarketServiceTestSuite))
}

func (s *StreetMarketServiceTestSuite) SetupSuite() {
	s.ctx = context.Background()
	s.assert = assert.New(s.T())

	s.settings = new(settings.Settings)
	err := grok.FromYAML("../../config.test.yaml", s.settings)
	s.assert.NoError(err)

	client := grok.NewMongoConnection(s.settings.Grok.Mongo.ConnectionString, s.settings.Grok.Mongo.CaFilePath)

	s.repository = repositories.NewStreetMarketRepository(client, s.settings.Grok.Mongo.Database)

	s.service = services.NewStreetMarketService(s.repository)
}

func (s *StreetMarketServiceTestSuite) TestInsert() {
	r, err := s.service.Create(s.ctx, createStreetMarket())

	s.assert.NoError(err)
	s.assert.NotNil(r)
}

func (s *StreetMarketServiceTestSuite) TestInsertModelError() {
	r, err := s.service.Create(s.ctx, &models.StreetMarketPostRequest{})

	s.assert.Error(err)
	s.assert.Nil(r)
}

func (s *StreetMarketServiceTestSuite) TestInsertDuplicate() {
	model := createStreetMarket()
	r, err := s.service.Create(s.ctx, model)

	s.assert.NoError(err)
	s.assert.NotNil(r)

	r2, err := s.service.Create(s.ctx, model)

	s.assert.Error(err)
	s.assert.Nil(r2)
}

func (s *StreetMarketServiceTestSuite) TestUpdate() {
	r, err := s.service.Create(s.ctx, createStreetMarket())

	s.assert.NoError(err)
	s.assert.NotNil(r)

	up := updateStreetMarket()

	up.Name = "teste1"
	r2, err := s.service.Update(s.ctx, r.ID, up)

	s.assert.NoError(err)
	s.assert.NotNil(r2)
	s.assert.NotEqual(r.Name, r2.Name)
	s.assert.Equal(r.ID, r2.ID)
}

func (s *StreetMarketServiceTestSuite) TestUpdateModelError() {
	r2, err := s.service.Update(s.ctx, "", &models.StreetMarketPutRequest{})

	s.assert.Error(err)
	s.assert.Nil(r2)
}

func (s *StreetMarketServiceTestSuite) TestUpdateFindError() {
	r2, err := s.service.Update(s.ctx, "", updateStreetMarket())

	s.assert.Error(err)
	s.assert.Nil(r2)
}

func (s *StreetMarketServiceTestSuite) TestGetByID() {
	r, err := s.service.Create(s.ctx, createStreetMarket())

	s.assert.NoError(err)
	s.assert.NotNil(r)

	get, err := s.service.GetByID(s.ctx, r.ID)

	s.assert.NoError(err)
	s.assert.NotNil(get)
}

func (s *StreetMarketServiceTestSuite) TestGetByIDNotFound() {
	get, err := s.service.GetByID(s.ctx, "")

	s.assert.Error(err)
	s.assert.Nil(get)
}

func (s *StreetMarketServiceTestSuite) TestDeleteByID() {
	r, err := s.service.Create(s.ctx, createStreetMarket())

	s.assert.NoError(err)
	s.assert.NotNil(r)

	err = s.service.DeleteByID(s.ctx, r.ID)

	s.assert.NoError(err)

	get, err := s.service.GetByID(s.ctx, r.ID)

	s.assert.Error(err)
	s.assert.Nil(get)
}

func (s *StreetMarketServiceTestSuite) TestDeleteByFilter() {
	r, err := s.service.Create(s.ctx, createStreetMarket())

	s.assert.NoError(err)
	s.assert.NotNil(r)

	err = s.service.DeleteByFilter(s.ctx, &models.FilterStreetMarketDelete{
		Registry: r.Registry,
	})

	s.assert.NoError(err)

	get, err := s.service.GetByID(s.ctx, r.ID)

	s.assert.Error(err)
	s.assert.Nil(get)
}

func (s *StreetMarketServiceTestSuite) TestDeleteByFilterError() {
	r, err := s.service.Create(s.ctx, createStreetMarket())

	s.assert.NoError(err)
	s.assert.NotNil(r)

	err = s.service.DeleteByFilter(s.ctx, &models.FilterStreetMarketDelete{})

	s.assert.Error(err)

	get, err := s.service.GetByID(s.ctx, r.ID)

	s.assert.NoError(err)
	s.assert.NotNil(get)
}

func (s *StreetMarketServiceTestSuite) TestFilter() {
	model := &models.StreetMarketPostRequest{
		Registry: primitive.NewObjectID().Hex(),
		StreetMarket: models.StreetMarket{
			Long:           -46.550164,
			Lat:            -23.558733,
			Name:           primitive.NewObjectID().Hex(),
			CensusSectors:  primitive.NewObjectID().Hex(),
			WeighingArea:   primitive.NewObjectID().Hex(),
			DistrictID:     primitive.NewObjectID().Hex(),
			District:       primitive.NewObjectID().Hex(),
			SubCityHallID:  primitive.NewObjectID().Hex(),
			SubCityHall:    primitive.NewObjectID().Hex(),
			Region5:        primitive.NewObjectID().Hex(),
			Region8:        primitive.NewObjectID().Hex(),
			AddressLine:    primitive.NewObjectID().Hex(),
			BuildingNumber: primitive.NewObjectID().Hex(),
			Neighborhood:   primitive.NewObjectID().Hex(),
			Reference:      primitive.NewObjectID().Hex(),
		},
	}

	r, err := s.service.Create(s.ctx, model)

	s.assert.NoError(err)
	s.assert.NotNil(r)

	distance := int64(1000)

	filter, err := s.service.Filter(s.ctx, &models.FilterStreetMarket{
		Page:         1,
		PerPage:      10,
		Lat:          &model.Lat,
		Long:         &model.Long,
		Distance:     &distance,
		District:     &model.District,
		Region5:      &model.Region5,
		Name:         &model.Name,
		Neighborhood: &model.Neighborhood,
		Registry:     &model.Registry,
	})

	s.assert.NoError(err)
	s.assert.NotNil(filter)
	s.assert.Equal(int64(1), filter.Total)
}

func (s *StreetMarketServiceTestSuite) TestFilterError() {

	r, err := s.service.Create(s.ctx, createStreetMarket())

	s.assert.NoError(err)
	s.assert.NotNil(r)

	filter, err := s.service.Filter(s.ctx, &models.FilterStreetMarket{})

	s.assert.Error(err)
	s.assert.Nil(filter)
}

func createStreetMarket() *models.StreetMarketPostRequest {
	return &models.StreetMarketPostRequest{
		Registry: primitive.NewObjectID().Hex(),
		StreetMarket: models.StreetMarket{
			Long:           -46.550164,
			Lat:            -23.558733,
			Name:           "teste",
			CensusSectors:  "teste",
			WeighingArea:   "teste",
			DistrictID:     "teste",
			District:       "teste",
			SubCityHallID:  "teste",
			SubCityHall:    "teste",
			Region5:        "teste",
			Region8:        "teste",
			AddressLine:    "teste",
			BuildingNumber: "teste",
			Neighborhood:   "teste",
			Reference:      "teste",
		},
	}
}

func updateStreetMarket() *models.StreetMarketPutRequest {
	return &models.StreetMarketPutRequest{
		StreetMarket: models.StreetMarket{
			Long:           -46.550164,
			Lat:            -23.558733,
			Name:           "teste",
			CensusSectors:  "teste",
			WeighingArea:   "teste",
			DistrictID:     "teste",
			District:       "teste",
			SubCityHallID:  "teste",
			SubCityHall:    "teste",
			Region5:        "teste",
			Region8:        "teste",
			AddressLine:    "teste",
			BuildingNumber: "teste",
			Neighborhood:   "teste",
			Reference:      "teste",
		},
	}
}
