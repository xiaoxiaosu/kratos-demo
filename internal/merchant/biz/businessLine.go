package biz

import (
	"context"
)

type BusinessLine struct {
	Name string
	Tag  string
}

type BusinessLineRepo interface {
	Save(context.Context, *BusinessLine) (bool, error)
	Update(context.Context, *BusinessLine) (bool, error)
	Get(context.Context, int64) (*BusinessLine, error)
	ListAll(context.Context) ([]*BusinessLine, error)
}

type BusinessLineUseCase struct {
	repo BusinessLineRepo
}

func NewBusinessLineUseCase(repo BusinessLineRepo) *BusinessLineUseCase {
	return &BusinessLineUseCase{repo: repo}
}

func (bc *BusinessLineUseCase) CreateBusinessLine(ctx context.Context, b *BusinessLine) (bool, error) {
	return bc.repo.Save(ctx, b)
}

func (bc *BusinessLineUseCase) UpdateBusinessLine(ctx context.Context, b *BusinessLine) (bool, error) {
	return bc.repo.Update(ctx, b)
}

func (bc *BusinessLineUseCase) GetBusinessLine(ctx context.Context, id int64) (*BusinessLine, error) {
	return bc.repo.Get(ctx, id)
}

func (bc *BusinessLineUseCase) ListBusinessLine(ctx context.Context) ([]*BusinessLine, error) {
	return bc.repo.ListAll(ctx)
}
