package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func plus(s *string) {
	b := []rune(*s)
	if b[48] == 57 {
		b[47] += 1
		b[48] = 0
		*s = string(b)
		*s = *s + "0"
	} else {
		b[48] += 1
		*s = string(b)
	}

}
func check(s string) []byte {
	req, _ := http.Get(s)
	// check err
	defer req.Body.Close()
	body, _ := io.ReadAll(req.Body)
	// check err
	return body
}
func eq(a, b []byte) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] && i < 2188 {
			fmt.Println(i)
			return false
		}
	}
	return true
}
func t() []byte {
	req, _ := http.Get("https://www.mosconsv.ru/ru/events.aspx?type=7308")
	defer req.Body.Close()
	body, _ := io.ReadAll(req.Body)

	// check err
	return body
}
func example1() {
	s := "https://www.mosconsv.ru/ru/event_p.aspx?id=178508"
	t1 := t()
	t2 := check(s)

	if eq(t1, t2) {
		fmt.Println("it are same")
	} else {
		fmt.Println("find")
		fmt.Println("smth")
		plus(&s)
	}
	time.Sleep(3 * time.Second)
}
