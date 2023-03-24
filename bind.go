package reportbinding

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func NewReportDatas() (*ReportDatas, error) { // ディレクトリにある"reportData.csv"を読み込みreportDatasを作成する
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

    return &reports, nil
}

func (r *ReportDatas)UniteReport() error { // reportDatasにあるpdfを結合する
    reportFilenames := make([]string, 0, len(*r))
    for _, report := range *r {
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

func (r *ReportDatas)AddPagenum(filename string, startPage int) error { // filenameの各ページにstartPage(>0)から始まるページ番号をふる。
    for i:=0; i< startPage-1; i++ {
        err := pdfcpu.InsertPagesFile(filename, "", []string{"1"}, true, nil)
        if err != nil {
            return err
        }
    }

    err := pdfcpu.AddTextWatermarksFile(filename, "", nil, false, "%p / %P", "sc:1.0 abs, points: 12, pos:bc, rot:0, fillc:#000000, ma:10", nil)
    if err != nil {
        return err
    }

    if startPage < 2 {
        return nil
    }
    err = pdfcpu.RemovePagesFile(filename, "", []string{fmt.Sprintf("1-%d", startPage-1)}, nil)
    if err != nil {
        return err
    }
    return nil
}

