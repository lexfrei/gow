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

var (
	playerRank = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "rank",
		},
		[]string{
			"user",
			"platform",
			"role",
		},
	)
	playerEndorsment = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "endorsment",
		},
		[]string{
			"user",
			"platform",
			"type",
		},
	)
)

func init() {
	prometheus.MustRegister(playerRank)
	prometheus.MustRegister(playerEndorsment)
}

func main() {
	u, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		gather(u)
		timerCh := time.Tick(1 * time.Minute)
		for range timerCh {
			gather(u)
		}

	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9420", nil))
}

func gather(u *url.URL) {
	p := parser.NewPlayerByLink(*u)
	p.Gather()
	playerRank.WithLabelValues(p.Name, p.Platform, "tank").Set(float64(p.Rank.Tank))
	playerRank.WithLabelValues(p.Name, p.Platform, "heal").Set(float64(p.Rank.Heal))
	playerRank.WithLabelValues(p.Name, p.Platform, "dd").Set(float64(p.Rank.DD))
	playerEndorsment.WithLabelValues(p.Name, p.Platform, "level").Set(float64(p.Endorsment.Level))
	playerEndorsment.WithLabelValues(p.Name, p.Platform, "sportsmanship").Set(p.Endorsment.Sportsmanship)
	playerEndorsment.WithLabelValues(p.Name, p.Platform, "shotcaller").Set(p.Endorsment.Shotcaller)
	playerEndorsment.WithLabelValues(p.Name, p.Platform, "teammate").Set(p.Endorsment.Teammate)
}
