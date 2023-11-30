package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

var (
	secret []byte
)

type Claims struct {
	UID     string `json:"uid"`
	Version int    `json:"version"`
	jwt.RegisteredClaims
}

func buildClaims(uid string, version int, ttl int64) Claims {
	now := time.Now()
	before := now.Add(-time.Second * 5)
	return Claims{
		UID:     uid,
		Version: version,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(ttl) * time.Second)), //Expiration time
			IssuedAt:  jwt.NewNumericDate(now),                                       //Issuing time
			NotBefore: jwt.NewNumericDate(before),                                    //Begin Effective time
		}}
}

func NewToken(uid string, version int, time int64) (string, error) {
	claims := buildClaims(uid, version, time)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(secret)
}

func getSecret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	}
}

func GetClaimFromToken(tokensString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokensString, &Claims{}, getSecret())
	if err != nil {
		return nil, err
	} else {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("invalid token")
	}
}

func ValidateToken(c *gin.Context) {
	tokenStr := c.Request.Header.Get("token")
	if tokenStr == "" {
		c.JSON(http.StatusUnauthorized, map[string]any{
			"code": http.StatusUnauthorized,
			"msg":  "token missing, please login",
		})
		c.Abort()
		return
	}
	claim, err := GetClaimFromToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]any{
			"code": http.StatusUnauthorized,
			"msg":  err.Error(),
		})
		c.Abort()
		return
	}
	c.Set("ID", claim.ID)
	c.Next()
}
