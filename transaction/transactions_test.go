package transactions

import (
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	type args struct {
		
		From   uint
		To     uint
		Amount int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateTransaction(tt.args.From, tt.args.To, tt.args.Amount)
		})
	}
}
