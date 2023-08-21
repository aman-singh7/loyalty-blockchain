package brand

import (
	"database/sql"

	domainBrand "github.com/aman-singh7/loyalty-blockchain/domain/brand"
)

func scanIntoBrand(rows *sql.Rows) (*domainBrand.Brand, error) {
	brand := new(domainBrand.Brand)

	err := rows.Scan(
		&brand.UID,
		&brand.Name,
		&brand.Email,
		&brand.PlatformUID,
		&brand.AccountType,
		&brand.PlatformUID,
	)

	return brand, err
}