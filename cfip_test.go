package cfip

import "testing"

func TestContains(t *testing.T) {
	ip := "103.21.247.0"
	c, e := Contains(ip)
	if e != nil {
		t.Fatalf("%+v", e)
		return
	}
	if c {
		t.Logf("Yes, %s is inside.", ip)
	} else {
		t.Logf("No, %s is not inside.", ip)
	}
}

func TestNotContains(t *testing.T) {
	ip := "103.21.248.0"
	c, e := Contains(ip)
	if e != nil {
		t.Fatalf("%+v", e)
		return
	}
	if c {
		t.Logf("Yes, %s is inside.", ip)
	} else {
		t.Logf("No, %s is not inside.", ip)
	}
}

func BenchmarkContains(b *testing.B) {
	ip := "103.21.248.0"
	_, e := Contains(ip)
	if e != nil {
		b.Fatalf("%+v", e)
		return
	}
}
