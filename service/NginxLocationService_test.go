package service

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerateLocationOutput(t *testing.T) {
	func() {
		result := generateLocationOutput("test-service", "8001")
		expected := `
location /test-service {
	proxy_pass http://test-service-swagger:8001;
}`
		if result != expected {
			t.Errorf("add() test returned an unexpected result: got %v want %v", result, expected)
		}
	}()
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Panic should be invoked")
			}
		}()
		generateLocationOutput("", "")
	}()
}

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
		assertDir("../service")
	}()
	func() {
		defer func() {
			if err := recover(); err != nil {
				t.Errorf("assertDir fails, cause of './testTmp folder creation failed")
			}
			os.Remove("./testTmp")
		}()
		assertDir("./testTmp")
	}()
	func() {
		defer func() {
			os.Remove("tmpTest")
			if r := recover(); r == nil {
				t.Error("Panic should be invoked")
			}
		}()
		os.Mkdir("tmpTest", 0555)
		assertDir("./tmpTest/a")
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
		write("test-service", "test-output")
	}()
	func() {
		content, err := ioutil.ReadFile("../nginx_confs/test-service.conf")
		if err != nil {
			t.Errorf("file doesnt exists")
		}
		if string(content) != "test-output" {
			t.Errorf("file content doesn't match to 'test-output'")
		}
	}()
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Panic should be invoked")
			}
		}()
		write("", "")
	}()
}
