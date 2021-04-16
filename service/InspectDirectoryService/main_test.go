package InspectDirectoryService

import (
	"os"
	"strings"
	"testing"
)

func TestIsMatchExtension(t *testing.T) {
	if !isMatchExtension("./test-service/TestApi.yaml", []string{"yaml", "yml"}) {
		t.Error("String pattern should match that extensions")
	}
	if isMatchExtension("./test-service/TestApi.tolm", []string{"yaml", "yml"}) {
		t.Error("String pattern shouldn't match that extensions")
	}
}

func TestGetDirectories(t *testing.T) {
	func() {
		defer func() {
			os.RemoveAll("./tmpDir/")
		}()
		os.MkdirAll("./tmpDir/a/", 0775)
		os.MkdirAll("./tmpDir/b/", 0775)
		os.MkdirAll("./tmpDir/c/", 0775)
		res := GetDirectories("./tmpDir/", []string{})
		exp := []string{"a", "b", "c"}
		if strings.Join(res, ",") != strings.Join(exp, ",") {
			t.Errorf("GetDirectories test returned an unexpected result: got %v want %v", res, exp)
		}
	}()

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Panic should be invoked")
			}
		}()
		GetDirectories("./tmpDir/", []string{})
	}()
}

func TestGetFiles(t *testing.T) {
	func() {
		defer func() {
			os.RemoveAll("./tmpDir/")
		}()
		os.MkdirAll("./tmpDir/a/", 0775)
		os.Create("./tmpDir/a/TestApi1.yaml")
		os.Create("./tmpDir/a/TestApi2.yml")
		os.Create("./tmpDir/a/TestApi3.toml")
		os.MkdirAll("./tmpDir/b/", 0775)
		os.Create("./tmpDir/b/TestApi4.yaml")
		os.Create("./tmpDir/b/TestApi5.toml")
		res := GetFiles("./tmpDir/a/", []string{"yaml", "yml"})
		exp := []string{"TestApi1.yaml", "TestApi2.yml"}
		if strings.Join(res, ",") != strings.Join(exp, ",") {
			t.Errorf("GetFiles test returned an unexpected result: got %v want %v", res, exp)
		}
		res = GetFiles("./tmpDir/b/", []string{"yaml", "yml"})
		exp = []string{"TestApi4.yaml"}
		if strings.Join(res, ",") != strings.Join(exp, ",") {
			t.Errorf("GetFiles test returned an unexpected result: got %v want %v", res, exp)
		}
	}()

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Panic should be invoked")
			}
		}()
		GetFiles("./tmpDir/", []string{})
	}()
}
