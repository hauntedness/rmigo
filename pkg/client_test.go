package pkg

import (
	"net/http"
	"testing"
)

func TestBaidu(t *testing.T) {
	LoadConfig()
	client := NewRTCClient()
	client.Request(http.MethodGet, "https://www.baidu.com", nil)
}
