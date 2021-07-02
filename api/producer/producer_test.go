package producer

import "testing"

func TestProducer(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "one",
			args: args{
				msg: "hello world",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Producer(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Producer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
