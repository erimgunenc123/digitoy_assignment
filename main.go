package main

import "fmt"

func main() {
	player1 := NewPlayer("bob")
	player2 := NewPlayer("alice")
	player3 := NewPlayer("mallory")
	player4 := NewPlayer("oscar")

	g := NewGame([4]*Player{player1, player2, player3, player4})
	g.DistributeTilesToPlayers()
	g.PrintBoardStates()
	minDist := 99999
	var distMethod string
	var topPlayer *Player
	for _, player := range []*Player{player1, player2, player3, player4} {
		dist, method := player.FinishingHandDistance()
		if dist < minDist {
			minDist = dist
			topPlayer = player
			distMethod = method
		}
	}
	println(fmt.Sprintf("\nClosest finishing board holder: %s  Distance: %d Doubles/Serial: %s", topPlayer.playerName, minDist, distMethod))
	println("*********Board*******")
	topPlayer.PrintBoard()

}
