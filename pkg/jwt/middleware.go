package jwt

import (
	"backend/config"
	"backend/pkg/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

// Auth is the middleware function for JWT authentication with role-based authorization
func Auth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var tokenString string
		authorization := ctx.Get("Authorization")

		// Check if the "Authorization" header contains the token
		if strings.HasPrefix(authorization, "Bearer ") {
			tokenString = strings.TrimPrefix(authorization, "Bearer ")
		} else if ctx.Cookies("token") != "" {
			// If the token is not in the "Authorization" header, check if it's in the cookies
			tokenString = ctx.Cookies("token")
		}

		if tokenString == "" {
			// Return an error response in JSON format when the token is missing
			return ctx.JSON(utils.ErrorResponse{
				Code:   fiber.StatusUnauthorized,
				Status: utils.StatusUnauthorized,
				Errors: "Token Null | Unauthorized",
			})
		}

		secretConfig := config.Config.GetString("JWT.TOKEN")

		// Parse and validate the token
		tokenByte, err := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
			if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
			}

			return []byte(secretConfig), nil
		})
		if err != nil {
			return ctx.JSON(utils.ErrorResponse{
				Code:   fiber.StatusUnauthorized,
				Status: utils.StatusUnauthorized,
				Errors: "Invalid Token Claim | Unauthorized",
			})
		}

		// Check if the token is valid
		_, ok := tokenByte.Claims.(jwt.MapClaims)
		if !ok || !tokenByte.Valid {
			return ctx.JSON(utils.ErrorResponse{
				Code:   fiber.StatusUnauthorized,
				Status: utils.StatusUnauthorized,
				Errors: "Invalid Token Claim | Unauthorized",
			})
		}

		// Token is valid and the user has the required role, continue to the next handler
		return ctx.Next()
	}
}
