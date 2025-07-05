package service

import (
	"crypto/rand"
	"encoding/base64"
	"trust-credit-back/environment"

	"golang.org/x/crypto/argon2"
)

type params struct {
    memory      uint32
    iterations  uint32
    parallelism uint8
    saltLength  uint32
    keyLength   uint32
}

var ( 
	p = &params {
	memory:      64 * 1024,
    iterations:  3,
    parallelism: 2,
    saltLength:  16,
    keyLength:   32,
	}

	pepper = environment.GetVariable("Pepper")
)


func GenerateHash(password string) string {
	peppered := password + pepper
	salt, err := generateSalt(int(p.saltLength))

	if err != nil {
		return ""
	}

	hash := argon2.IDKey([]byte(peppered), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	return base64.RawStdEncoding.EncodeToString(salt) + "&" + base64.RawStdEncoding.EncodeToString(hash)
}

func generateSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

