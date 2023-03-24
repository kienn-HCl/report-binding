package reportbinding

import (
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func GenCsv() error { // ディレクトリにある".pdf"で終わるファイルを読み取り"repsrtData.csv"を作成する。
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

func InitReportBinding() error {
    GenCsv()
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
