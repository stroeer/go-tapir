# go-tapir

![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/stroeer/go-tapir?color=%23f653a6&label=Release&style=flat-square)

This repository contains the generated Go packages for the **T**-online **API** **R**epository ([tapir](https://github.com/stroeer/tapir)) protocol buffer types, and the generated [gRPC](http://grpc.io/) code necessary for interacting with the t-online APIs.

The source proto files can be found in [tapir](https://github.com/stroeer/tapir/tree/main/stroeer).

## example

```go
package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"

	"github.com/stroeer/go-tapir/core/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	secure     = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	serverAddr = flag.String("s", "localhost:10000", "The server address in the format of host:port")
	id         = flag.Int64("id", 1, "The article id to request using gRPC")
)

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	if *secure {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := core.NewArticleServiceClient(conn)
	resp, err := client.GetArticle(context.TODO(), &core.GetArticleRequest{Id: *id})
	if err != nil {
		log.Fatalf("gRPC error: %v", err)
	}

	fmt.Printf("gRPC response: %v\n", resp.Id)
}
```