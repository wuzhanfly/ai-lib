// Copyright 2024 wuzhanfly <wuzhanfly@gmail.com>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package file

type RemoteFileSaveResult struct {
	SavePath              string
	RelativePath          string
	ThumbnailPath         string
	ThumbnailRelativePath string
	FileName              string
	FileSize              int64
	OriUrl                string
	OriFileName           string
}
