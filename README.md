# GT SDK Libraries for Golang
该SDK是多厂商推送工具集，目前包装的功能有：icon上传、图片上传。支持的厂商模块有OPPO、XM。

## Installation
go get github.com/GetuiLaboratory/getui-3rd-push-utils-go

## Usage
### 配置文件初始化

先在开发者自身项目合适路径下创建config.toml配置文件，配置文件内填写如下配置项

```conf
## 以下是各厂商配置参数
XmAppId="xxx"
XmAppKey="xxx"
XmAppSecret="xxx"
OppoAppId="xxx"
OppoAppKey="xxx"
OppoAppSecret="xxx"
OppoMasterSecret="xxx"
```

### 服务调用

```Go
import (
	"fmt"

	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/config"
	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/manufacturer"
	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/manufacturer/consts"
	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/manufacturer/result"
)

func main() {
    // 加载配置文件
	var conf config.Conf
	conf.GetConf("/Users/xxx/Projects/push-xx/config.toml")

    // OPPO、XM同时上传icon
	var files [2]manufacturer.File = [2]manufacturer.File{{consts.MANUFACTURER_XM, "/Users/xxx/Documents/PIC/icon1.jpg"}, {consts.MANUFACTURER_OPPO, "/Users/xxx/Documents/PIC/icon2.jpg"}}
	var resultMap map[string]result.Result = manufacturer.UploadIcons(files, conf)
	for ManufacturerName := range resultMap {
		fmt.Println(ManufacturerName, "的结果是", resultMap[ManufacturerName])
	}

    // OPPO、XM同时上传图片
	var files2 [2]manufacturer.File = [2]manufacturer.File{{consts.MANUFACTURER_XM, "/Users/xxx/Documents/PIC/pic1.jpg"}, {consts.MANUFACTURER_OPPO, "/Users/xxx/Documents/PIC/pic2.jpg"}}
	var resultMap2 map[string]result.Result = manufacturer.UploadPics(files2, conf)
	for ManufacturerName := range resultMap2 {
		fmt.Println(ManufacturerName, "的结果是", resultMap2[ManufacturerName])
	}

    // 单独上传XM icon
	fmt.Println(manufacturer.UploadIcon("/Users/xxx/Documents/PIC/icon1.jpg", consts.MANUFACTURER_XM, conf))
    // 单独上传XM 图片
	fmt.Println(manufacturer.UploadPic("/Users/xxx/Documents/PIC/pic1.jpg", consts.MANUFACTURER_XM, conf))
    // 单独上传OPPO icon
	fmt.Println(manufacturer.UploadIcon("/Users/xxx/Documents/PIC/icon2.jpg", consts.MANUFACTURER_OPPO, conf))
    // 单独上传OPPO 图片
	fmt.Println(manufacturer.UploadPic("/Users/xxx/Documents/PIC/pic2.jpg", consts.MANUFACTURER_OPPO, conf))
}
```

### 服务结果解析
上一步可以看出上传接口返回的都是个Map，Map的key是厂商名（OPPO、XM），value是一个Result对象。Result包含以下三个属性：
- code：结果码，0成功、1失败、6无效的厂商名称
- message：success、fail、invalid manufacturer name
- data：成功时，值为icon在各厂商的上传url结果（或者picId）；失败时，值是失败原因。

## 其他说明
由于该sdk本质只是各厂商api的包装，所以对于一些接口限制和返回处理，需要遵循各厂商的api文档。下面放出
[OPPO](https://open.oppomobile.com/wiki/doc#id=10693) 和[XM](https://dev.mi.com/console/doc/detail?pId=1163#_10_1) 的API在线文档供参考。

