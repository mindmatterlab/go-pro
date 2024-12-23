package secret

import (
	"context"

	"github.com/mindmatterlab/acex/internal/pkg/acexx"
	v1 "github.com/mindmatterlab/acex/pkg/api/usercenter/v1"
)

// Update updates a secret.
func (b *secretBiz) Update(ctx context.Context, rq *v1.UpdateSecretRequest) error {
	secret, err := b.ds.Secrets().Get(ctx, acexx.FromUserID(ctx), rq.Name)
	if err != nil {
		return err
	}

	if rq.Expires != nil {
		secret.Expires = *rq.Expires
	}
	if rq.Status != nil {
		secret.Status = *rq.Status
	}
	if rq.Description != nil {
		secret.Description = *rq.Description
	}

	return b.ds.Secrets().Update(ctx, secret)
}
