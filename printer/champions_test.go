package printer

import (
	"bytes"
	"interview-test/model"
	"testing"
	"time"
)

func TestPrintChampions(t *testing.T) {
	th, _ := time.Parse("4:05.000", "1:23.123")
	r := []model.Result{{1, "1", "test", 1, th}}
	var w bytes.Buffer
	PrintChampions(r, &w)
	got := w.String()
	want := ``

	if got == want {
		t.Errorf("PrintChampions() = \n%s\n, want \n%s", got, want)
	}

}
