package main

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

func main() {
	fmt.Println(timestamp.Timestamp{Seconds: time.Now().Unix()})
}
