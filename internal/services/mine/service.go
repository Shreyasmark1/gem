package mine

import (
	"context"
	appMiddleware "gem/internal/middlewares"
	"gem/internal/store"
	"gem/internal/store/repository"
	"gem/pkg/db"
	"gem/pkg/response"

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
	session, ok := ctx.Value(appMiddleware.SessionContextKey).(*appMiddleware.Session)

	if !ok || session == nil || session.UserID == "" {
		return nil, response.Unauthorized("missing user session")
	}

	ownerID, err := uuid.Parse(session.UserID)

	if err != nil {
		return nil, response.Unauthorized("invalid user id format")
	}

	mine, err := db.DbExe(m.store.Queries.CreateMine(ctx, repository.CreateMineParams{
		OwnerID:   ownerID,
		MineName:  body.MineName,
		CreatedBy: ownerID,
	}))

	if err != nil {
		return nil, response.InternalServerError("failed to create mine", err)
	}
	if mine == nil {
		return nil, response.InternalServerError("failed to create mine (nil result)")
	}

	return &SaveMine{
		MineID:   mine.MineID.String(),
		MineName: mine.MineName,
	}, nil
}

func (m *MineService) updateMine(ctx context.Context, body SaveMine) (*SaveMine, error) {
	mineID, err := uuid.Parse(body.MineID)

	if err != nil {
		return nil, response.BadRequest("invalid mine id format in request")
	}

	session, ok := ctx.Value(appMiddleware.SessionContextKey).(*appMiddleware.Session)

	if !ok || session == nil || session.UserID == "" {
		return nil, response.Unauthorized("missing user session")
	}

	ownerID, err := uuid.Parse(session.UserID)

	if err != nil {
		return nil, response.Unauthorized("invalid user id format")
	}

	existingMine, err := db.DbExe(m.store.Queries.GetMine(ctx, mineID))
	if err != nil {
		return nil, response.InternalServerError("failed to fetch mine", err)
	}
	if existingMine == nil {
		return nil, response.NotFound("mine not found")
	}

	if existingMine.OwnerID != ownerID {
		return nil, response.Unauthorized("you do not have permission to update this mine")
	}

	mine, err := db.DbExe(m.store.Queries.UpdateMine(ctx, repository.UpdateMineParams{
		MineID:    mineID,
		MineName:  body.MineName,
		UpdatedBy: uuid.NullUUID{UUID: ownerID, Valid: true},
	}))

	if err != nil {
		return nil, response.InternalServerError("failed to update mine", err)
	}
	if mine == nil {
		return nil, response.InternalServerError("failed to update mine (nil result)")
	}

	return &SaveMine{
		MineID:   mine.MineID.String(),
		MineName: mine.MineName,
	}, nil
}
