package jwtsj

import (
	"github.com/shijiu-xf/go-base/comm/uuidsj"
	"github.com/shijiu-xf/go-base/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Token struct {
	Token     string
	SessionId string
	Expire    int64
	UserUid   string
	Platform  string
}

type MyCustomClaims struct {
	Platform  string
	SessionId string
	UserUid   string
	jwt.RegisteredClaims
}

type JwtInfoMd *config.JwtInfo

var (
	globalJwtCfg JwtInfoMd
)

// SetGlobalJwtCfg 通过此方法设置jwt的全局配置
func SetGlobalJwtCfg(jc JwtInfoMd) {
	globalJwtCfg = jc
}

// GetGlobalJwtCfg 通过此方法获取jwt的全局配置
func GetGlobalJwtCfg() JwtInfoMd {
	return globalJwtCfg
}

// CreateCustomToken 创建自定义的 Token
// userUid 用户的UUID
// platform 请求的平台
// access 是否将token的SessionId保存到Redis
// expireSe 有效时长（秒）
// subject token的主题
// audience token的观众（token可以使用的平台数组）
func CreateCustomToken(jwtInfoMd JwtInfoMd, userUid string) (newToken *Token, err error) {
	myClaims := MyCustomClaims{
		Platform:  jwtInfoMd.Platform,
		SessionId: uuidsj.NewUUIDv4MD5Str(),
		UserUid:   userUid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtInfoMd.Expires) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    jwtInfoMd.Issuer,
			Subject:   jwtInfoMd.Subject,
			ID:        uuidsj.NewUUIDv4Str(),
			Audience:  jwtInfoMd.Audience,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	tokenStr, err := token.SignedString([]byte(jwtInfoMd.Key))
	newToken.Platform = jwtInfoMd.Platform
	newToken.Token = tokenStr
	newToken.Expire = myClaims.ExpiresAt.Unix()
	newToken.UserUid = userUid
	newToken.SessionId = myClaims.SessionId
	return newToken, err
}

// CreateConfigToken 根据配置文件创建一个新的 Token
// userUid 用户的id
func CreateConfigToken(userUid string) (newToken *Token, err error) {
	myClaims := MyCustomClaims{
		Platform:  globalJwtCfg.Platform,
		SessionId: uuidsj.NewUUIDv4MD5Str(),
		UserUid:   userUid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(globalJwtCfg.Expires) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    globalJwtCfg.Issuer,
			Subject:   globalJwtCfg.Subject,
			ID:        uuidsj.NewUUIDv4Str(),
			Audience:  globalJwtCfg.Audience,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	tokenStr, err := token.SignedString([]byte(globalJwtCfg.Key))
	newToken = new(Token)
	newToken.Platform = myClaims.Platform
	newToken.Token = tokenStr
	newToken.Expire = myClaims.ExpiresAt.Unix()
	newToken.UserUid = myClaims.UserUid
	newToken.SessionId = myClaims.SessionId
	return newToken, err
}

// ParseJwtByGlKey 通过全局jwt配置的key解析jwt
func ParseJwtByGlKey(tokenStr string) (*jwt.Token, error) {
	var key string
	key = globalJwtCfg.Key
	parseTk, err := jwt.Parse(tokenStr, func(j *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	return parseTk, err
}

// ParseJwtByCmKey 通过传入的key解析jwt
func ParseJwtByCmKey(tokenStr string, key string) (*jwt.Token, error) {
	parseTk, err := jwt.Parse(tokenStr, func(j *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	return parseTk, err
}
