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

	metricPath = "section:nth-child(2) > div > div[data-category-id=\"%s\"] > div:nth-child(1) > div > table > tbody > tr[data-stat-id=\"%s\"] > td:nth-child(2)"
	namePath   = "section:nth-child(2) > div > div > div > select > option[value=\"%s\"]"
)

// #quickplay > section:nth-child(2) > div > div.flex-container\@md-min.m-bottom-items > div > div > select
var heros = map[string]string{
	"ana":        "0x02E000000000013B",
	"ashe":       "0x02E0000000000200",
	"baptiste":   "0x02E0000000000221",
	"bastion":    "0x02E0000000000015",
	"brigitte":   "0x02E0000000000195",
	"dva":        "0x02E000000000007A",
	"doomfist":   "0x02E000000000012F",
	"echo":       "0x02E0000000000206",
	"genji":      "0x02E0000000000029",
	"hanzo":      "0x02E0000000000005",
	"junkrat":    "0x02E0000000000065",
	"lucio":      "0x02E0000000000079",
	"mccree":     "0x02E0000000000042",
	"mei":        "0x02E00000000000DD",
	"mercy":      "0x02E0000000000004",
	"moira":      "0x02E00000000001A2",
	"orisa":      "0x02E000000000013E",
	"pharah":     "0x02E0000000000008",
	"reaper":     "0x02E0000000000002",
	"reinhardt":  "0x02E0000000000007",
	"roadhog":    "0x02E0000000000040",
	"sigma":      "0x02E000000000023B",
	"soldier":    "0x02E000000000006E",
	"sombra":     "0x02E000000000012E",
	"symmetra":   "0x02E0000000000016",
	"torb":       "0x02E0000000000006",
	"tracer":     "0x02E0000000000003",
	"widowmaker": "0x02E000000000000A",
	"winston":    "0x02E0000000000009",
	"ball":       "0x02E00000000001CA",
	"zarya":      "0x02E0000000000068",
	"zen":        "0x02E0000000000020",
}

var combat = map[string]string{
	"TotalDMG":        "0x08600000000001BF",
	"BarrierDMG":      "0x0860000000000515",
	"Deaths":          "0x086000000000002A",
	"Eliminations":    "0x086000000000001F",
	"EnvKills":        "0x0860000000000362",
	"FinalBlows":      "0x086000000000002B",
	"HeroDMG":         "0x08600000000004B7",
	"MeleeFinalBlows": "0x086000000000037F",
	"Multikills":      "0x0860000000000345",
	"ObjKills":        "0x086000000000031C",
	"ObjTime":         "0x086000000000031D",
	"MeleeAccuracy":   "0x08600000000005AC",
	"SoloKills":       "0x086000000000002D",
	"OnFire":          "0x08600000000003CC",
	"WeaponAccuracy":  "0x086000000000002F",
}
