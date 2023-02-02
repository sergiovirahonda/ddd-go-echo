package config

import echojwt "github.com/labstack/echo-jwt/v4"

var AuthValidation = echojwt.WithConfig(echojwt.Config{
	SigningKey: []byte("secret"),
})
