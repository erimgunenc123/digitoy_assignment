package main

import "math/rand"

type Game struct {
	tileSet   [105]Tile
	players   [4]*Player
	indicator Tile
}

func NewGame(players [4]*Player) *Game {
	tileSet, indicator := NewTileSet()
	return &Game{
		tileSet:   tileSet,
		players:   players,
		indicator: indicator,
	}
}

func (g *Game) PrintBoardStates() {
	for _, player := range g.players {
		player.PrintBoard()
	}
}

func (g *Game) DistributeTilesToPlayers() {
	firstPlayerIdx := rand.Intn(4)
	tempTileSet := g.tileSet[:]
	for idx, player := range g.players {
		playerBoard := []Tile{}
		tileCount := 14
		if idx == firstPlayerIdx {
			tileCount = 15
		}
		for i := 0; i < tileCount; i++ {
			curTileIdx := rand.Intn(len(tempTileSet))
			playerBoard = append(playerBoard, tempTileSet[curTileIdx])
			tempTileSet = append(tempTileSet[:curTileIdx], tempTileSet[curTileIdx+1:]...)
		}
		player.SetBoard(playerBoard)
	}

}
