// time.Time構造体の日付だけを扱うラッパー関数群
package date

import "time"

// 時刻が0のtime.Timeを返します。
func NewDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

// time.Timeを日付のみ比較します。
func EqualDate(x, y time.Time) bool {
	return x.Year() == y.Year() &&
		x.Month() == y.Month() &&
		x.Day() == y.Day()
}

// time.Timeを"yyyy/mm/dd"形式に変換した文字列を返します。
func ToStringDate(x time.Time) string {
	return x.Format("2006/01/02")
}

// TODO: time.Parse関数のラッパーを作成する
