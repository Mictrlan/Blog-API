package upload

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

const (
	// PictureDir - save pictures file
	PictureDir = "picture"

	// VideoDir - save videos file
	VideoDir = "video"

	// OtherDir - files other than video and picture
	OtherDir = "other"

	//FileUploadDir - the root directory of the upload files
	FileUploadDir = "./pkg/upload/files"

	// FileKey - key of the file
	FileKey = "file"
)

var (
	fileMap = map[string]string{}

	picture = []string{".jpg", ".png", ".jpeg", ".gif", ".bmp"}
	video   = []string{".avi", ".wmv", ".mpg", ".mpeg", ".mpe", ".mov", ".rm", ".ram", ".swf", ".mp4", ".rmvb", ".asf", ".divx", ".vob"}
	fileDir = FilePath()
)

// FilePath  - default fileDir
func FilePath() map[string]string {
	for _, suffix := range picture {
		fileMap[suffix] = PictureDir
	}

	for _, suffix := range video {
		fileMap[suffix] = VideoDir
	}

	return fileMap
}

// ClassifyBySuffix - return "pictrue", "vedio", "other"
func ClassifyBySuffix(suffix string) string {
	if dir := fileDir[suffix]; dir != "" {
		return dir
	}

	return "other"
}

// MD5 return a hash value
func MD5(file []byte) (string, error) {
	sum := md5.New()

	sum.Write(file)

	MD5Str := hex.EncodeToString(sum.Sum(nil))
	return MD5Str, nil
}

// CopyFile create a new file by path
func CopyFile(path string, file []byte) error {
	cur, err := os.Create(path)
	defer cur.Close()
	if err != nil {
		return err
	}

	_, err = cur.Write(file)
	return err
}

// CheckDir Verify directory existence， if directory dosen't exists then create it
// Stat returns a FileInfo describing the named file.
// IsNotExist if file or directory not exists return true
// MkdirAll creates a directory named path
func CheckDir(path ...string) error {
	for _, name := range path {
		_, err := os.Stat(FileUploadDir + "/" + name)
		if err != nil {
			if os.IsNotExist(err) {
				err = os.MkdirAll(FileUploadDir+"/"+name, 0777)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
