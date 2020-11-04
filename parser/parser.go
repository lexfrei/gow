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

func NewPlayerByLink(u url.URL) *Player {
	return &Player{
		url:    u,
		Heroes: map[string][]Stat{},
	}
}

func (p *Player) Gather() {
	res, err := http.Get(p.url.String())
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	p.Name = doc.Find(userName).Text()
	p.Platform = doc.Find(platform).Text()

	var rawString string

	doc.Find(srPath).Each(func(i int, s *goquery.Selection) {
		rawString, e := s.Find(".competitive-rank-tier.competitive-rank-tier-tooltip").Attr("data-ow-tooltip-text")
		if e {
			switch rawString {
			case "Tank Skill Rating":
				p.Rank.Tank = stringToInt(s.Text())
			case "Damage Skill Rating":
				p.Rank.DD = stringToInt(s.Text())
			case "Support Skill Rating":
				p.Rank.Heal = stringToInt(s.Text())
			}
		}
	})

	rawString = doc.Find(endorsmentLvl).Text()
	if rawString != "" {
		i, err := strconv.Atoi(rawString)
		if err != nil {
			log.Println(err)
		}
		p.Endorsment.Level = i
	}

	rawString, exists := doc.Find(endorsmentShotcaller).Attr("data-value")
	if exists {
		rawEndorsment, err := strconv.ParseFloat(rawString, 64)
		if err == nil {
			p.Endorsment.Shotcaller = rawEndorsment
		}
	}

	rawString, exists = doc.Find(endorsmentTeammate).Attr("data-value")
	if exists {
		rawEndorsment, err := strconv.ParseFloat(rawString, 64)
		if err == nil {
			p.Endorsment.Teammate = rawEndorsment
		}
	}

	rawString, exists = doc.Find(endorsmentSportsmanship).Attr("data-value")
	if exists {
		rawEndorsment, err := strconv.ParseFloat(rawString, 64)
		if err == nil {
			p.Endorsment.Sportsmanship = rawEndorsment
		}
	}

	p.parseStats(doc)
}

func (p *Player) parseStats(s *goquery.Document) {
	var sel *goquery.Selection
	var str string
	var switcher = []bool{true, false}

	heroes := make(map[string]string)
	stats := make(map[string]string)

	s.Find("section:nth-child(2)").Find("option").Each(func(i int, s *goquery.Selection) {
		code, e := s.Attr("value")
		if e {
			heroes[code] = s.Text()
		}
	})

	for code, name := range heroes {
		fmt.Println(name)
		s.Find(baseComp).Find(fmt.Sprintf("div[data-category-id=\"%s\"]", code)).Find("table.DataTable").Each(func(i int, s *goquery.Selection) {
			fmt.Printf("\t%s\n", s.Find(".stat-title").Text())
			s.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
				stat, e := s.Attr("data-stat-id")
				if e {
					stats[stat] = s.Find("td:nth-child(1)").Text()
				}
				fmt.Printf("\t\t%s\t%s\n", s.Find("td:nth-child(1)").Text(), s.Find("td:nth-child(2)").Text())
			})
		})
	}

	fmt.Printf("Total stat code found:\t%d\n", len(stats))

	for _, isComp := range switcher {
		if isComp {
			sel = s.Find(baseComp)
		} else {
			sel = s.Find(baseQP)
		}

		for _, heroCode := range heroes {
			heroName, e := sel.Find(fmt.Sprintf(namePath, heroCode)).Attr("option-id")
			if !e {
				continue
			}

			var value float64
			sel.Find(fmt.Sprintf(statPath, heroCode)).Each(func(i int, s *goquery.Selection) {
				var stat Stat
				// id, exists := s.Attr("data-stat-id")
				// if !exists {
				// 	stat.Name = id
				// 	return
				// }

				stat.Name = s.Find("td:nth-child(1)").Text()
				str = s.Find("td:nth-child(2)").Text()

				switch {
				case strings.Contains(str, "%"):
					value = stringToFloat64(strings.Trim(str, "%"))
				case strings.Contains(str, ":"):
					value = timeToSec(str)
				default:
					value = stringToFloat64(str)
				}

				if isComp {
					stat.Value.Competitive = value
				} else {
					stat.Value.QP = value
				}

				p.Heroes[heroName] = append(p.Heroes[heroName], stat)
			})
		}

	}

	return
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
	// no reason to check this err
	u, _ = strconv.ParseFloat(s, 64)
	return
}

func stringToInt(s string) (i int) {
	// no reason to check this err
	i, _ = strconv.Atoi(s)
	return
}
