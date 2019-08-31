package stats

import (
	"testing"
	"time"
)

var url = "http://192.168.0.28:420/Api.json"

func Test_getStats(t *testing.T) {
	type args struct {
		timeOut time.Duration
		url     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Rig-1",
			args: args{
				timeOut: 500 * time.Millisecond,
				url:     url,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getStats(tt.args.timeOut, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Stats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("stats: %+v", got)
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Stats() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func Test_PrtgStats(t *testing.T) {

	type args struct {
		timeOut time.Duration
		url     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Rig-1",
			args: args{
				timeOut: 500 * time.Millisecond,
				url:     "http://192.168.0.28:420/Api.json",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrtgStats(tt.args.timeOut, tt.args.url)
		})
	}
}
