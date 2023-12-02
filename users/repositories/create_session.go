package repositories_users

import (
	"context"
	model_user "personal-budget/users/models"
	"personal-budget/util"
)

// CreateSession implements UserAuthentication.
func (auth *authentication) CreateSession(data model_user.Session) (model_user.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `insert into sessions (refresh_token, user_id, user_agent, client_ip)
		values ($1, $2, $3, $4) returning id, refresh_token,user_id, user_agent, client_ip, created_at, updated_at`
	var sess model_user.Session
	err := auth.DB.QueryRowContext(ctx, stmt,
		data.Refresh,
		data.UserID,
		data.UserAgent,
		data.ClientIP,
	).Scan(
		&sess.ID,
		&sess.Refresh,
		&sess.UserID,
		&sess.UserAgent,
		&sess.ClientIP,
		&sess.CreatedAt,
		&sess.UpdatedAt,
	)
	return sess, err
}
