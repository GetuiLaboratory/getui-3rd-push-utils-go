package oppo

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/config"
	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/manufacturer/consts"
	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/manufacturer/result"
	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/util"
)

const AUTH_URL string = "https://api.push.oppomobile.com/server/v1/auth"
const UPLOAD_ICON_URL string = "https://api-media.push.heytapmobi.com/server/v1/media/upload/small_picture"
const UPLOAD_PIC_URL string = "https://api-media.push.heytapmobi.com/server/v1/media/upload/big_picture"

func Auth(conf config.Conf) {
	if !needAuth() {
		return
	}
	var reqParams map[string]string = make(map[string]string)
	timeUnix := time.Now().UnixMilli()
	reqParams["app_key"] = conf.OppoAppKey
	reqParams["timestamp"] = fmt.Sprintf("%d", timeUnix)
	reqParams["sign"] = sign(conf.OppoAppKey, timeUnix, conf.OppoMasterSecret)
	var result = util.PostForm(AUTH_URL, reqParams, nil)
	var resultMap map[string]interface{}
	json.Unmarshal([]byte(result), &resultMap)

	if util.IsEqual(resultMap["code"].(float64), 0) {
		var authToken = resultMap["data"].(map[string]interface{})["auth_token"].(string)
		util.Set(consts.AUTH_TOKEN, consts.MANUFACTURER_OPPO, authToken, 24*3600)
	}
}

func needAuth() bool {
	_, result := util.Get(consts.AUTH_TOKEN, consts.MANUFACTURER_OPPO)
	return !result
}

func getAuthToken() string {
	authToken, ok := util.Get(consts.AUTH_TOKEN, consts.MANUFACTURER_OPPO)
	if ok {
		return authToken
	}
	return ""
}

func sign(appKey string, timestamp int64, masterSecret string) string {
	sum := sha256.Sum256([]byte(fmt.Sprintf("%s%d%s", appKey, timestamp, masterSecret)))
	return fmt.Sprintf("%x", sum)
}

func UploadPic(file string, isIcon bool, conf config.Conf) result.Result {
	var uploadUrl string
	if isIcon {
		uploadUrl = UPLOAD_ICON_URL
		if x, found := util.Get(consts.BIG_KEY_ICON+consts.MANUFACTURER_OPPO, file); found {
			return result.Success(x)
		}
	} else {
		uploadUrl = UPLOAD_PIC_URL
		if x, found := util.Get(consts.BIG_KEY_PIC+consts.MANUFACTURER_OPPO, file); found {
			return result.Success(x)
		}
	}

	Auth(conf)

	var reqParams map[string]string = make(map[string]string)
	reqParams["auth_token"] = getAuthToken()
	reqParams["picture_ttl"] = "2592000"

	files := []util.UploadFile{
		{Name: "file", Filepath: file},
	}

	var uploadResult = util.PostFile(uploadUrl, reqParams, files, nil)

	var resultMap map[string]interface{}
	json.Unmarshal([]byte(uploadResult), &resultMap)
	if util.IsEqual(resultMap["code"].(float64), 0) {
		var fileUrl string
		if isIcon {
			fileUrl = resultMap["data"].(map[string]interface{})["small_picture_id"].(string)
			util.Set(consts.BIG_KEY_ICON+consts.MANUFACTURER_OPPO, file, fileUrl, 2592000)
		} else {
			fileUrl = resultMap["data"].(map[string]interface{})["big_picture_id"].(string)
			util.Set(consts.BIG_KEY_PIC+consts.MANUFACTURER_OPPO, file, fileUrl, 2592000)
		}
		return result.Success(fileUrl)
	} else {
		return result.Fail(uploadResult)
		// return manufacturer.Fail(resultMap["message"].(string))
	}
}
