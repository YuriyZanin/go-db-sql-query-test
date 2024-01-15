package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.Nil(t, err, "db open error")
	clientID := 1

	// напиши тест здесь
	cl, err := selectClient(db, clientID)
	require.Nil(t, err, "db query error")
	require.NotNil(t, cl, "expected not empty client")
	assert.Equal(t, clientID, cl.ID, "expectd id %d got %d", clientID, cl.ID)
	assert.NotEmpty(t, cl.FIO, "expected FIO not empty")
	assert.NotEmpty(t, cl.Login, "expected login not empty")
	assert.NotEmpty(t, cl.Birthday, "expected birthday not empty")
	assert.NotEmpty(t, cl.Email, "expected email not empty")
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.Nil(t, err, "db open error")

	clientID := -1

	// напиши тест здесь
	cl, err := selectClient(db, clientID)
	require.Error(t, err, sql.ErrNoRows, "expected no rows error")
	require.NotNil(t, cl, "expected cl not nil")
	assert.Equal(t, 0, cl.ID, "expected id 0 got %s", cl.ID)
	assert.Empty(t, cl.FIO, "expected empty FIO")
	assert.Empty(t, cl.Login, "expected empty login")
	assert.Empty(t, cl.Birthday, "expected empty birthday")
	assert.Empty(t, cl.Email, "expected empty email")
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.Nil(t, err, "db open error")

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	id, err := insertClient(db, cl)
	require.Nil(t, err, "db query error")
	require.NotNil(t, id)

	fromDb, err := selectClient(db, id)
	require.Nil(t, err, "db query error")
	require.NotNil(t, fromDb, "expected not nil")
	assert.Equal(t, cl.FIO, fromDb.FIO,
		"expected fio %s got %s", cl.FIO, fromDb.FIO)
	assert.Equal(t, cl.Login, fromDb.Login,
		"expected login %s got %s", cl.Login, fromDb.Login)
	assert.Equal(t, cl.Birthday, fromDb.Birthday,
		"expected birthday %s got %s", cl.Birthday, fromDb.Birthday)
	assert.Equal(t, cl.Email, fromDb.Email,
		"expected email %s got %s", cl.Email, fromDb.Email)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.Nil(t, err, "db open error")

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	id, err := insertClient(db, cl)
	require.Nil(t, err, "db query error")
	require.NotNil(t, id, "expected id not nil")

	_, err = selectClient(db, id)
	require.Nil(t, err, "db query error")

	err = deleteClient(db, id)
	require.Nil(t, err, "error not expected")

	_, err = selectClient(db, id)
	require.Error(t, err, sql.ErrNoRows, "error no rows expected")
}
