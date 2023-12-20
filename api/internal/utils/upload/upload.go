package upload

import (
	daoPlatform "api/internal/dao/platform"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

type UploadParam struct {
	Dir        string //上传的文件目录
	Expire     int64  //签名有效时间戳。单位：秒
	ExpireTime int64  //签名有效时间。单位：秒
	MinSize    int    //限制上传的文件大小。单位：字节
	MaxSize    int    //限制上传的文件大小。单位：字节。本地上传（uploadOfLocal.go）需要同时设置配置文件api/manifest/config/config.yaml中的server.clientMaxBodySize字段
}

type NotifyInfo struct {
	Url      string //地址
	Width    uint   //宽度
	Height   uint   //高度
	Size     uint   //大小。单位：比特
	MimeType string //文件类型
}

type Upload interface {
	Upload() (notifyInfo NotifyInfo, err error)                          // 本地上传
	Sign(param UploadParam) (signInfo map[string]interface{}, err error) // 获取签名（H5直传用）
	Config(param UploadParam) (config map[string]interface{}, err error) // 获取配置信息（APP直传前调用）
	Sts(param UploadParam) (stsInfo map[string]interface{}, err error)   // 获取Sts Token（APP直传用）
	Notify() (notifyInfo NotifyInfo, err error)                          // 回调
}

func CreateUploadParam(uploadType string) (param UploadParam) {
	param = UploadParam{
		Dir:        `common/` + gtime.Now().Format(`Ymd`) + `/`,
		Expire:     gtime.Now().Unix() + 15*60,
		ExpireTime: 15 * 60,
		MinSize:    0,
		MaxSize:    1024 * 1024 * 1024,
	}
	return
}

func NewUpload(ctx context.Context, uploadTypeOpt ...string) Upload {
	uploadType := ``
	if len(uploadTypeOpt) > 0 {
		uploadType = uploadTypeOpt[0]
	} else {
		uploadTypeVar, _ := daoPlatform.Config.ParseDbCtx(ctx).Where(daoPlatform.Config.Columns().ConfigKey, `uploadType`).Value(daoPlatform.Config.Columns().ConfigValue)
		uploadType = uploadTypeVar.String()
	}

	switch uploadType {
	case `uploadOfAliyunOss`:
		config, _ := daoPlatform.Config.Get(ctx, []string{`uploadOfAliyunOssHost`, `uploadOfAliyunOssBucket`, `uploadOfAliyunOssAccessKeyId`, `uploadOfAliyunOssAccessKeySecret`, `uploadOfAliyunOssCallbackUrl`, `uploadOfAliyunOssEndpoint`, `uploadOfAliyunOssRoleArn`})
		return NewUploadOfAliyunOss(ctx, config)
	// case `uploadOfLocal`:
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`uploadOfLocalUrl`, `uploadOfLocalSignKey`, `uploadOfLocalFileSaveDir`, `uploadOfLocalFileUrlPrefix`})
		return NewUploadOfLocal(ctx, config)
	}
}
