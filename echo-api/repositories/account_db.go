package repositories

import "gorm.io/gorm"

type accountRepositoryDB struct {
	db *gorm.DB
}

// Create implements AccountRepository.
func (r accountRepositoryDB) Create(account Account) (*Account, error) {
	tx := r.db.Create(&account)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &account, nil
}

// GetByUsername implements AccountRepository.
func (r accountRepositoryDB) GetByUsername(username string) (*Account, error) {
	account := Account{}
	tx := r.db.Find(&account, "username=?", username)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &account, nil
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	db.AutoMigrate(&Account{})
	return accountRepositoryDB{db: db}
}
