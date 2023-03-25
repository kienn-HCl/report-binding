package reportbinding

import (
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func GenCsv() error { // ディレクトリにある".pdf"で終わるファイルを読み取り"repsrtData.csv"を作成する。
    files, err := filepath.Glob("*.pdf")
    if err != nil {
        return err
    }

    reports := make([]ReportData, 0, len(files))

    for _, file := range files {
        ctx, err := pdfcpu.ReadContextFile(file)
        if err != nil {
            return err
        }
        reports = append(reports, ReportData{ctx.PageCount, ctx.Author, ctx.Title, file})
    }

    f, err := os.Create(csvFile)
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

    dirnames := []string{frontCoverDir, tabeleOfContentsDir, unitedReportDir, backCoverDir}
    for _, dirname := range dirnames {
        err := os.MkdirAll(dirname, 0755)
        if err != nil {
            return err
        }
    }
    return nil
}

