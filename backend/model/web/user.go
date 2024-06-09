package web

import "context"

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

type UserProfile struct {
	Id        int    `json:"id"`
	Url_image string `json:"url"`
	Name      string `json:"name"`
	Deskripsi string `json:"deskripsi"`
	UserId    int    `json:"userid"`
}

type UserLoginPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserProfileUpdatePayload struct {
	Url_image string `json:"url"`
	Name      string `json:"name"`
	Deskripsi string `json:"deskripsi"`
	UserId    int    `json:"userid" validate:"required"`
}

type UserUpdatePayload struct {
	Userid   int    `json:"userid" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type UserRegisterPayload struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required,min=8"`
	Password string `json:"password" validate:"required,min=8"`
	VPasword string `json:"vpassword" validate:"required,eqfield=Password"`
	Email    string `json:"email" validate:"required,email"`
}

type UserProfileRegisterPayload struct {
	Url_image string `json:"url"`
	Name      string `json:"name"`
	Deskripsi string `json:"deskripsi"`
	UserId    int    `json:"userid" validate:"required"`
}

type UserService interface {
	GetUsername(ctx context.Context, username string) (*User, error)
	CreateUsers(ctx context.Context, us *UserRegisterPayload) (int, error)
	CreateUsersProfile(ctx context.Context, userId int) error
	GetByID(ctx context.Context, id int) (*User, error)
	UpdateUserProfile(ctx context.Context, pf *UserProfileUpdatePayload) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetUserProfile(ctx context.Context, id int) (*UserProfile, error)
}
