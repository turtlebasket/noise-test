package main

import (
	"context"
	"fmt"
	"time"

	"github.com/perlin-network/noise"
	"github.com/turtlebasket/noise-test/types"
)

func main() {
	node, err := noise.NewNode(noise.WithNodeBindPort(9001))
	node.RegisterMessage(types.PMessage{}, types.UnmarshalPMessage)
	node.RegisterMessage(types.PMessageResponse{}, types.UnmarshalPMessageResponse)

	if err != nil {
		panic(err)
	}

	defer node.Close()

	pMessage := types.PMessage{
		To:     "abc",
		Amount: 0.0001,
	}

	node.Ping(context.Background(), "127.0.0.1:9000")

	if err := node.Listen(); err != nil {
		panic(err)
	}

	go func() {
		for {

			fmt.Printf("Sending %s\n", fmt.Sprint(pMessage))

			ctx, _ := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
			res, err := node.RequestMessage(ctx, "127.0.0.1:9000", pMessage)
			if err != nil {
				panic(err)
			}

			fmt.Println(fmt.Sprint(res))

			time.Sleep(time.Duration(5 * time.Second))
		}
	}()

}
