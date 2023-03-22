package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

type PdfData struct {
    Title       string `csv:"Title"`
    Author      string `csv:"Author"`
    PageCount   int    `csv:"PageCount"`
    Filename    string `csv:"Filename"`
}

func genCsv() error {
    files, err := filepath.Glob("./*.pdf")
    if err != nil {
        return err
    }

    pdfs := make([]PdfData, 0, len(files))

    for _, file := range files {
        // aaa, _ := api.InfoFile(file, nil, nil)
        // println(file)
        // println(aaa[1], aaa[5], aaa[6])
        con, err := pdfcpu.ReadContextFile(file)
        if err != nil {
            return err
        }
        pdfs = append(pdfs, PdfData{con.Title, con.Author, con.PageCount, file})
    }

    f, err := os.Create("pdfData.csv")
    if err != nil {
        return err
    }
    defer f.Close()

    err = gocsv.MarshalFile(&pdfs, f)
    if err != nil {
        return err
    }
    return nil
}

func readCsv() ([]PdfData, error) {
    f, err := os.Open("pdfData.csv")
    if err != nil {
        return nil, err
    }
    defer f.Close()

    var pdfs []PdfData
    err = gocsv.UnmarshalFile(f, &pdfs)
    if err != nil {
        return nil, err
    }

    return pdfs, nil
}

func unitePdf(pdfs []PdfData) error {
    pdfFilename := make([]string, 0, len(pdfs))
    for _, pdf := range pdfs {
        pdfFilename = append(pdfFilename, pdf.Filename)
    }

    err := os.MkdirAll("Unite", 0755)
    if err != nil {
        return err
    }
    err = pdfcpu.MergeCreateFile(pdfFilename, "Unite/united.pdf", nil)
    if err != nil {
        return err
    }
    return nil
}

func addPagenum(startPage int) error {
    for i:=0; i< startPage; i++ {
        err := pdfcpu.InsertPagesFile("Unite/united.pdf", "", []string{"1"}, true, nil)
        if err != nil {
            return err
        }
    }

    err := pdfcpu.AddTextWatermarksFile("Unite/united.pdf", "", nil, false, "%p / %P", "sc:1.0 abs, points: 12, pos:bc, rot:0, fillc:#000000, ma:10", nil)
    if err != nil {
        return err
    }

    err = pdfcpu.RemovePagesFile("Unite/united.pdf", "", []string{fmt.Sprintf("1-%d", startPage)}, nil)
    if err != nil {
        return err
    }
    return nil
}

func main () {
    // err := genCsv()
    // if err != nil {
    //     log.Fatal(err)
    // }

    pdfs, err := readCsv()
    if err != nil {
        log.Fatal(err)
    }

    //for _, pdf := range pdfs {
    //    fmt.Println(pdf.Filename)
    //}
    pdfNum := len(pdfs)

    err = unitePdf(pdfs)
    if err != nil {
        log.Fatal(err)
    }

    err = addPagenum(2 + pdfNum/10)
    if err != nil {
        log.Fatal(err)
    }

    // pagenum, _ := api.PageCountFile("Unite/united.pdf")
    // println(pagenum)
}
