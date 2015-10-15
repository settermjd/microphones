package microphones

import (
	. "github.com/franela/goblin"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Csv Files", func() {
		mfr := MicrophoneFileReader{filename: "./data.csv"}
		g.It("Be able to initialise a microphone file reader", func() {
			g.Assert(mfr.filename).Equal("./data.csv")
		})
		g.It("Be able to read in all the contents of the data file", func() {
			mfr.LoadMicrophones()
			g.Assert(len(mfr.microphoneList)).Equal(1)
		})
		g.It("Be able to retrieve a list of all the microphones", func() {
			mfr.LoadMicrophones()
			micList := mfr.GetMicrophones()
			g.Assert(len(micList)).Equal(2)
			for i := 0; i < len(micList); i++ {
				g.Assert(reflect.TypeOf(micList[i]).Name()).Equal("Microphone")
				g.Assert(reflect.TypeOf(micList[i]).PkgPath()).Equal("github.com/settermjd/csvfiles")
			}
		})
		g.It("Can retrieve the properties from a Microphone", func() {
			mfr.LoadMicrophones()
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
