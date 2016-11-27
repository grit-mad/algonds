package main
import "fmt"
func divthree(c []int)(int){
    k:=0
    n:=0
    for k < len(c){
           n+=c[k]
        k++
       }    
    return n%3
    

}
func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
     var n,k int
    _, err := fmt.Scanf("%d %d", &n,&k)
    if(err==nil){
         n=n
    }

    j:=0
    loop_var:=0
    set := make([]int,0)
    new_set := make([]int,0)
    new_set1 := make([]int,0)
    new_set2 := make([]int,0)
    for j < n{
         _, err := fmt.Scanf("%d", &loop_var)
        if(err==nil){
            set=append(set,loop_var)
        }
        loop_var=0
        j++
    }
//    fmt.Println(set)
    for num:=0;num<len(set);num++{
//        new_set=append(set[:num], set[num+1:]...)
//       fmt.Println(set[:num])
//       fmt.Println(set[num+1:])
       new_set1=set[:num]
       new_set2=set[num+1:]
       new_set=append(new_set,new_set1...)
       new_set=append(new_set,new_set2...)
//       fmt.Println(new_set)
        if (divthree(new_set)==0){
            fmt.Println(len(new_set))
            break
        }
        new_set=new_set[:0]
        new_set1=new_set1[:0]
        new_set2=new_set2[:0]
//         fmt.Println(set)
   
    }
//    fmt.Println(set,new_set)
}
