package config

import (
	"github.com/gawwo/fake115-go/log"
	"sync"
	"time"
)

const (
	Version    = "1.0.1"
	ServerName = "fake115"
	AppVer     = "11.2.0"
	EndString  = "000000"
	UserAgent  = "Mozilla/5.0  115disk/11.2.0"

	spiderCheckIntervalInt = 20
	SpiderCheckInterval    = time.Second * spiderCheckIntervalInt

	LinkSep         = "|"
	DirTargetPrefix = "U_1_"
)

var (
	RetryTimes = 10
	UserId     = ""
	UserKey    = ""
	Step       = 1150
	// Cookie提供文件读取和命令行设置两种方式；
	// 文件读取提供默认设置和命令行设置两种方式；
	Cookie            = ""
	DefaultCookiePath = "cookies.txt"
	CookiePath        = ""

	TotalSize = 0
	FileCount = 0
	Debug     bool

	WorkerNum     = 10
	WorkerNumRate = 100

	ConsumerWaitGroup = sync.WaitGroup{}

	// 是否处于等待人机验证的状态
	// 不是一个重要的状态，且可能的操作中都是在处理网络请
	// 求处理之后，冲突可能极小，不加读写锁，随便改
	SpiderVerification      = false
	SpiderStatWaitAliveTime = 0
	SpiderStatWaitTimeout   = spiderCheckIntervalInt * 3

	Logger = log.InitLogger(ServerName, false)
)

var fakeHeaders = map[string]string{
	"User-Agent": UserAgent,
}

func GetFakeHeaders(withCookie bool) map[string]string {

	copyMap := map[string]string{}
	for k, v := range fakeHeaders {
		copyMap[k] = v
	}

	if withCookie {
		copyMap["Cookie"] = Cookie
	}

	return copyMap
}

var fakeRangeHeaders = map[string]string{
	"Range":      "bytes=0-131071",
	"User-Agent": UserAgent,
}

func GetFakeRangeHeaders() map[string]string {
	copyMap := map[string]string{}
	for k, v := range fakeRangeHeaders {
		copyMap[k] = v
	}
	return copyMap
}
