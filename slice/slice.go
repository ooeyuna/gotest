package slice
import (
	"fmt"
)
//append分片执行的copy操作
func Slice(){
	sli1 := []int{0,1,2,3,4}
	sli2 := append(sli1,5)
	sli3 := append(sli2,sli1...)
	sli1[0] = 447

	print(&sli1)
	print(&sli2)
	fmt.Println(sli3)
}

func print(sli *[]int){
	for i:=0;i<len(*sli);i++ {
		fmt.Print(" ",(*sli)[i]," ")
	}
	fmt.Println("")
}