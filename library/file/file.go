// Copyright 2024 wuzhanfly <wuzhanfly@gmail.com>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package file

import (
	"io"
	"os"
	"os/exec"

	"github.com/disintegration/imaging"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/encoding/gurl"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/wuzhanfly/ai-lib/library/helper"
	"github.com/wuzhanfly/ai-lib/library/snowflake"
)

func RemoteFileSave(url string, proxy ...string) (re *RemoteFileSaveResult, err error) {
	urlParse, err := gurl.ParseURL(url, 32)
	if err != nil {
		return nil, err
	}
	fileExt := gfile.Ext(urlParse["path"])
	oriFileName := gfile.Name(urlParse["path"])
	// 首先储存临时文件
	httpClient := ghttp.NewClient()
	if len(proxy) > 0 && proxy[0] != "" {
		httpClient.SetProxy(proxy[0])
	}
	response, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Close()
	defer response.Body.Close()
	id := snowflake.GenerateID()
	tmpFileName := gconv.String(id) + fileExt
	tmpDir := helper.FormatDirStr(g.Config().GetString("commonConf.fileTmpPath"))
	if !gfile.IsDir(tmpDir) {
		if err := gfile.Mkdir(tmpDir); err != nil {
			return nil, err
		}
	}
	tmpPath := tmpDir + tmpFileName
	file, err := os.Create(tmpPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return nil, err
	}
	// 对临时文件进行处理移动到目标目录
	sha1, err := gsha1.EncryptFile(tmpPath)
	if err != nil {
		return nil, err
	}
	fileSize := gfile.Size(tmpPath)
	// 先查看本地目录路径是否存在，不存在则创建
	savePath := helper.FormatDirStr(g.Config().GetString("commonConf.fileSavePath"))
	if !gfile.IsDir(savePath) {
		if err := gfile.Mkdir(savePath); err != nil {
			return nil, err
		}
	}
	relativePath := ""
	pathName1 := gstr.SubStr(sha1, 0, 3)
	pathName2 := gstr.SubStr(sha1, 3, 3)
	savePath += pathName1 + "/"
	relativePath += pathName1 + "/"
	if !gfile.IsDir(savePath) {
		if err := gfile.Mkdir(savePath); err != nil {
			return nil, err
		}
	}
	savePath += pathName2 + "/"
	relativePath += pathName2 + "/"
	if !gfile.IsDir(savePath) {
		if err := gfile.Mkdir(savePath); err != nil {
			return nil, err
		}
	}
	fileName := sha1 + fileExt
	thumbnailFileName := sha1 + "_thumbnail" + fileExt
	thumbnailPath := savePath + thumbnailFileName
	thumbnailRelativePath := relativePath + thumbnailFileName
	savePath += fileName
	relativePath += fileName
	// 如果已经存在这个文件就不在存了，直接删了临时文件就行了
	if !gfile.IsFile(savePath) {
		// 把临时文件转移到最终目录
		if err := gfile.Move(tmpPath, savePath); err != nil {
			// 防止跨文件系统移动文件报错，采用系统原生命令
			var cmd *exec.Cmd
			cmd = exec.Command("mv", tmpPath, savePath)
			_, err = cmd.Output()
			if err != nil {
				return nil, err
			}
		}
		// 把文件生成一张缩略图
		imageData, err := imaging.Open(savePath)
		if err == nil {
			resizeImage := imaging.Resize(imageData, 300, 0, imaging.Lanczos)
			err = imaging.Save(resizeImage, thumbnailPath)
			if err != nil {
				thumbnailPath = ""
				thumbnailRelativePath = ""
				glog.Line(true).Println("生成缩略图错误", savePath, thumbnailPath, err)
			}
		} else {
			thumbnailPath = ""
			thumbnailRelativePath = ""
			glog.Line(true).Println("生成缩略图错误", savePath, err)
		}
	} else {
		// 删除临时文件
		_ = gfile.Remove(tmpPath)
	}

	re = &RemoteFileSaveResult{
		SavePath:              savePath,
		RelativePath:          relativePath,
		ThumbnailPath:         thumbnailPath,
		ThumbnailRelativePath: thumbnailRelativePath,
		FileName:              fileName,
		FileSize:              fileSize,
		OriUrl:                url,
		OriFileName:           oriFileName,
	}
	return re, nil
}
