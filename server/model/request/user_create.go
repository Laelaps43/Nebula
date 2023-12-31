package request

type UserCreate struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Roles    []uint `json:"roles"`
}

type UserUpdate struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Roles []uint `json:"roles"`
}
