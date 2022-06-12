package service

import (
	"context"

	"github.com/sreesanthv/micro-finance-service/internal/domain"
)

type AccountService struct {
	repo domain.AccountRepo
	log  domain.Logger
}

func NewAccountService(repo domain.AccountRepo, log domain.Logger) *AccountService {
	return &AccountService{
		repo: repo,
		log:  log,
	}
}

func (r *AccountService) Create(ctx context.Context, a *domain.Account) (int, error) {
	return r.repo.Create(ctx, a)
}

func (r *AccountService) Get(ctx context.Context, id int) (*domain.Account, error) {
	return r.repo.Get(ctx, id)
}
func (r *AccountService) GetByEmail(ctx context.Context, email string) (*domain.Account, error) {
	return r.repo.GetByEmail(ctx, email)
}

func (r *AccountService) Update(ctx context.Context, a *domain.Account) error {
	return r.repo.Update(ctx, a)
}

func (r *AccountService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

func (r *AccountService) SaveTransaction(ctx context.Context, t *domain.Transaction) (int, error) {
	return r.repo.SaveTransaction(ctx, t)
}
