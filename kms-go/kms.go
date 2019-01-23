package kms

import "golang.org/x/net/websocket"

// KmsClient Export client
type KmsClient struct{}

// New  Create a kms client object
func New() *KmsClient {
	return &KmsClient{}
}

func Exec() {
	websocket.Dial()
}
