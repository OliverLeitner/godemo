package chucknorris

import (
	"github.com/gtl-hig/chuckjoke"
)

// one can defined the class and then below "insert" the functs in go...
type Export struct {
}

// all func names start with uppercase can be used outside...
// wtf
func (c Export) GetChuck() string {
	var mychuckquote, mychuckerror = chuckjoke.GetNextJoke()
	if mychuckerror == nil {
		return mychuckquote.Value
	}
	return "an error occured"
}
