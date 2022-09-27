package controllers

import (
	"context"
	"seqolah-qu/entities/dto"
	"seqolah-qu/services"
	"seqolah-qu/types"
)

type AuthController struct {
	authService *services.AuthServiceImpl
}

func NewAuthController(authService *services.AuthServiceImpl) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Register(data types.Request, context context.Context) (types.Response, error) {
	var body dto.RegisterDto
	data.Body(&body)

	resp := c.authService.Register(body)

	return types.Response{
		Message: "Register success",
		Data:    resp,
	}, nil
}

func (c *AuthController) Login(data types.Request, context context.Context) (types.Response, error) {
	var body dto.LoginDto
	data.Body(&body)

	resp := c.authService.Login(body)

	return types.Response{
		Message: "Login success",
		Data:    resp,
	}, nil
}
