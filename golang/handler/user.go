package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/tkhs1121/go-sns/service"
)

func Register(c *fiber.Ctx) error {
	link := c.Query("link")

	if err := AmazonLinkValidate(link); err != nil {
		fmt.Println("Validation Error")

		return err
	}

	userID, err := service.Register(link)

	if err != nil {
		fmt.Println("DB error")

		return err
	}

	token, err := GenerateJwt(userID)

	if err != nil {
		fmt.Println("Generate jwt error")

		return err
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.MaxAge = (60 * 60 * 24 * 365)

	c.Cookie(cookie)

	return c.SendStatus(200)
}
