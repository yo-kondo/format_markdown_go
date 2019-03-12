package info

import (
	"fmt"
	"format_markdown_go/cmd/date"
	"reflect"
	"testing"
)

func TestCreateInformation_Empty(t *testing.T) {
	in := []string{
		``,
	}
	got := createInformation(in)
	want := information{}

	if !structEqual(want, got) {
		fmt.Printf("want = %#v\n", want)
		fmt.Printf("got = %#v\n", got)
		t.Fatal("failed test")
	}
}

func TestCreateInformation1(t *testing.T) {
	in := []string{
		`# タイトル`,
		``,
		`## 基本情報`,
		``,
		`表紙 :\`,
		`![表紙](画像のパス)\`,
		`著者 : 著者\`,
		`出版社 : 出版社\`,
		`ISBN-10 : 9999999999\`,
		`ISBN-13 : 999-9999999999\`,
		`ASIN : XXXXXXXXXX\`,
		`発売日 : 2018/02/15\`,
		`Amazon : [amszonへのリンク](https://hogehoge.com)\`,
		`読了日 : 2019/03/11\`,
		`その他 : タグ1, タグ2`,
		``,
		`## 目次`,
	}
	got := createInformation(in)
	want := information{
		title: "タイトル",
		frontImage: []string{
			"![表紙](画像のパス)",
		},
		readingDate: date.NewDate(2019, 3, 11),
		author:      "著者",
		publisher:   "出版社",
		isbn10:      "9999999999",
		isbn13:      "999-9999999999",
		asin:        "XXXXXXXXXX",
		releaseDate: date.NewDate(2018, 2, 15),
		link:        "[amszonへのリンク](https://hogehoge.com)",
		tags:        "タグ1, タグ2",
	}

	if !structEqual(want, got) {
		fmt.Printf("want = %#v\n", want)
		fmt.Printf("got = %#v\n", got)
		t.Fatal("failed test")
	}
}

// information構造体の比較を行います。
func structEqual(x, y information) bool {
	if x.title != y.title {
		return false
	}
	if !reflect.DeepEqual(x.frontImage, y.frontImage) {
		return false
	}
	if !date.EqualDate(x.readingDate, y.readingDate) {
		return false
	}
	if x.author != y.author {
		return false
	}
	if x.publisher != y.publisher {
		return false
	}
	if x.isbn10 != y.isbn10 {
		return false
	}
	if x.isbn13 != y.isbn13 {
		return false
	}
	if x.asin != y.asin {
		return false
	}
	if !date.EqualDate(x.releaseDate, y.releaseDate){
		return false
	}
	if x.link != y.link {
		return false
	}
	if x.tags != y.tags {
		return false
	}
	return true
}
