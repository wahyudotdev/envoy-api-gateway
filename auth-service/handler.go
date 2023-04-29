package main

import (
	"auth-service/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"strings"
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

func (r AuthHandler) VerifyToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if strings.Contains(token, "Bearer") {
			t := strings.Split(token, " ")[1]
			d, err := helper.ParseJwt(t)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(Response{
					Error: err.Error(),
				})
			}
			return c.JSON(Response{
				Message: "success",
				Data:    d,
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Error: "invalid token",
		})
	}
}

func init() {
	pass, _ := helper.HashPassword("admin123")
	user := User{
		Id:       uuid.New().String(),
		Name:     "Admin",
		Email:    "admin@gmail.com",
		Password: pass,
	}
	db = append(db, user)
}
