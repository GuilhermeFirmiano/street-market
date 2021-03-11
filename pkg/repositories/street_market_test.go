package repositories_test

import (
	"context"
	"testing"

	"github.com/GuilhermeFirmiano/grok"
	"github.com/GuilhermeFirmiano/street-market/pkg/entities"
	"github.com/GuilhermeFirmiano/street-market/pkg/repositories"
	"github.com/GuilhermeFirmiano/street-market/pkg/settings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StreetMarketRepositoryTestSuite struct {
	suite.Suite
	assert     *assert.Assertions
	settings   *settings.Settings
	client     *mongo.Client
	repository *repositories.StreetMarketRepository
	ctx        context.Context
}

func TestStreetMarketRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(StreetMarketRepositoryTestSuite))
}

func (s *StreetMarketRepositoryTestSuite) SetupSuite() {
	s.ctx = context.Background()
	s.assert = assert.New(s.T())

	s.settings = new(settings.Settings)
	err := grok.FromYAML("../../config.test.yaml", s.settings)
	s.assert.NoError(err)

	s.client = grok.NewMongoConnection(s.settings.Grok.Mongo.ConnectionString, s.settings.Grok.Mongo.CaFilePath)

	s.repository = repositories.NewStreetMarketRepository(s.client, s.settings.Grok.Mongo.Database)
}

func (s *StreetMarketRepositoryTestSuite) TestInsert() {
	result, err := s.repository.Insert(s.ctx, &entities.StreetMarket{
		Registry: primitive.NewObjectID().Hex(),
		Location: entities.Point{
			Type:        "Point",
			Coordinates: []float64{-46.550164, -23.558733},
		},
	})

	s.assert.NoError(err)
	s.assert.NotNil(result)
}

func (s *StreetMarketRepositoryTestSuite) TestInsertDuplicate() {
	entity := &entities.StreetMarket{
		Registry: primitive.NewObjectID().Hex(),
		Location: entities.Point{
			Type:        "Point",
			Coordinates: []float64{-46.550164, -23.558733},
		},
	}

	result, err := s.repository.Insert(s.ctx, entity)

	s.assert.NoError(err)
	s.assert.NotNil(result)

	result2, err := s.repository.Insert(s.ctx, entity)

	s.assert.Error(err)
	s.assert.Nil(result2)
}

func (s *StreetMarketRepositoryTestSuite) TestInsertInvalidCoordinater() {
	entity := &entities.StreetMarket{
		Registry: primitive.NewObjectID().Hex(),
		Location: entities.Point{
			Type:        "Point",
			Coordinates: []float64{500, 500},
		},
	}

	result, err := s.repository.Insert(s.ctx, entity)

	s.assert.Error(err)
	s.assert.Nil(result)
}

func (s *StreetMarketRepositoryTestSuite) TestFindByID() {
	result, err := s.repository.Insert(s.ctx, &entities.StreetMarket{
		Registry: primitive.NewObjectID().Hex(),
		Location: entities.Point{
			Type:        "Point",
			Coordinates: []float64{-46.550164, -23.558733},
		},
	})

	s.assert.NoError(err)
	s.assert.NotNil(result)

	find, err := s.repository.FindByID(s.ctx, result.ID)

	s.assert.NoError(err)
	s.assert.NotNil(find)
}

func (s *StreetMarketRepositoryTestSuite) TestFindByIDNotFound() {
	find, err := s.repository.FindByID(s.ctx, primitive.NewObjectID())

	s.assert.Error(err)
	s.assert.Nil(find)
}

