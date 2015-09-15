package main

// 4m in real	1m45.438s
//	 user	0m0.036s
//	 sys	0m0.017s

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"../../"	// "github.com/mediawen/watson-go-sdk"
)

type Cfg struct {
	User string 			`json:"user"`
	Pass string 			`json:"pass"`
}

func loadCfg(name string) (*Cfg, error) {
	f, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	cfg := &Cfg{}

	err = json.Unmarshal(f, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf(
			"usage: stt out.srt model in.[wav|flac]\n" +
			"       stt -l\n")
		return
	}

	cfg, err := loadCfg("./stt.cfg.json")
	if err != nil {
		log.Fatal(err)
	}

	w := watson.New(cfg.User, cfg.Pass)

	ml, err := w.GetModels()
	if err != nil {
		log.Fatal(err)
	}

	if args[1] == "-l" {
		if len(args) > 2 {
			log.Fatal("too many args")
		}

		for _, m := range ml.Models {
			fmt.Printf("%s %-8d=> %s\n", m.Lang, m.Rate, m.Name)
		}

		return
	}

	out := args[1]
	model := args[2]
	in := args[3]
	ext := ""
	switch path.Ext(in) {
	case ".wav":
		ext = "wav"
	case ".flac":
		ext = "flac"
	case ".json":
		ext = "json"
	default:
		log.Fatal("stt: unknown file format: ", in)
	}

	found := false
	for _, m := range ml.Models {
		if m.Name == model {
			found = true
		}
	}
	if !found {
		log.Fatal("model not found")
	}

	is, err := os.Open(in)
	if err != nil {
		log.Fatal(err)
	}
	defer is.Close()

	os, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Close()

	tt, err := w.Recognize(is, model, ext)
	if err != nil {
		log.Fatal(err)
	}

	for _, w := range tt.Words {
		fmt.Printf("%v\n", w)
	}
}
