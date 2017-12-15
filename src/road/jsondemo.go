package main

import (
	"encoding/json"
	"log"
	"os"
	"fmt"
)

func main()  {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(2)
			return
		}
		for k := range v {
			if k != "Titles" {
				v[k] = nil
			}
		}
		return
		if err := enc.Encode(&v); err != nil {
			fmt.Println(6666666666666)
			log.Println(1)
		}
	}
}
