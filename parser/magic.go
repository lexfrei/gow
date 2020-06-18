package parser

const (
	userName = ".header-masthead"
	platform = "div.masthead-buttons.button-group.js-button-group > a.button.m-white-outline.m-sm.is-active"

	gamesWonTotal = "#overview-section > div > div.u-max-width-container.row.content-box.gutter-18 > div > div > p > span"

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
)

// heroList = "section:nth-child(2) > div > div.flex-container > div.flex-item > div > select > option"
var heroes = []string{
	"0x02E000000000013B", // Ana
	"0x02E0000000000200", // Ashe
	"0x02E0000000000221", // Baptiste
	"0x02E0000000000015", // Bastion
	"0x02E0000000000195", // Brigitte
	"0x02E000000000007A", // Dva
	"0x02E000000000012F", // Doomfist
	"0x02E0000000000206", // Echo
	"0x02E0000000000029", // Genji
	"0x02E0000000000005", // Hanzo
	"0x02E0000000000065", // Junkrat
	"0x02E0000000000079", // Lucio
	"0x02E0000000000042", // Mccree
	"0x02E00000000000DD", // Mei
	"0x02E0000000000004", // Mercy
	"0x02E00000000001A2", // Moira
	"0x02E000000000013E", // Orisa
	"0x02E0000000000008", // Pharah
	"0x02E0000000000002", // Reaper
	"0x02E0000000000007", // Reinhardt
	"0x02E0000000000040", // Roadhog
	"0x02E000000000023B", // Sigma
	"0x02E000000000006E", // Soldier
	"0x02E000000000012E", // Sombra
	"0x02E0000000000016", // Symmetra
	"0x02E0000000000006", // Torb
	"0x02E0000000000003", // Tracer
	"0x02E000000000000A", // Widowmaker
	"0x02E0000000000009", // Winston
	"0x02E00000000001CA", // Ball
	"0x02E0000000000068", // Zarya
	"0x02E0000000000020", // Zen
}
