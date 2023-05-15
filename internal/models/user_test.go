package models

import (
	"context"
	"testing"

	_ "github.com/glebarez/go-sqlite"
	"github.com/sirupsen/logrus"
)

func TestAddUser(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	InitDB(ctx)
	err := AddUser(GlobalDB, "root", "123456", 2)
	if err != nil {
		t.Error(err)
		return
	}
	user, err := QueryUserByToken(GlobalDB, "123456")
	if err != nil {
		t.Error(err)
		return
	}
	logrus.Infof("user:%v", *user)
}
