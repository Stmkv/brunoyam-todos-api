package users

type User struct {
	UID      string
	Name     string
	Email    string
	Password string
}

func NewUser(uid string, name string, email, password string) (*User, error) {
	checks := []struct {
		value string
		err   error
	}{
		{uid, ErrEmptyUID},
		{name, ErrEmptyName},
		{password, ErrEmptyPassword},
		{email, ErrEmptyEmail},
	}

	for _, c := range checks {
		if c.value == "" {
			return nil, c.err
		}
	}

	return &User{
		UID:      uid,
		Name:     name,
		Email:    email,
		Password: password,
	}, nil
}
