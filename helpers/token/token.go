package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)


// jwt签名结构
type JWT struct {
	SigningKey []byte
}

// 定义一些常量
var (
	TokenExpired     error  = errors.New("Token 已经过期")
	TokenNotValidYet error  = errors.New("Token 尚未激活")
	TokenMalformed   error  = errors.New("Token 格式错误")
	TokenInvalid     error  = errors.New("Token 无法解析")
	SignKey          string = "82040620FEFAC4511FC65000ADAB0F77"
)

// 载荷，加一些系统需要的信息
type CustomClaims struct {
	ID        string `json:"id"`
	Account   string `json:"account"`
	Nickname  string `json:"nickname"`
	RoleKey   string `json:"roke-key"`
	Organize  string `json:"organize"`
	jwt.StandardClaims
}

// 新建一个 jwt 实例
func NewJWT() *JWT {
	return &JWT{[]byte(GetSignKey())}
}

// 获取 signKey
func GetSignKey() string {
	return SignKey
}

func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// 生成 tokenConfig
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 tokenConfig
func (j *JWT) ResolveToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

// 生成令牌
func GenerateToken(infos map[string] string) string {
	j := NewJWT()
	claims := CustomClaims{
		ID:    infos["id"],
		Account:  infos["account"],
		Nickname: infos["nickname"],
		RoleKey:infos["roke-key"],
		Organize:infos["organize"],
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() + 0),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 100 * 3600), // 过期时间
			Issuer:    "bsit",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		return ""
	}else {
		return token
	}
}

