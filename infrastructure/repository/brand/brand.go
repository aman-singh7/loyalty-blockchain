package brand

import (
	"database/sql"
	"fmt"
	"log"

	domainBrand "github.com/aman-singh7/loyalty-blockchain/domain/brand"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	_, err := db.Exec(INSTALL_UUID_EXTENSION)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = db.Exec(CREATE_BRAND_TABLE)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(brandDomain *domainBrand.Brand) (*domainBrand.Brand, error) {
	rows, err := r.db.Query(INSERT_INTO_BRANDS,
		brandDomain.Name,
		brandDomain.Email,
		brandDomain.PlatformUID,
		brandDomain.AccountType,
		brandDomain.WalletAddress,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	brandId := ""
	for rows.Next() {
		rows.Scan(&brandId)
	}

	brandDomain.UID = brandId
	return brandDomain, nil
}

func (r *Repository) GetByUID(uid string) (*domainBrand.Brand, error) {
	rows, err := r.db.Query(GET_BY_UID, uid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanIntoBrand(rows)
	}

	return nil, fmt.Errorf("brand not found")
}

func (r *Repository) GetByPUID(puid string) (*domainBrand.Brand, error) {
	rows, err := r.db.Query(GET_BY_PUID, puid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanIntoBrand(rows)
	}

	return nil, fmt.Errorf("brand not found")
}

func (r *Repository) UpdateBrand(brandDomain *domainBrand.Brand) (*domainBrand.Brand, error) {
	rows, err := r.db.Query(UPDATE_BRAND, brandDomain.UID, brandDomain.Name, brandDomain.Email,
		brandDomain.AccountType, brandDomain.WalletAddress)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		return scanIntoBrand(rows)
	}

	return nil, fmt.Errorf("error occurred while updating brand")
}

func (r *Repository) BrandExists(puid string) (string, error) {
	brandId := ""
	err := r.db.QueryRow(CHECK_BRAND_EXISTS, puid).Scan(&brandId)

	return brandId, err
}

func (r *Repository) GetAddress(uid string) (string, error) {
	address := ""
	err := r.db.QueryRow(GET_WALLET_ADDRESS, uid).Scan(&address)

	return address, err
}