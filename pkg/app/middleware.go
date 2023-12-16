package app

import (
	"net/http"
	"strings"

	"github.com/DanilaNik/IU5_RIP2023/internal/service/role"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const prefix = "Bearer"

func (a *Application) RoleMiddleware(allowedRoles ...role.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractTokenFromHeader(c.Request)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("Key1234"), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		userID, ok := claims["userID"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set("UserID", int(userID))

		userRole, ok := claims["role"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if !isRoleAllowed(string(userRole), allowedRoles) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "отказано в доступе"})
			return
		}

		c.Next()
	}
}

func extractTokenFromHeader(req *http.Request) string {
	bearerToken := req.Header.Get("Authorization")
	if bearerToken == "" {
		return ""
	}

	if strings.Split(bearerToken, " ")[0] != prefix {
		return ""
	}

	return strings.Split(bearerToken, " ")[1]
}

func isRoleAllowed(userRole string, allowedRoles []role.Role) bool {
	for _, allowedRole := range allowedRoles {
		if userRole == string(allowedRole) {
			return true
		}
	}
	return false
}
