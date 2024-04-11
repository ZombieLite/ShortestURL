package script

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func IsEmptyString(s string) bool {
	if len(s) > 0 {
		return false
	}
	return true
}

func (d *Data) ReplaceLink() {
	if d.Link[0] != 'h' || d.Link[1] != 't' || d.Link[2] != 't' || d.Link[3] != 'p' {
		d.Link = "https://" + d.Link
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		resp, err := http.Get(d.Link)
		http.DefaultClient.Timeout = 5 * time.Second
		if err != nil {
			d.NewLink = "Некорректная ссылка или сайт не отвечает на запрос."
			return
		}

		var b = "https://" + d.Host + "/z-" + urlGenerator()
		fmt.Println(b)
		d.NewLink = resp.Request.URL.String()
	}()
	wg.Wait()

}

func urlGenerator() string {
	var b string

	for i := 0; i < 6; i++ {
		a := rand.Intn(122-65) + 65

		if 91 <= a && a <= 96 {
			i--
			continue
		}
		b += string(byte(a))
	}
	return b
}
