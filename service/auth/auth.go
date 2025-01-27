package auth

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/dhurley94/pritunl-client-electron/service/utils"
	"github.com/dropbox/godropbox/errors"
)

var Key = ""

func Init() (err error) {
	pth := utils.GetAuthPath()

	if _, e := os.Stat(pth); os.IsNotExist(e) {
		Key, err = utils.RandStr(64)
		if err != nil {
			return
		}

		err = ioutil.WriteFile(pth, []byte(Key), os.FileMode(0644))
		if err != nil {
			err = &WriteError{
				errors.Wrap(err, "auth: Failed to auth key"),
			}
			return
		}
	} else {
		data, e := ioutil.ReadFile(pth)
		if e != nil {
			err = &WriteError{
				errors.Wrap(e, "auth: Failed to auth key"),
			}
			return
		}

		Key = strings.TrimSpace(string(data))

		if Key == "" {
			err = os.Remove(pth)
			if err != nil {
				err = &WriteError{
					errors.Wrap(err, "auth: Failed to reset auth key"),
				}
				return
			}
			Init()
		}
	}

	return
}
