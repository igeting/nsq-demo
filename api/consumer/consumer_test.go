package consumer

import (
	"fmt"
	"testing"
	"time"
)

func TestConsumer(t *testing.T) {
	type args struct {
		task func(addr, msg string) error
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
				task: func(addr, msg string) error {
					fmt.Printf("<<from %s, msg:%s>>", addr, msg)
					return nil
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Consumer(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("Consumer() error = %v, wantErr %v", err, tt.wantErr)
			}
			time.Sleep(time.Second * 1)
		})
	}
}
