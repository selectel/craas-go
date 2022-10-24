package token

import "testing"

func Test_validateTokenTTL(t *testing.T) {
	type args struct {
		ttl TTL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid 12h",
			args: args{
				ttl: TTL12Hours,
			},
			wantErr: false,
		},
		{
			name: "valid 1y",
			args: args{
				ttl: TTL1Year,
			},
			wantErr: false,
		},
		{
			name: "invalid ttl",
			args: args{
				ttl: TTL("1d"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateTokenTTL(tt.args.ttl); (err != nil) != tt.wantErr {
				t.Errorf("validateTokenTTL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
