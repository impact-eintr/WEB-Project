package main

import (
	b64 "encoding/base64"
	"fmt"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {
	file, _ := os.Open("./sd.csv")

	df := dataframe.ReadCSV(file)
	s1 := df.Select([]string{"startx", "starty", "endx", "endy"})

	tmp := s1.Records()

	for _, row := range tmp {
		tmpres := ""
		for _, data := range row {
			sDec, _ := b64.StdEncoding.DecodeString(data)
			tmpres += string(sDec)
			tmpres += ","
		}
		fmt.Println(tmpres[:len(tmpres)-1])
	}

}
