// 向os ENV提供加载.env数据
package dotnev

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/pm-esd/config/ini/parser"
)

// DefaultName 默认文件名
var DefaultName = ".env"

// OnlyLoadExists 加载文件存在
var OnlyLoadExists bool

// Load 将解析.env文件数据加载到系统 ENV,使用:dotenv.Load("./", ".env")
func Load(dir string, filenames ...string) (err error) {
	if len(filenames) == 0 {
		filenames = []string{DefaultName}
	}

	for _, filename := range filenames {
		file := filepath.Join(dir, filename)
		if err = loadFile(file); err != nil {
			break
		}
	}
	return
}

// LoadExists 加载存在的文件
func LoadExists(dir string, filenames ...string) error {
	OnlyLoadExists = true

	return Load(dir, filenames...)
}

// LoadFromMap 从给定的字符串映射加载数据
func LoadFromMap(kv map[string]string) (err error) {
	for key, val := range kv {
		key = strings.ToUpper(key)
		err = os.Setenv(key, val)
		if err != nil {
			break
		}
	}
	return
}

// loadFile 加载并解析 .env 到系统变量中
func loadFile(file string) (err error) {
	fd, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) && OnlyLoadExists {
			return nil
		}
		return err
	}
	defer fd.Close()

	s := bufio.NewScanner(fd)
	p := parser.NewSimpled(parser.NoDefSection)

	if _, err = p.ParseFrom(s); err != nil {
		return
	}

	if mp, ok := p.SimpleData()[p.DefSection]; ok {
		err = LoadFromMap(mp)
	}
	return
}
