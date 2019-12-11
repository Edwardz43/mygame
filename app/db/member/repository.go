package member

// Repository ...
type Repository interface {
	GetOne(nameOrEmail string) (uint, error)
	Create(name, email, password string) (bool, error)
}
