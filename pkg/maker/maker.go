package maker

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"encoding/base32"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"golang.org/x/crypto/sha3"
	"log"
	"strconv"
)

type (
	Engine struct {
	}

	Generator interface {
		GenerateHashValue(
			secretKey string,
			uniqueID string,
			bitLen int,
		) (string, error)
		GenerateOTPCode(
			secret string,
			counter uint64,
		) (int, error)
		EncryptMessage(key, data []byte) ([]byte, error)
	}
)

func DefaultMaker() Generator {
	return Engine{}
}

func (dc Engine) GenerateHashValue(
	secretKey string,
	uniqueID string,
	bitLen int,
) (string, error) {
	secretByte, err := base32.StdEncoding.DecodeString(secretKey)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("StdEncoding.DecodeString: %w", err)
	}
	hash := hmac.New(sha3.New224, secretByte)
	_, err = hash.Write([]byte(uniqueID))
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("hash.Write: %w", err)
	}
	hmacBytes := hash.Sum(nil)

	if bitLen > 1 {
		return hex.EncodeToString(hmacBytes[:bitLen]), nil
	}

	return hex.EncodeToString(hmacBytes), nil
}

func (dc Engine) GenerateOTPCode(
	secret string,
	counter uint64,
) (int, error) {
	counterByte := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		counterByte[i] = byte(counter & 0xff)
		counter >>= 8
	}

	hmacBytes, err := dc.GenerateHashValue(secret, string(counterByte), 0)
	// "Dynamic truncation" in RFC 4226
	// http://tools.ietf.org/html/rfc4226#section-5.4
	offset := hmacBytes[len(hmacBytes)-1] & 0xf
	code := (int(hmacBytes[offset])&0x7f)<<24 |
		(int(hmacBytes[offset+1])&0xff)<<16 |
		(int(hmacBytes[offset+2])&0xff)<<8 |
		(int(hmacBytes[offset+3]) & 0xff)
	code = code % 1000000

	// padding the non 6-digits otp with zero value
	f := fmt.Sprintf("%%0%dd", 6)
	codeStr := fmt.Sprintf(f, code)
	newCode, err := strconv.ParseInt(codeStr, 10, 64)
	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf(" strconv.ParseInt newcode: %w", err)
	}

	return int(newCode), nil
}

func (d Engine) DeriveKey(password, salt []byte) ([]byte, []byte, error) {
	if salt == nil {
		salt = make([]byte, 32)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}

	key, err := scrypt.Key(password, salt, 32768, 8, 1, 32)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	return key, salt, nil
}

func (d Engine) EncryptMessage(key, data []byte) ([]byte, error) {
	key, salt, err := d.DeriveKey(key, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	ciphertext = append(ciphertext, salt...)

	return ciphertext, nil
}
