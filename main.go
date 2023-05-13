package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/HengMrZ/chat_azure/internal/api"
	"github.com/HengMrZ/chat_azure/internal/config"
	"github.com/sirupsen/logrus"
)

func main() {
	err := config.LoadConfig("./config.yaml")
	if err != nil {
		logrus.Fatal(err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/chat/completions", api.HandleCompletions)
	mux.HandleFunc("/v1/completions", api.HandleCompletions)
	mux.HandleFunc("/v1/models", api.HandleModels)

	port := 3389
	logrus.Infof("svc run on port:%v", port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", port), loggingMiddleware(mux))
	if err != nil {
		logrus.Fatal(err)
	}
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
