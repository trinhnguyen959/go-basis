package main

import (
	"log"
)

func main() {
	/*	defer fmt.Println("start")
				defer fmt.Println("middle")
				defer fmt.Println("end")

				res, err := http.Get("https://www.google.com/robots.txt")
				if err != nil {
					log.Fatal(err)
				}
				robots, err := io.ReadAll(res.Body)
				defer func(Body io.ReadCloser) {
					err := Body.Close()
					if err != nil {
						log.Fatal(err)
					}
				}(res.Body)
				fmt.Printf("%s", robots)

			a := "start"
			defer fmt.Println(a)
			// print start -> defer lay variable o luc func dc call chu khong phai luc return moi call
			a = "end"

		fmt.Println("\npanic")
		b, c := 1, 0
		ans := b / c
		fmt.Println(ans)*/

	// panic chay sau defer
	/*	fmt.Println("start")
		defer fmt.Println("this was a defer")
		panic("something bad happened")
		fmt.Println("end")*/

	/*	fmt.Println("start")
		defer func() {
			if err := recover(); err != nil {
				log.Println("Error: ", err)
			}
		}()

		panic("something bad happened")
		log.Println("end")*/
	log.Println("start")
	panicker()
	log.Println("end")

}

func panicker() {
	log.Println("about panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panic("something bad happened")
	log.Println("done panicking")
}
