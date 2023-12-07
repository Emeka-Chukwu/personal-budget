package usecase_user

import model_user "personal-budget/users/models"

// GetUserByEmail implements UsecaseUser.
func (us *usecaseuser) GetUserByEmail(email string) (model_user.UserResponse, error) {
	return us.repo.GetUserByEmail(email)
}
