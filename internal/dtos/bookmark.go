package dtos

import (

)

type BookmarkResponse struct {
	ID          string            `json:"id"`
	RekrutmenID string            `json:"rekrutmen_id"`
	Rekrutmen   RekrutmenResponse `json:"rekrutmen"`
}