package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
     var i string
    _, err := fmt.Scanf("%s", &i)
    if(err==nil){
         i=i
    }
    var j int
    _, err = fmt.Scanf("%d", &j)
    if(err==nil){
         j=j
    }
//    fmt.Println(i,j)
    noa:=0
    for lv:=0;lv<len(i);lv++{
        if (i[lv]=='a'){
            noa+=1
        }
        
    }
//    fmt.Println(noa)
    if(j%len(i)==0){
        fmt.Println((j/len(i))*noa)
    } else{
        noa2:=0
        for lv1:=0;lv1<j%len(i);lv1++{
            if (i[lv1]=='a'){
            noa2+=1
        }
        }
        fmt.Println(((j/len(i))*noa)+noa2)
    }
}
