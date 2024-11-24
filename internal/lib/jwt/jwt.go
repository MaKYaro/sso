package jwt

import (
	"time"

	"github.com/MaKYaro/sso/internal/domain/models"
	"github.com/golang-jwt/jwt"
)

func NewToken(
	user models.User,
	app models.App,
	dur time.Duration,
) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(dur).Unix()
	claims["app_id"] = app.ID

	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
