package main

import (
	"auth-service/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthHandler struct {
}

var db []User

func NewAuthHandler() AuthHandler {
	return AuthHandler{}
}

func (r AuthHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var reqBody User

		if err := c.BodyParser(&reqBody); err != nil {
			return err
		}

		for _, d := range db {
			if reqBody.Email == d.Email {
				err := helper.CheckPassword(reqBody.Password, d.Password)
				if err != nil {
					return c.Status(fiber.StatusUnauthorized).JSON(Response{
						Error: "invalid username or password",
					})
				}
				token, err := helper.SignJwt(d.Id, d.Roles)
				if err != nil {
					return c.Status(fiber.StatusUnauthorized).JSON(Response{
						Error: err.Error(),
					})
				}
				return c.JSON(Response{
					Message: "success",
					Data:    d,
					Token:   token,
				})
			}
		}

		return c.Status(fiber.StatusUnauthorized).JSON(Response{
			Error: "invalid username or password",
		})
	}
}

func init() {
	uPass1, _ := helper.HashPassword("user1")
	u1 := User{
		Id:       uuid.New().String(),
		Name:     "User 1",
		Email:    "user1@gmail.com",
		Password: uPass1,
		Roles:    []string{"book:read", "cart:read", "cart:write"},
	}
	db = append(db, u1)

	uPass2, _ := helper.HashPassword("user2")
	u2 := User{
		Id:       uuid.New().String(),
		Name:     "User 2",
		Email:    "user2@gmail.com",
		Password: uPass2,
		Roles:    []string{"book:read", "book:write", "cart:read", "cart:write"},
	}
	db = append(db, u2)
}
