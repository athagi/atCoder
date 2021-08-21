package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	var S string
	var K int
	fmt.Scan(&S, &K)

	a := len(S)
	if len(S) < 8 {
		for i := 0; i < 8-a; i++ {
			S += " "
		}
	}

	char := make(map[string]bool)
	for ia := 0; ia < len(S); ia++ {
		a := S[:ia] + S[ia+1:]
		for ib := 0; ib < len(a); ib++ {
			b := a[:ib] + a[ib+1:]
			for ic := 0; ic < len(b); ic++ {
				c := b[:ic] + b[ic+1:]
				for id := 0; id < len(c); id++ {
					d := c[:id] + c[id+1:]
					for ie := 0; ie < len(d); ie++ {
						e := d[:ie] + d[ie+1:]
						for ix := 0; ix < len(e); ix++ {
							f := e[:ix] + e[ix+1:]
							for ig := 0; ig < len(f); ig++ {
								g := f[:ig] + f[ig+1:]

								s := SpaceStringsBuilder(S[ia:ia+1] + a[ib:ib+1] + b[ic:ic+1] + c[id:id+1] + d[ie:ie+1] + e[ix:ix+1] + f[ig:ig+1] + g[0:])

								if _, ok := char[s]; !ok {
									char[s] = true
								}

							}
						}
					}
				}
			}
		}
	}

	dict := make([]string, len(char))
	i := 0
	for k, _ := range char {
		dict[i] = k
		i++
	}
	sort.Strings(dict)
	fmt.Println(dict[K-1])

}

func SpaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
