package v2

import (
	"github.com/hhhhkkk/mini-blog/v2/internal"
)

func Run() error {
	app, erro := internal.InitApp()
	if erro != nil {
		return erro
	}

	if err := app.Run(); err != nil {
		return err
	}
	return nil
}
