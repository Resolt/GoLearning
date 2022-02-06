package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// func doSomethingCool(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("timed out")
// 			return
// 		default:
// 			fmt.Println("doing something cool")
// 		}
// 		time.Sleep(500 * time.Millisecond)
// 	}

// }

// func main() {
// 	fmt.Println("Go Context Tutorial")
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2.4*float32(time.Second)))
// 	defer cancel()

// 	go doSomethingCool(ctx)

// 	fmt.Println("for the sake of comprehension")

// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("oh no, I've exceeded the deadline")
// 	}

// 	fmt.Println("for the sake of comprehension")
// }

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*800))
	defer cancel()
	req = req.WithContext(ctx)
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(out))
}
