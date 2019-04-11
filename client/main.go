package main

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/payfazz/test-grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:16969", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)
	g := gin.Default()

	g.GET("add/:a/:b", func(ctx *gin.Context) {
		a, _ := strconv.ParseInt(ctx.Param("a"), 10, 64)
		b, _ := strconv.ParseInt(ctx.Param("b"), 10, 64)
		resp, err := client.Add(ctx, &proto.Request{A: a, B: b})
		if err != nil {
			panic(err)
		}

		log.Println("client result for add ", resp)
	})

	g.GET("multiply/:a/:b", func(ctx *gin.Context) {
		a, _ := strconv.ParseInt(ctx.Param("a"), 10, 64)
		b, _ := strconv.ParseInt(ctx.Param("b"), 10, 64)
		resp, err := client.Multiply(ctx, &proto.Request{A: a, B: b})
		if err != nil {
			panic(err)
		}

		log.Println("client result for multiply ", resp)
	})

	log.Println("Starting to serve the client")
	if err := g.Run(":18080"); err != nil {
		panic(err)
	}
}
