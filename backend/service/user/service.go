package user

import (
	"context"
	"database/sql"
	"fullstack_toko/backend/exception"
	"fullstack_toko/backend/model/domain"
	"fullstack_toko/backend/model/web"
	"fullstack_toko/backend/utils"
)

type ServiceImpl struct {
	Repo Repository
	Db   *sql.DB
}

func NewService(repo Repository, db *sql.DB) *ServiceImpl {
	return &ServiceImpl{
		Repo: repo,
		Db:   db,
	}
}

func (s *ServiceImpl) GetUsername(ctx context.Context, username string) (*web.User, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer exception.CommitOrRollback(tx)

	us, err := s.Repo.GetUsername(ctx, tx, username)
	if err != nil {
		return nil, err
	}

	return utils.ConvertUserIntoWeb(us), nil
}

func (s *ServiceImpl) CreateUsers(ctx context.Context, us *web.UserRegisterPayload) (int, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return 0, err
	}
	defer exception.CommitOrRollback(tx)

	id, err := s.Repo.CreateUsers(ctx, tx, &domain.User{
		Name:     us.Name,
		Username: us.Username,
		Password: us.Password,
		Email:    us.Email,
	})
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (s *ServiceImpl) CreateUsersProfile(ctx context.Context, userId int) error {
	tx, err := s.Db.Begin()
	if err != nil {
		return err
	}
	defer exception.CommitOrRollback(tx)

	if err := s.Repo.CreateUsersProfile(ctx, tx, userId); err != nil {
		return err
	}

	return nil
}

func (s *ServiceImpl) GetByID(ctx context.Context, id int) (*web.User, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer exception.CommitOrRollback(tx)

	us, err := s.Repo.GetByID(ctx, tx, id)
	if err != nil {
		return nil, err
	}

	return utils.ConvertUserIntoWeb(us), nil
}

func (s *ServiceImpl) UpdateUserProfile(ctx context.Context, pf *web.UserProfileUpdatePayload) error {
	tx, err := s.Db.Begin()
	if err != nil {
		return err
	}
	defer exception.CommitOrRollback(tx)

	err = s.Repo.UpdateUserProfile(ctx, tx, &domain.UserProfile{
		Url_image: pf.Url_image,
		Name:      pf.Name,
		Deskripsi: pf.Deskripsi,
		UserId:    pf.UserId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceImpl) GetByEmail(ctx context.Context, email string) (*web.User, error) {
	tx, err := s.Db.Begin()
	if err != nil {
		return nil, err
	}
	defer exception.CommitOrRollback(tx)

	us, err := s.Repo.GetByEmail(ctx, tx, email)
	if err != nil {
		return nil, err
	}

	return utils.ConvertUserIntoWeb(us), nil
}

func (s *ServiceImpl) GetUserProfile(ctx context.Context, id int) (*web.UserProfile, error) {
	tx,err:=s.Db.Begin()
	if err != nil {
		return nil, err
	}

	defer exception.CommitOrRollback(tx)

	us,err:=s.Repo.GetUserProfile(ctx,tx,id)
	if err != nil {
		return nil, err
	}

	return utils.ConvertProfileIntoWeb(us),nil
}
