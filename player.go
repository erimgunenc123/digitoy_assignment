package main

import (
	"fmt"
	"math"
	"slices"
)

var White = "\033[97m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Red = "\033[31m"
var Blue = "\033[34m"
var Reset = "\033[0m"

type Player struct {
	playerName string
	board      []Tile
	okeyCount  int
	okeyTiles  []Tile
}

func NewPlayer(playerName string) *Player {
	return &Player{
		playerName: playerName,
		board:      []Tile{},
		okeyCount:  0,
		okeyTiles:  []Tile{},
	}
}

// bütün  12345 veya 555 gibi setleri bul
// minimum overlap eden en büyük setler bütününü al

func (p *Player) SetBoard(newBoard []Tile) {
	p.okeyCount = 0
	for _, tile := range newBoard {
		if tile.Type() == "okey" { // okey taşlarını boardda tutmaya gerek yok
			p.okeyCount += 1
			p.okeyTiles = append(p.okeyTiles, tile)
		} else {
			p.board = append(p.board, tile)
		}
	}

	sortTileSlice(p.board)
}

func (p *Player) PrintBoard() {
	println(fmt.Sprintf("%s's board: ", p.playerName))
	tiles := ""
	for _, tile := range append(p.board, p.okeyTiles...) {
		tiles += " " + tile.ToString()
	}
	println(tiles)

}

func (p *Player) FinishingHandDistance() (int, string) {
	doubles := p.FinishingHandDistanceForDoubles()
	sameNumberGroups := p.findSameNumberGroups()
	println(fmt.Sprintf("\n%s's same number groups:", p.playerName))
	for _, group := range sameNumberGroups {
		print("Group: ")
		groupStr := ""
		for _, tile := range group {
			groupStr += " " + tile.ToString()
		}
		println(groupStr)
	}

	sequentialGroups := p.findSequentialGroups()
	println(fmt.Sprintf("\n%s's sequential groups:", p.playerName))
	for _, group := range sequentialGroups {
		print("Group: ")
		groupStr := ""
		for _, tile := range group {
			groupStr += " " + tile.ToString()
		}
		println(groupStr)
	}

	serial := p.combinedSequentialAndSameNumberGroupsDistance(sequentialGroups, sameNumberGroups)

	if serial > doubles {
		return doubles, "doubles"
	} else {
		return serial, "serial"
	}
}

// seri giderken hem aynı renk sıralıların hem de farklı renk aynı sayıların iç içe bulunduğu bütün alt kümelerin
// skorlarını hesaplayıp minimum'u döndürüyor. alt kümeleri bulurken boyutu 3 ten büyük ve 14ten küçük olarak alıyor
func (p *Player) combinedSequentialAndSameNumberGroupsDistance(seq [][]Tile, same [][]Tile) int {
	topCombinedGroupDistance := 999999
	combined := append(seq, same...)
	allSubsets := powerset(combined)
	for _, subset := range allSubsets {
		dist := p.combinedGroupDistance(subset)
		if dist < topCombinedGroupDistance {
			topCombinedGroupDistance = dist
		}
	}
	return topCombinedGroupDistance
}

func (p *Player) combinedGroupDistance(combinedGroup []Tile) int {
	usedOkeyCount := 0
	totalOverlaps := 0
	totalTileCount := 0
	overlaps := map[string]int{} // tile code -> number of overlaps. example: y3 -> 2
	for _, tile := range combinedGroup {
		totalTileCount += 1
		if tile.Type() == "okey" {
			usedOkeyCount += 1
			continue
		}
		tileCode := fmt.Sprintf("%s%d", tile.Color(), tile.Value())
		if _, ok := overlaps[tileCode]; ok {
			overlaps[tileCode] += 1
			// 2 tane aynı taştan varsa 2yi geçene kadar overlap etmesin diye
			if overlaps[tileCode] > exactTileCount(tile.Color(), tile.Value(), p.board) {
				totalOverlaps += 1
			}
		} else {
			overlaps[tileCode] = 1
		}
	}
	excessOkeyCount := p.okeyCount - usedOkeyCount
	totalDistance := len(p.board) + p.okeyCount - excessOkeyCount + totalOverlaps - totalTileCount
	return totalDistance
}

func exactTileCount(color string, value uint8, group []Tile) int {
	count := 0
	for _, tile := range group {
		if tile.Color() == color && tile.Value() == value {
			count += 1
		}
	}
	return count
}

// FinishingHandDistanceForDoubles taşların sayılarını tutup her 2 şer çiftte uzaklığı 1 azaltıyor.
// okey taşlarını en son uzaklıktan çıkart, sonuçta herhangi bir taşın çifti olabilirler
func (p *Player) FinishingHandDistanceForDoubles() int {
	// 15 tile'a sahip olanda da 7 distance veriyorum 1 tanesini atacak zaten
	dist := 7

	doubleSet := map[string]int{}

	for _, tile := range p.board {
		tileStr := fmt.Sprintf("%s%d", tile.Color(), tile.Value())
		if count, ok := doubleSet[tileStr]; ok {
			if (count+1)%2 == 0 {
				dist -= 1
			}
			doubleSet[tileStr] = count + 1
		} else {
			doubleSet[tileStr] = 1
		}
	}

	newDist := dist - p.okeyCount
	if newDist < 0 {
		return 0
	}
	return newDist
}

// renk gruplarına bölüp her renk grubu için findSequencesInColorGroup çağırır
func (p *Player) findSequentialGroups() [][]Tile {
	sequences := [][]Tile{}
	colorGroups := map[string][]Tile{
		"Yellow": {},
		"Red":    {},
		"Black":  {},
		"Blue":   {},
	}

	for _, tile := range p.board {
		colorGroups[tile.Color()] = append(colorGroups[tile.Color()], tile)
	}
	for _, group := range colorGroups {
		if len(group) > 0 {
			sequences = append(sequences, p.findSequencesInColorGroup(group)...)
		}
	}
	return sequences
}

func (p *Player) findSequencesInColorGroup(tiles_ []Tile) [][]Tile {
	tiles := removeDuplicates(tiles_)
	// 12 13 1 tarzı overflow etmesi için
	if tiles_[0].Value() == 1 {
		tiles = append(tiles, tiles_[0])
	}
	if len(tiles_) > 1 {
		if tiles_[1].Value() == 1 {
			tiles = append(tiles, tiles_[1])
		}
	}
	sequences := [][]Tile{}
	for firstElemIdx := 0; firstElemIdx < len(tiles); firstElemIdx++ {
		currentSeq := []Tile{tiles[firstElemIdx]}
		lastElement := tiles[firstElemIdx]
		remainingOkeyTiles := p.okeyCount

		sequenceBreak := false
		for nextElemIdx := firstElemIdx + 1; nextElemIdx < len(tiles) && !sequenceBreak; nextElemIdx++ {
			switch sequenceDiff(lastElement, tiles[nextElemIdx]) {
			case 0:
				currentSeq = append(currentSeq, tiles[nextElemIdx])
				if len(currentSeq) >= 3 {
					sequences = append(sequences, currentSeq)
				}
				lastElement = tiles[nextElemIdx]
			default:
				if remainingOkeyTiles > 0 {
					currentSeq = append(currentSeq, p.okeyTiles[0])
					if len(currentSeq) >= 3 {
						sequences = append(sequences, currentSeq)
					}
					remainingOkeyTiles -= 1
					lastElement = ColoredTile{color: lastElement.Color(), val: lastElement.Value() + 1}
					nextElemIdx -= 1
				} else {
					sequenceBreak = true
				}
			}
		}

		switch remainingOkeyTiles {
		case 1: // başa 1 tane veya sona 1 tane okey taşı koy
			if len(currentSeq) >= 2 {
				if currentSeq[len(currentSeq)-1].Value() != 13 {
					sequences = append(sequences, append(currentSeq, p.okeyTiles[0]))
				}
				if currentSeq[0].Value() != 1 {
					sequences = append(sequences, append([]Tile{p.okeyTiles[0]}, currentSeq...))
				}
			}
		case 2: // başa 2 tane, sona 2 tane, başa ve sona 1er tane, başa 1 tane, sona 1 tane okey taşı koy
			if currentSeq[0].Value() > 2 {
				sequences = append(sequences, append(p.okeyTiles, currentSeq...))
			}
			if currentSeq[len(currentSeq)-1].Value() < 12 {
				sequences = append(sequences, append(currentSeq, p.okeyTiles...))
			}
			if currentSeq[0].Value() != 1 && currentSeq[len(currentSeq)-1].Value() != 13 {
				sequences = append(sequences, append(append([]Tile{p.okeyTiles[0]}, currentSeq...)), []Tile{p.okeyTiles[1]})
			}
			if len(currentSeq) > 1 {
				if currentSeq[0].Value() != 1 {
					sequences = append(sequences, append([]Tile{p.okeyTiles[0]}, currentSeq...))
				}
				if currentSeq[len(currentSeq)-1].Value() != 13 {
					sequences = append(sequences, append(currentSeq, p.okeyTiles[0]))
				}
			}
		}
	}
	return sequences
}

func removeDuplicates(seq []Tile) []Tile {
	result := []Tile{seq[0]}
	prevTile := seq[0]
	for i := 0; i < len(seq); i++ {
		if seq[i].Value() == prevTile.Value() {
			continue
		}
		result = append(result, seq[i])
		prevTile = seq[i]
	}
	return result
}

func powerset(sl [][]Tile) [][]Tile {
	powersets := [][]Tile{}
	for binarySet := 0; binarySet < int(math.Pow(2, float64(len(sl)))); binarySet++ {
		currentSet := []Tile{}
		for idx := 0; idx < len(sl); idx++ {
			if (binarySet & (1 << idx)) != 0 {
				currentSet = append(currentSet, sl[idx]...)
			}
		}
		if len(currentSet) >= 3 && len(currentSet) <= 14 {
			powersets = append(powersets, currentSet)
		}
	}
	return powersets
}

// powerset'in findSameNumberGroups'ta kullanımı için ayrıca yazdım üsttekini generic şekle getirmek fazla zahmetli
func powersetExtendedTileset(sl []Tile) [][]Tile {
	powersets := [][]Tile{}
	for binarySet := 0; binarySet < int(math.Pow(2, float64(len(sl)))); binarySet++ {
		currentSet := []Tile{}
		for idx := 0; idx < len(sl); idx++ {
			if (binarySet & (1 << idx)) != 0 {
				currentSet = append(currentSet, sl[idx])
			}
		}
		powersets = append(powersets, currentSet)
	}
	return powersets
}

// Kırmızı1-Siyah1-Mavi1
// renkleri value -> []tile şeklinde gruplayıp sonrasında her grubun powerset'ini oluşturup valid olanları groups'a ekler
func (p *Player) findSameNumberGroups() [][]Tile {
	groups := [][]Tile{}
	valuesToUniqueColors := map[uint8][]Tile{}
	for _, tile := range p.board {
		if colorSlice, ok := valuesToUniqueColors[tile.Value()]; ok {
			if !containsColor(colorSlice, tile.Color()) {
				newSlice := append(colorSlice, tile)
				valuesToUniqueColors[tile.Value()] = newSlice
			}
		} else {
			valuesToUniqueColors[tile.Value()] = []Tile{tile}
		}
	}
	for _, uniqueColors := range valuesToUniqueColors {
		subsets := powersetExtendedTileset(uniqueColors)
		for _, subset := range subsets {
			if len(subset) >= 3 && len(subset) < 14 {
				groups = append(groups, subset)
			}
			if (len(subset) == 1 && p.okeyCount == 2) || (len(subset) == 2 && p.okeyCount > 0) {
				groups = append(groups, append(subset, p.okeyTiles...))
			}
			if len(subset) == 3 && p.okeyCount > 0 {
				groups = append(groups, append(subset, p.okeyTiles[0]))
			}
		}

	}

	return groups
}

func containsColor(tileSlice []Tile, color string) bool {
	for _, tile := range tileSlice {
		if color == tile.Color() {
			return true
		}
	}
	return false
}

func sequenceDiff(prevTile Tile, nextTile Tile) int {
	if prevTile.Color() != nextTile.Color() {
		return -1
	}
	if (prevTile.Value() == nextTile.Value()-1) || (prevTile.Value() == 13 && nextTile.Value() == 1) {
		return 0
	}
	return int(nextTile.Value() - prevTile.Value() - 1)
}

var colorOrder = map[string]int{
	"Yellow": 3,
	"Blue":   2,
	"Black":  1,
	"Red":    0,
}

func sortTileSlice(sl []Tile) {
	// renk ve sayıya göre sort
	slices.SortFunc(sl, func(a, b Tile) int {
		if a.Color() == b.Color() {
			if a.Value() < b.Value() {
				return -1
			} else if a.Value() == b.Value() {
				return 0
			} else {
				return 1
			}
		} else {
			if colorOrder[a.Color()] > colorOrder[b.Color()] {
				return -1
			} else {
				return 1
			}
		}
	})
}
