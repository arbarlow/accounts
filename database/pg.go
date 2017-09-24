package database

import (
	"strconv"
	"strings"

	_ "github.com/gemnasium/migrate/driver/postgres"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

var TOKEN_LENGTH = 32

func (a *Account) BeforeCreate(scope *gorm.Scope) error {
	t, err := GenerateRandomString(TOKEN_LENGTH)
	if err != nil {
		logrus.Errorf("confirm token generation error %v", err)
		return err
	}

	scope.SetColumn("confirmation_token", t)
	return nil
}

type PostgreSQL struct {
	Database
	conn string
	DB   *gorm.DB
}

func (p *PostgreSQL) Connect(conn string) error {
	p.conn = conn

	db, err := gorm.Open("postgres", conn)
	p.DB = db

	return err
}

func (p *PostgreSQL) Migrate() error {
	p.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err := p.DB.AutoMigrate(&Account{}).Error
	p.DB.Model(&Account{}).AddUniqueIndex("idx_email", "email")
	p.DB.Model(&Account{}).AddUniqueIndex("idx_phone", "phone_no")
	p.DB.Model(&Account{}).AddUniqueIndex("idx_confirm", "confirmation_token")
	p.DB.Model(&Account{}).AddUniqueIndex("idx_reset", "password_reset_token")
	return err
}

func (p *PostgreSQL) Close() error {
	return p.DB.Close()
}

func (p *PostgreSQL) Truncate() error {
	return p.DB.Exec("TRUNCATE accounts;").Error
}

func (p *PostgreSQL) List(count32 int32, token string) (accounts []*Account, next_token string, err error) {
	count := int(count32)
	if token == "" {
		token = "0"
	}

	offset, err := strconv.Atoi(token)
	if err != nil {
		return accounts, next_token, err
	}

	err = p.DB.Order("created_at DESC").Limit(count).Offset(offset).Find(&accounts).Error

	if err != nil {
		return accounts, next_token, err
	}

	if len(accounts) == int(count) {
		next_token = strconv.FormatInt(int64(offset+count+1), 10)
	}

	return accounts, next_token, err
}

func (p *PostgreSQL) Search(query string) (accounts []*Account, err error) {
	q := "%" + query + "%"
	err = p.DB.Order("created_at DESC").
		Where("name LIKE ?", q).
		Or("email LIKE ?", q).
		Or("phone_no LIKE ?", q).
		Or("username LIKE ?", q).
		Limit(20).
		Find(&accounts).Error
	return accounts, err
}

func (p *PostgreSQL) Get(id, email string) (*Account, error) {
	a := Account{ID: id, Email: email}
	var res *gorm.DB

	if a.ID != "" {
		res = p.DB.Where("id = ?", a.ID).First(&a)
	} else {
		res = p.DB.Where("email = ?", a.Email).First(&a)
	}

	if res.RecordNotFound() {
		return nil, ErrAccountNotFound
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return &a, nil
}

func (p *PostgreSQL) Create(a *Account, password string) error {
	err := p.DB.Create(a).Error

	if err != nil {
		return err
	}

	return nil
}

func (p *PostgreSQL) Update(a *Account) error {
	err := p.DB.Model(a).Updates(a).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgreSQL) GeneratePasswordToken(id, email string) (*Account, error) {
	a, err := p.Get(id, email)
	if err != nil {
		return nil, err
	}

	t, err := GenerateRandomString(TOKEN_LENGTH)
	if err != nil {
		logrus.Errorf("password token generation error %v", err)
		return nil, err
	}

	a.PasswordResetToken = t
	err = p.DB.Model(a).Updates(a).Error
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (p *PostgreSQL) UpdatePassword(token, hashed_password string) (*Account, error) {
	var a Account
	err := p.DB.Model(&a).
		Where("password_reset_token = ?", token).
		First(&a).Error

	if err != nil {
		return nil, err
	}

	a.HashedPassword = hashed_password
	err = p.DB.Model(&a).Updates(&a).Error
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (p *PostgreSQL) Confirm(token string) (*Account, error) {
	var a Account
	err := p.DB.Model(&a).
		Where("confirmation_token = ?", token).
		First(&a).Error

	if err != nil {
		return nil, err
	}

	a.ConfirmationToken = ""
	err = p.DB.Model(&a).Updates(&a).Error
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (p *PostgreSQL) Delete(ID string) error {
	a := Account{ID: ID}
	err := p.DB.Delete(&a).Error
	if err != nil {
		return err
	}

	return nil
}

func uniqueError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}
