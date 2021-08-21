package security

import (
	"context"

	"github.com/cbuelvasc/cinema-backend/model"
	"github.com/cbuelvasc/cinema-backend/repository"
	"github.com/cbuelvasc/cinema-backend/util"
)

type AuthValidator struct {
	userRepository repository.UserRepository
}

func NewAuthValidator(userRepository repository.UserRepository) *AuthValidator {
	return &AuthValidator{userRepository: userRepository}
}

func (authValidator *AuthValidator) ValidateCredentials(ctx context.Context, username, password string) (*model.User, bool) {
	user, err := authValidator.userRepository.FindByEmail(ctx, username)
	if err != nil || util.VerifyPassword(user.Password, password) != nil {
		return nil, false
	}
	return user, true
}
