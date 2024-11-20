package model


type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u *User) GetID() int {
	return u.ID
}

func (u *User) SetID(id int) {
	u.ID = id
}
