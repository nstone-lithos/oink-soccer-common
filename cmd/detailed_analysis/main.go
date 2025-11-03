package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	soccer "github.com/stein-f/oink-soccer-common"
)

type PlayerComparison struct {
	Name          string
	Position      string
	ControlRating int
	AttackRating  int
	DefenseRating int
	SpeedRating   int
	ControlScore  float64
	AttackScore   float64
	DefenseScore  float64
}

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	randSource := rand.New(source)

	fmt.Println("========================================================================")
	fmt.Println("DETAILED PLAYER COMPARISON WITH SPEED & SCALING")
	fmt.Println("========================================================================\n")

	// Compare Defenders
	fmt.Println("DEFENDER COMPARISON:")
	fmt.Println("------------------------------------------------------------------------")
	goretzka := analyzePlayer("L. Goretzka", "Defense", 82, 82, 81, 81, soccer.PlayerPositionDefense)
	robertson := analyzePlayer("A. Robertson", "Defense", 81, 61, 82, 84, soccer.PlayerPositionDefense)
	printPlayerComparison(goretzka, robertson)

	// Compare Attackers
	fmt.Println("\nATTACKER COMPARISON:")
	fmt.Println("------------------------------------------------------------------------")
	vardy := analyzePlayer("J. Vardy", "Attack", 70, 84, 52, 85, soccer.PlayerPositionAttack)
	ilicic := analyzePlayer("J. Iličić", "Attack", 84, 86, 40, 75, soccer.PlayerPositionAttack)
	printPlayerComparison(vardy, ilicic)

	// Now run full team comparison
	fmt.Println("\n\n========================================================================")
	fmt.Println("FULL TEAM ANALYSIS")
	fmt.Println("========================================================================\n")

	currentTeam := createCurrentTeam()
	optimizedTeam := createOptimizedTeam()

	analyzeFullTeam("CURRENT TEAM", currentTeam, randSource)
	fmt.Println()
	analyzeFullTeam("OPTIMIZED TEAM", optimizedTeam, randSource)
}

func analyzePlayer(name, pos string, ctl, att, def, spd int, position soccer.PlayerPosition) PlayerComparison {
	player := soccer.PlayerAttributes{
		ControlRating: ctl,
		AttackRating:  att,
		DefenseRating: def,
		SpeedRating:   spd,
		Position:      position,
	}

	return PlayerComparison{
		Name:          name,
		Position:      pos,
		ControlRating: ctl,
		AttackRating:  att,
		DefenseRating: def,
		SpeedRating:   spd,
		ControlScore:  player.GetControlScore(),
		AttackScore:   player.GetAttackScore(),
		DefenseScore:  player.GetDefenseScore(),
	}
}

func printPlayerComparison(p1, p2 PlayerComparison) {
	fmt.Printf("%-15s | CTL:%2d ATT:%2d DEF:%2d SPD:%2d\n",
		p1.Name, p1.ControlRating, p1.AttackRating, p1.DefenseRating, p1.SpeedRating)
	fmt.Printf("%-15s | CTL:%2d ATT:%2d DEF:%2d SPD:%2d\n",
		p2.Name, p2.ControlRating, p2.AttackRating, p2.DefenseRating, p2.SpeedRating)
	fmt.Println()

	// Calculate with speed
	fmt.Printf("With Speed Factored In:\n")
	fmt.Printf("%-15s | Control Score: %.0f  Attack Score: %.0f  Defense Score: %.0f\n",
		p1.Name, p1.ControlScore, p1.AttackScore, p1.DefenseScore)
	fmt.Printf("%-15s | Control Score: %.0f  Attack Score: %.0f  Defense Score: %.0f\n",
		p2.Name, p2.ControlScore, p2.AttackScore, p2.DefenseScore)
	fmt.Println()

	// Show scaling
	fmt.Printf("After Scaling Function:\n")
	fmt.Printf("%-15s | Control: %.0f→%.0f  Attack: %.0f→%.0f  Defense: %.0f→%.0f\n",
		p1.Name, p1.ControlScore, soccer.ScalingFunction(p1.ControlScore),
		p1.AttackScore, soccer.ScalingFunction(p1.AttackScore),
		p1.DefenseScore, soccer.ScalingFunction(p1.DefenseScore))
	fmt.Printf("%-15s | Control: %.0f→%.0f  Attack: %.0f→%.0f  Defense: %.0f→%.0f\n",
		p2.Name, p2.ControlScore, soccer.ScalingFunction(p2.ControlScore),
		p2.AttackScore, soccer.ScalingFunction(p2.AttackScore),
		p2.DefenseScore, soccer.ScalingFunction(p2.DefenseScore))
	fmt.Println()

	// Calculate differences
	ctlDiff := p2.ControlScore - p1.ControlScore
	attDiff := p2.AttackScore - p1.AttackScore
	defDiff := p2.DefenseScore - p1.DefenseScore

	ctlScaledDiff := soccer.ScalingFunction(p2.ControlScore) - soccer.ScalingFunction(p1.ControlScore)
	attScaledDiff := soccer.ScalingFunction(p2.AttackScore) - soccer.ScalingFunction(p1.AttackScore)
	defScaledDiff := soccer.ScalingFunction(p2.DefenseScore) - soccer.ScalingFunction(p1.DefenseScore)

	fmt.Printf("Differences (%s vs %s):\n", p2.Name, p1.Name)
	fmt.Printf("  Control: %+.0f (scaled: %+.0f)\n", ctlDiff, ctlScaledDiff)
	fmt.Printf("  Attack:  %+.0f (scaled: %+.0f)\n", attDiff, attScaledDiff)
	fmt.Printf("  Defense: %+.0f (scaled: %+.0f)\n", defDiff, defScaledDiff)
}

