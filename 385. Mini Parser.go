/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
func deserialize(s string) *NestedInteger {
    fmt.Println(s)
    num := 0
    flag := 1
    tf:=false
    if s[0]!= '['{
        i:=0
        if s[0] == '-'{
            flag = -1
            i = 1
        }
        for ;i<len(s);i++{
            num = num *10 + int(s[i]-'0')
        }
        tmp := new(NestedInteger)
        tmp.SetInteger(flag * num)
        return tmp
    }
    ret := new(NestedInteger)
    for i:=1;i<len(s)-1;i++{
        if s[i] == '['{
            t,j :=1,i+1
            for ;j<len(s) && t >0;j++{
                if s[j] == '['{
                    t ++
                }else if s[j] ==']'{
                    t--
                } 
            }
            ret.Add(*deserialize(s[i:j]))
            tf = false
            i = j
        }else if s[i] == ','{
            tmp := new(NestedInteger)
            tmp.SetInteger(flag*num)
            flag = 1
            num = 0
            tf = false 
            ret.Add(*tmp)
        }else if s[i] == '-'{
            flag = -1
        }else{
            num = num*10 + int(s[i]-'0')
            tf = true
        }
        
    }
    if tf{
        tmp := new(NestedInteger)
        tmp.SetInteger(flag * num)
        ret.Add(*tmp)
    }
    return ret
}