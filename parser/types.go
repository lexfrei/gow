package parser

import (
	"net/url"
)

type Player struct {
	Platform   string
	Name       string
	Rank       Rank
	url        url.URL
	Endorsment Endorsment
	Comp       []Hero
	QP         []Hero
}

type Rank struct {
	DD   int
	Tank int
	Heal int
}

type Endorsment struct {
	Level         int
	Shotcaller    float64
	Teammate      float64
	Sportsmanship float64
}

type Hero struct {
	Name   string
	Combat Combat
}

type Combat struct {
	TotalDMG        uint32
	BarrierDMG      uint32
	Deaths          uint32
	Eliminations    uint32
	EnvKills        uint32
	FinalBlows      uint32
	HeroDMG         uint32
	MeleeFinalBlows uint32
	Multikills      uint32
	ObjKills        uint32
	SoloKills       uint32
	// Time in secs
	OnFire  uint32
	ObjTime uint32
	// Percents
	WeaponAccuracy uint8
	MeleeAccuracy  uint8
}
