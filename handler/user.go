package handler

import (
	"log"

	"github.com/danilaudza/go-fiber/database"
	"github.com/danilaudza/go-fiber/model"
	"github.com/danilaudza/go-fiber/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []model.User

	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	// Validasi request
	validate := validator.New()
	errValidate := validate.Struct(user)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate,
		})
	}

	newUser := model.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phonee:  user.Phonee,
	}

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user model.User
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// userResponse := response.UserResponse{
	// 	ID:        user.ID,
	// 	Name:      user.Name,
	// 	Email:     user.Email,
	// 	Address:   user.Address,
	// 	Phonee:    user.Phonee,
	// 	CreatedAt: user.CreatedAt,
	// 	UpdatedAt: user.UpdatedAt,
	// }

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user model.User

	userId := ctx.Params("id")
	// Check Available User
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// Update User
	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	user.Address = userRequest.Address
	user.Phonee = userRequest.Phonee
	errUpdate := database.DB.Save(&user).Error

	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal server Error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerUpdateEmail(ctx *fiber.Ctx) error {
	userRequest := new(request.UserEmailRequest)

	if err := ctx.BodyParser(userRequest); err != nil {

		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user model.User
	var isEmailUserExist model.User

	userId := ctx.Params("id")

	err := database.DB.First(&user, "id = ?", userId).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errEmail := database.DB.First(&isEmailUserExist, "email = ?", userRequest.Email).Error

	if errEmail == nil {
		return ctx.Status(402).JSON(fiber.Map{
			"message": "email used",
			"error":   errEmail,
		})
	}

	user.Email = userRequest.Email

	errUpdate := database.DB.Save(&user).Error

	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user model.User

	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errDelete := database.DB.Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal Server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "user was deleted",
	})
}
