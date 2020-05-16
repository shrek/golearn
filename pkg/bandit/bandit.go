// bandit implementation

package bandit

import (
	"fmt"
	//"math"
	"time"

	"golang.org/x/exp/rand"
	//"gonum.org/v1/gonum/stat/distuv"
)

var (
	seed    = time.Now().UTC().Unix()
	randSrc = rand.NewSource(uint64(seed))
)

// compare greedy with epsilon
func GreedyVsEpsilon(numRuns int, alpha float64) {
	// create a set of actions

	actions := NewActions(10, float64(0), float64(2), float64(1))
	maxAct := float64(-100)
	maxActInd := 0
	for i, a := range actions {
		fmt.Printf("%d - mean: %f, variance: %f\n", i, a.mean, a.vari)
		if maxAct < a.mean {
			maxAct = a.mean
			maxActInd = i
		}
	}
	fmt.Printf("best action: %d, %f\n", maxActInd, maxAct)
	greedy := NewChoice(actions, float64(0))
	explore1 := NewChoice(actions, float64(0.1))
	explore01 := NewChoice(actions, float64(0.01))
	ucb := NewChoice(actions, float64(0))
	ucb.useUCB = true
	ucb.ucbC = float64(2)
	gradient := NewChoice(actions, float64(0))
	gradient.useGradient = true
	gradient.stepSize = alpha

	greedy.Run(numRuns)
	explore1.Run(numRuns)
	explore01.Run(numRuns)
	ucb.Run(numRuns)
	gradient.Run(numRuns)

	//fmt.Println("%v", greedy.Rewards)
	//fmt.Println("%v", greedy.AvRewards)

	names := []string{"greedy", "e=0.1", "e=0.01", "ucb,c=2", "grad,a=0.1"}

	if err := plotSeries("/tmp/greedy.png", names, greedy.AvRewards, explore1.AvRewards, explore01.AvRewards, ucb.AvRewards, gradient.AvRewards); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("success")
	}

}

func Run(numRuns int, alpha float64) {
	GreedyVsEpsilon(numRuns, alpha)
}

func init() {
	rand.Seed(uint64(seed))
}
