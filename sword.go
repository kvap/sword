// ---------------------------------------------------------------------------
// Constantin S. Pan <kvapen@gmail.com> wrote this file. As long as you retain
// this notice you can do whatever you want with this stuff. If we meet some
// day, and you think this stuff is worth it, you can buy me a bottle of cider
// in return.
// 	Constantin S. Pan
// ---------------------------------------------------------------------------

// The implementation is based on the ideas proposed in
// https://stackoverflow.com/questions/8870261/how-to-split-text-without-spaces-into-list-of-words

package sword

import (
	"os"
	"bufio"
	"math"
)

func readLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := bufio.NewScanner(f)
	var lines []string
	for reader.Scan() {
		lines = append(lines, reader.Text())
	}

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// The one and only cutting tool in this package.
type Blade struct {
	words []string
	costs map[string]float64
	maxWord int
}

// Train loads the dictionary of words from file named filename.
// The file should contain the words one per line, most frequent first.
func (d *Blade) Train(filename string) (err error) {
	d.words, err = readLines(filename)
	if err != nil {
		return
	}
	d.costs = make(map[string]float64)
	d.maxWord = 0

	for i, word := range d.words {
		d.costs[word] = math.Log(float64(i + 1) * math.Log(float64(len(d.words))))
		if len(word) > d.maxWord {
			d.maxWord = len(word)
		}
	}
	return nil
}

func (d *Blade) cost(word string) float64 {
	if cost, ok := d.costs[word]; ok {
		return cost
	} else {
		return math.Inf(+1)
	}
}

// Carve cuts the given string s into pieces.
// Only use Carve after Train-ing, or bad things will happen.
func (d *Blade) Carve(s string) []string {
	type solution struct {
		cost float64
		lastlen int
	}

	solutions := []solution{{0, 0}}

	bestSolution := func(i int) (sol solution) {
		sol.cost = math.Inf(+1)
		sol.lastlen = 0

		// try to fit words of length 1 to d.MaxWord
		for suflen := 1; suflen <= i && suflen <= d.maxWord; suflen++ {
			baselen := i - suflen
			suffix := s[baselen:baselen+suflen]
			cost := solutions[baselen].cost + d.cost(suffix)
			if cost < sol.cost {
				sol.cost = cost
				sol.lastlen = suflen
			}
		}

		return
	}

	for l := 1; l <= len(s); l++ {
		bs := bestSolution(l)
		solutions = append(solutions, bs)
	}

	var words []string
	l := len(s)
	for l > 0 {
		sol := &solutions[l]
		if sol.lastlen > 0 {
			words = append(words, s[l-sol.lastlen:l])
			l -= sol.lastlen
		} else {
			// Failed to split the whole string,
			// try on character less.
			l--
		}
	}

	reverse := func(ss []string) {
		for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
			ss[i], ss[j] = ss[j], ss[i]
		}
	}
	reverse(words)

	return words
}
