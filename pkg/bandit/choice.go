package bandit

import (
	//"fmt"
	"math"
	//"time"

	"golang.org/x/exp/rand"
	//"gonum.org/v1/gonum/stat/distuv"
)

// a Choice is a set of actions from which one is chosen
type Choice struct {
	Actions []*Action
	// rewards received so far
	Rewards []float64
	// average of reward received so far
	AvRewards []float64
	// sum of rewards received so far
	rSum float64
	// epsilon - exploration likelihood
	Epsilon float64
	// Upper Confidence Bound - c
	ucbC float64
	// use UCB ?
	useUCB bool
	// use gradient
	useGradient bool
	stepSize    float64 // alpha in gradient
}

func NewChoice(actions []*Action, epsilon float64) *Choice {
	c := Choice{}
	c.Actions = actions
	c.Epsilon = epsilon
	return &c
}

// get the action with the highest reward found so far
func (c *Choice) greedyAction() *Action {
	max := -1 * math.MaxFloat32
	var maxAct *Action
	for _, a := range c.Actions {
		if max < a.Q {
			max = a.Q
			maxAct = a
		}
	}
	return maxAct
}

// get the action according to Upper Confidence Bound
func (c *Choice) ucbAction() *Action {
	max := -1 * math.MaxFloat32
	var maxAct *Action
	t := float64(len(c.Rewards))
	for _, a := range c.Actions {
		if a.N == 0 {
			return a
		}
		ucb := a.Q + c.ucbC*math.Sqrt(math.Log(t)/float64(a.N))
		if max < ucb {
			max = ucb
			maxAct = a
		}
	}
	return maxAct
}

// get a random action
func (c *Choice) randomAction() *Action {
	r := rand.Intn(len(c.Actions))
	return c.Actions[r]
}

func (c *Choice) gradientSelectAction() *Action {
	r := rand.Float64()
	sum := float64(0)
	for i, a := range c.Actions {
		if i == len(c.Actions)-1 {
			return a
		}
		if sum < r && r <= sum+a.P {
			return a
		}
		sum = sum + a.P
	}
	panic("shouldnt hit this")
}

func (c *Choice) selectAction() *Action {
	if c.useGradient {
		return c.gradientSelectAction()
	}

	// prob e choose random
	// 1-e choose greedy
	r := rand.Intn(10000) + 1
	if float64(r)/float64(10000) <= c.Epsilon {
		// choose random action
		if c.useUCB {
			return c.ucbAction()
		} else {
			return c.randomAction()
		}
	}
	// choose greedy action
	return c.greedyAction()
}

// gradient bandit
func (c *Choice) updateProbabilities() {
	if !c.useGradient {
		return
	}
	d := float64(0)
	for _, a := range c.Actions {
		d = d + math.Exp(a.H)
	}
	for _, a := range c.Actions {
		a.P = math.Exp(a.H) / d
	}
}

func (c *Choice) updatePreferences(selected *Action, r, rAv float64) {
	if !c.useGradient {
		return
	}
	selected.H = selected.H + c.stepSize*(r-rAv)*(1-selected.P)
	for _, a := range c.Actions {
		if a != selected {
			a.H = a.H - c.stepSize*(r-rAv)*a.P
		}
	}
}

// do a single step
func (c *Choice) Do() {
	// for gradient bandit, update action prob
	c.updateProbabilities()
	// select action
	a := c.selectAction()
	// do the action
	r := a.Reward()
	oldN := len(c.Rewards)
	// record result and update the estimates
	c.Rewards = append(c.Rewards, r)
	a.updateEstimate(r)
	oldSum := c.rSum
	oldAv := oldSum / float64(oldN)
	c.updatePreferences(a, r, oldAv)
	c.rSum += r
	c.AvRewards = append(c.AvRewards, c.rSum/float64(len(c.Rewards)))
}

// do k steps
func (c *Choice) Run(k int) {
	for i := 0; i < k; i++ {
		c.Do()
	}
}
