package main
import "fmt"
func shortestelement(c []int)(int){
    shortest_element:=c[0]
    k:=0
    for k < len(c){
           if (shortest_element>=c[k]){
               shortest_element=c[k]
               }
        k++

       }    
    
    return shortest_element
    

}
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
    new_slice := make([]int,0)
    for j < i{
         _, err := fmt.Scanf("%d", &n)
        if(err==nil){
            c=append(c,n)
        }
        n=0
        j++
    }
//    fmt.Println(i) 
    s:=0
    s=shortestelement(c)  
//    fmt.Println(c)
//   fmt.Println(len(c))       
    for len(c) > 0 {
        fmt.Println(len(c))
        s=shortestelement(c)
//        fmt.Println(s) 
        for num:=0;num<len(c);num++{
            c[num]-=s
            }
//        fmt.Println(c)
        for num:=0;num<len(c);num++{
            if c[num]!=0{
               new_slice=append(new_slice,c[num])
//                    fmt.Println(c)
//               fmt.Println("innerloop")
            }
            }
        c=new_slice
        new_slice = new_slice[:0]
//       fmt.Println(c)
//        fmt.Println(new_slice)
//        fmt.Println("outerloop")
    }  


}
