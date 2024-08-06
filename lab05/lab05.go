package main

import (
	crand "crypto/rand"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

type RandFlag struct{ minVal, maxVal int64 }

func (r *RandFlag) String() string { return "rand flag" }
func (r *RandFlag) Set(s string) error {
	vals := strings.Split(s, ",")
	minVal, _ := strconv.ParseInt(vals[0], 10, 10)
	maxVal, _ := strconv.ParseInt(vals[1], 10, 10)
	r.minVal = minVal
	r.maxVal = maxVal
	return nil
}
func setRandFlag(name string, s string, usage string) *RandFlag {
	vals := strings.Split(s, ",")
	minVal, _ := strconv.ParseInt(vals[0], 10, 10)
	maxVal, _ := strconv.ParseInt(vals[1], 10, 10)
	sf := RandFlag{minVal, maxVal}
	flag.CommandLine.Var(&sf, name, usage)
	return &sf
}

var randFlag = setRandFlag("rflag", "1,2", "-rflag string")

var sourceFlag = flag.String("source", "", "string")
var actionFlag = flag.String("action", "", "string")
var sortNumberFlag = flag.Int("sortnumber", 5, "int")
var sortFlag = flag.String("sortkey", "", "string")
var randTypeFlag = flag.Bool("randtype", true, "bool")

func ReadFromReader(r io.Reader) error {
	action := *actionFlag
	if action == "duplicates" {
		Info.Println("Action flag: duplicates")
	} else if action == "filter" {
		Info.Println("Action flag: filter")
	} else if action == "" {
		return errors.New("Action flag: - (no source action)")
	} else {
		return errors.New("Not supported action")
	}
	res := make([]string, 0)
	buf := make([]byte, 4048)
	for {
		_, err := r.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		r := strings.Split(strings.ReplaceAll(string(buf[:]), "\r\n", "\n"), "\n")
		res = append(res, r...)
	}
	if action == "duplicates" {
		m := make(map[string]bool)
		l := []string{}
		for _, el := range res {
			ok, _ := m[el]
			if !ok {
				m[el] = true
				l = append(l, el)
			}
		}
		fmt.Println("Removed duplicates")
		for _, el := range l {
			fmt.Println(el)
		}
	} else {
		l := []string{}
		for _, el := range res {
			if !strings.HasPrefix(el, "te") {
				l = append(l, el)
			}
		}
		fmt.Println("Filtered by prefix 'te'")
		for _, el := range l {
			fmt.Println(el)
		}
	}
	return nil
}

func CountJson(filename string) {
	var res map[string]interface{}
	fileContent, err := os.Open(filename)
	if err != nil {
		Error.Println("Error while opening the file ", filename)
		return
	}
	defer fileContent.Close()
	byteResult, err := io.ReadAll(fileContent)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(byteResult, &res)
	sortKey := *sortFlag
	sortnumber := *sortNumberFlag
	order := make([]string, 0)
	records := make([]string, 0)
	for _, el := range res {
		for y, v := range el.([]interface{}) {
			t := strconv.Itoa(y)
			for a, b := range v.(map[string]interface{}) {
				t += fmt.Sprintf(" %v: %v, ", a, b)
				if a == sortKey {
					order = append(order, b.(string))
				}
			}
			records = append(records, t)
		}
	}
	sort.Strings(order)
	fmt.Println("Sorted by", sortKey, "showing top", sortnumber)
	temp := sortKey + ": "
	j := 0
	for _, el := range order {
		tempHelp := temp + el
		for _, en := range records {
			if strings.Contains(en, tempHelp) {
				fmt.Println(en)
				j += 1
			}
		}
		if j == sortnumber {
			break
		}
	}
}

func GetRandom() int64 {
	randd := *randFlag
	maxVal := randd.maxVal
	minVal := randd.minVal
	randVer := *randTypeFlag
	if randVer {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		return r.Int63n(maxVal-minVal) + minVal
	} else {
		bg := big.NewInt(maxVal - minVal)
		n, err := crand.Int(crand.Reader, bg)
		if err != nil {
			panic(err)
		}
		return n.Int64() + minVal
	}
}

func init() {
	Trace = log.New(os.Stdout, "TRACE: ", log.Ldate|log.Ltime)
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
	Error = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
}

func main() {
	Info.Println("info")
	flag.Parse()
	source := *sourceFlag
	if source == "console" {
		Info.Println("Source flag: console")
		ReadFromReader(os.Stdin)
	} else if source == "file" {
		Info.Println("Source flag: file")
		sr, err := os.Open("test.txt")
		if err != nil {
			Error.Println(err)
			os.Exit(1)
		}
		defer sr.Close()
		ReadFromReader(sr)
	} else if source == "" {
		Info.Println("Source flag: - (no source flag)")
		os.Exit(1)
	} else {
		Error.Println("Not supported source")
		os.Exit(1)
	}
	CountJson("songs.json")
	fmt.Println(GetRandom())
}
