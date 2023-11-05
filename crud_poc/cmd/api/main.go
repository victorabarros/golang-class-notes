package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	connectionString := "postgresql://posts:p0stgr3s@db:5432/posts"
	_, err := NewConnection(connectionString)
	if err != nil {
		panic(err)
	}

	g := gin.Default()
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	g.Run(":3000")
}

func NewConnection(ctx context.Context, connectionString string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cfg, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("create connection pool: %w", err)
	}

	Conn, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return Conn, nil
}
