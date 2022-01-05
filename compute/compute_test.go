package compute

import "testing"

func TestAdd(t *testing.T) {
	a := 10
	b := 20

	want := 30
	actual := Add(a, b)
	if want != actual {
		t.Errorf("[Add 参数:%d %d ][期望:%d][实际:%d]\n", a, b, want, actual)
	}
}

func TestMul(t *testing.T) {
	a := 10
	b := 20

	want := 200
	actual := Mul(a, b)
	if want != actual {
		t.Errorf("[Mul 参数:%d %d ][期望:%d][实际:%d]\n", a, b, want, actual)
	}
}

func TestDiv(t *testing.T) {
	a := 10
	b := 20

	want := 2
	actual := Div(b, a)
	if want != actual {
		t.Errorf("[div 参数:%d %d ][期望:%d][实际:%d]\n", a, b, want, actual)
	}
}
