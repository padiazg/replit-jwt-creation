package main

import (
	"fmt"
  "github.com/golang-jwt/jwt/v4"
)

var (
  secretKey = "tyk-portal-jwt-signing-secret"
)

func main() {
  token, _ := GenerateJWT("$2a$10$weyvDhZ.4kksgVcPpOR9IeVvZBGS8Z1nKHvX6Gn8NBRWdVfqHSPCq", "0")
	fmt.Printf("token: %s\n", token)
 //  token, _ = GenerateJWT("dummy-api-token", "0")
	// fmt.Printf("token: %s\n", token)  
}

func GenerateJWT(userID, providerID string) (string, error) {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserID":     userID,
		"ProviderID": providerID,
		// "exp": time.Now().Add(time.Minute * 30).Unix(),
	})
  
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}