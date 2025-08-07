package test

import (
	"strings"
	"testing"

	urlsextractor "github.com/isadri/gitactions-dashboard/pkg/urls_extractor"
)

func TestExtractor(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{`<a href="#"></a>`, []string{"#"}},
		{`<a href="/home"></a>`, []string{"/home"}},
		{`<a href="/repo?name=test"></a>`, []string{"/repo?name=test"}},
		{`<a href="/repo?name=test&id=342"></a>`, []string{"/repo?name=test&id=342"}},
		{`<a href="1"></a><a href="2"></a>`, []string{"1", "2"}},
		{`<a href="/repos/jobs?repo=test&"></a><a href="/repos/jobs/logs?repo=test&jobid=33"></a>`, []string{"/repos/jobs?repo=test&", "/repos/jobs/logs?repo=test&jobid=33"}},
	}

	for _, test := range tests {
		result, err := urlsextractor.Extract(strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Extractor(%s), expect %s, get error, reason: %s",
				test.input, test.want, err)
			continue
		}
		if len(result) != len(test.want) {
			t.Errorf("%v != %v", result, test.want)
			continue
		}
		for i, item := range result {
			if item != test.want[i] {
				t.Errorf("%s != %s", item, test.want[i])
			}
		}
	}
}
