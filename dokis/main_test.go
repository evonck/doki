package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestConvert(t *testing.T) {
	pathFolder = "."
	indexPath = "./"
	output = "./docs"
	Convert(nil)
	log.Print(indexPath)
	_, err := os.Stat("./docs/template.html")
	convey.Convey("err should be nil", t, func() {
		convey.So(err, convey.ShouldEqual, nil)
	})
	_, err = ioutil.ReadFile("./docs/index.html")
	convey.Convey("err should be nil", t, func() {
		convey.So(err, convey.ShouldEqual, nil)
	})
	os.RemoveAll("./docs")
}
