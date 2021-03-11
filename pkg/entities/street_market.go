package entities

import (
	"time"

	"github.com/GuilhermeFirmiano/grok"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//StreetMarket ...
type StreetMarket struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	UpdatedAt      time.Time          `bson:"update_at,omitempty"`
	CreatedAt      time.Time          `bson:"created_at,omitempty"`
	CensusSectors  string             `bson:"census_sectors,omitempty"`
	WeighingArea   string             `bson:"weighing_area,omitempty"`
	DistrictID     string             `bson:"district_id,omitempty"`
	District       string             `bson:"district,omitempty"`
	SubCityHallID  string             `bson:"sub_city_hall_id,omitempty"`
	SubCityHall    string             `bson:"sub_city_hall,omitempty"`
	Region5        string             `bson:"region_5,omitempty"`
	Region8        string             `bson:"region_8,omitempty"`
	Name           string             `bson:"name,omitempty"`
	Registry       string             `bson:"registry,omitempty"`
	AddressLine    string             `bson:"address_line,omitempty"`
	BuildingNumber string             `bson:"building_number,omitempty"`
	Neighborhood   string             `bson:"neighborhood,omitempty"`
	Reference      string             `bson:"reference,omitempty"`
	Location       Point              `bson:"location,omitempty"`
}

//FilterStreetMarketResult ...
type FilterStreetMarketResult struct {
	grok.PaginationResult
	StreetMarket []*StreetMarket `bson:"street-market"`
}

//FilterStreetMarketDelete ...
type FilterStreetMarketDelete struct {
	Registry string
}

//FilterStreetMarket ...
type FilterStreetMarket struct {
	Page         int64
	PerPage      int64
	Lat          *float64
	Long         *float64
	Distance     *int64
	District     *string
	Region5      *string
	Name         *string
	Neighborhood *string
	Registry     *string
}

// Point ...
type Point struct {
	Type        string    `bson:"type"`
	Coordinates []float64 `bson:"coordinates"`
}
