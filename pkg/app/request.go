package app

import (
	"ginDemo/pkg/logging"
	"github.com/astaxie/beego/validation"
)

func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Log.Info(err.Key, err.Message)
	}

	return
}
