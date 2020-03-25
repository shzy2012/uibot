package main

import (
	"fmt"

	"github.com/shzy2012/uibot"
)

//Args 参数
type Args struct {
	uibot.Args
}

func main() {

	//传递给UiBot Worker的参数
	args := Args{}
	args.Name = "报销"
	args.Cmd = "报销"
	args.Data = `{"自定义数据":"测试数据"}`

	task := &uibot.TaskBody{}
	task.FlowCode = "xxx"   //从commander平台获取
	task.Args = args        //传递给worker的数据
	task.IsNowRun = 1       //立即运行
	task.EnvName = "泛微OA环境" //一定要写

	//实例化uibot客户端
	uibotClient := uibot.NewUiBot("cmdURL", "appKey", "appSecret")
	respByte, err := uibotClient.TaskCreate(task) //创建任务
	if err != nil {
		fmt.Printf("[TaskCreate]=>%s\n", err.Error())
	}

	fmt.Printf("[TaskCreate]=>%s\n", respByte)
}
