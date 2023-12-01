package token

import (
	"personal-budget/util"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	userId := uuid.New()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(userId, duration, true)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifiyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, userId, payload.UserId)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)

	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	userId := uuid.New()
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, payload, err := maker.CreateToken(userId, -time.Minute, false)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifiyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

// func TestInvalidJWTTokenAlgNone(t *testing.T) {
// 	payload, err := NewPayload(util.RandomOwner(), time.Minute)
// 	require.NoError(t, err)

// 	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
// 	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
// 	require.NoError(t, err)

// 	maker, err := NewJWTMaker(util.RandomString(32))
// 	require.NoError(t, err)

// 	payload, err = maker.VerifiyToken(token)
// 	require.Error(t, err)
// 	require.EqualError(t, err, ErrInvalidToken.Error())
// 	require.Nil(t, payload)

// }
