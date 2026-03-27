package hasher

import "golang.org/x/crypto/bcrypt"

type BcryptHasher struct{}

func New() *BcryptHasher {
	return &BcryptHasher{}
}

func (h *BcryptHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (h *BcryptHasher) Compare(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

type Hasher interface {
	Hash(password string) (string, error)
	Compare(hash, password string) bool
}
