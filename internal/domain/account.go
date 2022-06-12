package domain

import "context"

type Account struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	Pan         string `json:"pan"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Sex         string `json:"sex"`
	Nationality string `json:"nationality"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type AccountRepo interface {
	Create(context.Context, *Account) (int, error)
	Get(context.Context, int) (*Account, error)
	GetByEmail(context.Context, string) (*Account, error)
	Update(context.Context, *Account) error
	Delete(context.Context, int) error
	SaveTransaction(context.Context, *Transaction) (int, error)
}

type AccountService interface {
	Create(context.Context, *Account) (int, error)
	Get(context.Context, int) (*Account, error)
	GetByEmail(context.Context, string) (*Account, error)
	Update(context.Context, *Account) error
	Delete(context.Context, int) error
	SaveTransaction(context.Context, *Transaction) (int, error)
}
