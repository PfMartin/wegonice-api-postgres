package db

import (
	"context"
	"testing"
	"time"

	"github.com/PfMartin/wegonice-api/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomSession(t *testing.T) Session {
	createdUser := createRandomUser(t)

	arg := CreateSessionParams{
		ID:           uuid.New(),
		Email:        createdUser.Email,
		RefreshToken: util.RandomString(10),
		UserAgent:    util.RandomString(6),
		ClientIp:     util.RandomString(6),
		IsBlocked:    false,
		ExpiresAt:    time.Now(),
	}

	session, err := testQueries.CreateSession(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, session)

	require.Equal(t, session.ID, arg.ID)
	require.Equal(t, session.Email, arg.Email)
	require.Equal(t, session.RefreshToken, arg.RefreshToken)
	require.Equal(t, session.UserAgent, arg.UserAgent)
	require.Equal(t, session.ClientIp, arg.ClientIp)
	require.Equal(t, session.IsBlocked, arg.IsBlocked)
	require.WithinDuration(t, session.ExpiresAt, arg.ExpiresAt, time.Second)

	return session
}

func TestCreateSession(t *testing.T) {
	createRandomSession(t)
}

func TestGetSession(t *testing.T) {
	createdSession := createRandomSession(t)

	gotSession, err := testQueries.GetSession(context.Background(), createdSession.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotSession)

	require.Equal(t, gotSession.ID, createdSession.ID)
	require.Equal(t, gotSession.Email, createdSession.Email)
	require.Equal(t, gotSession.RefreshToken, createdSession.RefreshToken)
	require.Equal(t, gotSession.UserAgent, createdSession.UserAgent)
	require.Equal(t, gotSession.ClientIp, createdSession.ClientIp)
	require.Equal(t, gotSession.IsBlocked, createdSession.IsBlocked)
	require.WithinDuration(t, gotSession.ExpiresAt, createdSession.ExpiresAt, time.Second)
}
