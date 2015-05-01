package goruntines
import (
	"log"
)

type Queue struct {
	channel chan int
	count int
	cap int
}

func Init(max int) *Queue {
	return &Queue{make(chan int,max),0,max}
}

func (this *Queue) Add(i int) {
	this.channel <- i
	this.count = this.count + 1
	log.Print("add queue")
}

func (this *Queue) Get() int {
	defer func(){
		this.count = this.count - 1}()
	defer log.Println("queue get")
	return <- this.channel
}

func (this *Queue) Status(){
	log.Print("count:",this.count,",cap:",this.cap)
}

func Main(){
	q := Init(10)
	go func(q *Queue) {
		for i := 0; i<20; i++ {
			q.Add(i)
			q.Status()
			log.Println("index:",i)
		}
	}(q)
	for i := 0; i<20; i++ {
		a := q.Get()
		log.Println("get:",a,",index:",i)
	}
}