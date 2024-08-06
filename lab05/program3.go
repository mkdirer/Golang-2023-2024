package main

//example use: go run program3.go --w=album --f=songs.json --n=5

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Not enough argumets. Use --help for help")
		return
	}
	var wordFlag = flag.String("w", "", "Word to sort by")
	var filenameFlag = flag.String("f", "", "Filename")
	var nFlag = flag.Int("n", 10, "Number of records to print, default 10")
	flag.Parse()

	file, err := os.ReadFile(*filenameFlag)
	if err != nil {
		log.Fatal(err)
	}
	var structure map[string]interface{}
	err = json.Unmarshal(file, &structure)

	fmt.Println(len(structure))
	for k, v := range structure {
		fmt.Println(k)
		var list []interface{} = v.([]interface{})
		if *wordFlag == "rank" {
			sort.SliceStable(list, func(i, j int) bool {
				return list[i].(map[string]interface{})[*wordFlag].(float64) < list[j].(map[string]interface{})[*wordFlag].(float64)
			})
		} else {
			sort.SliceStable(list, func(i, j int) bool {
				return list[i].(map[string]interface{})[*wordFlag].(string) < list[j].(map[string]interface{})[*wordFlag].(string)
			})
		}
		for i := 0; i < *nFlag; i++ {
			fmt.Println(list[i])
		}
	}
}
