package jef

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomSelecter(list []string, count int, step bool, print bool) map[string]int {
	if print {
		fmt.Print("\033[H\033[2J") // clear the screen
	}

	mapper := map[string]int{}
	index := []string{}

	for ch := range listToRandomChannel(list, count, step) {
		if print {
			fmt.Printf("\033[0;0H") // move the cursor to row 0, column 0 of the terminal
		}

		_, ok := mapper[ch]
		if ok {
			mapper[ch]++
		} else {
			mapper[ch] = 1
		}
		deduplicate(&index, ch)

		if print {
			for _, v := range index {
				fmt.Printf("%s: %d\n", v, mapper[v])
			}
		}
	}
	return mapper
}

func deduplicate(set *[]string, x string) {
	for _, v := range *set {
		if v == x {
			return
		}
	}
	*set = append(*set, x)
}

func listToRandomChannel(ss []string, count int, step bool) <-chan string {
	ch := make(chan string, count)
	go func() {
		defer close(ch)

		rand.Seed(time.Now().UnixNano())
		for i := 0; i < count; i++ {
			ch <- ss[rand.Intn(len(ss))]
			if step {
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
	return ch
}
