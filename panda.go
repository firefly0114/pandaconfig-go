package pandaconfig

import (
	"errors"
	"reflect"
)

type Parser struct {
	cfgValue     reflect.Value
	fileCfgValue reflect.Value
	cfgFile      string
}

func (p *Parser) InitConfig(v any) error {
	var err error
	if err = p.SetConfigValue(v); err != nil {
		return err
	}

	// TODO: 处理配置加载逻辑
	return nil
}

func (p *Parser) SetConfigValue(v any) error {
	value := reflect.ValueOf(v)
	if value.Kind() != reflect.Pointer || value.Kind() != reflect.Struct {
		return errors.New("Pointer is not struct pointer")
	}
	p.cfgValue = value
	return nil
}
