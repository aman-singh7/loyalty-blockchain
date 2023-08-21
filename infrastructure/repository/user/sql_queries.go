package user

const (
	INSTALL_UUID_EXTENSION = `CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`

	CREATE_USER_TABLE = `CREATE TABLE IF NOT EXISTS users (
		uid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
		name TEXT,
		email TEXT,
		puid TEXT UNIQUE,
		account_type INTEGER,
		wallet_address TEXT
	)`

	INSERT_INTO_USERS = `INSERT INTO users
	(name, email, puid, account_type, wallet_address)
	values
	($1, $2, $3, $4, $5) RETURNING uid
	`

	GET_BY_PUID = `SELECT * FROM users WHERE puid = $1`
	GET_BY_UID  = `SELECT * FROM users WHERE uid = $1`

	UPDATE_USER       = `UPDATE users SET name = $2, email = $3, account_type = $4, wallet_address = $5 WHERE uid = $1 RETURNING *`
	CHECK_USER_EXISTS = `SELECT uid FROM users WHERE puid = $1`

	GET_WALLET_ADDRESS = `SELECT wallet_address FROM users WHERE uid = $1`
)
