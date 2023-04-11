package pandaconfig

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"gopkg.in/yaml.v3"
)

type Parser struct {
	cfgValue     reflect.Value
	fileCfgValue reflect.Value
}

func (p *Parser) InitConfig(v any) error {
	var err error
	if err = p.SetConfigPtrValue(v); err != nil {
		return err
	}
	// TODO: 处理配置加载逻辑
	if err = p.loadConfigFile(); err != nil {
		return err
	}
	return nil
}

func (p *Parser) SetConfigPtrValue(v any) error {
	value := reflect.ValueOf(v)
	if value.Kind() != reflect.Pointer || value.Elem().Kind() != reflect.Struct {
		return errors.New("Pointer is not struct pointer")
	}
	p.cfgValue = value
	return nil
}

func (p *Parser) getConfigFile() string {
	path := flag.String("config", "", "config file")
	return *path
}

func (p *Parser) loadConfigFile() error {

	var (
		data []byte
		err  error
	)

	file := p.getConfigFile()
	if file == "" {
		return nil
	}

	fileType := filepath.Ext(file)

	if data, err = os.ReadFile(file); err != nil {
		return err
	}

	switch fileType {
	case "json":
		{
			err = json.Unmarshal(data, p.fileCfgValue.Elem().Interface())
		}
	case "yaml":
		{
			err = yaml.Unmarshal(data, p.fileCfgValue.Elem().Interface())
		}
	default:
		err = fmt.Errorf("not supported file type %s, use `json` or `yaml`", fileType)
	}

	return err

}
