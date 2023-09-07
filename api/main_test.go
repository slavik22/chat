package api

import (
	db "github.com/slavik22/chat/db/sqlc"
	"github.com/slavik22/chat/util"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		SecretKey:           util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
