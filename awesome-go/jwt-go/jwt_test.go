package jwt_go

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Test_JWTMe_GenerateToken(t *testing.T) {
	userInfo := &UserInfo{UserName: "betty", UserCode: "abc123", Phone: "12123431112"}
	claims := CustomClaims{
		UserName:       userInfo.UserName,
		UserCode:       userInfo.UserCode,
		Phone:          userInfo.Phone,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(TokenExpireTime).Unix()},
	}
	jwtm := NewJWTMe(JWTSigningKey, jwt.SigningMethodHS256, string(JWTDesKey[:]))
	jwtm.SetClaims(claims)
	token, _ := jwtm.GenerateToken()
	token, _ = jwtm.Encrypt(token)
	fmt.Println(token)
}

func Test_JWTMe_ParseAndValidate(t *testing.T) {
	tokenString := "w3UiIeh5lDVLZhu0NgIuVZwS7PMquZT+x7tZTASGE3Ca0nuiungFCZa+wRJ6NrB5itweaHfqLZZFcByCi3QYifjA5Jo+6qg8sEZX4qiOSP/yNGzow5MwzMeGy1CaNsBYTU/3Uzh+XU+0vVl2V2DpSwsvCIETz6prDDjvOUNkqVSwbSM+zx03NouSwAwfUqCcR+tDJOyOrtEQQgqIe0du40Z+uixFLQfLrqZf4t93qTIB/hn7yXtUdsrMYgcUp1ee"
	jwtm := NewJWTMe(JWTSigningKey, jwt.SigningMethodHS256, string(JWTDesKey[:]))
	tokenString, _ = jwtm.Decrypt(tokenString)
	token, _ := jwtm.ParseAndValidate(tokenString, new(CustomClaims))
	userInfo := jwtm.MapStructureUserInfo(token)
	fmt.Println(userInfo)
}
