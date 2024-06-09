package user

import (
	"context"
	"database/sql"
	"fmt"
	"fullstack_toko/backend/model/domain"
	"log"
)

type Repository interface {
	GetUsername(ctx context.Context, tx *sql.Tx, username string) (*domain.User, error)
	GetByID(ctx context.Context, tx *sql.Tx, id int) (*domain.User, error)
	CreateUsers(ctx context.Context, tx *sql.Tx, us *domain.User) (int, error)
	CreateUsersProfile(ctx context.Context, tx *sql.Tx, userId int) error
	UpdateUserProfile(ctx context.Context, tx *sql.Tx, pf *domain.UserProfile) error
	GetByEmail(ctx context.Context, tx *sql.Tx, email string) (*domain.User, error)
	GetUserProfile(ctx context.Context, tx *sql.Tx, id int) (*domain.UserProfile, error)
}

type RepositoryImpl struct{}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) GetUsername(ctx context.Context, tx *sql.Tx, username string) (*domain.User, error) {

	us := &domain.User{}
	err := tx.QueryRowContext(ctx, "select * from users where username = ? ", username).Scan(&us.Id, &us.Name, &us.Username, &us.Password, &us.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("cannot find username")
		}
		return nil, err
	}
	return us, nil

}

func (r *RepositoryImpl) CreateUsers(ctx context.Context, tx *sql.Tx, us *domain.User) (int, error) {
	result, err := tx.ExecContext(ctx, "insert into users(name,username,password,email)  values(?,?,?,?)", us.Name, us.Username, us.Password, us.Email)
	if err != nil {

		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *RepositoryImpl) CreateUsersProfile(ctx context.Context, tx *sql.Tx, userId int) error {
	_, err := tx.ExecContext(ctx, "insert into users_profile (userid,deskripsi) values(?,?)", userId,"description not setting")
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryImpl) GetByID(ctx context.Context, tx *sql.Tx, id int) (*domain.User, error) {
	us := &domain.User{}
	err := tx.QueryRowContext(ctx, "select * from users where id = ? ", id).Scan(&us.Id, &us.Name, &us.Username, &us.Password, &us.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("cannot find users by id")
		}
		return nil, err
	}

	return us, nil

}

func (r *RepositoryImpl) UpdateUserProfile(ctx context.Context, tx *sql.Tx, pf *domain.UserProfile) error {
	result, err := tx.ExecContext(ctx, "update users_profile set url_image = ?,name = ? ,deskripsi = ? where userid= ?", pf.Url_image, pf.Name, pf.Deskripsi, pf.UserId)
	if err != nil {
		log.Println("error message:", err)
		return err
	}

	if affected, err := result.RowsAffected(); affected == 0 || err != nil {
		return fmt.Errorf("no rows affeected")
	}

	return nil
}

func (r *RepositoryImpl) GetByEmail(ctx context.Context, tx *sql.Tx, email string) (*domain.User, error) {
	us := &domain.User{}
	err := tx.QueryRowContext(ctx, "select * from users where email = ? ", email).Scan(&us.Id, &us.Name, &us.Username, &us.Password, &us.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("cannot find users by email")
		}
		return nil, err
	}

	return us, nil
}

func (r *RepositoryImpl) GetUserProfile(ctx context.Context, tx *sql.Tx, id int) (*domain.UserProfile, error) {
	u := &domain.UserProfile{}
	err := tx.QueryRowContext(ctx, "select id,url_image,name,deskripsi,userid from users_profile where userid = ?", id).Scan(&u.Id, &u.Url_image, &u.Name, &u.Deskripsi, &u.UserId)
	if err==sql.ErrNoRows{
		return nil,fmt.Errorf("cannot fine userId:%d",id)
	}
	return u, err
}
