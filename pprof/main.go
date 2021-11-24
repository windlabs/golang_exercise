package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"strings"
	"time"
)

func main(){
    //logicCode()
	//args()
	//myFlag()
	prof()
}

func prof(){
	var isCpuPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCpuPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCpuPprof {
		cpuFile, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed. err:%v\n", err)
			return
		}
		fmt.Printf("create cpu pprof success.")
		pprof.StartCPUProfile(cpuFile)
		defer func() {
			pprof.StopCPUProfile()
			cpuFile.Close()
		}()
	}

	for i :=0;i<60; i++{
		go split("a.b.c.d.e.f.g",".")
	}
	if isMemPprof {
		memFile, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof file failed, err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(memFile)
		defer func(){
			memFile.Close()
		}()

	}
	time.Sleep(time.Second * 20)

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

func logicCode(){
	c := make(chan int)

	for {
		select {
		case v:= <-c:
			fmt.Printf("chan c value:%#v", v)
		default:
			//time.Sleep(time.Millisecond*500)
		}
	}
}

func myFlag(){
    name := flag.String("name", "刘先生", "请输入名字。")
	flag.Parse()
	fmt.Printf("名字：%s", *name)
}
func args(){
	fmt.Printf("%v\n", os.Args)
	fmt.Printf("%v, %v\n", os.Args[1], os.Args[2])
	fmt.Printf("%T\n", os.Args)
}
