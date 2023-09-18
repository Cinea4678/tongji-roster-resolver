package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"tongji-roster-resolver/pkg/resolver"
)

var (
	inputFile = flag.String("input", "", "点名册的路径")
)

func main() {
	flag.Parse()

	if inputFile == nil {
		fmt.Println("缺少参数input，请添加后重试。")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if strings.HasSuffix(*inputFile, "xls") {
		fmt.Println("我们尚不支持由Excel 2003或更早版本创建的XLS文件，很抱歉。您可以使用新版本Excel将其转换为XLS格式后再使用本工具。")
	}

	fmt.Print(resolver.ResolveFile(*inputFile))
}
