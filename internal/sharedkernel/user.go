package sharedkernel

type User struct {
	id    string
	login string
}

func NewUser(login string) *User {
	return &User{
		id:    NewUUID(),
		login: login,
	}
}

func RestoreUser(id, login string) *User {
	return &User{
		id:    id,
		login: login,
	}
}

func (u *User) Login() string {
	return u.login
}

func (u *User) ID() string {
	return u.id
}
