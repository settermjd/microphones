package microphones

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type MicrophoneFileReader struct {
	filename       string
	microphoneList []Microphone
}

func (mfr *MicrophoneFileReader) LoadMicrophones() {
	data, err := ioutil.ReadFile(mfr.filename)
	s := string(data)
	r := csv.NewReader(strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	} else {
		records, err := r.ReadAll()
		if err != nil {
			fmt.Println(err)
		} else {
			for i := 0; i < len(records); i++ {
				price, _ := strconv.ParseFloat(records[i][3], 64)
				mic := Microphone{
					name:        records[i][0],
					brand:       records[i][1],
					description: records[i][2],
					price:       price,
					url:         records[i][4],
					micType:     records[i][5],
					micStyle:    records[i][6],
				}
				mfr.microphoneList = append(mfr.microphoneList, mic)
			}
		}
	}
}

func (mfr *MicrophoneFileReader) GetMicrophones() []Microphone {
	return mfr.microphoneList
}

type Microphone struct {
	name, brand, description, url, micType, micStyle string
	price                                            float64
}
