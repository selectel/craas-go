package tokenv2

import "testing"

func Test_makeQueryString(t *testing.T) {
	limit := new(int)
	*limit = 1
	offset := new(int)
	*offset = 1
	type args struct {
		o Opts
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "all params",
			args: args{
				Opts{
					Limit:     limit,
					Offset:    offset,
					SortField: "name",
					SortType:  "asc",
					Search:    "tags",
					ScopeMode: "rw",
				},
			},
			want: "limit=1&offset=1&scope_mode=rw&search=tags&sort_field=name&sort_type=asc",
		},
		{
			name: "limit=1 and offset=1",
			args: args{
				Opts{
					Limit:  limit,
					Offset: offset,
				},
			},
			want: "limit=1&offset=1",
		},
		{
			name: "search=test",
			args: args{
				Opts{
					Search: "test",
				},
			},
			want: "search=test",
		},
		{
			name: "empty",
			args: args{
				Opts{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeQueryString(tt.args.o); got != tt.want {
				t.Errorf("makeQueryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
