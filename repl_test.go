package main

import (
	"testing"
	"reflect"
)
func TestCleanInput(t *testing.T){
	
	cases:=[]struct {
		input string 
		want []string 
	}{
		{
			input:"  hello  world   ", 
			want:[]string{"hello", "world"},
		},
		{
			input:"buba chari biscuit", 
			want:[]string{"buba", "chari", "biscuit"},
		},
	};

	for _, c:= range cases{
		got:=cleanInput(c.input)
		if !reflect.DeepEqual(c.want, got){
			t.Fatalf("expected: %v, got: %v", c.want, got);
		}
	}

}