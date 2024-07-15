// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package pgstore

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Activity struct {
	ID        uuid.UUID        `json:"id"`
	TripID    uuid.UUID        `json:"trip_id"`
	Title     string           `json:"title"`
	OccoursAt pgtype.Timestamp `json:"occours_at"`
}

type Link struct {
	ID     uuid.UUID `json:"id"`
	TripID uuid.UUID `json:"trip_id"`
	Title  string    `json:"title"`
	Url    string    `json:"url"`
}

type Participant struct {
	ID          uuid.UUID `json:"id"`
	TripID      uuid.UUID `json:"trip_id"`
	Email       string    `json:"email"`
	IsConfirmed bool      `json:"is_confirmed"`
}

type Trip struct {
	ID          uuid.UUID        `json:"id"`
	Destination string           `json:"destination"`
	OwnerEmail  string           `json:"owner_email"`
	OwnerName   string           `json:"owner_name"`
	IsConfirmed bool             `json:"is_confirmed"`
	StartsAt    pgtype.Timestamp `json:"starts_at"`
	EndsAt      pgtype.Timestamp `json:"ends_at"`
}
