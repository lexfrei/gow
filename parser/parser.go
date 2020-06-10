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

	s := doc.Find(HealSR).Text()
	if s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Println(err)
		}
		p.Rank.Heal = i
	}

	s = doc.Find(TankSR).Text()
	if s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Println(err)
		}
		p.Rank.Tank = i
	}

	d, e := doc.Find(EndorsmentShotcaller).Attr("data-value")
	if e {
		c, err := strconv.ParseFloat(d, 64)
		if err == nil {
			p.Endorsment.Shotcaller = c
		}
	}

	d, e = doc.Find(EndorsmentTeammate).Attr("data-value")
	if e {
		c, err := strconv.ParseFloat(d, 64)
		if err == nil {
			p.Endorsment.Teammate = c
		}
	}

	d, e = doc.Find(EndorsmentShotcaller).Attr("data-value")
	if e {
		c, err := strconv.ParseFloat(d, 64)
		if err == nil {
			p.Endorsment.Sportsmanship = c
		}
	}

}
