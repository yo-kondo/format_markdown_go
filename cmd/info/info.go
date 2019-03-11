package info

import (
	"log"
	"strings"
	"time"
)

// 基本情報
type information struct {
	// タイトル
	title string
	// 表紙の画像パス
	frontImage []string
	// 読了日
	readingDate time.Time
	// 著者
	author string
	// 出版社
	publisher string
	// ISBN-10
	isbn10 string
	// ISBN-13
	isbn13 string
	// ASIN
	asin string
	// 発売日
	releaseDate time.Time
	// Amazon URL
	link string
	// タグ
	tags string
}

// 基本情報を変換したリストを返します。
func ChangeInformation(input []string) []string {
	// info := createInformation(input)

	// TODO: 未作成
	return nil
}

// 変換前の基本情報を構造体にセットします。
func createInformation(input []string) information {
	info := information{}

	// 基本情報を抜き出す
	for _, line := range input {
		// タイトル
		if strings.HasPrefix(line, "# ") {
			info.title = strings.TrimPrefix(line, "# ")
		}

		// 表紙の画像パス
		if strings.HasPrefix(line, "![表紙](") {
			info.frontImage = append(info.frontImage, strings.TrimSuffix(line, `\`))
		}

		// 読了日
		if strings.HasPrefix(line, "読了日 : ") {
			dateStr := strings.TrimPrefix(line, "読了日 : ")
			dateStr = strings.TrimSuffix(dateStr, `\`)

			date, err := time.Parse("2006/01/02", dateStr)
			if err != nil {
				log.Fatal(info.title, "の読了日が日付に変換できませんでした。: ", err)
			}
			info.readingDate = date
		}

		// 著者
		if strings.HasPrefix(line, "著者 : ") {
			author := strings.TrimPrefix(line, "著者 : ")
			author = strings.TrimSuffix(author, `\`)
			info.author = author
		}

		// 出版社
		if strings.HasPrefix(line, "出版社 : ") {
			publisher := strings.TrimPrefix(line, "出版社 : ")
			publisher = strings.TrimSuffix(publisher, `\`)
			info.publisher = publisher
		}

		// ISBN-10
		if strings.HasPrefix(line, "ISBN-10 : ") {
			isbn10 := strings.TrimPrefix(line, "ISBN-10 : ")
			isbn10 = strings.TrimSuffix(isbn10, `\`)

			if isbn10 == "－" {
				info.isbn10 = "-"
			} else {
				info.isbn10 = isbn10
			}
		}

		// ISBN-13
		if strings.HasPrefix(line, "ISBN-13 : ") {
			isbn13 := strings.TrimPrefix(line, "ISBN-13 : ")
			isbn13 = strings.TrimSuffix(isbn13, `\`)

			if isbn13 == "－" {
				info.isbn13 = "-"
			} else {
				info.isbn13 = isbn13
			}
		}

		// ASIN
		if strings.HasPrefix(line, "ASIN : ") {
			asin := strings.TrimPrefix(line, "ASIN : ")
			asin = strings.TrimSuffix(asin, `\`)

			if asin == "－" {
				info.asin = "-"
			} else {
				info.asin = asin
			}
		}

		// 発売日
		if strings.HasPrefix(line, "発売日 : ") {
			dateStr := strings.TrimPrefix(line, "発売日 : ")
			dateStr = strings.TrimSuffix(dateStr, `\`)

			date, err := time.Parse("2006/01/02", dateStr)
			if err != nil {
				log.Fatal(info.title, "の発売日が日付に変換できませんでした。: ", err)
			}
			info.releaseDate = date
		}

		// Amazon URL
		if strings.HasPrefix(line, "Amazon : ") {
			link := strings.TrimPrefix(line, "Amazon : ")
			link = strings.TrimSuffix(link, `\`)
			info.link = link
		}

		// タグ
		if strings.HasPrefix(line, "その他 : ") {
			tags := strings.TrimPrefix(line, "その他 : ")
			tags = strings.TrimSuffix(tags, `\`)
			info.tags = tags
		}
	}
	return info
}
