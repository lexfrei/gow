package parser

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var heroes []string = []string{ana, ashe, baptiste, bastion, brigitte, dva, doomfist, echo, genji, hanzo, junkrat, lucio, mccree, mei, mercy, moira, orisa, pharah, reaper, reinhardt, roadhog, sigma, soldier, sombra, symmetra, torb, tracer, widowmaker, winston, ball, zarya, zen}

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

	p.Name = doc.Find(userName).Text()
	p.Platform = doc.Find(platform).Text()

	rankDD := doc.Find(ddSR).Text()
	if rankDD != "" {
		i, err := strconv.Atoi(rankDD)
		if err != nil {
			log.Println(err)
		}
		p.Rank.DD = i
	}

	s := doc.Find(healSR).Text()
	if s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Println(err)
		}
		p.Rank.Heal = i
	}

	s = doc.Find(tankSR).Text()
	if s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Println(err)
		}
		p.Rank.Tank = i
	}

	s = doc.Find(endorsmentLvl).Text()
	if s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Println(err)
		}
		p.Endorsment.Level = i
	}

	d, e := doc.Find(endorsmentShotcaller).Attr("data-value")
	if e {
		c, err := strconv.ParseFloat(d, 64)
		if err == nil {
			p.Endorsment.Shotcaller = c
		}
	}

	d, e = doc.Find(endorsmentTeammate).Attr("data-value")
	if e {
		c, err := strconv.ParseFloat(d, 64)
		if err == nil {
			p.Endorsment.Teammate = c
		}
	}

	d, e = doc.Find(endorsmentShotcaller).Attr("data-value")
	if e {
		c, err := strconv.ParseFloat(d, 64)
		if err == nil {
			p.Endorsment.Sportsmanship = c
		}
	}

	p.Stats = append(p.Stats, parseStats(doc, true)...)
	p.Stats = append(p.Stats, parseStats(doc, false)...)
}

func parseStats(s *goquery.Document, isComp bool) []Stat {
	var sel *goquery.Selection
	var str string
	var stats []Stat

	if isComp {
		sel = s.Find(baseComp)
	} else {
		sel = s.Find(baseQP)
	}

	for _, code := range heroes {

		heroName, e := sel.Find(fmt.Sprintf(namePath, code)).Attr("option-id")
		if !e {
			continue
		}
		sel.Find(fmt.Sprintf(statPath, code)).Each(func(i int, s *goquery.Selection) {
			var stat Stat
			stat.Hero = heroName
			stat.IsComp = isComp
			stat.Name = s.Find("td:nth-child(1)").Text()

			str = s.Find("td:nth-child(2)").Text()
			switch {
			case strings.Contains(str, "%"):
				stat.Value = stringToFloat64(strings.Trim(str, "%"))
			case strings.Contains(str, ":"):
				stat.Value = timeToSec(str)
			default:
				stat.Value = stringToFloat64(str)
			}

			stats = append(stats, stat)
		})
	}
	return stats
}

func timeToSec(s string) (time float64) {
	switch len(s) {
	case 8:
		time = float64((((int(s[0])-48)*10+int(s[1])-48)*60+((int(s[3])-48)*10+(int(s[4])-48)))*60 + (int(s[6])-48)*10 + int(s[7]) - 48)
	case 5:
		time = float64(((int(s[0])-48)*10+(int(s[1])-48))*60 + (int(s[3])-48)*10 + int(s[4]) - 48)
	case 2:
		time = float64((int(s[0])-48)*10 + int(s[1]) - 48)
	default:
		time = 0
	}
	return
}

func stringToFloat64(s string) (u float64) {
	u, err := strconv.ParseFloat(s, 64)
	if err != nil {
		u = 0
		return
	}
	return
}
