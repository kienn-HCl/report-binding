package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

type ReportData struct {
    PageCount   int    `csv:"PageCount"`
    Author      string `csv:"Author"`
    Title       string `csv:"Title"`
    Filename    string `csv:"Filename"`
}

type ReportDatas []ReportData

func genCsv() error { // ディレクトリにある".pdf"で終わるファイルを読み取り"repsrtData.csv"を作成する。
    files, err := filepath.Glob("./*.pdf")
    if err != nil {
        return err
    }

    reports := make(ReportDatas, 0, len(files))

    for _, file := range files {
        // aaa, _ := api.InfoFile(file, nil, nil)
        // println(file)
        // println(aaa[1], aaa[5], aaa[6])
        ctx, err := pdfcpu.ReadContextFile(file)
        if err != nil {
            return err
        }
        reports = append(reports, ReportData{ctx.PageCount, ctx.Author, ctx.Title, file})
    }

    f, err := os.Create("reportData.csv")
    if err != nil {
        return err
    }
    defer f.Close()

    err = gocsv.MarshalFile(&reports, f)
    if err != nil {
        return err
    }
    return nil
}

func initReportBinding() error {
    genCsv()
    err := os.MkdirAll("./Hyougi", 0755)
    if err != nil {
        return err
    }
    err = os.MkdirAll("./Mokuzi", 0755)
    if err != nil {
        return err
    }
    err = os.MkdirAll("./UraByougi", 0755)
    if err != nil {
        return err
    }
    return nil
}

func newReportDatas() (ReportDatas, error) { // ディレクトリにある"reportData.csv"を読み込みreportDatasを作成する
    f, err := os.Open("reportData.csv")
    if err != nil {
        return nil, err
    }
    defer f.Close()

    var reports ReportDatas
    err = gocsv.UnmarshalFile(f, &reports)
    if err != nil {
        return nil, err
    }

    return reports, nil
}

func (r ReportDatas)uniteReport() error { // reportDatasにあるpdfを結合する
    reportFilenames := make([]string, 0, len(r))
    for _, report := range r {
        reportFilenames = append(reportFilenames, report.Filename)
    }

    err := os.MkdirAll("UniteReport", 0755)
    if err != nil {
        return err
    }
    err = pdfcpu.MergeCreateFile(reportFilenames, "UniteReport/uniteReport.pdf", nil)
    if err != nil {
        return err
    }
    return nil
}

func (r ReportDatas)addPagenum(startPage int) error {
    for i:=0; i< startPage; i++ {
        err := pdfcpu.InsertPagesFile("UniteReport/uniteReport.pdf", "", []string{"1"}, true, nil)
        if err != nil {
            return err
        }
    }

    err := pdfcpu.AddTextWatermarksFile("UniteReport/uniteReport.pdf", "", nil, false, "%p / %P", "sc:1.0 abs, points: 12, pos:bc, rot:0, fillc:#000000, ma:10", nil)
    if err != nil {
        return err
    }

    err = pdfcpu.RemovePagesFile("UniteReport/uniteReport.pdf", "", []string{fmt.Sprintf("1-%d", startPage)}, nil)
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

    reports, err := newReportDatas()
    if err != nil {
        log.Fatal(err)
    }

    //for _, pdf := range reports {
    //    fmt.Println(pdf.Filename)
    //}
    reportsNum := len(reports)

    err = reports.uniteReport()
    if err != nil {
        log.Fatal(err)
    }

    err = reports.addPagenum(2 + reportsNum/10)
    if err != nil {
        log.Fatal(err)
    }

    // pagenum, _ := api.PageCountFile("Unite/united.pdf")
    // println(pagenum)
}
