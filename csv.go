package microphones

import (
	"encoding/csv"
	"errors"
	"fmt"
	"strconv"
)

type MicrophoneFileReader struct {
	filename       string
	microphoneList []Microphone
}

func (mfr *MicrophoneFileReader) LoadMicrophones(reader *csv.Reader) (bool, error) {
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	for i := 0; i < len(records); i++ {
		price, err := strconv.ParseFloat(records[i][3], 64)

		if err != nil {
			return false, errors.New("Not able to parse price to float")
		}

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

	return true, nil
}

func (mfr *MicrophoneFileReader) GetMicrophones() []Microphone {
	return mfr.microphoneList
}

type Microphone struct {
	name, brand, description, url, micType, micStyle string
	price                                            float64
}
