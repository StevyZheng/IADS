package config

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/olebedev/config"
	"gopkg.in/ini.v1"
	"iads/lib/stringx"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/****************************************/
//config file like ifcfg-eth0
/****************************************/
type CommonConfigParser struct {
	filePath string
	buffer   map[string]string
}

func NewCommonConfigParser(filename string) *CommonConfigParser {
	return &CommonConfigParser{
		filePath: filename,
	}
}

func (e *CommonConfigParser) Read() (int, error) {
	e.buffer = make(map[string]string)
	file, err := os.Open(e.filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nowRowStr := stringx.Trim(scanner.Text(), "\n")
		tmpList := strings.Split(nowRowStr, "=")
		listLen := len(tmpList)
		if listLen > 1 {
			e.buffer[tmpList[0]] = tmpList[1]
		}
	}
	return len(e.buffer), err
}

func (e *CommonConfigParser) save() (int, error) {
	var (
		err error
		fp  *os.File
	)
	mapLen := len(e.buffer)
	if e.buffer == nil || mapLen <= 0 {
		log.Fatal("buffer is nil")
		return mapLen, err
	} else {
		fp, err = os.OpenFile(e.filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if fp == nil {
			log.Fatal("open file failed.")
			return mapLen, err
		}
		for k, v := range e.buffer {
			rowStr := fmt.Sprintf("%s=%s\n", k, v)
			_, err = fp.WriteString(rowStr)
		}
	}
	return mapLen, err
}

func (e *CommonConfigParser) GetValue(key string) (string, error) {
	var (
		ret string
		err error
	)
	if e.buffer == nil {
		log.Fatal(errors.New("buffer is nil"))
	} else {
		if _, ok := e.buffer[key]; ok {
			ret = e.buffer[key]
		} else {
			ret = ""
			log.Fatal(errors.New("nokey"))
		}
	}
	return ret, err
}

func (e *CommonConfigParser) SetValue(key string, value string) error {
	var (
		err error
	)
	_, err = e.Read()
	e.buffer[key] = value
	_, err = e.save()
	return err
}

/***************************** like ini ***********************************/
type SectionConfigParserError struct {
	errorInfo string
}
type SectionConfigParser struct {
	confParser *ini.File
	filePath   string
}

func (e *SectionConfigParserError) Error() string { return e.errorInfo }

func NewSectionConfigParser(filename string) *SectionConfigParser {
	e := &SectionConfigParser{
		filePath: filename,
	}
	conf, err := ini.Load(e.filePath)
	if err != nil {
		e.confParser = nil
	}
	e.confParser = conf
	return e
}

func (e *SectionConfigParser) GetString(section string, key string) string {
	if e.confParser == nil {
		log.Fatal("confParser is nil")
	}
	s := e.confParser.Section(section)
	if s == nil {
		log.Fatal("get section failed.")
	}
	return s.Key(key).String()
}

func (e *SectionConfigParser) GetInt32(section string, key string) int32 {
	if e.confParser == nil {
		log.Fatal("confParser is nil")
	}
	s := e.confParser.Section(section)
	if s == nil {
		log.Fatal("get section failed.")
	}
	valueInt, _ := s.Key(key).Int()
	return int32(valueInt)
}

func (e *SectionConfigParser) GetUint32(section string, key string) uint32 {
	if e.confParser == nil {
		log.Fatal("confParser is nil")
	}
	s := e.confParser.Section(section)
	if s == nil {
		log.Fatal("get section failed.")
	}
	valueInt, _ := s.Key(key).Uint()
	return uint32(valueInt)
}

func (e *SectionConfigParser) GetInt64(section string, key string) int64 {
	if e.confParser == nil {
		log.Fatal("confParser is nil")
	}
	s := e.confParser.Section(section)
	if s == nil {
		log.Fatal("get section failed.")
	}
	valueInt, _ := s.Key(key).Int64()
	return valueInt
}

func (e *SectionConfigParser) GetUint64(section string, key string) uint64 {
	if e.confParser == nil {
		log.Fatal("confParser is nil")
	}
	s := e.confParser.Section(section)
	if s == nil {
		log.Fatal("get section failed.")
	}
	valueInt, _ := s.Key(key).Uint64()
	return valueInt
}

func (e *SectionConfigParser) GetFloat32(section string, key string) float32 {
	if e.confParser == nil {
		log.Fatal("confParser is nil")
	}
	s := e.confParser.Section(section)
	if s == nil {
		log.Fatal("get section failed.")
	}
	valueFloat, _ := s.Key(key).Float64()
	return float32(valueFloat)
}

func (e *SectionConfigParser) GetFloat64(section string, key string) float64 {
	if e.confParser == nil {
		log.Fatal("confParser is nil")
	}
	s := e.confParser.Section(section)
	if s == nil {
		log.Fatal("get section failed.")
	}
	valueFloat, _ := s.Key(key).Float64()
	return valueFloat
}

func (e *SectionConfigParser) SetValue(section string, key string, value string) error {
	if e.confParser == nil {
		log.Fatal("confParser is nil")
	}
	e.confParser.Section(section).Key(key).SetValue(value)
	err := e.confParser.SaveTo(e.filePath)
	return err
}

/****************************  yaml *******************************/
type YamlConfigParser struct {
	cfg      *config.Config
	filePath string
}

func (e *YamlConfigParser) Load(filename string) error {

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	yamlString := string(file)
	e.cfg, err = config.ParseYaml(yamlString)
	e.filePath = filename
	return err
}

//keyStr like "development.users.0.name"
func (e *YamlConfigParser) GetStringValue(keyStr string) (string, error) {
	value, err := e.cfg.String(keyStr)
	return value, err
}

func (e *YamlConfigParser) GetBoolValue(keyStr string) (bool, error) {
	value, err := e.cfg.Bool(keyStr)
	return value, err
}

func (e *YamlConfigParser) GetIntValue(keyStr string) (int, error) {
	value, err := e.cfg.Int(keyStr)
	return value, err
}

func (e *YamlConfigParser) GetFloat64Value(keyStr string) (float64, error) {
	value, err := e.cfg.Float64(keyStr)
	return value, err
}

func (e *YamlConfigParser) GetListValue(keyStr string) ([]interface{}, error) {
	value, err := e.cfg.List(keyStr)
	return value, err
}

func (e *YamlConfigParser) GetMapValue(keyStr string) (map[string]interface{}, error) {
	value, err := e.cfg.Map(keyStr)
	return value, err
}

func (e *YamlConfigParser) SetValue(keyStr string, value string) error {
	err := config.Set(e.cfg, keyStr, value)
	return err
}

func (e *YamlConfigParser) SaveToFile(filename string) error {
	yaml, err := config.RenderYaml(e.cfg)
	err = ioutil.WriteFile(filename, []byte(yaml), os.ModePerm)
	return err
}

func (e *YamlConfigParser) SaveSelf(filename string) error {
	yaml, err := config.RenderYaml(e.cfg)
	err = ioutil.WriteFile(e.filePath, []byte(yaml), os.ModePerm)
	return err
}

func (e *YamlConfigParser) CreateYamlFile(filename string, mapBuffer map[string]interface{}) error {
	yaml, err := config.RenderYaml(mapBuffer)
	err = ioutil.WriteFile(e.filePath, []byte(yaml), os.ModePerm)
	return err
}
