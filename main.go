package main

import (
	"context"
	"fmt"
	server "github.com/rolancia/thing"
	"os"
	"sync"
)

var (
	payloadChannelContextKey string = "payloadChan"
)

func main() {
	dispatcherCtx, cancel := context.WithCancel(context.Background())
	payloadChan := make(chan Payload)

	defer func() {
		close(payloadChan)
	}()

	ctx := context.WithValue(context.Background(), payloadChannelContextKey, payloadChan)
	svr := server.NewServer(ctx, TcpFrameConfig(), &EventHandler{}, &PostHandler{})

	go DispatcherHeader(payloadChan, dispatcherCtx)

	ip := os.Getenv("SERVER_IP")
	port := os.Getenv("SERVER_PORT")

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := server.Serve(svr, ip+":"+port, server.ErrorActionSave)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}()

	wg.Wait()
	cancel()
}
