package brand

const (
	INSTALL_UUID_EXTENSION = `CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`

	CREATE_BRAND_TABLE = `CREATE TABLE IF NOT EXISTS brand (
		uid UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
		name TEXT,
		email TEXT,
		puid TEXT UNIQUE,
		account_type INTEGER,
		wallet_address TEXT
	)`

	INSERT_INTO_BRANDS = `INSERT INTO brands
	(name, email, puid, account_type, wallet_address)
	values
	($1, $2, $3, $4) RETURNING id
	`

	GET_BY_PUID = `SELECT * FROM brands WHERE puid = $1`
	GET_BY_UID  = `SELECT * FROM brands WHERE uid = $1`

	UPDATE_BRAND       = `UPDATE brands SET name = $2, email = $3, account_type = $4, wallet_address = $5 WHERE uid = $1 RETURNING *`
	CHECK_BRAND_EXISTS = `SELECT uid FROM brands WHERE puid = $1`

	GET_WALLET_ADDRESS = `SELECT wallet_address FROM brands WHERE uid = $1`
)
