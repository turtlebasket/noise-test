package main

import (
	"fmt"

	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/gossip"
	"github.com/perlin-network/noise/kademlia"
	"github.com/turtlebasket/noise-test/types"
	"go.uber.org/zap"
)

func main() {
	// Create logger
	logger, _ := zap.NewDevelopment(zap.AddStacktrace(zap.PanicLevel))

	defer logger.Sync()

	node, err := noise.NewNode(noise.WithNodeLogger(logger), noise.WithNodeBindPort(9000))

	if err != nil {
		panic(err)
	}

	defer node.Close()

	overlay := kademlia.New()
	hub := gossip.New(overlay)

	// Start the network.
	node.Bind(
		overlay.Protocol(),
		hub.Protocol(),
	)

	node.RegisterMessage(types.PMessage{}, types.UnmarshalPMessage)
	node.RegisterMessage(types.PMessageResponse{}, types.UnmarshalPMessageResponse)

	node.Handle(func(ctx noise.HandlerContext) error {

		obj, err := ctx.DecodeMessage()
		if err != nil {
			return nil
		}

		msg, ok := obj.(types.PMessage)
		if !ok {
			return nil
		}

		fmt.Printf("Got a message from Bob: '%s'\n", fmt.Sprint(msg))

		return ctx.SendMessage(types.PMessageResponse{
			To:     msg.To,
			Amount: msg.Amount,
			Status: types.Verified,
		})
	})

	if err := node.Listen(); err != nil {
		panic(err)
	}
}
