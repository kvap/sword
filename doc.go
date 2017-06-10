/*
sword is a word cutter in Go.

It uses a dictionary of known words to split lines into words.

Usage example:

	import (
		"github.com/kvap/sword"
	)
	func main() {
		var blade sword.Blade
		blade.Train("english.words.txt")
		words := b.Carve("helloworld")
		…
	}
*/
package sword
