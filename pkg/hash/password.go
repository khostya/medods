package hash

import (
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	cost int
}

func NewBcrypt(cost int) Bcrypt {
	return Bcrypt{cost: cost}
}

func (h Bcrypt) Hash(s string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(s), h.cost)
	if err != nil {
		return "", err
	}
	return string(hashed), err
}

func (h Bcrypt) Equals(hashed string, s string) bool {
	eq := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(s))
	return eq == nil
}
