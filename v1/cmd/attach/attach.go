package attach

import (
	"errors"

	"github.com/hhhhkkk/mini-blog/v1/internal"
)

var START_FAIL = errors.New("服务生成失败...")
var OTHER_ERROR = errors.New("启动失败")

func Run() error {
	app, err := internal.InitApp()
	if err != nil {
		return START_FAIL
	}
	if err = app.Run(); err != nil {
		return errors.New(OTHER_ERROR.Error() + err.Error())
	}
	return nil
}
