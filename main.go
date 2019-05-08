package main

import (
	"fmt"

	"github.com/ory/hydra/sdk/go/hydra/client"
)

func main() {
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{})
	fmt.Printf("c=%+v\n", c)
}
