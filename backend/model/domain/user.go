package domain

type User struct {
	Id       int
	Name     string
	Username string
	Password string
	Email    string
}

type UserProfile struct {
	Id        int
	Url_image string
	Name      string
	Deskripsi string
	UserId    int
}
