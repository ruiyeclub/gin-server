package mylogs

import (
	"io"
	"log"
	"os"
	"time"
)

func init() {
	f, err := GetLogFile()
	if err != nil {
		return
	}
	// 组合一下即可，os.Stdout代表标准输出流
	multiWriter := io.MultiWriter(f, os.Stdout)
	log.SetOutput(multiWriter)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func GetLogFile() (*os.File, error) {
	dir := "./logs"
	b, mkdirErr := DirExists(dir)
	if !b && mkdirErr == nil {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			log.Println("logs目录创建失败~")
			return nil, err
		} else {
			log.Println("logs目录创建成功~")
		}
	}
	file := "./logs/" + time.Now().Format("2006-01-02") + "_info" + ".log"
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		// 关闭
		defer f.Close()
		return nil, err
	}
	return f, err
}

// 判断目录是否存在
func DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
