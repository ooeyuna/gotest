package goruntines
import "log"

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