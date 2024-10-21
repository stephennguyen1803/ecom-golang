package impl

import (
	"context"
	"ecom-project/internal/database"
)

type sUserLogin struct {
	query *database.Queries
}

func NewUserLogin(query *database.Queries) *sUserLogin {
	return &sUserLogin{
		query: query,
	}
}

func (s *sUserLogin) Login(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) Register(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) VerifyOTP(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context) error {
	return nil
}
