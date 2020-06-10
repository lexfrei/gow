package parser

import "net/url"

type Player struct {
	Platform string
	Name     string
	Rank     Rank
	url      url.URL
}

type Rank struct {
	DD   int
	Tank int
	Heal int
}
