package repository

import (
	"github.com/teixie/platon/helpers"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strings"
)

func SaveAs(file *multipart.FileHeader, dst string, repos string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 计算文件MD5，用于获取文件在仓库中的目的地址
	body, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}
	md5 := helpers.MD5(string(body))

	// 文件在仓库中的目的地址
	reposDst := strings.TrimRight(repos, "/") + "/" + md5[0:2] + "/" + md5

	// 如果文件不存在仓库中，则将文件保存到仓库
	if _, err := os.Stat(reposDst); err != nil && os.IsNotExist(err) {
		out, err := os.Create(reposDst)
		if err != nil {
			return err
		}
		defer out.Close()

		if _, err = out.Write(body); err != nil {
			return err
		}
	}

	// 创建软链接
	return os.Symlink(reposDst, dst)
}
