// Declare a new struct type to hold information about a tennis player,
// including the number of matches played and the number won.  Add a method to
// this type that calculates the win ratio for the player.  Create a new
// player, and output the win ratio for them.
//
// Template available at: http://play.golang.org/p/jnBw-jtE3n
package main

// Add your imports here.
import "fmt"

// Declare a struct type `player` to maintain information about a player.
type player struct {
	name string
	numMatchesPlayed int
	numMatchesWon int
}

// Declare a method that calculates the win ratio for the player.  Note that
// you'll likely need to convert one or more values to floats, which can be
// done like: float32(intValue)
func (p player) winRatio() float32 {
	return float32(p.numMatchesWon) / float32(p.numMatchesPlayed)
}

func main() {
	// Create a new player, and output their win ratio.
	p1 := player{"Mike", 100, 32}
	winRatio := p1.winRatio()
	fmt.Printf("%s: %f\n", p1.name, winRatio)

	// If you're feeling adventurous, try creating a slice of multiple players
	// and iterating over the slice, displaying the player name and win ratio.
	players := []player{{"Melissa", 150, 93}, {"Edward", 84, 32}}
	for _, player := range players {
		fmt.Printf("%s: %f\n", player.name, player.winRatio())
	}
}
