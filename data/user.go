package data

import (
	"errors"
	"time"

	"github.com/RobertoSuarez/art-digital/db"
	"github.com/RobertoSuarez/art-digital/types"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type StatusUser string

const (
	Blocked StatusUser = "blocked"
	Actived StatusUser = "actived"
)

type TypeUser string

const (
	NormalUser TypeUser = "normal_user"
	SuperUser  TypeUser = "superuser"
)

type User struct {
	gorm.Model
	Name      string
	Email     string
	Password  string
	Status    StatusUser
	Type      TypeUser
	CountryID uint
	Country   Country
	Birthday  time.Time
}

func RegisterUser(userApi *types.APIUser) error {

	hashpassword, err := HashPassword(userApi.Password)
	if err != nil {
		return err
	}

	user := User{
		Name:      userApi.Name,
		Email:     userApi.Email,
		Password:  hashpassword,
		Status:    StatusUser(userApi.Status),
		Type:      TypeUser(userApi.Type),
		CountryID: userApi.CountryID,
		Birthday:  userApi.Birthday,
	}

	err = db.DB.Create(&user).Error
	if err != nil {
		return err
	}

	// establecemos el id
	userApi.ID = user.ID
	userApi.Password = hashpassword

	return nil
}

// Login se le pasan las credenciales, y retorna el usuario, si esta correcta
// las credenciales y si existe algun erro.
func Login(login types.Login) (types.APIUser, bool, error) {
	var user User
	result := db.DB.First(&user, "email = ?", login.Email).Omit("Password")
	if result.RowsAffected < 1 {
		return types.APIUser{}, false, result.Error
	}

	// checkear la pass
	ok := CheckPasswordHash(login.Password, user.Password)

	if ok {

		userAPi := types.APIUser{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Status:    string(user.Status),
			Type:      string(user.Type),
			CountryID: user.CountryID,
			Birthday:  user.Birthday,
		}
		return userAPi, ok, nil
	} else {
		return types.APIUser{}, ok, errors.New("las credenciales no eran correctas")
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
