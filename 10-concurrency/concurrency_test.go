package concurrency

import (
	"reflect"
	"testing"
	"time"
)

// @Title
// @Description
// @Author
// @Update

func mockWebsiteChecker(url string) bool {
	if url == "waat://whatisthisshit.poop" {
		return false
	}
	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < 100; i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"waat://whatisthisshit.poop",
		"http://google.com",
		"http://comick.fun",
	}

	want := map[string]bool{
		"waat://whatisthisshit.poop": false,
		"http://google.com":          true,
		"http://comick.fun":          true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}

}
