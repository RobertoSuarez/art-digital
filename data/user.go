package data

import (
	"errors"

	"time"

	"github.com/RobertoSuarez/art-digital/db"
	"github.com/RobertoSuarez/art-digital/types"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type StatusUser string

const (
	Actived StatusUser = "actived"
	Blocked StatusUser = "blocked"
)

type TypeUser string

const (
	NormalUser TypeUser = "normal_user"
	Superuser  TypeUser = "superuser"
)

type User struct {
	gorm.Model
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Status    StatusUser
	TypeUser  TypeUser
	CountryID uint
	Country   Country
	Birthday  time.Time
}

func RegisterUser(userAPI *types.UserAPI) error {

	err := userAPI.Validate()
	if err != nil {
		return err
	}

	hash, err := HashPassword(userAPI.Password)
	if err != nil {
		return err
	}

	userAPI.Password = hash

	user := User{
		Name:      userAPI.Name,
		Email:     userAPI.Email,
		Password:  userAPI.Password,
		Status:    StatusUser(userAPI.Status),
		TypeUser:  TypeUser(userAPI.Type),
		CountryID: userAPI.CountryID,
		Birthday:  userAPI.Birthday,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return errors.New("correo electronico duplicado")
			}
		}
	}

	userAPI.Password = ""
	userAPI.ID = user.ID

	return nil
}

func Login(login *types.Login) (types.UserAPI, error) {

	// TODO: Validar que exista el registro del usuario.

	var user User
	result := db.DB.Where("email = ?", login.Email).First(&user)
	if result.RowsAffected < 1 {
		// no existe el registro
		return types.UserAPI{}, errors.New("el registro no existe")
	}

	// TODO: Revisar el password con el hash
	ok := CheckPasswordHash(login.Password, user.Password)
	if ok {
		userAPI := types.UserAPI{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Password:  "",
			Status:    string(user.Status),
			Type:      string(user.TypeUser),
			CountryID: user.CountryID,
			Birthday:  user.Birthday,
		}
		return userAPI, nil
	} else {
		return types.UserAPI{}, errors.New("credenciales incorrectas")
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
