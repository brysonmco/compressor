package auth

import (
	"aidanwoods.dev/go-paseto"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type Auth struct {
	PrivateKey paseto.V4AsymmetricSecretKey
	PublicKey  paseto.V4AsymmetricPublicKey
}

func NewAuth(
	privateKeyHex string,
) (*Auth, error) {
	privateKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(privateKeyHex)
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.Public()

	return &Auth{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}, nil
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

func (a *Auth) GenerateAccessToken(id int64, role string) (string, error) {
	token := paseto.NewToken()
	now := time.Now()
	token.SetIssuedAt(now)
	token.SetNotBefore(now)
	token.SetExpiration(now.Add(time.Hour))
	token.SetSubject(strconv.FormatInt(id, 10))
	token.Set("role", role)

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
		return 0, fmt.Errorf("token is not yet valid")
	}

	exp, err := parsedToken.GetExpiration()
	if err != nil {
		return 0, err
	}
	if exp.Before(time.Now()) {
		return 0, fmt.Errorf("expired_token")
	}

	return id, nil
}

func (a *Auth) GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (a *Auth) HashRefreshToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}
