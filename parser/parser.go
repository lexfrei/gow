package parser

import (
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func NewPlayerByLink(u url.URL) *Player {
	return &Player{
		url: u,
	}
}

func (p *Player) Gather() {
	res, err := http.Get(p.url.String())
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	p.Name = doc.Find(UserName).Text()
	p.Platform = doc.Find(Platform).Text()
	rankDD := doc.Find(DDSR).Text()
	if rankDD != "" {
		i, err := strconv.Atoi(rankDD)
		if err != nil {
			log.Println(err)
		}
		p.Rank.DD = i
	}

	rankHeal := doc.Find(HealSR).Text()
	if rankHeal != "" {
		i, err := strconv.Atoi(rankHeal)
		if err != nil {
			log.Println(err)
		}
		p.Rank.Heal = i
	}

	rankTank := doc.Find(TankSR).Text()
	if rankTank != "" {
		i, err := strconv.Atoi(rankTank)
		if err != nil {
			log.Println(err)
		}
		p.Rank.Tank = i
	}
}
