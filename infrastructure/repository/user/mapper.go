package user

import (
	"database/sql"

	domainUser "github.com/aman-singh7/loyalty-blockchain/domain/user"
)

func scanIntoUser(rows *sql.Rows) (*domainUser.User, error) {
	user := new(domainUser.User)

	err := rows.Scan(
		&user.UID,
		&user.Name,
		&user.Email,
		&user.PlatformUID,
		&user.AccountType,
		&user.PlatformUID,
	)

	return user, err
}
