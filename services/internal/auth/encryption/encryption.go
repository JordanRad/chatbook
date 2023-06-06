package encryption

import "golang.org/x/crypto/bcrypt"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . Encryption
type Encryption interface {
	EncryptPassword(passwordString string) (string, error)
	CheckPassword(hashString string, passwordString string) bool
}

type Encrypter struct{}

// Compile time assertion that this service implements the generated interface
var _ Encryption = (*Encrypter)(nil)

// This methods takes a plain string as a parameter and
// returns the encrypted value of the same string
func (e *Encrypter) EncryptPassword(passwordString string) (string, error) {
	password := []byte(passwordString)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// This method compares a given hash string with
// a plain string and returns if they are equal or not
func (e *Encrypter) CheckPassword(hashString string, passwordString string) bool {
	password := []byte(passwordString)
	hash := []byte(hashString)
	err := bcrypt.CompareHashAndPassword(hash, password)

	return err == nil
}
