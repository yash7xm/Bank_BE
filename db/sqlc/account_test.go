package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner: "yash",
		Balance: 100,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

}