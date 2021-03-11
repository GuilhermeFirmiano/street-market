package repositories

import (
	"context"
	"time"

	"github.com/GuilhermeFirmiano/street-market/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

//StreetMarketRepository ...
type StreetMarketRepository struct {
	collection *mongo.Collection
}

const (
	//EarthRadiusByM ...
	EarthRadiusByM = 6378100
)

//NewStreetMarketRepository ...
func NewStreetMarketRepository(client *mongo.Client, database string) *StreetMarketRepository {
	repository := &StreetMarketRepository{client.Database(database).Collection("street-market")}
	repository.createIndex()

	return repository
}

//Insert ...
func (repository *StreetMarketRepository) Insert(ctx context.Context, r *entities.StreetMarket) (*entities.StreetMarket, error) {
	result, err := repository.collection.InsertOne(ctx, r)

	if err != nil {
		return nil, err
	}

	r.ID = result.InsertedID.(primitive.ObjectID)

	return r, nil
}

//UpdateAllFields ...
func (repository *StreetMarketRepository) UpdateAllFields(ctx context.Context, ID primitive.ObjectID, entity *entities.StreetMarket) error {
	entity.UpdatedAt = time.Now()

	return repository.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": ID},
		bson.M{"$set": entity},
	).Err()
}

//FindByID ...
func (repository *StreetMarketRepository) FindByID(ctx context.Context, ID primitive.ObjectID) (*entities.StreetMarket, error) {
	var r *entities.StreetMarket

	filter := bson.M{"_id": ID}

	err := repository.collection.FindOne(ctx, filter).Decode(&r)

	return r, err
}

// Filter ...
func (repository *StreetMarketRepository) Filter(ctx context.Context, entity *entities.FilterStreetMarket) (*entities.FilterStreetMarketResult, error) {
	result := new(entities.FilterStreetMarketResult)
	result.StreetMarket = []*entities.StreetMarket{}

	filter := bson.M{}

	if entity.Distance != nil {
		filter["location"] = bson.M{
			"$geoWithin": bson.M{
				"$centerSphere": []interface{}{[]float64{*entity.Long, *entity.Lat}, calculateRadius(*entity.Distance)},
			},
		}
	}

	if entity.Name != nil {
		filter["name"] = entity.Name
	}

	if entity.Neighborhood != nil {
		filter["neighborhood"] = entity.Neighborhood
	}

	if entity.Region5 != nil {
		filter["region_5"] = entity.Region5
	}

	if entity.District != nil {
		filter["district"] = entity.District
	}

	if entity.Registry != nil {
		filter["registry"] = entity.Registry
	}

	result.Total, _ = repository.collection.CountDocuments(ctx, filter)
	result.Pages = result.Total / entity.PerPage

	cursor, err := repository.collection.Find(ctx, filter, options.Find().SetLimit(entity.PerPage).SetSkip((entity.Page-1)*entity.PerPage))

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &result.StreetMarket)

	return result, err
}

//DeleteByFilter ...
func (repository *StreetMarketRepository) DeleteByFilter(ctx context.Context, filterDelete *entities.FilterStreetMarketDelete) error {
	filter := bson.M{
		"registry": filterDelete.Registry,
	}

	_, err := repository.collection.DeleteOne(ctx, filter)

	return err
}

//DeleteByID ...
func (repository *StreetMarketRepository) DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{
		"_id": id,
	}

	_, err := repository.collection.DeleteOne(ctx, filter)

	return err
}

func (repository *StreetMarketRepository) createIndex() {
	ctx, close := context.WithTimeout(context.Background(), 10*time.Second)
	defer close()

	unique := mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "registry", Value: bsonx.Int32(-1)}},
		Options: options.Index().SetName("unique").SetUnique(true),
	}

	repository.collection.Indexes().CreateOne(ctx, unique)

	location := mongo.IndexModel{
		Keys:    bsonx.MDoc{"location": bsonx.String("2dsphere")},
		Options: options.Index().SetBackground(true).SetName("location"),
	}

	repository.collection.Indexes().CreateOne(ctx, location)

}

func calculateRadius(r int64) float64 {
	return float64(r) / EarthRadiusByM
}
