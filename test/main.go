package main

import (
	"fmt"
)

type TransInfo struct {
	ID int32
}

type Fragment interface{
	Exec(transInfo *TransInfo) error
}

type GetPodAction struct{}

func (g GetPodAction) Exec(transInfo *TransInfo) error {
	return nil
}

func main() {
	var a Fragment = new(GetPodAction)
	fmt.Println(a)
}




