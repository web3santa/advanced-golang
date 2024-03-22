package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Use(cors.Default())
	server.GET("/progress", func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "text/event-stream")
		ctx.Writer.Header().Set("Cashe-Control", "no-cache")
		ctx.Writer.Flush()

		for i := 0; i < 100; i++ {

			ctx.Writer.Write([]byte(fmt.Sprintf("id: %d\n", i)))
			ctx.Writer.Write([]byte("event: onPregress\n"))
			data, _ := json.Marshal(gin.H{
				"progressPercentage": i + 1,
			})
			ctx.Writer.Write([]byte(fmt.Sprintf("id: %d\n", data)))
			ctx.Writer.Write([]byte("\n"))
			time.Sleep(time.Second / 10)

		}

		ctx.Writer.Write([]byte("id: 100 \n"))
		ctx.Writer.Write([]byte("event: done \n"))
		ctx.Writer.Write([]byte("data: {} \n"))
		ctx.Writer.Write([]byte("\n"))
		ctx.Writer.Flush()
	})

	if err := server.Run("0.0.0.0:8000"); err != nil {
		panic(err)
	}

}
