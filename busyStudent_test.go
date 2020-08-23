package busy_student

import "testing"

func Test_busyStudent(t *testing.T) {
	type args struct {
		startTime []int
		endTime   []int
		queryTime int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				startTime: []int{1, 2, 3},
				endTime:   []int{3, 2, 7},
				queryTime: 4,
			},
			want: 1,
		},
		{
			name: "Example2",
			args: args{
				startTime: []int{4},
				endTime:   []int{4},
				queryTime: 4,
			},
			want: 1,
		},
		{
			name: "Example3",
			args: args{
				startTime: []int{4},
				endTime:   []int{4},
				queryTime: 5,
			},
			want: 0,
		},
		{
			name: "Example4",
			args: args{
				startTime: []int{1, 1, 1, 1},
				endTime:   []int{1, 3, 2, 4},
				queryTime: 7,
			},
			want: 0,
		},
		{
			name: "Example5",
			args: args{
				startTime: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
				endTime:   []int{10, 10, 10, 10, 10, 10, 10, 10, 10},
				queryTime: 5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := busyStudent(tt.args.startTime, tt.args.endTime, tt.args.queryTime); got != tt.want {
				t.Errorf("busyStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}
