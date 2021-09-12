package jwt_go

import (
	"encoding/base64"
	"encoding/hex"
	"time"

	"github.com/dgrijalva/jwt-go"

	"ultimate-go/awesome-go/mapstructure"
	"ultimate-go/build-in-package/crypto/encrypter_decrypter_cipher"
)

var (
	JWTSigningKey   = "jwt.key"
	JWTDesKey, _    = hex.DecodeString("6368616e676520746869732070617373")
	TokenExpireTime = time.Hour * 720
)

type JWTMe struct {
	JWTSigningKey string
	SigningMethod jwt.SigningMethod
	JWTDesKey     string
	Claims        jwt.Claims
}
type CustomClaims struct {
	UserName string `json:"user_name"`
	UserCode string `json:"user_code"`
	Phone    string `json:"phone"`
	jwt.StandardClaims
}
type UserInfo struct {
	UserName string `json:"user_name"`
	UserCode string `json:"user_code"`
	Phone    string `json:"phone"`
}

func NewJWTMe(signingKey string, signingMethod jwt.SigningMethod, desKey string) *JWTMe {
	return &JWTMe{
		JWTSigningKey: signingKey,
		SigningMethod: signingMethod,
		JWTDesKey:     desKey,
		Claims:        nil,
	}
}
func (j *JWTMe) SetClaims(claims jwt.Claims) {
	j.Claims = claims
	return
}
func (j *JWTMe) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(j.SigningMethod, j.Claims)
	return token.SignedString([]byte(j.JWTSigningKey))
}
func (j *JWTMe) ParseAndValidate(tokenString string, claims jwt.Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.JWTSigningKey, nil
	})
}
func (j *JWTMe) Encrypt(token string) (string, error) {
	ciphertext := encrypter_decrypter_cipher.CBCEncrypt([]byte(token))
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
func (j *JWTMe) Decrypt(token string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}
	plaintext := encrypter_decrypter_cipher.CBCDecrypt(ciphertext)
	return plaintext, nil
}
func (j JWTMe) MapStructureUserInfo(token *jwt.Token) *UserInfo {
	userInfo := &UserInfo{}
	mapstructure.MapStructureJsonTagDecode(token.Claims, userInfo)
	return userInfo
}
