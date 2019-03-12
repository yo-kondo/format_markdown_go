package date

import (
	"testing"
	"time"
)

func TestNewDate(t *testing.T) {
	got := time.Date(2019, 01, 02, 0, 0, 0, 0, time.Local)
	want := NewDate(2019, 01,02)

	if got.Year() != want.Year() {
		t.Fatal("Yearが同一ではありません。")
	}
	if got.Month() != want.Month() {
		t.Fatal("Monthが同一ではありません。")
	}
	if got.Day() != want.Day() {
		t.Fatal("Dayが同一ではありません。")
	}
	if got.Hour() != want.Hour() {
		t.Fatal("Hourが同一ではありません。")
	}
	if got.Minute() != want.Minute() {
		t.Fatal("Minuteが同一ではありません。")
	}
	if got.Second() != want.Second() {
		t.Fatal("Secondが同一ではありません。")
	}
	if got.Nanosecond() != want.Nanosecond() {
		t.Fatal("Nanosecondが同一ではありません。")
	}
}

func TestEqualDate(t *testing.T) {
	got := time.Date(2019, 01, 02, 3, 4, 5, 6, time.Local)
	want := NewDate(2019, 01,02)

	if !EqualDate(got, want) {
		t.Fatal("時刻が違う比較で失敗しました。")
	}
}

func TestToStringDate(t *testing.T) {
	got := "2019/01/02"
	want := NewDate(2019, 01,02)
	if got != ToStringDate(want) {
		t.Fatal("日付の文字列が正しくありません。")
	}

}