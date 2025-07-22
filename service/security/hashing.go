package security

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

type HashedPassword struct {
	Salt string
	Hash string
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


func GenerateHash(password string) (HashedPassword, error) {
	peppered := password + pepper
	salt, err := generateSalt(int(p.saltLength))

	if err != nil {
		return HashedPassword{}, err
	}

	hash := argon2.IDKey([]byte(peppered), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	return HashedPassword{
		Salt: base64.RawStdEncoding.EncodeToString(salt),
		Hash: base64.RawStdEncoding.EncodeToString(hash),
	}, nil
}

func CompareHash(hashed HashedPassword, password string) bool {
	salt, err := base64.RawStdEncoding.DecodeString(hashed.Salt)
	if err != nil {
		return false
	}

	peppered := password + pepper
	hash := argon2.IDKey([]byte(peppered), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	return base64.RawStdEncoding.EncodeToString(hash) == hashed.Hash
}

func generateSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

