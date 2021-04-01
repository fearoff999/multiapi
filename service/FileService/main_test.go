package FileService

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDirExists(t *testing.T) {
	if !dirExists("../service/") {
		t.Errorf("dirExists fails, cause of '../service/' folder exists")
	}
	if dirExists("../service1/") {
		t.Errorf("dirExists fails, cause of '../service1/' folder doesnt exists")
	}
	func() {
		defer func() {
			os.Remove("tmpTest")
			if r := recover(); r == nil {
				t.Error("Panic should be invoked")
			}
		}()
		os.Mkdir("tmpTest", 0000)
		dirExists("./tmpTest/a")
		t.Error("Panic should be invoked")
	}()
}

func TestAssertDir(t *testing.T) {
	func() {
		defer func() {
			if err := recover(); err != nil {
				t.Errorf("assertDir fails, cause of '../service/ folder exists already")
			}
		}()
		AssertDir("../service")
	}()
	func() {
		defer func() {
			if err := recover(); err != nil {
				t.Errorf("assertDir fails, cause of './testTmp folder creation failed")
			}
			os.Remove("./testTmp")
		}()
		AssertDir("./testTmp")
	}()
	func() {
		defer func() {
			os.Remove("tmpTest")
			if r := recover(); r == nil {
				t.Error("Panic should be invoked")
			}
		}()
		os.Mkdir("tmpTest", 0555)
		AssertDir("./tmpTest/a")
		t.Error("Panic should be invoked")
	}()
}

func TestWrite(t *testing.T) {
	func() {
		defer func() {
			if err := recover(); err != nil {
				t.Errorf("write fails")
			}
		}()
		Write("./nginx_config/", "test-service.conf", "test-output")
	}()
	func() {
		content, err := ioutil.ReadFile("./nginx_config/test-service.conf")
		if err != nil {
			t.Errorf("file doesnt exists")
		}
		if string(content) != "test-output" {
			t.Errorf("file content doesn't match to 'test-output'")
		}
		os.Remove("./nginx_config/test-service.conf")
		os.Remove("./nginx_config/")
	}()
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Panic should be invoked")
			}
		}()
		Write("", "", "")
	}()
}
