package chineseconvert

import (
	"testing"
)

func TestS2T(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{name: "s2t", args: "中国", want: "中國"},
		{name: "s2t", args: "123", want: "123"},
		{name: "s2t", args: "abc", want: "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := S2T(tt.args); got != tt.want {
				t.Errorf("S2T() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestT2S(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{name: "s2t", args: "中國", want: "中国"},
		{name: "s2t", args: "123", want: "123"},
		{name: "s2t", args: "abc", want: "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := T2S(tt.args); got != tt.want {
				t.Errorf("T2S() = %v, want %v", got, tt.want)
			}
		})
	}
}
