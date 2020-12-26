package board

import (
	"reflect"
	"testing"
)

func TestSize_MakeMaze(t *testing.T) {
	type fields struct {
		width  int
		height int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Map
	}{
		{name: "board shap",
			fields: fields{width: 3, height: 3},
			want:   &Map{Maze: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Size{
				Width:  tt.fields.width,
				Height: tt.fields.height,
			}
			if got := s.MakeField(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Size.MakeField() = %v, want %v", got, tt.want)
			}
		})
	}
}
