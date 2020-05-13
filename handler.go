package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func HDLR() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return func(c *gin.Context) {
		h.ContextHandler(c, c.Writer, c.Request)
	}
}
