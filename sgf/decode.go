package sgf

import (
	"strconv"

	"github.com/dgf/gotv/model"
)

type decode func(g *model.Game, p string)

// see https://en.wikipedia.org/wiki/Smart_Game_Format#About_the_format
var decoder = map[string]decode{

	// Add Black: locations of Black stones to be placed on the board prior to the first move
	//"AB": func(g *model.Game, p string) {},

	// Annotations: name of the person commenting the game.
	//"AN": func(g *model.Game, p string) {},

	// Application: application that was used to create the SGF file (e.g. CGOban2,...).
	//"AP": func(g *model.Game, p string) {},

	// Add White: locations of White stones to be placed on the board prior to the first move.
	//"AW": func(g *model.Game, p string) {},

	// Black Rank: rank of the Black player.
	"BR": func(g *model.Game, p string) {
		g.Black.Rank = model.ToRank(p)
	},

	// Black Team: name of the Black team.
	"BT": func(g *model.Game, p string) {
		g.Black.Name = p
	},

	// Copyright: copyright information.
	//"CP": func(g *model.Game, p string) {},

	// Date: date of the game.
	"DT": func(g *model.Game, p string) {
		g.Date = model.ToDate(p)
	},

	// Event: name of the event (e.g. 58th Honinbō Title Match).
	//"EV": func(g *model.Game, p string) {},

	// File format: version of SGF specification governing this SGF file.
	//"FF": func(g *model.Game, p string) {},

	// Game: type of game represented by this SGF file. A property value of 1 refers to Go.
	//"GM": func(g *model.Game, p string) {},

	// Game Name: name of the game record.
	"GN": func(g *model.Game, p string) {
		g.Name = p
	},

	// Handicap: the number of handicap stones given to Black. Placement of the handicap stones are set using the AB property.
	"HA": func(g *model.Game, p string) {
		h, _ := strconv.Atoi(p)
		g.Handicap = model.Handicap(h)
	},

	// Komi: komi.
	"KM": func(g *model.Game, p string) {
		k, _ := strconv.ParseFloat(p, 64)
		g.Komi = model.Komi(k)
	},

	// Opening: information about the opening (Fuseki), rarely used in any file.
	//"ON": func(g *model.Game, p string) {},

	// Overtime: overtime system.
	//"OT": func(g *model.Game, p string) {},

	// Black Name: name of the black player.
	"PB": func(g *model.Game, p string) {
		g.Black.Name = p
	},

	// Player: color of player to start.
	"PC": func(g *model.Game, p string) {},

	// Place: place where the game was played (e.g.: Tokyo).
	"PL": func(g *model.Game, p string) {
		g.Place = p
	},

	// White Name: name of the white player.
	"PW": func(g *model.Game, p string) {
		g.White.Name = p
	},

	// Result: result, usually in the format "B+R" (Black wins by resign) or "B+3.5" (black wins by 3.5).
	"RE": func(g *model.Game, p string) {
		g.Result = p
	},

	// Round: round (e.g.: 5th game).
	//"RO": func(g *model.Game, p string) {},

	// Rules: ruleset (e.g.: Japanese).
	//"RU": func(g *model.Game, p string) {},

	// Source: source of the SGF file.
	//"SO": func(g *model.Game, p string) {},

	// Size: size of the board, non-square boards are not supported.
	"SZ": func(g *model.Game, p string) {
		s, _ := strconv.Atoi(p)
		g.Size = model.Size(s)
	},

	// Time limit: time limit in seconds.
	//"TM": func(g *model.Game, p string) {},

	// User: name of the person who created the SGF file.
	//"US": func(g *model.Game, p string) {},

	// White Rank: rank of the White player.
	"WR": func(g *model.Game, p string) {
		g.White.Rank = model.ToRank(p)
	},

	// White Team: name of the White team.
	"WT": func(g *model.Game, p string) {
		if len(g.White.Name) == 0 {
			g.White.Name = p
		}
	},
}
