package main

import (
	"github.com/alexflint/go-arg"
	"io"
	"log"
	"os"
	"text/template"
)

const BUILD_DIR = "build/"
const TMPL_DIR = "templates/"

type Tool struct {
	Name string
	Url  string
}

var tool Tool
var buildFiles = []string{"README.md", "Makefile"}
var files = []string{".travis.yml"}
var folders = []string{"latest"}

func main() {
	arg.MustParse(&tool)

	for _, buildFile := range buildFiles {
		if err := tool.buildFile(TMPL_DIR+buildFile, BUILD_DIR+buildFile); err != nil {
			log.Fatal("Could not build file", err)
		}
	}

	for _, folder := range folders {
		if err := tool.copyFolder(TMPL_DIR+folder, BUILD_DIR+folder); err != nil {
			log.Fatal("Could not copy folder", err)
		}
	}

	for _, file := range files {
		if err := tool.copyFile(TMPL_DIR+file, BUILD_DIR+file); err != nil {
			log.Fatal("Could not copy file", err)
		}
	}

	log.Println("Done.")
}

func (tool Tool) buildFile(src, dst string) error {
	t, err := template.ParseFiles(src)
	if err != nil {
		return err
	}

	readme, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer readme.Close()
	t.Execute(readme, tool)

	return nil
}

func (tool Tool) copyFolder(src, dst string) error {
	err := os.MkdirAll(dst, os.ModePerm)
	if err != nil {
		return err
	}

	dir, _ := os.Open(src)
	objects, err := dir.Readdir(-1)

	for _, obj := range objects {
		srcFile := src + "/" + obj.Name()
		dstFile := dst + "/" + obj.Name()

		if obj.IsDir() {
			return tool.copyFolder(srcFile, dstFile)
		}

		if err := tool.copyFile(srcFile, dstFile); err != nil {
			return err
		}
	}

	return nil
}

func (tool Tool) copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}

	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Close()
}
