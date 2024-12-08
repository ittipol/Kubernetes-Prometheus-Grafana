package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"
)

type Device struct {
	ID       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}

type handler struct {

	// Prometheus metrics
	metrics *metrics

	// App configuration object
	config *Config

	rdb *redis.Client
}

// var counter int = 0

func main() {
	// Load app config from yaml file.
	var c Config
	c.loadConfig("config.yml")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "password",
		DB:       0, // use default DB
	})

	// Create Prometheus registry
	reg := prometheus.NewRegistry()
	m := NewMetrics(reg)

	// Create Prometheus HTTP server to expose metrics
	pMux := http.NewServeMux()
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	pMux.Handle("/metrics", promHandler)

	// Start an HTTP server to expose Prometheus metrics in the background.
	metricsPort := fmt.Sprintf(":%d", c.MetricsPort)
	fmt.Printf("metrics start %s", metricsPort)
	go func() {
		log.Fatal(http.ListenAndServe(metricsPort, pMux))

		// router := mux.NewRouter()
		// router.Handle("/metrics", promhttp.Handler())
		// log.Fatal(http.ListenAndServe(metricsPort, router))
	}()

	// go simulateTraffic(m)

	// Initialize Gin handler.
	h := handler{config: &c, metrics: m, rdb: rdb}

	app := fiber.New()

	app.Get("/healthz", h.getHealth)
	app.Get("/api/devices", h.redisSet)

	appPort := fmt.Sprintf(":%d", c.AppPort)
	log.Fatal(app.Listen(appPort))
}

// getHealth returns the status of the application.
func (h *handler) getHealth(c fiber.Ctx) error {
	return c.SendStatus(200)
}

func getDevices(c fiber.Ctx) error {
	sleep(1000)
	dvs := []Device{
		{1, "5F-33-CC-1F-43-82", "2.1.6"},
		{2, "EF-2B-C4-F5-D6-34", "2.1.6"},
	}

	return c.JSON(dvs)
}

func (h *handler) redisSet(c fiber.Ctx) error {

	ctx := context.Background()

	value := rand.Intn(10000)
	key := fmt.Sprintf("key:%d", value)

	err := h.rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	return c.SendString("Redis set...")
}

// func simulateTraffic(m *metrics) {
// 	for {
// 		now := time.Now()
// 		sleep(1000)
// 		m.duration.WithLabelValues("/fake", "200").Observe(time.Since(now).Seconds())
// 	}
// }

func sleep(ms int) {
	rand.Seed(time.Now().UnixNano())
	now := time.Now()
	n := rand.Intn(ms + now.Second())
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// s3Connect initializes the S3 session.
// func (h *handler) s3Connect() {
// 	// Get credentials to authorize with AWS S3 API.
// 	crds := credentials.NewStaticCredentials(h.config.S3Config.User, h.config.S3Config.Secret, "")

// 	// Create S3 config.
// 	s3c := aws.Config{
// 		Region:           &h.config.S3Config.Region,
// 		Endpoint:         &h.config.S3Config.Endpoint,
// 		S3ForcePathStyle: &h.config.S3Config.PathStyle,
// 		Credentials:      crds,
// 	}

// 	// Establish a new session with the AWS S3 API.
// 	h.sess = session.Must(session.NewSessionWithOptions(session.Options{
// 		SharedConfigState: session.SharedConfigEnable,
// 		Config:            s3c,
// 	}))
// }

// dbConnect creates a connection pool to connect to Postgres.
// func (h *handler) dbConnect() {
// 	url := fmt.Sprintf("postgres://%s:%s@%s:5432/%s",
// 		h.config.DbConfig.User, h.config.DbConfig.Password, h.config.DbConfig.Host, h.config.DbConfig.Database)

// 	// Connect to the Postgres database.
// 	dbpool, err := pgxpool.New(context.Background(), url)
// 	if err != nil {
// 		log.Fatalf("Unable to create connection pool: %s", err)
// 	}

// 	h.dbpool = dbpool
// }
