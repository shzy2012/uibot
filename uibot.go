package uibot

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/shzy2012/common/network"
	"github.com/shzy2012/common/tool"
)

//getSign 签名
func getSign(secret, nonce string, time int64) string {
	//签名=sha1(固定字符串+AppSecret+随机字符串+Unix时间戳)
	var signStr = fmt.Sprintf("afd8426953b54e23b925a63dff4bf7ed%s%s%v", secret, nonce, time)
	h := sha1.New()
	io.WriteString(h, signStr)
	return fmt.Sprintf("%x", h.Sum(nil))
}

//UiBot UiBot
type UiBot struct {
	appKey    string
	appSecret string
	cmdUrl    string
}

//NewUiBot 实例化
func NewUiBot(cmdUrl, appKey, secret string) *UiBot {
	return &UiBot{
		appKey:    appKey,
		appSecret: secret,
		cmdUrl:    cmdUrl,
	}
}

//TaskBody body 创建任务
type TaskBody struct {
	FlowCode    string      `json:"flowCode"`    //流程代码
	Args        interface{} `json:"args"`        //任务参数,可以自定义
	CallbackURL string      `json:"callbackUrl"` //回调地址
	WorkerName  string      `json:"workerName"`  //worker名称，环境名称与worker名称必须指定一个
	EnvName     string      `json:"envName"`     //环境名称，环境名称与worker名称必须指定一个
	IsNowRun    int         `json:"isNow"`       //是否立即运行 0否 1是
}

//Args 默认UiBot的参数,结合实际情况,自定义该结构体
type Args struct {
	Name string      `json:"name"` //命令名称
	Cmd  string      `json:"cmd"`  //命令
	Data interface{} `json:"data"` //数据
}

/***********************
	查询Ui Bot机器人
************************/

//TaskQuery 查询任务
func (x *UiBot) TaskQuery(taskID int) ([]byte, error) {

	//每次请求都需要重新生成签名
	nonce := tool.GetRandomString(16)
	timestamp := time.Now().UTC().Unix() // / 1e6
	sign := getSign(x.appSecret, nonce, timestamp)

	//拼接URL
	url := fmt.Sprintf("%s/task/status/%v?appKey=%s&nonce=%s&timestamp=%v&sign=%s", x.cmdUrl, taskID, x.appKey, nonce, timestamp, sign)
	return network.HTTPGet(url)
}

//TaskCreate 创建任务
func (x *UiBot) TaskCreate(task *TaskBody) ([]byte, error) {

	//每次请求都需要重新生成签名
	nonce := tool.GetRandomString(16)
	timestamp := time.Now().UTC().Unix() // / 1e6
	sign := getSign(x.appSecret, nonce, timestamp)

	//拼接URL
	url := fmt.Sprintf("%s/task/create?appKey=%s&nonce=%s&timestamp=%v&sign=%s", x.cmdUrl, x.appKey, nonce, timestamp, sign)
	input, err := json.Marshal(task)
	if err != nil {
		return nil, err
	}

	return network.HTTPost(url, input)
}
