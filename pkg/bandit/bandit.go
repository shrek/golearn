// bandit implementation

package bandit

import (
	"fmt"
	"math"
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

var (
	seed    = time.Now().UTC().Unix()
	randSrc = rand.NewSource(uint64(seed))
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
	return nil
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
		fmt.Println("success\n")
	}

}

func Run(numRuns int, alpha float64) {
	GreedyVsEpsilon(numRuns, alpha)
}

func init() {
	rand.Seed(uint64(seed))
}
