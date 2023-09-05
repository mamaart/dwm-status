package application

import (
	"os/exec"
	"strconv"
	"testing"
	"time"
)

func Test_getBrightness(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{name: "255", want: 255},
		{name: "10", want: 10},
		{name: "100", want: 100},
	}
	for _, tt := range tests {
		want := SetBrightness(tt.want)
		got, err := getBrightness()
		time.Sleep(time.Second)
		if err != nil {
			t.Errorf("getBrightness() failed with error %s", err)
		}
		if got != want {
			t.Errorf("getBrightness() = %v, want %v", got, want)
		}
	}
}

func SetBrightness(val int) int {
	exec.Command("brightnessctl", "set", strconv.Itoa(val)).Run()
	return (val * 100) / 255
}
