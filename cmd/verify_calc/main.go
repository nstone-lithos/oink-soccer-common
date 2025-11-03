package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("MANUAL CALCULATION VERIFICATION")
	fmt.Println("=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=")
	fmt.Println()

	// Vardy
	fmt.Println("J. Vardy (CTL:70 ATT:84 DEF:52 SPD:85)")
	vardyCtl := (70.0*4 + 85) / 5
	vardyAtt := (84.0*3 + 85) / 4
	vardyDef := (52.0*5 + 85) / 6
	fmt.Printf("  Control: (70*4 + 85)/5 = %.2f → rounds to %.0f\n", vardyCtl, math.Round(vardyCtl))
	fmt.Printf("  Attack:  (84*3 + 85)/4 = %.2f → rounds to %.0f\n", vardyAtt, math.Round(vardyAtt))
	fmt.Printf("  Defense: (52*5 + 85)/6 = %.2f → rounds to %.0f\n", vardyDef, math.Round(vardyDef))
	fmt.Println()

	// Iličić
	fmt.Println("J. Iličić (CTL:84 ATT:86 DEF:40 SPD:75)")
	ilicicCtl := (84.0*4 + 75) / 5
	ilicicAtt := (86.0*3 + 75) / 4
	ilicicDef := (40.0*5 + 75) / 6
	fmt.Printf("  Control: (84*4 + 75)/5 = %.2f → rounds to %.0f\n", ilicicCtl, math.Round(ilicicCtl))
	fmt.Printf("  Attack:  (86*3 + 75)/4 = %.2f → rounds to %.0f\n", ilicicAtt, math.Round(ilicicAtt))
	fmt.Printf("  Defense: (40*5 + 75)/6 = %.2f → rounds to %.0f\n", ilicicDef, math.Round(ilicicDef))
	fmt.Println()

	// Goretzka
	fmt.Println("L. Goretzka (CTL:82 ATT:82 DEF:81 SPD:81)")
	goretzkaCtl := (82.0*4 + 81) / 5
	goretzkaAtt := (82.0*3 + 81) / 4
	goretzkaDef := (81.0*5 + 81) / 6
	fmt.Printf("  Control: (82*4 + 81)/5 = %.2f → rounds to %.0f\n", goretzkaCtl, math.Round(goretzkaCtl))
	fmt.Printf("  Attack:  (82*3 + 81)/4 = %.2f → rounds to %.0f\n", goretzkaAtt, math.Round(goretzkaAtt))
	fmt.Printf("  Defense: (81*5 + 81)/6 = %.2f → rounds to %.0f\n", goretzkaDef, math.Round(goretzkaDef))
	fmt.Println()

	// Robertson
	fmt.Println("A. Robertson (CTL:81 ATT:61 DEF:82 SPD:84)")
	robertsonCtl := (81.0*4 + 84) / 5
	robertsonAtt := (61.0*3 + 84) / 4
	robertsonDef := (82.0*5 + 84) / 6
	fmt.Printf("  Control: (81*4 + 84)/5 = %.2f → rounds to %.0f\n", robertsonCtl, math.Round(robertsonCtl))
	fmt.Printf("  Attack:  (61*3 + 84)/4 = %.2f → rounds to %.0f\n", robertsonAtt, math.Round(robertsonAtt))
	fmt.Printf("  Defense: (82*5 + 84)/6 = %.2f → rounds to %.0f\n", robertsonDef, math.Round(robertsonDef))
	fmt.Println()

	// Scaling verification
	fmt.Println("\nSCALING VALUES FROM CSV:")
	scalingMap := map[int]int{
		46: 4,
		58: 8,
		67: 22,
		73: 30,
		81: 53,
		82: 58,
		83: 60,
		84: 61,
	}

	for rating, scaled := range scalingMap {
		fmt.Printf("  %d → %d\n", rating, scaled)
	}
	fmt.Println()

	// Summary
	fmt.Println("COMPARISON SUMMARY:")
	fmt.Println("-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-")
	fmt.Printf("Vardy Control:   %.0f → %d (scaled)\n", math.Round(vardyCtl), scalingMap[int(math.Round(vardyCtl))])
	fmt.Printf("Iličić Control:  %.0f → %d (scaled)\n", math.Round(ilicicCtl), scalingMap[int(math.Round(ilicicCtl))])
	fmt.Printf("Difference:      +%.0f → +%d (scaled)\n",
		math.Round(ilicicCtl)-math.Round(vardyCtl),
		scalingMap[int(math.Round(ilicicCtl))]-scalingMap[int(math.Round(vardyCtl))])
	fmt.Println()

	fmt.Printf("Vardy Attack:    %.0f → %d (scaled)\n", math.Round(vardyAtt), scalingMap[int(math.Round(vardyAtt))])
	fmt.Printf("Iličić Attack:   %.0f → %d (scaled)\n", math.Round(ilicicAtt), scalingMap[int(math.Round(ilicicAtt))])
	fmt.Printf("Difference:      %.0f → %d (scaled)\n",
		math.Round(ilicicAtt)-math.Round(vardyAtt),
		scalingMap[int(math.Round(ilicicAtt))]-scalingMap[int(math.Round(vardyAtt))])
}
