package parser

import (
	"fmt"
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

	qps := doc.Find(baseQP)
	for _, code := range heros {
		var h Hero
		name, e := qps.Find(fmt.Sprintf(namePath, code)).Attr("option-id")
		if e {
			h.Name = name
		}

		h.Combat.TotalDMG = stringToUint32(qps.Find(fmt.Sprintf(metricPath, code, combat["TotalDMG"])).Text())
		h.Combat.BarrierDMG = stringToUint32(qps.Find(fmt.Sprintf(metricPath, code, combat["TotalDMG"])).Text())
		h.Combat.Deaths = stringToUint32(qps.Find(fmt.Sprintf(metricPath, code, combat["Deaths"])).Text())
		h.Combat.Eliminations = stringToUint32(qps.Find(fmt.Sprintf(metricPath, code, combat["Eliminations"])).Text())
		h.Combat.EnvKills = stringToUint32(qps.Find(fmt.Sprintf(metricPath, code, combat["EnvKills"])).Text())
		h.Combat.FinalBlows = stringToUint32(qps.Find(fmt.Sprintf(metricPath, code, combat["FinalBlows"])).Text())
		h.Combat.HeroDMG = stringToUint32(qps.Find(fmt.Sprintf(metricPath, code, combat["HeroDMG"])).Text())
		h.Combat.MeleeFinalBlows = stringToUint32(qps.Find(fmt.Sprintf(metricPath, code, combat["MeleeFinalBlows"])).Text())
		h.Combat.Multikills = stringToUint32(qps.Find(fmt.Sprintf(metricPath, code, combat["Multikills"])).Text())
		h.Combat.ObjKills = stringToUint32(qps.Find(fmt.Sprintf(metricPath, code, combat["ObjKills"])).Text())
		h.Combat.SoloKills = stringToUint32(qps.Find(fmt.Sprintf(metricPath, code, combat["SoloKills"])).Text())
		h.Combat.OnFire = timeToSec(qps.Find(fmt.Sprintf(metricPath, code, combat["OnFire"])).Text())
		h.Combat.ObjTime = timeToSec(qps.Find(fmt.Sprintf(metricPath, code, combat["ObjTime"])).Text())
		h.Combat.WeaponAccuracy = stringToPecents(qps.Find(fmt.Sprintf(metricPath, code, combat["WeaponAccuracy"])).Text())
		h.Combat.MeleeAccuracy = stringToPecents(qps.Find(fmt.Sprintf(metricPath, code, combat["MeleeAccuracy"])).Text())
	}

}

func timeToSec(s string) (time uint32) {
	switch len(s) {
	case 8:
		time = uint32((((int(s[0])-48)*10+int(s[1])-48)*60+((int(s[3])-48)*10+(int(s[4])-48)))*60 + (int(s[6])-48)*10 + int(s[7]) - 48)
	case 5:
		time = uint32(((int(s[0])-48)*10+(int(s[1])-48))*60 + (int(s[3])-48)*10 + int(s[4]) - 48)
	case 3:
		time = uint32((int(s[0])-48)*10 + int(s[1]) - 48)
	default:
		time = 0
	}
	return
}

func stringToPecents(s string) (percs uint8) {
	data, err := strconv.ParseUint(s[0:1], 10, 8)
	if err != nil {
		percs = 0
		return
	}
	percs = uint8(data)
	return
}

func stringToUint32(s string) (u uint32) {
	data, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		u = 0
		return
	}
	u = uint32(data)
	return
}
