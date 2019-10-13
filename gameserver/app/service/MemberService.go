package service

import (
	"github.com/Edwardz43/mygame/gameserver/db"
	"github.com/Edwardz43/mygame/gameserver/db/member"
	"github.com/Edwardz43/mygame/gameserver/db/member/repository"
	"golang.org/x/crypto/bcrypt"
)

// MemberService ...
type MemberService struct {
	Repo member.Repository
}

// GetLoginInstance returns instance of lobby service.
func GetLoginInstance() *MemberService {
	return &MemberService{
		// Repo: repository.NewMysqlGameResultRepository(db.Connect()),
		Repo: repository.GetMemberInstance(db.ConnectGorm()),
	}
}

// Login ...
func (service *MemberService) Login(nameOrEmail string) (uint, error) {
	logger.Printf("parameters [%d]", nameOrEmail)
	return service.Repo.GetOne(nameOrEmail)
}

// Register ...
func (service *MemberService) Register(name, email, password string) (bool, error) {

	pw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Printf("Password crypto error : [%v]", err)
		return false, err
	}

	logger.Printf("parameters [%v;%v;%v]", name, email, pw)
	ok, err := service.Repo.Create(name, email, string(pw))

	if !ok && err != nil {
		logger.Printf("Register error : [%v]", err)
		return false, err
	}

	return ok, nil
}
