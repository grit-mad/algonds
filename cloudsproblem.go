package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
     var i int
    _, err := fmt.Scanf("%d", &i)
    if(err==nil){
         i=i
    }

    j:=0
    n:=0
    c := make([]int,0)
//    new_slice := make([]int,0)
    for j < i{
         _, err := fmt.Scanf("%d", &n)
        if(err==nil){
            c=append(c,n)
        }
        n=0
        j++
    }
//    fmt.Println(c)
    noofjumps:=0
    index:=0
    for index<(len(c)-1){
        if (index+2 <len(c)&&c[index+2]==0){
            noofjumps+=1
            index+=2
        } else {
            noofjumps+=1
            index+=1
        }
    }
 //   if c[len(c)-1]==
fmt.Println(noofjumps)
}
