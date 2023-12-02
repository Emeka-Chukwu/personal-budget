package repositories_account

import (
	"context"
	"personal-budget/util"

	"github.com/google/uuid"
)

// Delete implements AccountRepository.
func (repo *accountRepository) Delete(id uuid.UUID) error {
	_, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmp := `delete from accounts where id=$1`
	_, err := repo.DB.Exec(stmp, id)
	return err
}
