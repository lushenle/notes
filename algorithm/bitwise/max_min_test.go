package bitwise

import "testing"

func Test_min(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testPositive01",
			args: args{a: 10, b: 20},
			want: 10,
		},
		{
			name: "testPositive02",
			args: args{a: 1, b: 21},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testPositive01",
			args: args{a: 10, b: 20},
			want: 20,
		},
		{
			name: "testPositive02",
			args: args{a: 1, b: 21},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min1(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testPositive01",
			args: args{a: 10, b: 20},
			want: 10,
		},
		{
			name: "testPositive02",
			args: args{a: 1, b: 21},
			want: 1,
		},
		{
			name: "testNegative01",
			args: args{a: -10, b: -20},
			want: -20,
		},
		{
			name: "testNegative02",
			args: args{a: -1, b: -21},
			want: -21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min1(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max1(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testPositive01",
			args: args{a: 10, b: 20},
			want: 20,
		},
		{
			name: "testPositive02",
			args: args{a: 1, b: 21},
			want: 21,
		},
		{
			name: "testNegative01",
			args: args{a: -10, b: -20},
			want: -10,
		},
		{
			name: "testNegative02",
			args: args{a: -1, b: -21},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max1(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max1() = %v, want %v", got, tt.want)
			}
		})
	}
}
