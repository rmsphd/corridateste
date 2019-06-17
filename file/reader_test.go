package file

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func ExampleReader() {
	file := strings.NewReader("This is a test header\nThis is a test string")

	got, _ := Reader(file)
	fmt.Printf("%v", got)

	// Output: [[This is a test string]]
}

func TestReader(t *testing.T) {
	file := strings.NewReader("This is a test header\nThis is a test string")
	want := [][]string{{"This", "is", "a", "test", "string"}}
	got, err := Reader(file)
	if err != nil {
		t.Errorf("Reader() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Reader() = %v, want %v", got, want)
	}
}

func TestReader_Empty(t *testing.T) {
	file := strings.NewReader("This is a test header\n")
	want := [][]string{}
	got, err := Reader(file)
	if err != nil {
		t.Errorf("Reader() error = %v", err)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Reader() = %v, want %v", got, want)
	}
}

type args struct{}

func (r args) Read(b []byte) (int, error) {
	return 0, errors.New("test")
}

func TestReader_Error(t *testing.T) {
	a := args{}
	got, err := Reader(a)
	if err == nil {
		t.Errorf("Reader() error = %v", err)
		return
	}
	if got != nil {
		t.Errorf("Reader() = %v", got)
	}
}
