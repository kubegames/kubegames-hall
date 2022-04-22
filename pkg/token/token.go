package token

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	playerModel "github.com/kubegames/kubegames-hall/app/model/player"
	"github.com/kubegames/kubegames-hall/internal/pkg/des"
	"github.com/kubegames/kubegames-hall/internal/pkg/log"
	"google.golang.org/grpc/metadata"
)

var (
	jwtkey        = []byte("games.hall.com")
	Authorization = "authorization"
	desEngine     = des.Engine([]byte("gamehall"), []byte("gamehall"))
)

//claims
type Claims struct {
	Text string
	jwt.StandardClaims
}

//getPlayer
func GetPlayerByContext(ctx context.Context) (player *playerModel.Player, token string, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Errorf("token is nil")
		return nil, "", errors.New("token is nil")
	}

	//get tokens
	tokenString := md.Get(Authorization)
	if len(tokenString) <= 0 {
		log.Errorf("token is nil")
		return nil, "", errors.New("token is nil")
	}

	//autch
	player, err = GetPlayerByToken(tokenString[0])
	if err != nil {
		return nil, "", err
	}
	return player, tokenString[0], nil
}

//get player token
func GetPlayerByToken(tokenString string) (*playerModel.Player, error) {
	claims := new(Claims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("permissions error")
	}
	//new player
	player := new(playerModel.Player)
	if err := desEngine.Decrypt(claims.Text, &player); err != nil {
		return nil, err
	}
	return player, nil
}

//new player token
func NewPlayerToken(player *playerModel.Player) (string, error) {
	cipherText, err := desEngine.Encrypt(player)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Text: cipherText,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "games",
			Subject:   "games token",
		},
	})

	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//new context by token grpc used
func NewContextByToken(ctx context.Context, token string) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.Pairs(Authorization, token))
}
