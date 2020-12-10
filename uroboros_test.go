package uroboros

import (
	"testing"
	"time"
)

func TestUroboros_Can(t *testing.T) {
	u := New(3, time.Minute)
	tests := []struct {
		name string
		u    *Uroboros
		tn   string
		want bool
	}{
		{
			name: "2010-07-13T06:15:05Z03:00->true",
			u:    u,
			tn:   "2010-07-13T06:15:05+03:00",
			want: true,
		},
		{
			name: "2010-07-13T06:15:15Z03:00->true",
			u:    u,
			tn:   "2010-07-13T06:15:15+03:00",
			want: true,
		},
		{
			name: "2010-07-13T06:15:25Z03:00->true",
			u:    u,
			tn:   "2010-07-13T06:15:25+03:00",
			want: true,
		},
		{
			name: "2010-07-13T06:15:35Z03:00->false",
			u:    u,
			tn:   "2010-07-13T06:15:35+03:00",
			want: false,
		},
		{
			name: "2010-07-13T06:16:05Z03:00->true",
			u:    u,
			tn:   "2010-07-13T06:16:05+03:00",
			want: true,
		},
		{
			name: "2010-07-13T06:16:13Z03:00->false",
			u:    u,
			tn:   "2010-07-13T06:16:13+03:00",
			want: false,
		},
		{
			name: "2010-07-13T06:16:25Z03:00->true",
			u:    u,
			tn:   "2010-07-13T06:16:25+03:00",
			want: true,
		},
	}
	for _, tt := range tests {
		nextTime, err := time.Parse(time.RFC3339, tt.tn)
		if err != nil {
			t.Error(err)
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			got := tt.u.Can(nextTime)
			if got != tt.want {
				t.Errorf("Can() = %v, want %v", got, tt.want)
			}
		})
	}
}
