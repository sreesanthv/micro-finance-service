package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sreesanthv/micro-finance-service/internal/domain"
)

type AccountRepo struct {
	db  *pgxpool.Pool
	log domain.Logger
}

func NewAccontRepo(db *pgxpool.Pool, log domain.Logger) *AccountRepo {
	return &AccountRepo{
		db:  db,
		log: log,
	}
}

func (r *AccountRepo) Create(ctx context.Context, a *domain.Account) (int, error) {
	var id int
	sql := `INSERT INTO accounts(name, location, pan, address, phone, sex, nationality, email, password) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	err := r.db.QueryRow(ctx, sql, a.Name, a.Location, a.Pan, a.Address, a.Phone, a.Sex, a.Nationality, a.Email, a.Password).Scan(&id)
	if err != nil {
		r.log.Errorf("Error in creating account:", err)
		return 0, err
	}

	return id, nil
}

func (r *AccountRepo) Get(ctx context.Context, id int) (*domain.Account, error) {
	act := new(domain.Account)
	q := `SELECT id, name, location, pan, address, phone, sex, nationality, email, password
		FROM accounts
			WHERE id = $1 AND deleted_at IS NULL`
	err := r.db.QueryRow(ctx, q, id).Scan(
		&act.ID,
		&act.Name,
		&act.Location,
		&act.Pan,
		&act.Address,
		&act.Phone,
		&act.Sex,
		&act.Nationality,
		&act.Email,
		&act.Password,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			r.log.Errorf("Error in querying account:", err)
		}
		return nil, err
	}

	return act, nil
}

func (r *AccountRepo) GetByEmail(ctx context.Context, email string) (*domain.Account, error) {
	act := new(domain.Account)
	q := `SELECT id, name, location, pan, address, phone, sex, nationality, email, password
		FROM accounts
			WHERE email = $1 AND deleted_at IS NULL`
	err := r.db.QueryRow(ctx, q, email).Scan(
		&act.ID,
		&act.Name,
		&act.Location,
		&act.Pan,
		&act.Address,
		&act.Phone,
		&act.Sex,
		&act.Nationality,
		&act.Email,
		&act.Password,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			r.log.Errorf("Error in querying account:", err)
		}
		return nil, err
	}

	return act, nil
}

func (r *AccountRepo) Update(ctx context.Context, a *domain.Account) error {
	fmt.Println("here")
	sql := `UPDATE accounts
		SET name = $1, location = $2, pan = $3, address = $4, phone = $5, sex = $6, nationality = $7, updated_at = NOW()
			WHERE id = $8`
	_, err := r.db.Exec(ctx, sql, a.Name, a.Location, a.Pan, a.Address, a.Phone, a.Sex, a.Nationality, a.ID)
	if err != nil {
		r.log.Errorf("Error in updating account:", err)
		return err
	}

	return nil
}

func (r *AccountRepo) Delete(ctx context.Context, id int) error {
	sql := "UPDATE accounts SET deleted_at = NOW() WHERE id = $1"
	_, err := r.db.Exec(ctx, sql, id)
	if err != nil {
		r.log.Errorf("Error in updating account:", err)
		return err
	}

	return nil
}

func (r *AccountRepo) SaveTransaction(ctx context.Context, t *domain.Transaction) (int, error) {
	var id int
	sql := `INSERT INTO transactions(account_id, debit, credit, narration) 
		VALUES($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRow(ctx, sql, t.AccountId, t.Debit, t.Credit, t.Narration).Scan(&id)
	if err != nil {
		r.log.Errorf("Error in creating account:", err)
		return 0, err
	}

	return id, nil
}
