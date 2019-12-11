package repository

import (
	"fmt"
	"regexp"

	"github.com/Edwardz43/mygame/app/db/models"
	"github.com/jinzhu/gorm"
)

// MemberRepository ...
type MemberRepository struct {
	db *gorm.DB
}

// GetMemberInstance ...
func GetMemberInstance(db *gorm.DB) *MemberRepository {
	return &MemberRepository{db: db}
}

// GetOne finds the member and returns the member ID
func (repo MemberRepository) GetOne(nameOrEmail string) (uint, error) {
	matched, err := regexp.Match(
		"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
		[]byte(nameOrEmail))
	if err != nil {
		return 0, err
	}
	var member models.Member
	var column string

	if matched {
		column = "email=?"
	} else {
		column = "name=?"
	}
	d := repo.db.Where(column, nameOrEmail).Find(&member)
	return member.ID, d.Error
}

// Create creates a new member
func (repo MemberRepository) Create(name, email, password string) (bool, error) {

	// validate name
	matched, err := regexp.Match("[a-zA-Z0-9]+", []byte(name))
	if err != nil {
		return false, err
	}

	if !matched {
		return false, fmt.Errorf("invalid name")
	}

	// validate email
	matched, err = regexp.Match(
		"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
		[]byte(email))
	if err != nil {
		return false, err
	}

	if !matched {
		return false, fmt.Errorf("invalid email address")
	}

	member := &models.Member{
		Name:     name,
		Email:    email,
		Password: password,
	}
	d := repo.db.Create(member)

	if d.Error != nil {
		return false, d.Error
	}

	return true, nil
}
