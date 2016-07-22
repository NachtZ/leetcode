package main
import "fmt"
func lengthOfLastWord(s string) int {
    if len(s) == 0{
        return 0
    }
    count :=0
    idx:=0
    for idx =len(s)-1;idx>=0&&s[idx] == ' ';idx--{

    }
    for i := idx; i>=0&&s[i]!=' ';i--{
        count ++;
    }
    return count;
}

func main(){
    fmt.Println(lengthOfLastWord("a "))
    fmt.Print("Hello")
}
