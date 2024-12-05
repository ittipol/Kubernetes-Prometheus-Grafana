package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
}

func main() {
	// Load app config from yaml file.
	var c Config
	c.loadConfig("config.yml")

	reg := prometheus.NewRegistry()
	m := NewMetrics(reg)

	pMux := http.NewServeMux()
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	pMux.Handle("/metrics", promHandler)

	go func() {
		log.Fatal(http.ListenAndServe(":8081", pMux))
	}()

	// go simulateTraffic(m)

	h := handler{config: &c, metrics: m}

	app := fiber.New()

	app.Get("/healthz", h.getHealth)
	app.Get("/api/devices", getDevices)

	log.Fatal(app.Listen(":8080"))
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

func simulateTraffic(m *metrics) {
	for {
		now := time.Now()
		sleep(1000)
		m.duration.WithLabelValues("/fake", "200").Observe(time.Since(now).Seconds())
	}
}

func sleep(ms int) {
	rand.Seed(time.Now().UnixNano())
	now := time.Now()
	n := rand.Intn(ms + now.Second())
	time.Sleep(time.Duration(n) * time.Millisecond)
}
