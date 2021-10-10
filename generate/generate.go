package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"io/ioutil"
	"strings"
	"log"
	"html/template"
	"os"
	"path"
)

func textFromPath(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err == nil {
		s := string(b)
		return strings.TrimSpace(s), nil
	}
	return "", err
}


type Ketel struct {
	Merk *Merk
	Ketelnaam string
	Stappenplan string
	Ketelherkenning string
	Video string
	Path string
	Filename string
}


type Merk struct {
	Naam string
	Logo string
	Ketels []*Ketel
	Path string
	Filename string
}

func ytEmbedUrl(url string) string {
	const prefix = "https://www.youtube.com/watch?v="
	if len(url) > len(prefix) {
		if url[:len(prefix)] == prefix {
			return "https://www.youtube.com/embed/" + url[len(prefix):]
		}
	}
	return url
}


func main() {
	var err error



	var merken []*Merk
	var ketels []*Ketel
	var cur *Ketel
	var curMerk *Merk

	err = filepath.Walk("../ketels", func(filepath string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", filepath, err)
			return err
		}

		if info.IsDir() {
			dir, path := path.Split(filepath)
			if dir == "../" {
			} else if dir == "../ketels/" {
				// Dit is een merk
				if curMerk != nil {
					merken = append(merken, curMerk)
				}
				curMerk = &Merk{
					Filename: path,
					Path:filepath}
			} else {
				if cur != nil {
					ketels = append(ketels, cur)
				}
				cur = &Ketel{
					Merk: curMerk,
					Ketelnaam: "ketelnaam",
					Filename: path,
					Path:filepath}

				curMerk.Ketels = append(curMerk.Ketels, cur)
			}
		}

		if info.Name() == "logo.svg" {
			curMerk.Logo = filepath
		}

		if info.Name() == "merknaam.txt" {
			curMerk.Naam, _ = textFromPath(filepath)
		}

		if info.Name() == "ketel.png" {
			cur.Ketelherkenning = filepath
		}

		if info.Name() == "ketel.png" {
			cur.Ketelherkenning = filepath
		}
		if info.Name() == "naam.txt" {
			cur.Ketelnaam, _ = textFromPath(filepath)
		}

		if info.Name() == "handleiding.png" {
			cur.Stappenplan = filepath
		}
		if info.Name() == "video" {
			vidUrl, _ := textFromPath(filepath)
			cur.Video = ytEmbedUrl(vidUrl)
		}

		return nil
	})

	if curMerk != nil {
		if cur != nil {
			curMerk.Ketels = append(curMerk.Ketels, cur)
		}
		merken = append(merken, curMerk)
	}
	if cur != nil {
		ketels = append(ketels, cur)
	}


	if err != nil {
		fmt.Printf("error walking the path %v\n", err)
		return
	}

	// for _, k := range ketels {
	// 	//		fmt.Printf("%s, %#v\n", k.Ketelnaam, k.Merk)
	// }

	templ, err := template.ParseFiles("ketel-template.html")
	if err != nil {
		log.Fatalf("Could not open template")
	}

	for ki := range ketels {
		p := "out/out/" + ketels[ki].Path
		err := os.MkdirAll(p, fs.ModePerm)
		if err != nil {
			log.Fatalf("Error creating directory")
		}
		f, err := os.Create(p + "/" + ketels[ki].Filename + ".html")
		if err == nil {
			templ.Execute(f, ketels[ki])
		}
		f.Close()
	}

	merktempl, err := template.ParseFiles("merk-template.html")
	if err != nil {
		log.Fatalf("Could not open merk template")
	}

	for mi := range merken {
		p := "out/out/" + merken[mi].Path
		err := os.MkdirAll(p, fs.ModePerm)
		if err != nil {
			log.Fatalf("Error creating directory")
		}
		f, err := os.Create(p + "/" + merken[mi].Filename + ".html")
		if err == nil {
			merktempl.Execute(f, merken[mi])
		}
		f.Close()
	}

	{
		overzichttempl, err := template.ParseFiles("merkkeuze-template.html")
		if err != nil {
			log.Fatalf("Could not open merkkeuze template")
		}
		merkOverzicht := struct{
			Merken []*Merk
		}{merken}

		p := "out/out/" + "../ketels/overzicht.html"
		f, err := os.Create(p)
		if err == nil {
			overzichttempl.Execute(f, merkOverzicht)
		}
		f.Close()
	}


}
