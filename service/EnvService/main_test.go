package EnvService

import "testing"

func TestReplaceExtension(t *testing.T) {
	cases := map[string]string{
		"1.yaml":   "1",
		"1.1.yaml": "1.1",
		"1.1":      "1",
	}
	for path, exp := range cases {
		res := replaceExtension(path)
		if res != exp {
			t.Errorf("Result %v not equal to expected %v", res, exp)
		}
	}
}

func TestBuildFilePathsString(t *testing.T) {
	res := buildFilePathsString([]string{
		"1.yaml",
		"Globus.Chat.yaml",
	}, "test")
	exp := "[{ name: \"1\", url: \"/test/api/1.yaml\" }, { name: \"Globus.Chat\", url: \"/test/api/Globus.Chat.yaml\" }]"
	if res != exp {
		t.Errorf("Result \n%v\n not equal to expected \n%v\n", res, exp)
	}
}

func TestGetEnvVariableString(t *testing.T) {
	res := GetEnvVariableString("test-service", []string{
		"1.yaml",
		"Globus.Chat.yaml",
	})
	exp := "TEST-SERVICE_URLS=[{ name: \"1\", url: \"/test-service/api/1.yaml\" }, { name: \"Globus.Chat\", url: \"/test-service/api/Globus.Chat.yaml\" }]"
	if res != exp {
		t.Errorf("Result \n%v\n not equal to expected \n%v\n", res, exp)
	}
}
