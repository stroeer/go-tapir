package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"

	core "buf.build/gen/go/stroeer/tapir/grpc/go/stroeer/core/v1/corev1grpc"
	article "buf.build/gen/go/stroeer/tapir/protocolbuffers/go/stroeer/core/v1"
)

var (
	secure     = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	serverAddr = flag.String("s", "localhost:8080", "The server address in the format of host:port")
	id         = flag.Int64("id", 1, "The article id to request using gRPC")
	auth       = flag.String("auth", "", "Authorization token")
)

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	if *secure {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	if *auth != "" {
		md := metadata.Pairs("Authorization", *auth)
		ctx = metadata.NewOutgoingContext(ctx, md)
	}

	client := core.NewArticleServiceClient(conn)
	resp, err := client.GetArticle(ctx, &article.GetArticleRequest{Id: *id})
	if err != nil {
		log.Fatalf("gRPC error: %v", err)
	}

	fmt.Printf("gRPC response: %v\n", resp)
}
