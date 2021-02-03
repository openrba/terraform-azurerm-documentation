package main

import (
	"reflect"
	"testing"
)

func TestNewDocGenerator(t *testing.T) {
	tests := []struct {
		name string
		want DocGenerator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDocGenerator(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDocGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_docGenerator_AppendMenuItem(t *testing.T) {
	type args struct {
		title string
		path  string
	}
	tests := []struct {
		name string
		g    *docGenerator
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.AppendMenuItem(tt.args.title, tt.args.path)
		})
	}
}

func Test_docGenerator_WriteIndex(t *testing.T) {
	tests := []struct {
		name string
		g    *docGenerator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.g.WriteIndex()
		})
	}
}
