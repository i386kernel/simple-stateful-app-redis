package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
	"net/http"
	"os"
	"time"
)

var redisURL = os.Getenv("REDIS_URL")
var redisPass = os.Getenv("REDIS_PASSWORD")

var rclient = &redis.Client{}

var conversionMAP = map[string]string{
	"ASR": "-3h",    // North America Atlantic standard time
	"EST": "-5h",    // North american eastern standard time
	"BST": "+1h",    // British summer time
	"IST": "+5h30m", // Indian standard time
	"HKT": "+8h",    // Hong Kong Time
	"ART": "-3h",    // Argentina Time
}

func addUp(num1, num2 int) int {
	return num1 + num2
}

func redisConnect() {
	ctx := context.Background()
	fmt.Println("Go Redis Client")

	reo := &redis.Options{
		Addr:     redisURL,
		Password: redisPass,
		DB:       0,
	}
	rclient = redis.NewClient(reo)
	pong, err := rclient.Ping(ctx).Result()
	fmt.Println(pong, err)
}

func redisSET() {
	ctx := context.Background()
	for k, v := range conversionMAP {
		err := rclient.Set(ctx, k, v, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

func redisGET(key string) (string, error) {
	ctx := context.Background()
	val, err := rclient.Get(ctx, key).Result()
	if err != nil {
		return " ", err
	}
	fmt.Println(val)
	return val, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	timeZone := r.URL.Query().Get("tz")
	if timeZone == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Error 400: 'tz' query parameter not found")
		return
	}
	w.WriteHeader(http.StatusOK)
	tf, err := redisGET(timeZone)
	if err != nil {
		fmt.Fprintf(w, "unable to get the timezone %s\n", err)
	}
	fmt.Fprintf(w, "Time in timezone %v is %s\n", timeZone, tf)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Error 404: The requested url does not exist.")
}

type handlerFunc func(w http.ResponseWriter, r *http.Request)

// Middleware take http handler, validates something and then re-directs from there.
func loggingMiddleware(handler handlerFunc) handlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s = %s - %s", time.Now().Format("2020-06-28 13:00"), r.Method, r.URL.String())
		handler(w, r)
	}
	return fn
}
func init() {
	redisConnect()
	redisSET()
}

func main() {
	http.HandleFunc("/convert", loggingMiddleware(handler))
	http.HandleFunc("/", loggingMiddleware(notFoundHandler))
	log.Printf("%s - Starting server on port: 8888", time.Now().Format("12-12-1983 20:18"))
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}
