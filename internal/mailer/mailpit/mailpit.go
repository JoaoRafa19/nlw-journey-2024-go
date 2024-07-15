package mailpit

import (
	"context"
	"fmt"
	"time"

	"github.com/JoaoRafa19/nlw-journey-2024-go/internal/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wneessen/go-mail"
)

type store interface {
	GetTrip(context.Context, uuid.UUID) (pgstore.Trip, error)
}

type Mailpit struct {
	store store
	mailpithost string
}

func NewMailpit(pool *pgxpool.Pool, mailpithost string) Mailpit {
	return Mailpit{
		store: pgstore.New(pool),
		mailpithost: mailpithost,
	}
}

func (mp Mailpit) SendConfirmTripEmailToTripOwner(tripId uuid.UUID) error {
	ctx := context.Background()
	trip, err := mp.store.GetTrip(ctx, tripId)
	if err != nil {
		return fmt.Errorf("mailpit: failed to get trip for SendConfirmTripEmailToTripOwner: %w", err)
	}

	msg := mail.NewMsg()
	if err := msg.From("mailpit@journey.com"); err != nil {
		return fmt.Errorf("mailpit: failed to set  From in email for SendConfirmTripEmailToTripOwner: %w", err)
	}
	if err := msg.To(trip.OwnerEmail); err != nil {
		return fmt.Errorf("mailpit: failed to set To in email to SendConfirmTripEmailToTripOwner: %w", err)
	}

	msg.Subject("Confirme sua viagem")
	msg.SetBodyString(mail.TypeTextPlain, fmt.Sprintf(
		`
		Olá! %s!
		Sua viagem par %s que começa no dia %s precisa ser confirmada.
		Clique no botão a baixo  para confirmar.
		`,
		trip.OwnerName, trip.Destination, trip.StartsAt.Time.Format(time.DateOnly),
	))

	client, err := mail.NewClient(mp.mailpithost, mail.WithTLSPolicy(mail.NoTLS), mail.WithPort(1025))
	if err != nil {
		return fmt.Errorf("mailpit: failed to create email client to SendConfirmTripEmailToTripOwner: %w", err)
	}

	if err := client.DialAndSend(msg); err != nil {
		return fmt.Errorf("mailpit: failed to send email to SendConfirmTripEmailToTripOwner: %w", err)
	}

	return nil

}
