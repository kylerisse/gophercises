package blackjack

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// Cli for running a game
type Cli struct {
	game *game
}

// NewCli creates a new game
func NewCli(numDecks int) *Cli {
	return &Cli{
		game: newGame(numDecks),
	}
}

// Run the game
func (c *Cli) Run() {
	c.game.run()
}

func (g *game) run() {
	for {
		if g.shoe.shuffleTime {
			g.printShuffle()
		}
		g.nextHand()
		if g.checkNaturalBlackjack() {
			continue
		}
		if g.playerTurn() {
			continue
		}
		g.dealerTurn()
	}
}

func (g *game) dealerTurn() {
	for {
		if g.dealer.points() > 16 {
			break
		}
		g.dealer.add(g.shoe.draw())
	}
	if g.dealer.points() > 21 {
		fmt.Println("DEALER BUST")
		g.win()
		return
	}
	if g.player.points() > g.dealer.points() {
		g.win()
		return
	}
	if g.dealer.points() == g.player.points() {
		g.push()
		return
	}
	g.lose()
}

func (g *game) playerTurn() bool {
	for {
		if g.player.points() >= 21 {
			break
		}
		g.printDealerHidden()
		g.printPlayer()
		fmt.Printf(" [H]it  [S]tay  [Q]uit  # ")
		reader := bufio.NewReader(os.Stdin)
		a, _ := reader.ReadString('\n')
		b := strings.ToUpper(a)[0]
		if b == 'H' {
			g.player.add(g.shoe.draw())
			continue
		}
		if b == 'S' {
			break
		}
		if b == 'Q' {
			g.lose()
			os.Exit(0)
		}
		fmt.Println("**Invalid** " + string(b))
	}
	if g.player.points() > 21 {
		fmt.Println("PLAYER BUST")
		g.lose()
		return true
	}
	return false
}

func (g *game) checkNaturalBlackjack() bool {
	if g.dealer.points() == 21 {
		fmt.Println("DEALER BLACKJACK")
		if g.player.points() == 21 {
			g.push()
			return true
		}
		g.lose()
		return true
	}
	if g.player.points() == 21 {
		fmt.Println("PLAYER BLACKJACK")
		g.win()
		return true
	}
	return false
}

func (g *game) push() {
	g.printDealer()
	g.printPlayer()
	fmt.Println(" PUSH")
	g.pushes++
	g.printScore()
	time.Sleep(3 * time.Second)
}

func (g *game) lose() {
	g.printDealer()
	g.printPlayer()
	fmt.Println(" LOSE")
	g.losses++
	g.printScore()
	time.Sleep(3 * time.Second)
}

func (g *game) win() {
	g.printDealer()
	g.printPlayer()
	fmt.Println(" WIN")
	g.wins++
	g.printScore()
	time.Sleep(1 * time.Second)
}

func (g *game) printScore() {
	time.Sleep(1 * time.Second)
	fmt.Printf(" ----  Wins: %d  Losses: %d  Pushes: %d\n",
		g.wins, g.losses, g.pushes)
}

func (g *game) printShuffle() {
	fmt.Println(" - shuffling shoe - ")
	g.shoe = *newShoe(g.numDecks)
	time.Sleep(1 * time.Second)
	fmt.Printf(" %v %v %v \n", cardBack, cardBack, cardBack)
	fmt.Println(" - done - ")
}

func (g *game) printDealerHidden() {
	fmt.Printf("Dealer: %v %v\n", g.dealer.cards[0].String(), cardBack)
}

func (g *game) printDealer() {
	fmt.Printf("Dealer: %v (%d)\n", g.dealer.String(), g.dealer.points())
}

func (g *game) printPlayer() {
	fmt.Printf("Player: %v (%d)\n", g.player.String(), g.player.points())
}
