package repository

import (
	"context"
	"log"
	"yayawallet-webhook/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type WebhookRepository interface {
	Save(payload models.WebhookPayload) error
}

type webhookRepo struct {
	u   *mongo.Collection
	ctx context.Context
}

func NewWebhookRepository(u *mongo.Collection, ctx context.Context) WebhookRepository {
	return &webhookRepo{
		u:   u,
		ctx: ctx,
	}
}

func (r *webhookRepo) Save(payload models.WebhookPayload) error {
	_, err := r.u.InsertOne(r.ctx, payload)
	if err != nil {
		return err
	}
	log.Printf("Received webhook payload: %+v\n", payload)
	return nil
}
