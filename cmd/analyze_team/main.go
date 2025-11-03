package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	soccer "github.com/stein-f/oink-soccer-common"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	randSource := rand.New(source)

	// Current Y Formation Team
	currentTeam := soccer.GameLineup{
		Team: soccer.Team{
			ID:         "current",
			CustomName: "Current Y",
			Formation:  soccer.FormationTypeY,
		},
		Players: []soccer.SelectedPlayer{
			// GK
			{
				ID:   "casteels",
				Name: "K. Casteels",
				Attributes: soccer.PlayerAttributes{
					GoalkeeperRating: 83,
					DefenseRating:    14,
					ControlRating:    80,
					AttackRating:     0,
					SpeedRating:      86,
					Position:         soccer.PlayerPositionGoalkeeper,
				},
				SelectedPosition: soccer.PlayerPositionGoalkeeper,
			},
			// DEF
			{
				ID:   "goretzka",
				Name: "L. Goretzka",
				Attributes: soccer.PlayerAttributes{
					GoalkeeperRating: 0,
					DefenseRating:    81,
					ControlRating:    82,
					AttackRating:     82,
					SpeedRating:      81,
					Position:         soccer.PlayerPositionDefense,
				},
				SelectedPosition: soccer.PlayerPositionDefense,
			},
			// MID
			{
				ID:   "cruyff",
				Name: "Cruyff",
				Attributes: soccer.PlayerAttributes{
					GoalkeeperRating: 0,
					DefenseRating:    52,
					ControlRating:    88,
					AttackRating:     87,
					SpeedRating:      91,
					Position:         soccer.PlayerPositionMidfield,
				},
				SelectedPosition: soccer.PlayerPositionMidfield,
			},
			// ATT
			{
				ID:   "benyedder",
				Name: "W. Ben Yedder",
				Attributes: soccer.PlayerAttributes{
					GoalkeeperRating: 0,
					DefenseRating:    39,
					ControlRating:    77,
					AttackRating:     84,
					SpeedRating:      82,
					Position:         soccer.PlayerPositionAttack,
				},
				SelectedPosition: soccer.PlayerPositionAttack,
			},
			// ATT
			{
				ID:   "vardy",
				Name: "J. Vardy",
				Attributes: soccer.PlayerAttributes{
					GoalkeeperRating: 0,
					DefenseRating:    52,
					ControlRating:    70,
					AttackRating:     84,
					SpeedRating:      85,
					Position:         soccer.PlayerPositionAttack,
				},
				SelectedPosition: soccer.PlayerPositionAttack,
			},
		},
	}

	// Optimized Y Formation Team
	optimizedTeam := soccer.GameLineup{
		Team: soccer.Team{
			ID:         "optimized",
			CustomName: "Optimized Y",
			Formation:  soccer.FormationTypeY,
		},
		Players: []soccer.SelectedPlayer{
			// GK - Keep Casteels
			{
				ID:   "casteels",
				Name: "K. Casteels",
				Attributes: soccer.PlayerAttributes{
					GoalkeeperRating: 83,
					DefenseRating:    14,
					ControlRating:    80,
					AttackRating:     0,
					SpeedRating:      86,
					Position:         soccer.PlayerPositionGoalkeeper,
				},
				SelectedPosition: soccer.PlayerPositionGoalkeeper,
			},
			// DEF - A. Robertson (upgrade)
			{
				ID:   "robertson",
				Name: "A. Robertson",
				Attributes: soccer.PlayerAttributes{
					GoalkeeperRating: 0,
					DefenseRating:    82,
					ControlRating:    81,
					AttackRating:     61,
					SpeedRating:      84,
					Position:         soccer.PlayerPositionDefense,
				},
				SelectedPosition: soccer.PlayerPositionDefense,
			},
			// MID - Keep Cruyff (elite midfielder)
			{
				ID:   "cruyff",
				Name: "Cruyff",
				Attributes: soccer.PlayerAttributes{
					GoalkeeperRating: 0,
					DefenseRating:    52,
					ControlRating:    88,
					AttackRating:     87,
					SpeedRating:      91,
					Position:         soccer.PlayerPositionMidfield,
				},
				SelectedPosition: soccer.PlayerPositionMidfield,
			},
			// ATT - Keep Ben Yedder
			{
				ID:   "benyedder",
				Name: "W. Ben Yedder",
				Attributes: soccer.PlayerAttributes{
					GoalkeeperRating: 0,
					DefenseRating:    39,
					ControlRating:    77,
					AttackRating:     84,
					SpeedRating:      82,
					Position:         soccer.PlayerPositionAttack,
				},
				SelectedPosition: soccer.PlayerPositionAttack,
			},
			// ATT - J. Iličić (upgrade for control)
			{
				ID:   "ilicic",
				Name: "J. Iličić",
				Attributes: soccer.PlayerAttributes{
					GoalkeeperRating: 0,
					DefenseRating:    40,
					ControlRating:    84,
					AttackRating:     86,
					SpeedRating:      75,
					Position:         soccer.PlayerPositionAttack,
				},
				SelectedPosition: soccer.PlayerPositionAttack,
			},
		},
	}

	fmt.Println(strings.Repeat("=", 72))
	fmt.Println("TEAM COMPARISON - Y Formation (1-1-2)")
	fmt.Println(strings.Repeat("=", 72))

	analyzeTeam("CURRENT TEAM", currentTeam, randSource)
	fmt.Println()
	analyzeTeam("OPTIMIZED TEAM", optimizedTeam, randSource)
}

func analyzeTeam(name string, lineup soccer.GameLineup, randSource *rand.Rand) {
	controlScore := soccer.CalculateTeamControlScore(randSource, lineup)
	defenseScore := soccer.CalculateTeamDefenseScore(randSource, lineup)

	fmt.Printf("\n%s:\n", name)
	fmt.Println(strings.Repeat("-", 72))

	// List players
	for _, player := range lineup.Players {
		fmt.Printf("%-3s | %-15s | CTL:%2d ATT:%2d DEF:%2d SPD:%2d\n",
			player.SelectedPosition,
			player.Name,
			player.Attributes.ControlRating,
			player.Attributes.AttackRating,
			player.Attributes.DefenseRating,
			player.Attributes.SpeedRating)
	}

	fmt.Println(strings.Repeat("-", 72))
	fmt.Printf("TEAM CONTROL SCORE:  %.2f\n", controlScore)
	fmt.Printf("TEAM DEFENSE SCORE:  %.2f\n", defenseScore)
	fmt.Printf("COMBINED RATING:     %.2f\n", (controlScore+defenseScore)/2)
}
