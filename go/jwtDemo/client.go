package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"go/jwtDemo/demo"
	"net/http"
)

func authPage(writer http.ResponseWriter) {
	token, err := demo.generateJWT()
	if err != nil {
		return
	}
	client := &http.Client{}
	request, _ := http.NewRequest("POST", "<http://localhost:8080/>", nil)
	request.Header.Set("Token", token)
	_, _ = client.Do(request)

}

func extractClaims(_ http.ResponseWriter, request *http.Request) (string, error) {
	if request.Header["Token"] != nil {
		tokenString := request.Header["Token"][0]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("there's an error with the signing method")
			}
			return demo.sampleSecretKey, nil

		})

		if err != nil {
			return "Error Parsing Token: ", err
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			username := claims["username"].(string)
			return username, nil
		}

	}
	return "unable to extract claims", nil
}
