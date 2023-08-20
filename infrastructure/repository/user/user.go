package user

import (
	"database/sql"
	"fmt"
	"log"

	domainUser "github.com/aman-singh7/loyalty-blockchain/domain/user"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	_, err := db.Exec(INSTALL_UUID_EXTENSION)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = db.Exec(CREATE_USER_TABLE)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(userDomain *domainUser.User) (*domainUser.User, error) {
	rows, err := r.db.Query(INSERT_INTO_USERS,
		userDomain.Name,
		userDomain.Email,
		userDomain.PlatformUID,
		userDomain.AccountType,
		userDomain.WalletAddress,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userId := ""
	for rows.Next() {
		rows.Scan(&userId)
	}

	userDomain.UID = userId
	return userDomain, nil
}

func (r *Repository) GetByUID(uid string) (*domainUser.User, error) {
	rows, err := r.db.Query(GET_BY_UID, uid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanIntoUser(rows)
	}

	return nil, fmt.Errorf("user not found")
}

func (r *Repository) GetByPUID(puid string) (*domainUser.User, error) {
	rows, err := r.db.Query(GET_BY_PUID, puid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanIntoUser(rows)
	}

	return nil, fmt.Errorf("user not found")
}

func (r *Repository) UpdateUser(userDomain *domainUser.User) (*domainUser.User, error) {
	rows, err := r.db.Query(UPDATE_USER, userDomain.UID, userDomain.Name, userDomain.Email,
		userDomain.AccountType, userDomain.WalletAddress)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanIntoUser(rows)
	}

	return nil, fmt.Errorf("error occurred while updating user")
}

func (r *Repository) UserExists(puid string) (string, error) {
	userId := ""
	err := r.db.QueryRow(CHECK_USER_EXISTS, puid).Scan(&userId)

	return userId, err
}
