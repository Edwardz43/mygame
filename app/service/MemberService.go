package service

import (
	"github.com/Edwardz43/mygame/app/db/member/repository"
	"github.com/Edwardz43/mygame/app/db/member"
	
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
		Repo: repository.GetMemberInstance(dbGormConn),
	}
}

// Login returns member id if success.
func (service *MemberService) Login(nameOrEmail string) (uint, error) {
	logger.Printf("parameters [%v]", nameOrEmail)
	return service.Repo.GetOne(nameOrEmail)
}

// Register add new member and return true if success.
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
