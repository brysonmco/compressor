package auth

import (
	"aidanwoods.dev/go-paseto"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type Auth struct {
	PrivateKey paseto.V4AsymmetricSecretKey
	PublicKey  paseto.V4AsymmetricPublicKey
}

func NewAuth() *Auth {
	privateKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := privateKey.Public()

	return &Auth{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}

func (a *Auth) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func (a *Auth) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a *Auth) GenerateAccessToken(id int64) (string, error) {
	token := paseto.NewToken()
	now := time.Now()
	token.SetIssuedAt(now)
	token.SetNotBefore(now)
	token.SetExpiration(now.Add(time.Hour))
	token.SetSubject(strconv.FormatInt(id, 10))

	return token.V4Sign(a.PrivateKey, nil), nil
}

func (a *Auth) ValidateAccessToken(token string) (int64, error) {
	parsedToken, err := paseto.NewParser().ParseV4Public(a.PublicKey, token, nil)
	if err != nil {
		return 0, err
	}

	subject, err := parsedToken.GetSubject()
	if err != nil {
		return 0, err
	}

	id, err := strconv.ParseInt(subject, 10, 64)
	if err != nil {
		return 0, err
	}

	issued, err := parsedToken.GetIssuedAt()
	if err != nil {
		return 0, err
	}
	if issued.After(time.Now()) {
		return 0, err
	}

	return id, nil
}

func (a *Auth) GenerateRefreshToken(id int64) (string, error) {
	return "", nil
}

func (a *Auth) ValidateRefreshToken() (string, error) {
	return "", nil
}
