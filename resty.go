package ocpc_media

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/gtkit/logger"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

const (
	timeout       = 10
	tryNum        = 3
	retryInterval = 1

	maxIdleConns          = 10
	idleConnTimeout       = 30 * time.Second
	tlsHandshakeTimeout   = 10 * time.Second
	responseHeaderTimeout = 20 * time.Second
)

var restyClient *resty.Client

func newResty() {
	// 设置全局的 resty 客户端
	json := jsoniter.ConfigCompatibleWithStandardLibrary // jsoniter 库

	client := resty.New()
	client.SetTimeout(timeout * time.Second)                                   // 设置超时时间为 5 秒钟
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})           //nolint:gosec // 关闭证书校验
	client.SetRetryCount(tryNum).SetRetryWaitTime(retryInterval * time.Second) // 设置最大重试次数为 3 次，重试间隔时间为 1 秒钟
	client.SetTransport(&http.Transport{
		MaxIdleConnsPerHost:   maxIdleConns,          // 对于每个主机，保持最大空闲连接数为 10
		IdleConnTimeout:       idleConnTimeout,       // 空闲连接超时时间为 30 秒
		TLSHandshakeTimeout:   tlsHandshakeTimeout,   // TLS 握手超时时间为 10 秒
		ResponseHeaderTimeout: responseHeaderTimeout, // 等待响应头的超时时间为 20 秒
	})
	client.JSONMarshal = json.Marshal
	client.JSONUnmarshal = json.Unmarshal
	if logger.Zlog() != nil {
		// 设置日志
		client.SetLogger(logger.RestyLogger())
	}

	restyClient = client
}

// Client 全局的 resty 客户端.
func Client() *resty.Client {
	return restyClient
}

// R 返回一个新的请求对象.
func R() *resty.Request {
	return restyClient.R()
}

// 使用教程: https://blog.csdn.net/qq_29799655/article/details/130831278
