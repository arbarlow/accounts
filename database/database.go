package database

import (
	"errors"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrAccountExists   = errors.New("account already exists")
	ErrNoDatabase      = errors.New("no database connection details, (i.e POSTGRESQL_URL)")
	ErrNoPasswordGiven = errors.New("a password is required")
)

type Database interface {
	List(count int32, token string) ([]*Account, string, error)
	Get(id, email string) (*Account, error)
	Search(query string) ([]*Account, error)
	Create(a *Account, password string) error
	Update(a *Account) error
	Delete(ID string) error
	GeneratePasswordToken(id, email string) (*Account, error)
	UpdatePassword(string, string) (*Account, error)
	Migrate() error
	Truncate() error
	Close() error
}

type Account struct {
	ID                 string `sql:"type:uuid;primary key;default:uuid_generate_v1mc()"`
	Name               string `gorm:"default:null"`
	Email              string `gorm:"default:null"`
	Username           string `gorm:"default:null"`
	PhoneNo            string `gorm:"default:null"`
	HashedPassword     string `db:"hashed_password"`
	ConfirmationToken  string `gorm:"default:null"`
	PasswordResetToken string `gorm:"default:null"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (a *Account) HashPassword(password string) error {
	if password == "" {
		return ErrNoPasswordGiven
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	a.HashedPassword = string(hash[:])

	return nil
}

func (a *Account) ComparePasswordToHash(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(a.HashedPassword), []byte(password))
}

func DatabaseFromEnv() Database {
	var conn Database
	var err error

	pg := os.Getenv("POSTGRESQL_URL")
	if pg != "" {
		pgConn := &PostgreSQL{}
		err = pgConn.Connect(pg)
		conn = pgConn
	}

	if conn == nil {
		panic(ErrNoDatabase)
	}

	if err != nil {
		panic(err)
	}

	return conn
}
