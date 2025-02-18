package day16

import "testing"

func Test16A(t *testing.T) {
	p1, p2, _ := Execute("../input/16a")
	if p1 != 7036 {
		t.Errorf("[1] Expected 7036, Found %d", p1)
	} else if p2 != 45 {
		t.Errorf("[2] Expected 45, Found %d", p2)

	}
}
func Test16B(t *testing.T) {
	p1, p2, _ := Execute("../input/16b")
	if p1 != 11048 {
		t.Errorf("[1] Expected 11048, Found %d", p1)
	} else if p2 != 64 {
		t.Errorf("[2] Expected 64, Found %d", p2)
	}
}

func Test16C(t *testing.T) {
	p1, _, _ := Execute("../input/16c")
	if p1 != 4013 {
		t.Errorf("[1] Expected 4013, Found %d", p1)
	}
}

// func TestReal(t *testing.T) {
// 	p1, p2, _ := Execute("../input/16")
// 	if p1 != 85420 {
// 		t.Errorf("[1] Expected 85420, Found %d", p1)
// 	} else if p2 != -2 {
// 		t.Errorf("[2] Expected ?, Found %d", p2)
// 	}
// }
