package dude

import "testing"

func TestSayMyName(t *testing.T) {
	expect := "Alpha"
	ob := SayMyName{Name: expect}
	if ob.SayIt() != expect {
		t.Fail()
	}
}
