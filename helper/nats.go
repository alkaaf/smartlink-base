package helper

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nats-io/nats.go"
	"os"
)

var conn *nats.Conn

func NatsConn() *nats.Conn {
	godotenv.Load()
	if conn == nil {
		servers := os.Getenv("NATS_HOSTS")
		conn, err = nats.Connect(servers)
		fmt.Println(err)
	}
	return conn
}
