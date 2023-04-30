package main

import (
	"auth-service/helper"
	"fmt"
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
				token, err := helper.SignJwt(d.Id)
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
	for i := 1; i < 5; i++ {
		pass, _ := helper.HashPassword(fmt.Sprintf("user%d", i))
		user := User{
			Id:       uuid.New().String(),
			Name:     fmt.Sprintf("User %d", i),
			Email:    fmt.Sprintf("user%d@gmail.com", i),
			Password: pass,
		}
		db = append(db, user)
	}
}
