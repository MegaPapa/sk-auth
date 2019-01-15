package entity

import (
	"sk-auth/auth/crypto"
	"time"
)

// System user which uses for authentication in all systems.
// Here, 'Mandatory' means this info is mandatory for user creating.
// Nickname - Mandatory. Unique. User's nickname for all systems.
// Email - Mandatory. Unique. User's email.
// Password - Mandatory. Hashed password.
// FirstName - Not mandatory. User's first name.
// SecondName - Not mandatory. User's second name.
// Gender - Not mandatory. User gender. Can be M or F. Not a "Sign up info".
// PhoneNumber - Not mandatory. Phone number should be configurated in settings and isn't a "Sign up info".
// CreatedTime - Not mandatory. Timestamp when user was created.
// AuthTokens - Not mandatory. "History". All tokens that was used by this user. Don't used for json serialization.
// Roles - Mandatory. We don't need to put this field to json, so we have only bson mapping.
type User struct {
	Id          int64        `json:"id" bson:"_id"`
	Nickname    string       `json:"nickname" bson:"nickname" binding:"required"`
	Email       string       `json:"email" bson:"email" binding:"required"`
	Password    string       `json:"password" bson:"password" binding:"required"`
	FirstName   string       `json:"firstName" bson:"firstName"`
	LastName    string       `json:"lastName" bson:"lastName"`
	Gender      string       `json:"gender" bson:"gender"`
	PhoneNumber string       `json:"phoneNumber" bson:"phoneNumber"`
	CreatedTime time.Time    `json:"createdTime" bson:"createdTime"`
	AuthTokens  []*AuthToken `bson:"tokens"`
	Roles       []*UserRole  `bson:"roles"`
}

const (
	_UNDEFINED_PASSWORD = "undefined"
)

// Factory function for User entity
func CreateUser(nickname, email, password string) {
	user := new(User)
	user.Nickname = nickname
	user.Email = email
	user.CreatedTime = time.Now()
	encryptedPassword, err := crypto.EncryptPassword(password)
	// If we has some problems with encryption
	// set undefined value for password
	if err == nil {
		user.Password = encryptedPassword
	} else {
		user.Password = _UNDEFINED_PASSWORD
	}
	user.Roles = append(user.Roles, USER_ROLE)
}
