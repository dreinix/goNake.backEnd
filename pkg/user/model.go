package user

type User struct {
	ID       int    `json:"ID"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   string `json:"status"`
}
