package main

import "testing"

func TestCheckColisionWitharg(t *testing.T) {
	type args struct {
		board [8][8]int

	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				board: [8][8]int{
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},

				},

			},
			want: true,
		},
		{
			name: "test2",
			args: args{

				board: [8][8]int{
					{1, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 1, 0},
					{0, 0, 0, 1, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 1, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1},
					{0, 1, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 1, 0, 0, 0},
					{0, 0, 1, 0, 0, 0, 0, 0},


				},
				
			},
			want: true,
			
		},
		{
			name: "row",
			args: args{
				board: [8][8]int{
					{1, 1, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},

				},

			},
			want: false,
		},

				{
			name: "col",
			args: args{
				board: [8][8]int{
					{1, 0, 0, 0, 0, 0, 0, 0},
					{1, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},

				},

			},
			want: false,
		},
				{
			name: "diagnal",
			args: args{
				board: [8][8]int{
					{1, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 1},

				},

			},
			want: false,
		},



	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckColisionWitharg(tt.args.board); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}

}
