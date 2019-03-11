package resume

import (
	"testing"
)


func TestFilterSlice(t *testing.T)  {
	a := []string{"1", "2", "3", "", "", ""}

	a = FilterSlice(a)

	if 3 != len(a) {
		t.Error("TestFilterSlice: len(a) must be 3")
	}
}

func TestIsZhiLianResume(t *testing.T) {
	s := "aka 智联招聘"
	if !IsZhiLianResume(s) {
		t.Error("TestIsZhiLianResume error")
	}
}