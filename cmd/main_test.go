package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetMarkdownFile(t *testing.T) {
	got := getMarkdownFile("../test/test01")
	want := []string{
		`..\test\test01\dir01\file02.md`,
		`..\test\test01\file01.md`,
	}

	if !reflect.DeepEqual(want, got) {
		fmt.Printf("%#v\n", want)
		fmt.Printf("%#v\n", got)
		t.Fatal("failed test")
	}
}

func TestAddHeaderAfterEmptyLine(t *testing.T) {
	in := []string{
		`# head1`,
		``,
		`## head2`,
		`line1`,
		`line2`,
		``,
		`## head3`,
		``,
		`line3`,
		``,
	}
	want := []string{
		`# head1`,
		``,
		`## head2`,
		``,
		`line1`,
		`line2`,
		``,
		`## head3`,
		``,
		`line3`,
		``,
	}
	got := addHeaderAfterEmptyLine(in)
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("%#v\n", want)
		fmt.Printf("%#v\n", got)
		t.Fatal("failed test")
	}
}

func TestDeleteQuoteEmpty(t *testing.T) {
	in := []string{
		`# head1`,
		``,
		`line1`,
		`line2`,
		``,
		`> line3`,
		`> `,
		`> line4`,
		`>`,
		`> `,
	}
	want := []string{
		`# head1`,
		``,
		`line1`,
		`line2`,
		``,
		`> line3`,
		`>`,
		`> line4`,
		`>`,
		`>`,
	}
	got := deleteQuoteEmpty(in)
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("%#v\n", want)
		fmt.Printf("%#v\n", got)
		t.Fatal("failed test")
	}
}

func TestDeleteDuplicationEmpty(t *testing.T) {
	in := []string{
		`# head1`,
		``,
		``,
		`line1`,
		``,
		``,
	}
	want := []string{
		`# head1`,
		``,
		`line1`,
		``,
	}
	got := deleteDuplicationEmpty(in)
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("%#v\n", want)
		fmt.Printf("%#v\n", got)
		t.Fatal("failed test")
	}
}
