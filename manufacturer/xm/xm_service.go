package xm

import (
	"encoding/json"

	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/config"
	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/manufacturer/consts"
	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/manufacturer/result"
	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/util"
)

const UPLOAD_PIC_URL string = "https://api.xmpush.xiaomi.com/media/upload/image"

func UploadPic(file string, isIcon bool, conf config.Conf) result.Result {
	var reqParams map[string]string = make(map[string]string)
	if isIcon {
		if x, found := util.Get(consts.BIG_KEY_ICON+consts.MANUFACTURER_XM, file); found {
			return result.Success(x)
		}
		reqParams["is_icon"] = "true"
	} else {
		if x, found := util.Get(consts.BIG_KEY_PIC+consts.MANUFACTURER_XM, file); found {
			return result.Success(x)
		}
		reqParams["is_icon"] = "false"
	}

	var headers map[string]string = make(map[string]string)
	headers["Authorization"] = "key=" + conf.XmAppSecret
	reqParams["is_global"] = "false"
	files := []util.UploadFile{
		{Name: "file", Filepath: file},
	}

	var uploadResult = util.PostFile(UPLOAD_PIC_URL, reqParams, files, headers)

	var resultMap map[string]interface{}
	json.Unmarshal([]byte(uploadResult), &resultMap)
	if util.IsEqual(resultMap["code"].(float64), 0) {
		var fileUrl string
		if isIcon {
			fileUrl = resultMap["data"].(map[string]interface{})["icon_url"].(string)
			util.Set(consts.BIG_KEY_ICON+consts.MANUFACTURER_XM, file, fileUrl, -1)
		} else {
			fileUrl = resultMap["data"].(map[string]interface{})["pic_url"].(string)
			util.Set(consts.BIG_KEY_PIC+consts.MANUFACTURER_XM, file, fileUrl, -1)
		}
		return result.Success(fileUrl)
	} else {
		return result.Fail(uploadResult)
		// return manufacturer.Fail(resultMap["message"].(string))
	}
}
