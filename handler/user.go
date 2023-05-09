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
}

func newUserPresentationRepository() IUserPresentation {
	return &controller.UserRepository{}
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
		Log:        "All Data decrypted successfully",
		Model:      result,
	})
}
