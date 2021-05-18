package jsonarg

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"reflect"
)

type JSONArgs struct {
	TmpDir     string `json:"tmp_dir" pos_short:"t" overridable:"true"`
	Tmp2Dir    string `json:"tmp2_dir" pos_short:"2" overridable:"true"`
	FinalDir   string `json:"final_dir" pos_short:"d" overridable:"true"`
	Filename   string `json:"filename" pos_short:"f" overridable:"true"`
	Size       int64  `json:"size" pos_short:"k" overridable:"true"`
	PlotMemo   string `json:"plot_memo" pos_short:"m" overridable:"true"`
	PlotId     string `json:"plot_id" pos_short:"i" overridable:"true"`
	Buffer     int64  `json:"buffer" pos_short:"b" overridable:"true"`
	StripeSize int64  `json:"stripe_size" pos_short:"s" overridable:"true"`
	NumThreads int64  `json:"num_threads" pos_short:"r" overridable:"true"`
	NoBitField bool   `json:"nobitfield" pos_short:"e"`
}

type Override map[string]interface{}

var args JSONArgs
var overrides []Override
var PosArgs []string

func init() {
	k := reflect.ValueOf(args)
	for i := 0; i < k.NumField(); i++ {
		override := Override{}
		if k.Type().Field(i).Tag.Get("overridable") == "true" {
			help := fmt.Sprintf("Override value for %s", k.Type().Field(i).Tag.Get("json"))
			if k.Type().Field(i).Type == reflect.TypeOf("") {
				override[k.Type().Field(i).Tag.Get("pos_short")] = flag.String(k.Type().Field(i).Tag.Get("pos_short"), "", help)
			} else if k.Type().Field(i).Type == reflect.TypeOf(int64(0)) {
				override[k.Type().Field(i).Tag.Get("pos_short")] = flag.Int64(k.Type().Field(i).Tag.Get("pos_short"), 0, help)
			}
		}
		overrides = append(overrides, override)
	}
}

func Parse(jsonFile *string) {
	blob, err := ioutil.ReadFile(*jsonFile)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(blob, &args)
	if err != nil {
		fmt.Println("error:", err)
	}

	v := reflect.ValueOf(args)
	for i := 0; i < v.NumField(); i++ {
		var arg []string
		pos_short := v.Type().Field(i).Tag.Get("pos_short")

		for _, override := range overrides {
			update := false
			if _, ok := override[pos_short]; ok {
				flag := override[pos_short]
				if fmt.Sprint(reflect.TypeOf(flag)) == "*string" && *flag.(*string) != "" {
					reflect.ValueOf(&args).Elem().Field(i).SetString(*flag.(*string))
					update = true
				} else if fmt.Sprint(reflect.TypeOf(flag)) == "*int64" && *flag.(*int64) != 0 {
					reflect.ValueOf(&args).Elem().Field(i).SetInt(*flag.(*int64))
					update = true
				}
			}
			if update {
				v = reflect.ValueOf(args)
			}
		}

		if v.Type().Field(i).Type == reflect.TypeOf("") {
			arg = append(arg, fmt.Sprintf("-%s", pos_short))
			arg = append(arg, fmt.Sprintf("%s", v.Field(i).Interface()))
		} else if v.Type().Field(i).Type == reflect.TypeOf(false) {
			if v.Field(i).Interface() == false {
				arg = append(arg, "-"+pos_short)
			}
		} else if v.Type().Field(i).Type == reflect.TypeOf(0) {
			arg = append(arg, fmt.Sprintf("-%s", pos_short))
			arg = append(arg, fmt.Sprintf("%d", int(v.Field(i).Interface().(int64))))
		} else {
			arg = append(arg, fmt.Sprintf("-%s", pos_short))
			arg = append(arg, fmt.Sprintf("%v", v.Field(i).Interface()))
		}
		if arg[0] != "" {
			PosArgs = append(PosArgs, arg...)
		}
	}
	PosArgs = append([]string{"create"}, PosArgs...)
}
