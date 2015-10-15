package microphones

import (
	. "github.com/franela/goblin"
	"io/ioutil"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Csv Files", func() {
		mfr := MicrophoneFileReader{}
		csvfile, _ := ioutil.ReadFile("./data.csv")
		data := string(csvfile)
		g.It("Be able to retrieve a list of all the microphones", func() {
			loaded, _ := mfr.LoadMicrophones(data)
			g.Assert(loaded).IsTrue("Microphones were not successfully loaded")
			g.Assert(len(mfr.microphoneList)).Equal(1)
			micList := mfr.GetMicrophones()
			g.Assert(len(micList)).Equal(1)
			for i := 0; i < len(micList); i++ {
				g.Assert(reflect.TypeOf(micList[i]).Name()).Equal("Microphone")
				g.Assert(reflect.TypeOf(micList[i]).PkgPath()).Equal("github.com/settermjd/microphones")
			}
		})
		g.It("Can retrieve the properties from a Microphone", func() {
			mfr.LoadMicrophones(data)
			micList := mfr.GetMicrophones()
			micOne := micList[0]
			g.Assert(len(micOne.name) != 0).IsTrue("Microphone name should be set")
			g.Assert(len(micOne.brand) != 0).IsTrue("Microphone brand should be set")
			g.Assert(len(micOne.description) != 0).IsTrue("Microphone description should be set")
			g.Assert(micOne.price != 0).IsTrue("Microphone price should be set")
			g.Assert(len(micOne.url) != 0).IsTrue("Microphone url should be set")
			g.Assert(len(micOne.micType) != 0).IsTrue("Microphone micType should be set")
			g.Assert(len(micOne.micStyle) != 0).IsTrue("Microphone micStyle should be set")
		})
	})
}
