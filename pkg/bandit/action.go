// bandit implementation

package bandit

import (
	//"fmt"
	//"math"
	//"time"

	//"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

// An specific possible action - with a reward which is normally distributed
// here the reward distribution is stationary - ie , not uncertain
type Action struct {
	mean   float64
	vari   float64
	normal *distuv.Normal

	// number of times this action was selected
	N int
	// estimate of the reward from the past actions
	Q float64
	// preference for this action in the gradient bandit
	H float64
	// probability of choosing this action in gradient bandit
	P float64
}

func NewAction(mean, vari float64) *Action {
	a := Action{
		mean: mean,
		vari: vari,
		normal: &distuv.Normal{
			Mu:    mean,
			Sigma: vari,
			Src:   randSrc,
		},
	}
	a.N = 0
	//a.Q = a.Reward()
	a.Q = 0
	return &a
}

// do an action - which returns a reward
func (a *Action) Reward() float64 {
	return a.normal.Rand()
}

// update the reward estimate
func (a *Action) updateEstimate(r float64) {
	a.N++
	a.Q = a.Q + (r-a.Q)/float64(a.N)
}

// create a list of actions
func NewActions(numActions int, mean, vari float64, actionVari float64) []*Action {
	actions := []*Action{}
	meanSelect := distuv.Normal{
		Mu:    mean,
		Sigma: vari,
		Src:   randSrc,
	}
	for i := 0; i < numActions; i++ {
		actionMean := meanSelect.Rand()
		action := NewAction(actionMean, actionVari)
		actions = append(actions, action)
	}
	return actions
}
