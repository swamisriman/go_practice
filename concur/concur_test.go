package concur

import (
	"reflect"
	"testing"
	"time"
)

func MockWebsiteChecker(url string) bool {
	if url == "www.google.com" {
		time.Sleep(10000 * time.Millisecond)
		return true
	} else {
		return false
	}
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"www.google.com",
		"ww.google.co",
		"www.foogle.com",
	}

	expected := map[string]bool{
		"www.google.com": true,
		"ww.google.co":   false,
		"www.foogle.com": false,
	}

	actual := CheckWebsites(MockWebsiteChecker, websites)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v but actual: %v", expected, actual)
	}
}

func SlowMockWebsiteChecker(url string) bool {
	time.Sleep(1000 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	listLen := 1000
	urlList := make([]string, listLen)
	for i := 0; i < listLen; i++ {
		urlList[i] = "random url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(SlowMockWebsiteChecker, urlList)
	}
}
