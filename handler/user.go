package handler

import (
	"dependency_inversion/controller"
	api_models "dependency_inversion/models/api"

	"github.com/gofiber/fiber/v2"
	"swap-gitlab.apk-group.net/apkswap/apk_go/helper/error_helper"
	response_models "swap-gitlab.apk-group.net/apkswap/apk_go/models/api"
)

type IUserPresentation interface {
	GetUser(id string) api_models.User
	UpdateUser(api_models.User)
}

func newUserPresentationRepository() IUserPresentation {
	return controller.NewUserController()
}

func GetUser(c *fiber.Ctx) error {
	defer error_helper.Recover_without_panic("getUser", c)
	var (
		id = c.Params("id")
	)
	result := newUserPresentationRepository().GetUser(id)
	return c.Status(fiber.StatusOK).JSON(response_models.Response_model{
		Is_success: true,
		Message:    "Successfull Operation",
		Log:        "User fetched successfully",
		Model:      result,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	defer error_helper.Recover_without_panic("UpdateUser", c)
	var (
		user = api_models.User{}
	)
	c.BodyParser(&user)
	newUserPresentationRepository().UpdateUser(user)
	return c.Status(fiber.StatusOK).JSON(response_models.Response_model{
		Is_success: true,
		Message:    "Successfull Operation",
		Log:        "User updated successfully",
		Model:      nil,
	})
}
