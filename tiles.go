package main

import (
	"fmt"
	"log"
	"math/rand"
)

var allTiles map[uint8]Tile

func init() {
	allTiles = make(map[uint8]Tile)
	for i := 0; i < 52; i++ {
		var color string
		switch i / 13 {
		case 0:
			color = "Yellow"
		case 1:
			color = "Blue"
		case 2:
			color = "Black"
		case 3:
			color = "Red"
		default:
			panic("?")
		}
		allTiles[uint8(i)] = ColoredTile{
			color:  color,
			val:    uint8(i%13) + 1,
			tileNo: uint8(i),
		}
	}
	allTiles[52] = JokerTile{}
}

func NewTileSet() (tileSetArr [105]Tile, indicator Tile) {
	var okeyTile Tile
	var indicatorTile Tile
	// joker açmayana kadar rastgele seç
	for { // buraya hiçbir şekilde sonsuz loopa girmesin diye counter eklenebilir, basit kalsın diye eklemedim
		okeyTileIdx := uint8(rand.Intn(53))
		if allTiles[okeyTileIdx].Type() != "joker" {
			indicatorTile = allTiles[okeyTileIdx]
			if okeyTileIdx%13 == 12 { // 13 açılırsa 1 okey olur
				okeyTile = allTiles[okeyTileIdx-12]
			} else {
				okeyTile = allTiles[okeyTileIdx+1]
			}
			break
		}
	}
	log.Printf("Indicator tile selected: %s %d", indicatorTile.Color(), indicatorTile.Value())
	log.Printf("Okey tile selected: %s %d", okeyTile.Color(), okeyTile.Value())

	var tileSet []Tile
	for _, tile := range allTiles {
		if tile.Type() == "joker" {
			// havuza 2 tane okey belirlenen taş değerinde taş ekle
			tileSet = append(tileSet, []Tile{
				&JokerTile{
					color:  okeyTile.Color(),
					val:    okeyTile.Value(),
					tileNo: 52,
				},
				&JokerTile{
					color:  okeyTile.Color(),
					val:    okeyTile.Value(),
					tileNo: 52,
				}}...)
		} else if tile.Color() == okeyTile.Color() && tile.Value() == okeyTile.Value() {
			// havuza 2 tane okey tipinde taş ekle
			tileSet = append(tileSet, []Tile{
				&OkeyTile{
					color:  okeyTile.Color(),
					val:    okeyTile.Value(),
					tileNo: 53, // gereksiz
				},
				&OkeyTile{
					color:  okeyTile.Color(),
					val:    okeyTile.Value(),
					tileNo: 53,
				}}...)
		} else if tile.Color() == indicatorTile.Color() && tile.Value() == indicatorTile.Value() {
			// sadece 1 gösterge taş havuzunda
			tileSet = append(tileSet, &ColoredTile{
				color:  tile.Color(),
				val:    tile.Value(),
				tileNo: tile.TileNo(),
			})
		} else {
			// normal taşlar için 2 şer çift havuzda
			tileSet = append(tileSet, []Tile{
				&ColoredTile{
					color:  tile.Color(),
					val:    tile.Value(),
					tileNo: tile.TileNo(),
				},
				&ColoredTile{
					color:  tile.Color(),
					val:    tile.Value(),
					tileNo: tile.TileNo(),
				}}...)
		}
	}
	copy(tileSetArr[:], tileSet)
	return tileSetArr, &ColoredTile{
		color:  indicatorTile.Color(),
		val:    indicatorTile.Value(),
		tileNo: indicatorTile.TileNo(),
	}
}

type Tile interface {
	Value() uint8
	Type() string
	Color() string
	TileNo() uint8
	ToString() string
}

type ColoredTile struct {
	color  string
	val    uint8
	tileNo uint8
}

func (t ColoredTile) Value() uint8 {
	return t.val
}
func (t ColoredTile) Type() string {
	return "colored"
}
func (t ColoredTile) Color() string {
	return t.color
}

func (t ColoredTile) ToString() string {
	switch t.Color() {
	case "Yellow":
		return fmt.Sprintf("  %s%d%s", Yellow, t.Value(), Reset)
	case "Red":
		return fmt.Sprintf("  %s%d%s", Red, t.Value(), Reset)
	case "Blue":
		return fmt.Sprintf("  %s%d%s", Blue, t.Value(), Reset)
	case "Black":
		return fmt.Sprintf("  %s%d%s", White, t.Value(), Reset)
	}
	return "?"
}

func (t ColoredTile) TileNo() uint8 {
	return t.tileNo
}

type OkeyTile struct {
	color  string
	val    uint8
	tileNo uint8
}

func (t OkeyTile) ToString() string {
	return Green + " OKEY" + Reset
}

func (t OkeyTile) Value() uint8 {
	return t.val
}

func (t OkeyTile) Type() string {
	return "okey"
}

func (t OkeyTile) Color() string {
	return t.color
}

func (t OkeyTile) TileNo() uint8 {
	return t.tileNo
}

type JokerTile struct {
	color  string
	val    uint8
	tileNo uint8
}

func (t JokerTile) ToString() string {
	switch t.Color() {
	case "Yellow":
		return fmt.Sprintf("  %sJoker%d%s", Yellow, t.Value(), Reset)
	case "Red":
		return fmt.Sprintf("  %sJoker%d%s", Red, t.Value(), Reset)
	case "Blue":
		return fmt.Sprintf("  %sJoker%d%s", Blue, t.Value(), Reset)
	case "Black":
		return fmt.Sprintf("  %sJoker%d%s", White, t.Value(), Reset)
	}
	return "?"
}

func (t JokerTile) Value() uint8 {
	return t.val
}

func (t JokerTile) Type() string {
	return "joker"
}

func (t JokerTile) Color() string {
	return t.color
}

func (t JokerTile) TileNo() uint8 {
	return t.tileNo
}
