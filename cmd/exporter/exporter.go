package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/lexfrei/gow/parser"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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
	stats = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "stat",
		},
		[]string{
			"user",
			"platform",
			"type",
			"stat",
			"hero",
		},
	)
)

func main() {
	r := prometheus.NewRegistry()
	r.MustRegister(playerRank)
	r.MustRegister(playerEndorsment)
	r.MustRegister(stats)

	u, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		getStats(u)
		timerCh := time.Tick(1 * time.Minute)
		for range timerCh {
			getStats(u)
		}

	}()

	handler := promhttp.HandlerFor(r, promhttp.HandlerOpts{})
	http.Handle("/metrics", handler)
	log.Fatal(http.ListenAndServe(":9420", nil))
}

func getStats(u *url.URL) {
	p := parser.NewPlayerByLink(*u)
	p.Gather()
	playerRank.WithLabelValues(p.Name, p.Platform, "tank").Set(float64(p.Rank.Tank))
	playerRank.WithLabelValues(p.Name, p.Platform, "heal").Set(float64(p.Rank.Heal))
	playerRank.WithLabelValues(p.Name, p.Platform, "dd").Set(float64(p.Rank.DD))
	playerEndorsment.WithLabelValues(p.Name, p.Platform, "level").Set(float64(p.Endorsment.Level))
	playerEndorsment.WithLabelValues(p.Name, p.Platform, "sportsmanship").Set(p.Endorsment.Sportsmanship)
	playerEndorsment.WithLabelValues(p.Name, p.Platform, "shotcaller").Set(p.Endorsment.Shotcaller)
	playerEndorsment.WithLabelValues(p.Name, p.Platform, "teammate").Set(p.Endorsment.Teammate)
	// TODO: Optimize
	for name, gameStats := range p.Heroes {
		for _, stat := range gameStats {
			if stat.Value.Competitive != 0 {
				stats.WithLabelValues(p.Name, p.Platform, "competitive", normalize(stat.Name), normalize(name)).Set(stat.Value.Competitive)
			}
			if stat.Value.QP != 0 {
				stats.WithLabelValues(p.Name, p.Platform, "qp", normalize(stat.Name), normalize(name)).Set(stat.Value.QP)
			}
		}
	}
}

func normalize(s string) string {
	isMn := func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
	}
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, s)
	return strings.ReplaceAll(result, " ", "_")
}
