// ---------------------------------------------------------------------------
// Constantin S. Pan <kvapen@gmail.com> wrote this file. As long as you retain
// this notice you can do whatever you want with this stuff. If we meet some
// day, and you think this stuff is worth it, you can buy me a bottle of cider
// in return.
// 	Constantin S. Pan
// ---------------------------------------------------------------------------

package sword

import (
	"strings"
	"testing"
)

func TestCarve(t *testing.T) {
	var b Blade
	b.Train("dict.txt")

	orig := "fat father offered her or them other error"
	glued := strings.Replace(orig, " ", "", -1)
	words := b.Carve(glued)
	recombined := strings.Join(words, " ")
	if orig != recombined {
		t.Fatalf("'%s' != '%s'", orig, recombined)
	}
}
