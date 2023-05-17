package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/HengMrZ/chat_azure/internal/api"
	"github.com/HengMrZ/chat_azure/internal/config"
	"github.com/HengMrZ/chat_azure/internal/models"
	"github.com/HengMrZ/chat_azure/internal/pkg"
	"github.com/sirupsen/logrus"
)

func InitAdminUser() error {
	_, err := models.QueryUserByName(models.GlobalDB, "root")
	if err != nil {
		sqlStmt := `
		CREATE TABLE "users" (
			"id" integer,
			"username" text NOT NULL UNIQUE,
			"token" text NOT NULL UNIQUE,
			"count" integer NOT NULL DEFAULT 0,
			"status" integer NOT NULL DEFAULT 1,
			"create_time" datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
			"update_time" datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY ("id")
		);
	`
		if _, err := models.GlobalDB.Exec(sqlStmt); err != nil {
			panic(err)
		}

		initAdminUser := "root"
		if v, exist := os.LookupEnv("INIT_ADMIN_USER"); exist {
			initAdminUser = v
		}

		initAdminToken := pkg.RndStr(16)
		if v, exist := os.LookupEnv("INIT_ADMIN_TOKEN"); exist {
			initAdminToken = v
		}
		err := models.AddUser(models.GlobalDB, initAdminUser, initAdminToken, 2)
		if err != nil {
			return err
		}
		logrus.Infof("init rootuser, user[%v], token[%v]", initAdminUser, initAdminToken)
	}
	return nil
}

func InitUsers() error {
	for _, it := range config.GlobalCfg.InitUsers {
		if user, err := models.QueryUserByName(models.GlobalDB, it.Username); err == nil {
			if user.Token != it.Token {
				if err = models.UpdateTokenByName(models.GlobalDB, it.Username, it.Token); err != nil {
					return err
				}
				logrus.Infof("User[%s]'s Token[%s] Updated!", it.Username, it.Token)
			} else {
				logrus.Infof("User[%s]'s Token[%s] Checked OK!", it.Username, it.Token)
			}
		} else {
			if err = models.AddUser(models.GlobalDB, it.Username, it.Token, 1); err != nil {
				return err
			}
			logrus.Infof("User[%s] with Token[%s] Created!", it.Username, it.Token)
		}
	}
	return nil
}

func main() {
	err := config.LoadConfig("./config.yaml")
	if err != nil {
		logrus.Fatal(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	models.InitDB(ctx)

	err = InitAdminUser()
	if err != nil {
		logrus.Fatal(err)
	}

	err = InitUsers()
	if err != nil {
		logrus.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/chat/completions", api.HandleCompletions)
	mux.HandleFunc("/v1/completions", api.HandleCompletions)
	mux.HandleFunc("/v1/models", api.HandleModels)
	mux.HandleFunc("/v1/adduser", api.AddUser)
	mux.HandleFunc("/v1/queryuser", api.QueryUser)
	mux.HandleFunc("/", api.HandleOptions)

	port := 8080
	if v, exist := os.LookupEnv("LISTEN_PORT"); exist {
		port, _ = strconv.Atoi(v)
	}
	logrus.Infof("svc run on port [:%v]", port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", port), loggingMiddleware(mux))
	if err != nil {
		logrus.Fatal(err)
	}
	cancel()
	time.Sleep(time.Second)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 在处理请求之前记录
		logrus.Infof("[%s] %s %s", time.Now().Format(time.RFC1123), r.Method, r.URL.Path)

		// 处理请求
		next.ServeHTTP(w, r)

		// 在处理请求之后记录
		logrus.Infof("[%s] Request handled: %s %s", time.Now().Format(time.RFC1123), r.Method, r.URL.Path)
	})
}
