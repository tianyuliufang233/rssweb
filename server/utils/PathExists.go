package utils

import (
	"os"
	"server/global"

	"go.uber.org/zap"
)

// 判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return false, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 创建目录
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.GVA_LOG.Debug("create directory" + v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				global.GVA_LOG.Error("create directory"+v, zap.Any(" err: ", err))
			}
		}
	}
	return err
}