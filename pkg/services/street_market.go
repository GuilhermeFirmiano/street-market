package services

import (
	"context"
	"time"

	"github.com/GuilhermeFirmiano/grok"
	"github.com/GuilhermeFirmiano/street-market/pkg/models"
	"github.com/GuilhermeFirmiano/street-market/pkg/repositories"
	"github.com/sirupsen/logrus"
)

//StreetMarketService ...
type StreetMarketService struct {
	repository *repositories.StreetMarketRepository
}

//NewStreetMarketService ...
func NewStreetMarketService(
	repository *repositories.StreetMarketRepository) *StreetMarketService {
	return &StreetMarketService{
		repository: repository,
	}
}

//Create ...
func (service *StreetMarketService) Create(ctx context.Context, model *models.StreetMarketPostRequest) (*models.StreetMarketResponse, error) {
	err := grok.Validator.Struct(model)

	if err != nil {
		logrus.WithError(err).
			Errorf("error model street market")
		return nil, grok.FromValidationErros(err)
	}

	entity := model.ToEntity()
	entity.CreatedAt = time.Now()
	entity.UpdatedAt = time.Now()

	result, err := service.repository.Insert(ctx, entity)

	if err != nil {
		return nil, err
	}

	return models.ParseEntityToModel(result), nil
}

//Update ...
func (service *StreetMarketService) Update(ctx context.Context, ID string, model *models.StreetMarketPutRequest) (*models.StreetMarketResponse, error) {
	err := grok.Validator.Struct(model)

	if err != nil {
		logrus.WithError(err).
			Errorf("error model contact")
		return nil, grok.FromValidationErros(err)
	}

	restul, err := service.repository.FindByID(ctx, grok.ObjectIDFromHex(ID))

	if err != nil {
		return nil, err
	}

	entity := model.ToEntity()
	entity.ID = restul.ID
	entity.Registry = restul.Registry

	err = service.repository.UpdateAllFields(ctx, grok.ObjectIDFromHex(ID), entity)

	if err != nil {
		return nil, err
	}

	return models.ParseEntityToModel(entity), nil
}

//GetByID ...
func (service *StreetMarketService) GetByID(ctx context.Context, ID string) (*models.StreetMarketResponse, error) {
	result, err := service.repository.FindByID(ctx, grok.ObjectIDFromHex(ID))

	if err != nil {
		return nil, err
	}

	return models.ParseEntityToModel(result), nil
}

// Filter ...
func (service *StreetMarketService) Filter(ctx context.Context, model *models.FilterStreetMarket) (*models.FilterStreetMarketResponse, error) {
	err := grok.Validator.Struct(model)

	if err != nil {
		logrus.WithError(err).
			Errorf("error model street market")
		return nil, grok.FromValidationErros(err)
	}

	entity := model.ToEntity()

	result, err := service.repository.Filter(ctx, entity)

	if err != nil {
		return nil, err
	}

	return models.ParseFilterStreetMarketResponse(result), nil
}

//DeleteByFilter ...
func (service *StreetMarketService) DeleteByFilter(ctx context.Context, filter *models.FilterStreetMarketDelete) error {
	err := grok.Validator.Struct(filter)

	if err != nil {
		logrus.WithError(err).
			Errorf("error model filter delete")
		return grok.FromValidationErros(err)
	}

	filterDelete := filter.ToEntity()

	return service.repository.DeleteByFilter(ctx, filterDelete)
}

//DeleteByID ...
func (service *StreetMarketService) DeleteByID(ctx context.Context, id string) error {
	return service.repository.DeleteByID(ctx, grok.ObjectIDFromHex(id))
}
