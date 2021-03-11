package models

import (
	"strconv"
	"time"

	"github.com/GuilhermeFirmiano/grok"
	"github.com/GuilhermeFirmiano/street-market/pkg/entities"
)

//StreetMarketPostRequest ...
type StreetMarketPostRequest struct {
	Registry string `validate:"required" json:"registry"`
	StreetMarket
}

//StreetMarketPutRequest ...
type StreetMarketPutRequest struct {
	StreetMarket
}

//StreetMarketResponse ...
type StreetMarketResponse struct {
	ID        string    `json:"id,omitempty"`
	UpdatedAt time.Time `json:"update_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Registry  string    `validate:"required" json:"registry"`
	StreetMarket
}

//StreetMarket ...
type StreetMarket struct {
	Long           float64 `validate:"required,longitude" json:"long"`
	Lat            float64 `validate:"required,latitude" json:"lat"`
	CensusSectors  string  `validate:"required" json:"census_sectors"`
	WeighingArea   string  `validate:"required" json:"weighing_area"`
	DistrictID     string  `validate:"required" json:"district_id"`
	District       string  `validate:"required" json:"district"`
	SubCityHallID  string  `validate:"required" json:"sub_city_hall_id"`
	SubCityHall    string  `validate:"required" json:"sub_city_hall"`
	Region5        string  `validate:"required" json:"region_5"`
	Region8        string  `validate:"required" json:"region_8"`
	Name           string  `validate:"required" json:"name"`
	AddressLine    string  `validate:"required" json:"address_line"`
	BuildingNumber string  `json:"building_number"`
	Neighborhood   string  `validate:"required" json:"neighborhood"`
	Reference      string  `json:"reference"`
}

//FilterStreetMarket ...
type FilterStreetMarket struct {
	Page         int64    `validate:"required" form:"page" json:"page"`
	PerPage      int64    `validate:"required" form:"per_page" json:"per_page"`
	Lat          *float64 `validate:"required_with=Long Distance" form:"lat" json:"lat"`
	Long         *float64 `validate:"required_with=Lat Distance" form:"long" json:"long"`
	Distance     *int64   `validate:"required_with=Lat Long" form:"distance" json:"distance"`
	District     *string  `form:"district" json:"district"`
	Region5      *string  `form:"region_5" json:"region_5"`
	Name         *string  `form:"name" json:"name"`
	Neighborhood *string  `form:"neighborhood" json:"neighborhood"`
	Registry     *string  `form:"registry" json:"registry"`
}

//FilterStreetMarketResponse ...
type FilterStreetMarketResponse struct {
	grok.PaginationResult
	StreetMarket []*StreetMarketResponse `json:"street-markets"`
}

//FilterStreetMarketDelete ...
type FilterStreetMarketDelete struct {
	Registry string `validate:"required" form:"registry"`
}

//ToEntity ...
func (f *FilterStreetMarketDelete) ToEntity() *entities.FilterStreetMarketDelete {
	return &entities.FilterStreetMarketDelete{
		Registry: f.Registry,
	}
}

//ToEntity ...
func (s *StreetMarketPutRequest) ToEntity() *entities.StreetMarket {
	return &entities.StreetMarket{
		CensusSectors:  s.CensusSectors,
		WeighingArea:   s.WeighingArea,
		DistrictID:     s.DistrictID,
		District:       s.District,
		SubCityHallID:  s.SubCityHallID,
		SubCityHall:    s.SubCityHall,
		Region5:        s.Region5,
		Region8:        s.Region8,
		Name:           s.Name,
		AddressLine:    s.AddressLine,
		Neighborhood:   s.Neighborhood,
		BuildingNumber: s.BuildingNumber,
		Reference:      s.Reference,
		Location: entities.Point{
			Type:        "Point",
			Coordinates: []float64{s.Long, s.Lat},
		},
	}
}

//ToEntity ...
func (f *FilterStreetMarket) ToEntity() *entities.FilterStreetMarket {
	return &entities.FilterStreetMarket{
		Page:         f.Page,
		PerPage:      f.PerPage,
		Lat:          f.Lat,
		Long:         f.Long,
		Distance:     f.Distance,
		District:     f.District,
		Region5:      f.Region5,
		Name:         f.Name,
		Neighborhood: f.Neighborhood,
		Registry:     f.Registry,
	}
}

//ParseCSVToModel ...
func ParseCSVToModel(data []string) (*StreetMarketPostRequest, error) {
	long, err := strconv.ParseFloat(data[1], 64)

	if err != nil {
		return nil, err
	}

	long = long / 1000000

	lat, err := strconv.ParseFloat(data[2], 64)

	if err != nil {
		return nil, err
	}

	lat = lat / 1000000

	return &StreetMarketPostRequest{
		StreetMarket: StreetMarket{
			Long:           long,
			Lat:            lat,
			CensusSectors:  data[3],
			WeighingArea:   data[4],
			DistrictID:     data[5],
			District:       data[6],
			SubCityHallID:  data[7],
			SubCityHall:    data[8],
			Region5:        data[9],
			Region8:        data[10],
			Name:           data[11],
			AddressLine:    data[13],
			BuildingNumber: data[14],
			Neighborhood:   data[15],
			Reference:      data[16],
		},
		Registry: data[12],
	}, nil
}

//ToEntity ...
func (s *StreetMarketPostRequest) ToEntity() *entities.StreetMarket {
	return &entities.StreetMarket{
		CensusSectors:  s.CensusSectors,
		WeighingArea:   s.WeighingArea,
		DistrictID:     s.DistrictID,
		District:       s.District,
		SubCityHallID:  s.SubCityHallID,
		SubCityHall:    s.SubCityHall,
		Region5:        s.Region5,
		Region8:        s.Region8,
		Name:           s.Name,
		Registry:       s.Registry,
		AddressLine:    s.AddressLine,
		Neighborhood:   s.Neighborhood,
		BuildingNumber: s.BuildingNumber,
		Reference:      s.Reference,
		Location: entities.Point{
			Type:        "Point",
			Coordinates: []float64{s.Long, s.Lat},
		},
	}
}

//ParseEntityToModel ...
func ParseEntityToModel(entity *entities.StreetMarket) *StreetMarketResponse {
	return &StreetMarketResponse{
		ID:        entity.ID.Hex(),
		UpdatedAt: entity.UpdatedAt,
		CreatedAt: entity.CreatedAt,
		Registry:  entity.Registry,
		StreetMarket: StreetMarket{
			Long:           entity.Location.Coordinates[0],
			Lat:            entity.Location.Coordinates[1],
			CensusSectors:  entity.CensusSectors,
			WeighingArea:   entity.WeighingArea,
			DistrictID:     entity.DistrictID,
			District:       entity.District,
			SubCityHallID:  entity.SubCityHallID,
			SubCityHall:    entity.SubCityHall,
			Region5:        entity.Region5,
			Region8:        entity.Region8,
			Name:           entity.Name,
			AddressLine:    entity.AddressLine,
			Neighborhood:   entity.Neighborhood,
			BuildingNumber: entity.BuildingNumber,
			Reference:      entity.Reference,
		},
	}
}

// ParseSliceEntityToSliceModel ...
func ParseSliceEntityToSliceModel(s []*entities.StreetMarket) []*StreetMarketResponse {
	slice := make([]*StreetMarketResponse, len(s))

	for i, m := range s {
		slice[i] = ParseEntityToModel(m)
	}

	return slice
}

//ParseFilterStreetMarketResponse ...
func ParseFilterStreetMarketResponse(c *entities.FilterStreetMarketResult) *FilterStreetMarketResponse {
	return &FilterStreetMarketResponse{
		PaginationResult: c.PaginationResult,
		StreetMarket:     ParseSliceEntityToSliceModel(c.StreetMarket),
	}
}
