package handler

import (
	"os"
	"regexp"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofiber/fiber/v2"
)

func AmazonLinkValidate(link string) error {

	err := validation.Validate(link,
		validation.Required,
		validation.Match(regexp.MustCompile("amazon")),
		is.URL)

	if err != nil {
		return err
	}

	return nil
}

type ClaimsWithScope struct {
	jwt.StandardClaims
	Scope string
}

// JWTを生成
func GenerateJwt(id uint) (string, error) {
	seacretKey := os.Getenv("SEACRETKEY")
	payload := ClaimsWithScope{}

	payload.Subject = strconv.Itoa(int(id))
	payload.ExpiresAt = time.Now().Add(time.Hour * 24 * 365).Unix()
	payload.Scope = ""

	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(seacretKey))
}

// JWTからユーザーIDの取得
func GetUserId(c *fiber.Ctx) (int, error) {
	seacretKey := os.Getenv("SEACRETKEY")
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(seacretKey), nil
	})

	if err != nil || !token.Valid {
		return 0, err
	}

	payload := token.Claims.(*ClaimsWithScope)

	id, _ := strconv.Atoi(payload.Subject)

	return id, nil
}
