package main

//应该在GOPATH下建立

import (
	"Golang/gostudy/5.init/lib1"
	"Golang/gostudy/5.init/lib2"
	// _"Golang/gostudy/5.init/lib1"匿名导入
	//无法使用当前包的方法，但是会执行当前包中的init()方法
	//import aa"fmt",给fmt包起一个别名，aa.Println()来直接调用
	//import ."fmt"当前fmt包中的全部方法导入当前本包的作用中，fmt包中的全部的方法可以直接用API调用
)

func main() {
	lib1.Lib1Test
	lib2.Lib2Test
}
