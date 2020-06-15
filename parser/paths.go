package parser

const (
	userName = ".header-masthead"
	platform = "div.masthead-buttons.button-group.js-button-group > a.button.m-white-outline.m-sm.is-active"

	tankSR = "div.masthead-player > div > div.competitive-rank > div:nth-child(1) > div:nth-child(2) > div.competitive-rank-level"
	ddSR   = "div.masthead-player > div > div.competitive-rank > div:nth-child(2) > div:nth-child(2) > div.competitive-rank-level"
	healSR = "div.masthead-player > div > div.competitive-rank > div:nth-child(3) > div:nth-child(2) > div.competitive-rank-level"

	endorsmentLvl           = "div.masthead-player > div > div.EndorsementIcon-tooltip > div.u-center"
	endorsmentShotcaller    = "div.masthead-player > div > div.EndorsementIcon-tooltip > div.endorsement-level > div > div > svg.EndorsementIcon-border.EndorsementIcon-border--shotcaller"
	endorsmentTeammate      = "div.masthead-player > div > div.EndorsementIcon-tooltip > div.endorsement-level > div > div > svg.EndorsementIcon-border.EndorsementIcon-border--teammate"
	endorsmentSportsmanship = "div.masthead-player > div > div.EndorsementIcon-tooltip > div.endorsement-level > div > div > svg.EndorsementIcon-border.EndorsementIcon-border--sportsmanship"

	baseComp = "#competitive"
	baseQP   = "#quickplay"

	statPath = "section > div > div[data-category-id=\"%s\"] > div > div > table > tbody > tr"
	namePath = "section:nth-child(2) > div > div > div > select > option[value=\"%s\"]"
	heroList = "section:nth-child(2) > div > div.flex-container > div.flex-item > div > select > option"

	// Magic numbers
	// Heroes
	ana        = "0x02E000000000013B"
	ashe       = "0x02E0000000000200"
	baptiste   = "0x02E0000000000221"
	bastion    = "0x02E0000000000015"
	brigitte   = "0x02E0000000000195"
	dva        = "0x02E000000000007A"
	doomfist   = "0x02E000000000012F"
	echo       = "0x02E0000000000206"
	genji      = "0x02E0000000000029"
	hanzo      = "0x02E0000000000005"
	junkrat    = "0x02E0000000000065"
	lucio      = "0x02E0000000000079"
	mccree     = "0x02E0000000000042"
	mei        = "0x02E00000000000DD"
	mercy      = "0x02E0000000000004"
	moira      = "0x02E00000000001A2"
	orisa      = "0x02E000000000013E"
	pharah     = "0x02E0000000000008"
	reaper     = "0x02E0000000000002"
	reinhardt  = "0x02E0000000000007"
	roadhog    = "0x02E0000000000040"
	sigma      = "0x02E000000000023B"
	soldier    = "0x02E000000000006E"
	sombra     = "0x02E000000000012E"
	symmetra   = "0x02E0000000000016"
	torb       = "0x02E0000000000006"
	tracer     = "0x02E0000000000003"
	widowmaker = "0x02E000000000000A"
	winston    = "0x02E0000000000009"
	ball       = "0x02E00000000001CA"
	zarya      = "0x02E0000000000068"
	zen        = "0x02E0000000000020"
)
