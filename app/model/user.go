package model

// import (
// 	validation "github.com/go-ozzo/ozzo-validation"
// 	"github.com/go-ozzo/ozzo-validation/is"
// 	"golang.org/x/crypto/bcrypt"
// )

// User ...
type User struct {
	ID                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

// Validate ...
// func (u *User) Validate() error {
// 	return validation.ValidateStruct(
// 		u,
// 		validation.Field(&u.Email, validation.Required, is.Email),
// 		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
// 	)
// }