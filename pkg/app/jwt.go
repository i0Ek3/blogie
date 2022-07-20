package app

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/blogie/util"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.RegisteredClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// GenerateToken generates the token by given claims and HS256
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		// NOTES: change StandardClaims to RegisteredClaims in jwt v4
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expireTime},
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	// NOTES: NewWithClaims support these three algorithms: HS256/HS384/HS512
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

// ParseToken parses into a tokenClaims and validates that token according Claims
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
