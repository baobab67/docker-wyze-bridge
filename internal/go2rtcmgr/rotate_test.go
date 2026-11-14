package go2rtcmgr

import "testing"

func TestParseRotate(t *testing.T) {
	m := parseRotate(" Front_Doorbell:90, porch:180,bad:45,junk,x:abc,garage:270")
	want := map[string]int{"front_doorbell": 90, "porch": 180, "garage": 270}
	if len(m) != len(want) {
		t.Fatalf("got %v", m)
	}
	for k, v := range want {
		if m[k] != v {
			t.Errorf("%s: got %d want %d", k, m[k], v)
		}
	}
}
