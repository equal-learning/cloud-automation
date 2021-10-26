/*
// Struct tag :: Metadata for fields of a struct, usefurl for example to serialization information among others.
Struct tags are small pieces of metadata attached to fields of a struct that provide instructions to other Go code that works with the struct.

*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	Name      string    `json:"name"`
	roles     []string  `json:"preferredFish,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

func main() {
	u := &User{
		Name:      "",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
