package manufacturer

import (
	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/config"
	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/manufacturer/consts"
	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/manufacturer/oppo"
	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/manufacturer/result"
	"github.com/GetuiLaboratory/getui-3rd-push-utils-go/manufacturer/xm"
)

type File struct {
	ManufacturerName string
	FilePath         string
}

func UploadPic(file string, manufacturerName string, conf config.Conf) result.Result {
	if manufacturerName == consts.MANUFACTURER_XM {
		return xm.UploadPic(file, false, conf)
	}
	if manufacturerName == consts.MANUFACTURER_OPPO {
		return oppo.UploadPic(file, false, conf)
	}
	return result.InvalidManufacturerName()
}

func UploadIcon(file string, manufacturerName string, conf config.Conf) result.Result {
	if manufacturerName == consts.MANUFACTURER_XM {
		return xm.UploadPic(file, true, conf)
	}
	if manufacturerName == consts.MANUFACTURER_OPPO {
		return oppo.UploadPic(file, true, conf)
	}
	return result.InvalidManufacturerName()
}

func UploadPics(files [2]File, conf config.Conf) map[string]result.Result {
	var results map[string]result.Result = make(map[string]result.Result)
	var i int
	for i = 0; i < 2; i++ {
		var file File = files[i]
		if file.ManufacturerName == consts.MANUFACTURER_XM {
			results[consts.MANUFACTURER_XM] = xm.UploadPic(file.FilePath, false, conf)
		}
		if file.ManufacturerName == consts.MANUFACTURER_OPPO {
			results[consts.MANUFACTURER_OPPO] = oppo.UploadPic(file.FilePath, false, conf)
		}
	}
	return results
}

func UploadIcons(files [2]File, conf config.Conf) map[string]result.Result {
	var results map[string]result.Result = make(map[string]result.Result)
	var i int
	for i = 0; i < 2; i++ {
		var file File = files[i]
		if file.ManufacturerName == consts.MANUFACTURER_XM {
			results[consts.MANUFACTURER_XM] = xm.UploadPic(file.FilePath, true, conf)
		}
		if file.ManufacturerName == consts.MANUFACTURER_OPPO {
			results[consts.MANUFACTURER_OPPO] = oppo.UploadPic(file.FilePath, true, conf)
		}
	}
	return results
}
