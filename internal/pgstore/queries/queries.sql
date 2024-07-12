-- name: InsertTip :one
INSERT INTO trips 
("destination", "owner_email", "owner_name", "starts_at", "ends_at") VALUES
( $1, $2, $3, $4, $5)
RETURNING "id";

-- name GetTrip :one
SELECT "id", "destination", "owner_email", "owner_name", "is_confirmed", "starts_at", "ends_at"
FROM trips
WHERE
    id = $1;

-- name: UpdateTrip :exec
UPDATE trips 
SET 
    "destination" = $1 ,
    "ends_at" = $2,
    "starts_at" = $3,
    "is_confirmed" = $4
WHERE 
    id = $5;

-- name: GetParticipant :one
SELECT 
    "id", "trip_id", "email", "is_confirmed"
FROM participants
WHERE 
    id = $1;

-- name: ConfirmedParticipants :exec
UPDATE participants
SET "is_confirmed"=true
WHERE id = $1;

-- name: GetParticipants :many
SELECT 
    "id", "trip_id", "email", "is_confirmed"
FROM participants
WHERE 
    trip_id = $1;
