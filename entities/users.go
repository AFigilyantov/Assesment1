package entity

type Users struct {
	WhiteList map[string]struct{}
}

func NewUsers() *Users {
	var u Users
	u.WhiteList = make(map[string]struct{})
	return &u
}

func (u *Users) AddNewUser(token string) {
	u.WhiteList[token] = struct{}{}
}
