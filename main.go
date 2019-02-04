package main

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"github.com/hermanschaaf/cedict"
)

const CEDICTDownloadUrl = "https://www.mdbg.net/chinese/export/cedict/cedict_1_0_ts_utf-8_mdbg.txt.gz"

func main() {
	resp, err := http.Get(CEDICTDownloadUrl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	cedictHttpReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cedictHttpReader.Close()

	c := cedict.New(cedictHttpReader)

	outFile, err := os.Create("./cc-cedict-parsed.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()
	writer := csv.NewWriter(outFile)

	header := []string{
		"Simplified",
		"Traditional",
		"Pinyin",
		"PinyinWithTones",
		"PinyinNoTones",
		"Definition",
	}
	writer.Write(header)

	for {
		err := c.NextEntry()
		if err != nil {
			fmt.Println(err)
			break
		}
		entry := c.Entry()

		line := []string{
			entry.Simplified,
			entry.Traditional,
			entry.Pinyin,
			entry.PinyinWithTones,
			entry.PinyinNoTones,
			entry.Definitions[0], // TODO include all defs
		}
		writer.Write(line)
	}

	// TODO: compute hash of file and add to filename

	fmt.Println("Success.")
}
