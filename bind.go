package reportbinding

import (
	"fmt"
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

    err := pdfcpu.MergeCreateFile(reportFilenames, "./UnitedReport/unitedReport.pdf", nil)
    if err != nil {
        return err
    }
    return nil
}

func (r *ReportDatas)GenTabeleOfContents() {
    // to do
}

func AddPagenum(filename string, startPage int) error { // filenameの各ページにstartPage(>0)から始まるページ番号をふる。
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

func getPdfFilenames(dirnames []string) ([]string, error) {
    files := make([]string, 0, len(dirnames))
    for _, dirname := range dirnames {
        file, err := filepath.Glob(dirname + "/*.pdf")
        if err != nil {
            return nil, err
        }
        if file == nil {
            return nil, fmt.Errorf("error bindReport: %v, no pdf file", dirname)
        }
        if len(file) > 1 {
            return nil, fmt.Errorf("error bindReport: %v, too many pdf files", dirname)
        }
        files = append(files, file[0])
    }
    return files, nil
}

func addBlankpage(outputfile string) error {
    pagenum, err := pdfcpu.PageCountFile(outputfile)
    if err != nil {
        return err
    }
    for i := pagenum%4; i!=3; i++ {
        err := pdfcpu.InsertPagesFile(outputfile, "", []string{fmt.Sprint(pagenum)}, false, nil)
        if err != nil {
            return err
        }
    }
    return nil
}

func BindReport() error {
    reports, err := NewReportDatas()
    if err != nil {
        return err
    }

    err = reports.UniteReport()
    if err != nil {
        return err
    }

    dirnames := []string{"./FrontCover", "./TableOfContents", "./UnitedReport", "./BackCover"}
    files, err := getPdfFilenames(dirnames)
    if err != nil {
        return err
    }

    frontCoverNum, err := pdfcpu.PageCountFile(files[0])
    if err != nil {
        return err
    }
    tabeleOfContentsNum, err := pdfcpu.PageCountFile(files[1])
    if err != nil {
        return err
    }
    err = AddPagenum(files[2], frontCoverNum + tabeleOfContentsNum + 1)
    if err != nil {
        return err
    }

    os.MkdirAll("./BoundReport", 0755)
    outputfile :=  "./BoundReport/boundReport.pdf"
    err = pdfcpu.MergeCreateFile(files[0:3], outputfile, nil)
    if err != nil {
        return err
    }

    err = addBlankpage(outputfile)
    if err != nil {
        return err
    }

    err = pdfcpu.MergeAppendFile([]string{files[3]}, outputfile, nil)
    if err != nil {
        return err
    }
    return nil
}
