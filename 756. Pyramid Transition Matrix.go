package main

import (
	"fmt"
	"time"
)

func pyramidTransition(bottom string, allowed []string) bool {
	m := make(map[string][]rune)
	//l := []string{"A", "B", "C", "D", "E", "F", "G"}
	for _, a := range allowed {
		tmp := string(a[0:2])
		m[tmp] = append(m[tmp], rune(a[2]))
	}
	py := make([][][]rune, len(bottom))
	for i := 0; i < len(py); i++ {
		py[i] = make([][]rune, len(bottom))
		py[0][i] = []rune{rune(bottom[i])}
	}
	for i := 0; i < len(bottom)-1; i++ {
		for j := 1; j < len(bottom)-i; j++ {
			t0, t1 := py[i][j-1], py[i][j]
			ap := []rune{}
			for x := 0; x < len(t0); x++ {
				for y := 0; y < len(t1); y++ {
					tmp := string(t0[x]) + string(t1[y])
					if len(m[tmp]) > 0 {
						ap = append(ap, m[tmp]...)
					}
				}
			}
			if len(ap) == 0 {
				return false
			}
			mp, f := make(map[rune]bool), []rune{}
			for _, v := range ap {
				if mp[v] == true {
					continue
				}
				mp[v] = true
				f = append(f, v)
			}
			py[i+1][j-1] = f
		}
	}
	return len(py[len(bottom)-1][0]) > 0
}

func main() {
	start := time.Now()
	res := pyramidTransition("BDBBAA", []string{"ACB", "ACA", "AAA", "ACD", "BCD", "BAA", "BCB", "BCC", "BAD", "BAB", "BAC", "CAB", "CCD", "CAA", "CCA", "CCC", "CAD", "DAD", "DAA", "DAC", "DCD", "DCC", "DCA", "DDD", "BDD", "ABB", "ABC", "ABD", "BDB", "BBD", "BBC", "BBA", "ADD", "ADC", "ADA", "DDC", "DDB", "DDA", "CDA", "CDD", "CBA", "CDB", "CBD", "DBA", "DBC", "DBD", "BDA"})
	end := time.Now()
	fmt.Println("run in ", end.Sub(start).Nanoseconds(), "ns")
	fmt.Println(res)

}
