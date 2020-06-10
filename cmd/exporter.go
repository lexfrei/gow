package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/lexfrei/gow/parser"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	u, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	p := parser.NewPlayerByLink(*u)

	pr := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "rank",
		},
		[]string{
			"user",
			"platform",
			"role",
		},
	)

	prometheus.MustRegister(pr)

	go func() {
		timerCh := time.Tick(1 * time.Minute)
		for range timerCh {
			p.Gather()
			log.Println("gathered!")
			pr.WithLabelValues(p.Name, p.Platform, "tank").Set(float64(p.Rank.Tank))
			pr.WithLabelValues(p.Name, p.Platform, "heal").Set(float64(p.Rank.Heal))
			pr.WithLabelValues(p.Name, p.Platform, "dd").Set(float64(p.Rank.DD))
		}

	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9420", nil))
}
