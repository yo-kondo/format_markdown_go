package main

import (
	"fmt"
	"format_markdown_go/cmd/info"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
)

// 設定ファイル
type Config struct {
	// 対象ディレクトリ
	TargetDir string
}

// エントリポイント
func main() {
	conf := readSettings("settings.toml")
	existDirectory(conf.TargetDir)
	files := getMarkdownFile(conf.TargetDir)

	for _, file := range files {
		lines := getAllText(file)

		lines = addHeaderAfterEmptyLine(lines)
		lines = deleteQuoteEmpty(lines)

		// 基本情報を整形
		lines = info.ChangeInformation(lines)
		lines = deleteDuplicationEmpty(lines)

		// TODO: debug
		fmt.Printf("targetDir: %#v", lines)
	}

	// TODO: 作成中
}

// 設定ファイル settings.toml を読み込みます。
func readSettings(settingFile string) Config {
	var conf Config
	if _, err := toml.DecodeFile(settingFile, &conf); err != nil {
		log.Fatal("設定ファイル settings.toml の読み込みでエラーが発生しました: ", err)
	}
	return conf
}

// 指定したディレクトリが存在するか確認します。
func existDirectory(dir string) {
	fInfo, err := os.Lstat(dir)
	if err != nil {
		log.Fatal("設定ファイルの targetDir で指定したディレクトリが存在しません。: ", err)
	}
	if !fInfo.IsDir() {
		log.Fatal("設定ファイルの targetDir は、ファイルではなくディレクトリを指定してください。: ", err)
	}
}

// 指定したディレクトリからMarkdownファイルのリストを返します。
func getMarkdownFile(dir string) []string {
	var files []string
	err1 := filepath.Walk(dir,
		func(path string, fInfo os.FileInfo, err2 error) error {
			if err2 != nil {
				return err2
			}
			if fInfo.IsDir() {
				return nil
			}
			if filepath.Ext(path) != ".md" {
				return nil
			}
			files = append(files, path)
			return nil
		})
	if err1 != nil {
		log.Fatal("ファイル名の取得でエラーが発生しました。: ", err1)
	}
	return files
}

// 指定したファイルを行単位で読み込みます。
func getAllText(fileName string) []string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("ファイルの読み込みでエラーが発生しました。: ", err)
	}
	return strings.Split(string(data), "\n")
}

var regHeader = regexp.MustCompile(`^#+ `)

// 見出し（#、##、###など）の下に空行を追加します。
func addHeaderAfterEmptyLine(lines []string) []string {
	var rtnLines []string

	for i, line := range lines {
		if regHeader.MatchString(line) {
			if lines[i+1] != "" {
				rtnLines = append(rtnLines, line)
				rtnLines = append(rtnLines, "")
				continue
			}
		}
		rtnLines = append(rtnLines, line)
	}
	return rtnLines
}

var regQuote = regexp.MustCompile(`> $`)

// >だけの行は後ろのスペースを除去します。
func deleteQuoteEmpty(lines []string) []string {
	var rtnLines []string

	for _, line := range lines {
		if regQuote.MatchString(line) {
			rtnLines = append(rtnLines, strings.TrimRight(line, " "))
		} else {
			rtnLines = append(rtnLines, line)
		}
	}
	return rtnLines
}

// 2つ連続した空行を1つにします。
func deleteDuplicationEmpty(lines []string) []string {
	var rtnLines []string
	isEmpty := false

	for _, line := range lines {

		if isEmpty && line == "" {
			continue
		}

		if line == "" {
			isEmpty = true
		} else {
			isEmpty = false
		}
		rtnLines = append(rtnLines, line)
	}
	return rtnLines
}
