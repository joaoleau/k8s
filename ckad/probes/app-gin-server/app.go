package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"io"
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
)


func working(max int) int {
	min := 0
	randWrk := rand.Intn(max-min) + min
	return randWrk
}


func start(ctx *gin.Context) {
	maxStr := ctx.Query("max")
	maxInt, err := strconv.Atoi(maxStr)
	if err != nil {
		panic(err)
	}
	wrk := working(maxInt)
	if wrk == 1 {
		ctx.JSON(200, gin.H{
			"startup": "Starting...",
		})
		fmt.Println("start: Starting...")
	} else {
		ctx.JSON(500, gin.H{
			"startup": "Not started",
		})
		fmt.Println("start: Not started")
	}
}


func readiness(ctx *gin.Context) {
	url := os.Getenv("KEY_APP_URL")
	if url == "" {
		fmt.Println("error: KEY_APP_URL is not set")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "KEY APP URL is not configured"})
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error: KEY APP not found -", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "KEY APP not found"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("readiness: Available for network")
		ctx.JSON(http.StatusOK, gin.H{"message": "Available for network"})
	} else {
		fmt.Printf("readiness: Not available for network (status: %d)\n", resp.StatusCode)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Not available for network"})
	}
}


func healthz(ctx *gin.Context) {
	maxStr := ctx.Query("max")
	maxInt, err := strconv.Atoi(maxStr)
	if err != nil {
		panic(err)
	}
	wrk := working(maxInt)
	if wrk == 1 {
		ctx.JSON(200, gin.H{
			"message": "Working...",
		})
		fmt.Println("liveness: Working...")
	} else {
		ctx.JSON(500, gin.H{
			"error": "Not Working...",
		})
		fmt.Println("liveness: Not Working...")
	}
}


func prestop(ctx *gin.Context) {
	time.Sleep(time.Duration(10))
	fmt.Println("Stoped")
}

func main() {
	
	gin.DefaultWriter = io.Discard
	g := gin.Default()

	fmt.Println("Starting App...")
	time.Sleep(time.Duration(5))

	g.GET("/start", start)
	g.GET("/readiness", readiness)
	g.GET("/healthz", healthz)
	g.GET("/prestop", prestop)

	g.Run(":3000")
}
