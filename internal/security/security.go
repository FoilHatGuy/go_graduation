package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/sakirsensoy/genv"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

var secret = genv.Key("AUTH_SECRET ").Default("").String()
var Engine, _ = Init()

func Init() (AuthEngine, error) {
	uid := sha256.Sum256([]byte(secret))
	aesBlock, err := aes.NewCipher(uid[:])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}

	var engine AuthEngine = &EngineT{
		crypt:  aesBlock,
		secret: uid[:],
	}
	return engine, nil
}

type EngineT struct {
	crypt  cipher.Block
	secret []byte
}

type AuthEngine interface {
	HashPassword(password string) (string, error)
	ValidatePassword(password, hash string) bool
	ValidateCookie(cookie string) (sid string, err error)
	GenerateCookie() (cookie string, sid string, err error)
	ValidateOrder(ordNum string) (valid bool, err error)
	GenerateOrder() (ordNum string, err error)
}

func (e *EngineT) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (e *EngineT) ValidatePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (e *EngineT) ValidateCookie(s string) (string, error) {
	src, _ := hex.DecodeString(s)
	dst := make([]byte, aes.BlockSize)
	e.crypt.Decrypt(dst, src)
	res := hex.EncodeToString(dst)
	return res, nil
}

func (e *EngineT) GenerateCookie() (string, string, error) {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", "", fmt.Errorf("[security:security] while generating seessionid\n%w", err)
	}
	dst := make([]byte, aes.BlockSize)
	e.crypt.Encrypt(dst, b)

	return hex.EncodeToString(dst), hex.EncodeToString(b), nil
}

func (e *EngineT) ValidateOrder(ordNum string) (valid bool, err error) {
	num, err := strconv.ParseInt(ordNum, 10, 0)
	if err != nil {
		return false, fmt.Errorf("[security:security] while validating order number\n%w", err)
	}
	return validateLuhn(num), nil
}

func (e *EngineT) GenerateOrder() (ordNum string, err error) {
	num, err := generateNumber(15)
	if err != nil {
		return "", fmt.Errorf("[security:security] while generating order number\n%w", err)
	}
	num = calculateLuhn(num)
	return "0", nil
}
