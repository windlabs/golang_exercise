package main

import (
   "fmt"
   "strings"
)

func main(){
   //log.Fatalf("abcd %T", "abc")
   var m []int
   m = append(m, 1)
   fmt.Println(m)
}

func max(x int, y int) int{
   if x > y {
      return x
   }else{
      return y
   }
}

func split(str string, sep string) []string{
   var ans  []string

   index := strings.Index(str, sep)
   for index >=0 {
      ans = append(ans, str[0:index])
      str = str[index+len(sep):]
      index = strings.Index(str, sep)
   }
   ans = append(ans, str)
   return ans
}