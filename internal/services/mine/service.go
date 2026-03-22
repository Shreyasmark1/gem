package mine

import (
	"context"
	"gem/internal/store"
	"gem/internal/store/repository"

	"github.com/google/uuid"
)

type MineService struct {
	store *store.Store
}

func NewMineService(store *store.Store) *MineService {
	return &MineService{
		store: store,
	}
}

func (m *MineService) SaveMine(ctx context.Context, body SaveMine) (*SaveMine, error) {
	if body.MineID == "" {
		return m.createMine(ctx, body)
	} else {
		return m.updateMine(ctx, body)
	}
}

func (m *MineService) createMine(ctx context.Context, body SaveMine) (*SaveMine, error) {
	// For now, assuming owner_id is passed in context or mocking it if context doesn't have it.
	// Since we don't have the explicit session middleware context key, we will extract it or use an empty UUID.
	ownerIDStr, _ := ctx.Value("user_id").(string)
	if ownerIDStr == "" {
		// fallback to a mock UUID if no auth middleware is wired yet
		ownerIDStr = "00000000-0000-0000-0000-000000000000"
	}
	ownerID, err := uuid.Parse(ownerIDStr)
	if err != nil {
		ownerID = uuid.Nil
	}

	mine, err := m.store.Queries.CreateMine(ctx, repository.CreateMineParams{
		OwnerID:   ownerID,
		MineName:  body.MineName,
		CreatedBy: ownerID,
	})
	if err != nil {
		return nil, err
	}
	return &SaveMine{
		MineID:   mine.MineID.String(),
		MineName: mine.MineName,
	}, nil
}

func (m *MineService) updateMine(ctx context.Context, body SaveMine) (*SaveMine, error) {
	mineID, err := uuid.Parse(body.MineID)
	if err != nil {
		return nil, err
	}
	ownerIDStr, _ := ctx.Value("user_id").(string)
	if ownerIDStr == "" {
		ownerIDStr = "00000000-0000-0000-0000-000000000000"
	}
	ownerID, err := uuid.Parse(ownerIDStr)
	if err != nil {
		ownerID = uuid.Nil
	}

	mine, err := m.store.Queries.UpdateMine(ctx, repository.UpdateMineParams{
		MineID:    mineID,
		MineName:  body.MineName,
		UpdatedBy: uuid.NullUUID{UUID: ownerID, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &SaveMine{
		MineID:   mine.MineID.String(),
		MineName: mine.MineName,
	}, nil
}
