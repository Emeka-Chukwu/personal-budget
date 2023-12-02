package usecase_account

import "github.com/google/uuid"

// Delete implements AccountUsecase.
func (us *accountUsecase) Delete(id uuid.UUID) error {
	return us.repo.Delete(id)
}
