package InitializeController

import (
	"os"
	"strings"
	"testing"
)

func TestScanDirs(t *testing.T) {
	defer func() {
		os.RemoveAll("./a/")
		os.RemoveAll("./b/")
		os.RemoveAll("./c/")
		os.RemoveAll("./nginx_config/")
	}()

	os.Mkdir("./nginx_config/", 0755)
	os.Mkdir("./a/", 0775)
	os.Create("./a/1.yaml")
	os.Create("./a/2.yaml")
	os.Create("./a/3.toml")
	os.Mkdir("./b/", 0775)
	os.Create("./b/1.yaml")
	os.Create("./b/2.yaml")
	os.Create("./b/3.toml")
	os.Mkdir("./c/", 0775)
	os.Create("./c/1.toml")
	os.Create("./c/2.toml")
	os.Create("./c/3.toml")

	dirs, files := scanDirs()

	if strings.Join(dirs, ",") != "a,b" {
		t.Errorf("%v %v", dirs, files)
	}

	res := ""
	for _, fileArr := range files {
		res += strings.Join(fileArr, ",")
	}
	exp := "./a/1.yaml,./a/2.yaml./b/1.yaml,./b/2.yaml"
	if res != exp {
		t.Errorf("%v %v", res, exp)
	}
}

func TestInitialize(t *testing.T) {
	defer func() {
		os.RemoveAll("./a/")
		os.RemoveAll("./b/")
		os.RemoveAll("./c/")
		os.RemoveAll("./nginx_config/")
		os.Remove(".env")
		os.Remove("docker-compose.yaml")
		os.RemoveAll("./html/")
		os.RemoveAll("./basic_auth")
	}()

	os.Mkdir("./nginx_config/", 0755)
	os.Mkdir("./a/", 0775)
	os.Create("./a/1.yaml")
	os.Create("./a/2.yaml")
	os.Create("./a/3.toml")
	os.Mkdir("./b/", 0775)
	os.Create("./b/1.yaml")
	os.Create("./b/2.yaml")
	os.Create("./b/3.toml")
	os.Mkdir("./c/", 0775)
	os.Create("./c/1.toml")
	os.Create("./c/2.toml")
	os.Create("./c/3.toml")

	Initialize()
}
