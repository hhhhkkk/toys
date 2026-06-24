package middleware

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dromara/carbon/v2"
	"github.com/gin-gonic/gin"
	"github.com/hhhhkkk/mini-blog/v1/config"
)

type MiniBlogErrorWriter struct {
	f io.Writer
}

var _ io.Writer = (*MiniBlogErrorWriter)(nil)

func (w *MiniBlogErrorWriter) Write(s []byte) (int, error) {
	return w.f.Write(s)
}

func Recovery(app config.Config) gin.HandlerFunc {
	date := carbon.Now().Format("Y-m-d")
	errorLog := fmt.Sprintf("%s-error-%s.log", app.Server.Name, date)
	return gin.RecoveryWithWriter(NewMiniBlogErrorWriter(app, errorLog), func(c *gin.Context, err any) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
	})
}

func NewMiniBlogErrorWriter(app config.Config, filename string) *MiniBlogErrorWriter {
	rootPath := config.GetRootPath()
	runtimePath := filepath.Join(rootPath, "runtime")

	_, err := os.Stat(runtimePath)
	if os.IsNotExist(err) {
		os.Mkdir(runtimePath, 0755)
	}

	errorLogPath := filepath.Join(runtimePath, filename)
	file, err := os.OpenFile(errorLogPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.FileMode(0644))
	if err != nil {
		panic(err)
	}
	return &MiniBlogErrorWriter{
		f: file,
	}
}
