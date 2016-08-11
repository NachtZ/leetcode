package main
import "fmt"
import "strconv"
var res []string
func help(s string, offset, dot int){
    if offset >= len(s){
        return
    }
    if dot == 3{
        if len(s) - offset >=2 && s[offset] == '0'{
            return
        }
        if i,_:=strconv.Atoi(string(s[offset:len(s)]));i <=255{
            res = append(res,s)
        }
        return 
    }
    if s[offset] == '0'{
        tmp := []byte(s)
        t := string(tmp[0:offset+1]) + "." + string(tmp[offset+1:])
        help(t,offset+2,dot+1)
        return
    }
    for i:=1;i<=3;i++{
        if offset+i >=len(s){
            return
        }
        if v,_:=strconv.Atoi(string(s[offset:offset+i]));v <=255{
            tmp :=[]byte(s)
            t := string(tmp[:offset+i])+"."+string(tmp[offset+i:])
            help(t,offset+i+1,dot+1)
        }
    }
}
func restoreIpAddresses(s string) []string {
    res = []string{}
    if len(s) <4 || len(s)>12{
        return res
    }
    help(s,0,0)
    return res

}
func main(){
    fmt.Println(restoreIpAddresses("25525511135"))
    fmt.Print("Hello")
}
