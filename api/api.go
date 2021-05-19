package api

import (
	"croacker.com/cemetery-park/data"
)

// API -
type API struct {
	// nominations      *data.NominationManager
	// nominationResult *data.NominationResultManager
	// participants     *data.ParticipantManager
}

func NewAPI(db *data.DB) *API {
	return &API{}
}
