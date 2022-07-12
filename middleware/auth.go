package middleware

import (
  "fmt"
  "strings"
  "time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	// "github.com/gofiber/fiber/v2"
)

var Salt = []byte("MaiSanGaSuki")

type UserDataClaims struct {
  Uname string `json:"uname"`
  jwt.StandardClaims
}

type Authorization struct {
  Authorization string `reqHeader:"Authorization"`
}

func CheckUser(name string, password string) bool {
  return (name == "smc" && password == "passwd@123")
}

func GenJWT(uname string) (string, error) {
  claims := UserDataClaims {
    uname,
    jwt.StandardClaims {
      ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(Salt)
}

func ValidateJWT(c *fiber.Ctx) error {
  bearer := new(Authorization)

  if err := c.ReqHeaderParser(bearer); err != nil {
    return fiber.ErrBadRequest
  }
  var tokenVal string

  if val := strings.Split(bearer.Authorization, " "); len(val) < 2 {
    fmt.Println(len(val))
    return fiber.ErrForbidden
  } else {
    tokenVal = val[1]
  }

  token, err := jwt.ParseWithClaims(
    tokenVal, 
    &UserDataClaims{}, 
    func (*jwt.Token) (interface{}, error) {
      return Salt, nil
  })


  fmt.Println(token)
  if claims, ok := token.Claims.(*UserDataClaims); !(ok && token.Valid && claims.Uname == "smc") {
    return fiber.ErrForbidden
  }

  if err != nil {
    fmt.Println(err)
    return fiber.ErrForbidden
  }

  return c.Next()

}
