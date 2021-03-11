package controllers

import (
	"github.com/GuilhermeFirmiano/grok"
	"github.com/GuilhermeFirmiano/street-market/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	grok.DefaultErrorMapping.Register(mongo.ErrNoDocuments, errors.ErrEntryNotFound)
	grok.DefaultErrorMapping.Register(mongo.WriteError{Code: 11000}, errors.ErrDuplicateStreetMarket)
}
