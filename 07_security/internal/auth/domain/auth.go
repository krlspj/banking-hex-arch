package domain

type AuthRepository interface {
	FindBy(string, string) (*Login, error)
}
