package structures

import (
	"testing"

	"github.com/rusq/fsadapter"
	"github.com/rusq/slack"
)

func TestIndex_Marshal(t *testing.T) {
	type args struct {
		fs fsadapter.FS
	}
	tests := []struct {
		name    string
		fields  ExportIndex
		args    args
		wantErr bool
	}{
		{
			"x",
			ExportIndex{Channels: []slack.Channel{{IsChannel: true}}},
			args{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx := &ExportIndex{
				Channels: tt.fields.Channels,
				Groups:   tt.fields.Groups,
				MPIMs:    tt.fields.MPIMs,
				DMs:      tt.fields.DMs,
				Users:    tt.fields.Users,
			}
			if err := idx.Marshal(tt.args.fs); (err != nil) != tt.wantErr {
				t.Errorf("Index.Marshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
