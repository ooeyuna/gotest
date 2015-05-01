package main
import (
	"./goruntines"
	"fmt"
	"log"
)

func main() {
	q := goruntines.Init(10)
	go func(q *goruntines.Queue) {
		for i := 0; i<20; i++ {
			a := q.Get()
			log.Println("get:",a,",index:",i)
		}
	}(q)
	for i := 0; i<20; i++ {
		q.Add(i)
		q.Status()
		log.Println("index:",i)
	}

	fmt.Println("main start")
}