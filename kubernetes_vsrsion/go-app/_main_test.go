package main_test

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func main() {

	//recordMetrics()

	fmt.Println("Serving requests on port xxxxx")

	// var dir string

	// flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	// flag.Parse()

	// router := mux.NewRouter()
	// router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	// router.Handle("/metrics", promhttp.Handler())
	// log.Fatal(http.ListenAndServe(":8009", router))

	go func() {
		fmt.Println("metrics start")

		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8010", nil)
	}()

	app := fiber.New()

	app.Get("/health_test", func(c *fiber.Ctx) error {
		return c.JSON("health_test OK")
	})

	app.Listen(":8009")

}
