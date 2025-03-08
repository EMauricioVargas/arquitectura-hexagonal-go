package out

import "github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/entities"

type UserRepository interface {
	FindByID(id string) (*entities.User, error)
}
