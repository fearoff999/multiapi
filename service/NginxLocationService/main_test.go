package NginxLocationService

import (
	"testing"
)

func TestGenerateLocationOutput(t *testing.T) {
	func() {
		result := GenerateLocationOutput("test-service", "8001")
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
		GenerateLocationOutput("", "")
	}()
}