func (s *StreetMarketRepositoryTestSuite) TestFilter() {
	result, err := s.repository.Insert(s.ctx, &entities.StreetMarket{
		Location: entities.Point{
			Type:        "Point",
			Coordinates: []float64{-46.5441173, -23.5370976},
		},
		Name:         primitive.NewObjectID().Hex(),
		Region5:      primitive.NewObjectID().Hex(),
		District:     primitive.NewObjectID().Hex(),
		Registry:     primitive.NewObjectID().Hex(),
		Neighborhood: primitive.NewObjectID().Hex(),
	})

	s.assert.NoError(err)
	s.assert.NotNil(result)

	distance := int64(1000)
	lat := float64(-23.5370976)
	long := float64(-46.5441173)

	r, err := s.repository.Filter(s.ctx, &entities.FilterStreetMarket{
		Page:         1,
		PerPage:      10,
		Lat:          &lat,
		Long:         &long,
		Distance:     &distance,
		Region5:      &result.Region5,
		Name:         &result.Name,
		District:     &result.District,
		Registry:     &result.Registry,
		Neighborhood: &result.Neighborhood,
	})

	s.assert.NoError(err)
	s.assert.NotNil(r)
	s.assert.Equal(r.Total, int64(1))
}

func (s *StreetMarketRepositoryTestSuite) TestFilterError() {
	client := grok.NewMongoConnection(s.settings.Grok.Mongo.ConnectionString, s.settings.Grok.Mongo.CaFilePath)
	repository := repositories.NewStreetMarketRepository(client, s.settings.Grok.Mongo.Database)

	ctx := context.Background()
	client.Disconnect(ctx)

	distance := int64(1000)
	lat := float64(-23.5370976)
	long := float64(-46.5441173)

	r, err := repository.Filter(ctx, &entities.FilterStreetMarket{
		Page:     1,
		PerPage:  10,
		Lat:      &lat,
		Long:     &long,
		Distance: &distance,
	})

	s.assert.Error(err)
	s.assert.Nil(r)
}

func (s *StreetMarketRepositoryTestSuite) TestDeleteByID() {
	result, err := s.repository.Insert(s.ctx, &entities.StreetMarket{
		Registry: primitive.NewObjectID().Hex(),
		Location: entities.Point{
			Type:        "Point",
			Coordinates: []float64{-46.550164, -23.558733},
		},
	})

	s.assert.NoError(err)
	s.assert.NotNil(result)

	find, err := s.repository.FindByID(s.ctx, result.ID)

	s.assert.NoError(err)
	s.assert.NotNil(find)

	err = s.repository.DeleteByID(s.ctx, find.ID)

	s.assert.NoError(err)

	find2, err := s.repository.FindByID(s.ctx, find.ID)

	s.assert.Error(err)
	s.assert.Nil(find2)
}

func (s *StreetMarketRepositoryTestSuite) TestDeleteFilter() {
	result, err := s.repository.Insert(s.ctx, &entities.StreetMarket{
		Registry: primitive.NewObjectID().Hex(),
		Location: entities.Point{
			Type:        "Point",
			Coordinates: []float64{-46.550164, -23.558733},
		},
	})

	s.assert.NoError(err)
	s.assert.NotNil(result)

	find, err := s.repository.FindByID(s.ctx, result.ID)

	s.assert.NoError(err)
	s.assert.NotNil(find)

	err = s.repository.DeleteByFilter(s.ctx, &entities.FilterStreetMarketDelete{
		Registry: result.Registry,
	})

	s.assert.NoError(err)

	find2, err := s.repository.FindByID(s.ctx, find.ID)

	s.assert.Error(err)
	s.assert.Nil(find2)
}

func (s *StreetMarketRepositoryTestSuite) TestUpdate() {
	result, err := s.repository.Insert(s.ctx, &entities.StreetMarket{
		Registry: primitive.NewObjectID().Hex(),
		Name:     "teste",
		Location: entities.Point{
			Type:        "Point",
			Coordinates: []float64{-46.550164, -23.558733},
		},
	})

	s.assert.NoError(err)
	s.assert.NotNil(result)

	err = s.repository.UpdateAllFields(s.ctx, result.ID, &entities.StreetMarket{
		Registry: result.Registry,
		Location: result.Location,
		Name:     "teste1",
	})

	s.assert.NoError(err)

	find, err := s.repository.FindByID(s.ctx, result.ID)

	s.assert.NoError(err)
	s.assert.NotNil(find)
	s.assert.NotEqual(result.Name, find.Name)

}
