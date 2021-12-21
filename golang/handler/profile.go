package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tkhs1121/go-sns/service"
	"github.com/tkhs1121/go-sns/util"
)

func GetRandProfile(c *fiber.Ctx) error {
	id, err := util.GetUserId(c)

	if err != nil {
		return err
	}

	randID, err := service.GetRandProfile(uint(id))

	if err != nil {
		return err
	}

	fmt.Println(randID)

	return c.SendStatus(200)
}

func UpdateRecommendation(c *fiber.Ctx) error {

	id, err := util.GetUserId(c)

	if err != nil {
		return err
	}

	link := c.Query("link")

	if err := util.AmazonLinkValidate(link); err != nil {
		fmt.Println("Validation Error")

		return err
	}

	if err := service.UpdateRecommendation(uint(id), link); err != nil {
		fmt.Println("Update Recomendation Error")

		return err
	}

	return c.SendStatus(200)
}
