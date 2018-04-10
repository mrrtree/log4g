package log4g

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"

	logrus "github.com/sirupsen/logrus"
)

var opts map[string]*option = make(map[string]*option)

type option struct {
	logruslvl  logrus.Level
	Xtype      string `json:"type"`
	FileName   string `json:"fileName"`
	Level      string `json:"level"`
	MaxLogSize int    `json:"maxLogSize"`
	Backups    int    `json:"backups"`
	Category   string `json:"category"`
}

var defaultOpt *option

func init() {
	defaultOpt = &option{
		logruslvl:  logrus.InfoLevel,
		Xtype:      "stdout",
		Level:      "INFO",
		MaxLogSize: 1024 * 1024 * 1024,
		Backups:    2,
		Category:   "default",
	}
}

func loadConfigure(confFile string) {
	var conf []option

	bytes, err := ioutil.ReadFile(confFile)
	if err != nil {
		if strings.HasSuffix(err.Error(), "no such file or directory") {
			return
		}
		log.Fatal("Load log4g configure fail with error ", err)
	}

	if err := json.Unmarshal(bytes, &conf); err != nil {
		log.Fatal("Parse log4g configure fail with error ", err)
	}

	for _, opt := range conf {
		logruslvl, err := logrus.ParseLevel(opt.Level)
		if err != nil {
			log.Fatal("Invalid log4g conf", err)
		}
		opt.logruslvl = logruslvl
		setOption(opt.Category, &opt)
	}
}

func setOption(t string, opt *option) {
	opts[t] = opt
}

func getOption(t string) *option {
	opt, ok := opts[t]
	if !ok {
		opt = defaultOpt
	}
	return opt
}