func analyzeFullTeam(name string, lineup soccer.GameLineup, randSource *rand.Rand) {
	controlScore := soccer.CalculateTeamControlScore(randSource, lineup)
	defenseScore := soccer.CalculateTeamDefenseScore(randSource, lineup)

	fmt.Printf("%s:\n", name)
	fmt.Println("------------------------------------------------------------------------")
	for _, player := range lineup.Players {
		attrs := player.Attributes
		selected := player.SelectedPosition

		// Calculate actual scores
		ctlScore := attrs.GetControlScore()
		attScore := attrs.GetAttackScore()
		defScore := attrs.GetDefenseScore()

		fmt.Printf("%-3s | %-15s | CTL:%2d→%.0f ATT:%2d→%.0f DEF:%2d→%.0f SPD:%2d\n",
			selected,
			player.Name,
			attrs.ControlRating, ctlScore,
			attrs.AttackRating, attScore,
			attrs.DefenseRating, defScore,
			attrs.SpeedRating)
	}
	fmt.Println("------------------------------------------------------------------------")
	fmt.Printf("TEAM CONTROL SCORE:  %.2f\n", controlScore)
	fmt.Printf("TEAM DEFENSE SCORE:  %.2f\n", defenseScore)
	fmt.Printf("COMBINED RATING:     %.2f\n", (controlScore+defenseScore)/2)
}

func createCurrentTeam() soccer.GameLineup {
	return soccer.GameLineup{
		Team: soccer.Team{
			ID:         "current",
			CustomName: "Current Y",
			Formation:  soccer.FormationTypeY,
		},
		Players: []soccer.SelectedPlayer{
			createPlayer("casteels", "K. Casteels", 80, 0, 14, 86, 83, soccer.PlayerPositionGoalkeeper),
			createPlayer("goretzka", "L. Goretzka", 82, 82, 81, 81, 0, soccer.PlayerPositionDefense),
			createPlayer("cruyff", "Cruyff", 88, 87, 52, 91, 0, soccer.PlayerPositionMidfield),
			createPlayer("benyedder", "W. Ben Yedder", 77, 84, 39, 82, 0, soccer.PlayerPositionAttack),
			createPlayer("vardy", "J. Vardy", 70, 84, 52, 85, 0, soccer.PlayerPositionAttack),
		},
	}
}

func createOptimizedTeam() soccer.GameLineup {
	return soccer.GameLineup{
		Team: soccer.Team{
			ID:         "optimized",
			CustomName: "Optimized Y",
			Formation:  soccer.FormationTypeY,
		},
		Players: []soccer.SelectedPlayer{
			createPlayer("casteels", "K. Casteels", 80, 0, 14, 86, 83, soccer.PlayerPositionGoalkeeper),
			createPlayer("robertson", "A. Robertson", 81, 61, 82, 84, 0, soccer.PlayerPositionDefense),
			createPlayer("cruyff", "Cruyff", 88, 87, 52, 91, 0, soccer.PlayerPositionMidfield),
			createPlayer("benyedder", "W. Ben Yedder", 77, 84, 39, 82, 0, soccer.PlayerPositionAttack),
			createPlayer("ilicic", "J. Iličić", 84, 86, 40, 75, 0, soccer.PlayerPositionAttack),
		},
	}
}

func createPlayer(id, name string, ctl, att, def, spd, gkp int, pos soccer.PlayerPosition) soccer.SelectedPlayer {
	return soccer.SelectedPlayer{
		ID:   id,
		Name: name,
		Attributes: soccer.PlayerAttributes{
			GoalkeeperRating: gkp,
			ControlRating:    ctl,
			AttackRating:     att,
			DefenseRating:    def,
			SpeedRating:      spd,
			Position:         pos,
		},
		SelectedPosition: pos,
	}
}

func round(x float64) float64 {
	return math.Round(x)
}
