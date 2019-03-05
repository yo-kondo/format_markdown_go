package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

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

	// TODO: 作成中

	// TODO: debug
	fmt.Printf("targetDir: %#v", files)
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
