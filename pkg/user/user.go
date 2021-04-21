package user

type Getter interface {
	GetAll() []User
}

type Adder interface {
	Add(score User)
}

type User struct {
	ID       int    `json:"ID"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"pass"`
}
