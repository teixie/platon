package repository

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/teixie/platon/helpers"
)

func SaveAs(file *multipart.FileHeader, lnk string, repos string) error {
	if _, err := os.Stat(lnk); err == nil || !os.IsNotExist(err) {
		return errors.New("file exists")
	}

	lnkPath := filepath.Dir(lnk)
	if err := os.MkdirAll(lnkPath, os.ModePerm); err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	body, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}
	md5 := helpers.MD5(string(body))

	dstPath := strings.TrimRight(repos, "/") + "/" + md5[0:2]
	if err := os.MkdirAll(dstPath, os.ModePerm); err != nil {
		return err
	}
	dst := dstPath + "/" + md5

	if _, err := os.Stat(dst); err != nil && os.IsNotExist(err) {
		out, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer out.Close()

		if _, err = out.Write(body); err != nil {
			return err
		}
	}

	return os.Symlink(dst, lnk)
}
