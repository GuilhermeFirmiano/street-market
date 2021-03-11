package errors

import (
	"net/http"

	"github.com/GuilhermeFirmiano/grok"
)

var (
	// ErrEntryNotFound ...
	ErrEntryNotFound = grok.NewError(http.StatusNotFound, "not found")
	// ErrDuplicateStreetMarket ...
	ErrDuplicateStreetMarket = grok.NewError(http.StatusConflict, "duplicate street market")
	//ErrFileExt ...
	ErrFileExt = grok.NewError(http.StatusConflict, "File extension ins't equal to .csv")
)
