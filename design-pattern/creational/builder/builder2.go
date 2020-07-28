package builder

type User struct {
	Name      string
	Role      string
	MinSalary int
	MaxSalary int
}

type UserBuilder struct {
	User
}

func (ub *UserBuilder) Build() User {
	return ub.User
}

func (ub *UserBuilder) Name(name string) *UserBuilder {
	ub.User.Name = name
	return ub
}

func (ub *UserBuilder) Role(role string) *UserBuilder {
	// verify the role is valid
	if role == "manager" {
		ub.User.MinSalary = 20000
		ub.User.MaxSalary = 60000
	}
	ub.User.Role = role
	return ub
}
