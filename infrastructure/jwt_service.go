package infrastructure

import (
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var JwtSecret = []byte("my_secret_key_but_will_be_replaced_by_env_later")

func GenerateJWT(userID primitive.ObjectID) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID.Hex(),
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(JwtSecret)
}
