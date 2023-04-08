package reportbinding

import (
	"embed"
	"fmt"
	"path"
	"path/filepath"

	"github.com/signintech/gopdf"
)

func genMokuzi(pdf *gopdf.GoPdf, pdfWidth, fontSize float64) error {
    err := pdf.SetFont("ipaexm", "", fontSize)
    if err != nil {
        return err
    }
    pdf.SetX(0.1*pdfWidth)
    pdf.SetLineWidth(0.1)
    pdf.CellWithOption(&gopdf.Rect{W: 0.8*pdfWidth,H: 2*fontSize}, "目次",gopdf.CellOption{Align: gopdf.Center, Float: gopdf.Bottom, Border: gopdf.Bottom})
    return nil
}

func genTitles(pdf *gopdf.GoPdf, pdfWidth, fontSize float64, rowsLimit int) error {
    reports, err := NewReportDatas()
    if err != nil {
        return err
    }
    reportsNum := len(*reports)
    titlesWidth := 0.8*pdfWidth - 1.4*fontSize
    pageNum := 4 + (reportsNum-1) / rowsLimit
    dotSize := 1.3
    pdf.SetLineWidth(dotSize)
    pdf.SetCustomLineType([]float64{dotSize,5}, dotSize)
    err = pdf.SetFont("ipaexm", "", fontSize)
    if err != nil {
        return err
    }

    for i := 0; i< reportsNum; i++ {
        if i%rowsLimit==0 && i!=0 {
            pdf.AddPage()
            pdf.SetXY(0.1*pdfWidth, 100)
            pdf.SetLineWidth(dotSize)
            pdf.SetCustomLineType([]float64{dotSize,5}, dotSize)
        }
        strs, err := pdf.SplitText((*reports)[i].Title, titlesWidth)
        if err!=nil {
            return err
        }
        for i := 0; i < len(strs)-1; i++ {
            pdf.CellWithOption(nil,strs[i], gopdf.CellOption{Float: gopdf.Bottom})
            pdf.SetY(pdf.GetY()+fontSize)
        }
        pdf.Cell(nil, strs[len(strs)-1])
        pdf.Line(pdf.GetX(),pdf.GetY() + 0.5*fontSize, 0.1*pdfWidth + titlesWidth, pdf.GetY() + 0.5*fontSize)

        pdf.SetX(0.9*pdfWidth - fontSize)
        pdf.CellWithOption(&gopdf.Rect{W: fontSize,H: fontSize}, fmt.Sprint(pageNum), gopdf.CellOption{Align: gopdf.Right})
        pageNum += (*reports)[i].PageCount

        pdf.SetXY(0.1*pdfWidth, pdf.GetY()+1.5*fontSize)
        //pdf.MultiCellWithOption(&gopdf.Rect{W: 0.8*pdfWidth, H: pdfWidth}, (*reports)[i].Author, gopdf.CellOption{Align: gopdf.Right, Float: gopdf.Bottom})
        strs, err = pdf.SplitText((*reports)[i].Author, 0.8*pdfWidth)
        if err!=nil {
            return err
        }
        for i := 0; i < len(strs)-1; i++ {
            pdf.CellWithOption(&gopdf.Rect{W: 0.8*pdfWidth, H: fontSize},strs[i], gopdf.CellOption{Align: gopdf.Right, Float: gopdf.Bottom})
            pdf.SetY(pdf.GetY()+fontSize)
        }
        pdf.CellWithOption(&gopdf.Rect{W: 0.8*pdfWidth, H: fontSize},strs[len(strs)-1], gopdf.CellOption{Align: gopdf.Right, Float: gopdf.Bottom})
        pdf.SetY(pdf.GetY()+0.5*fontSize)
    }

    return nil
}

//go:embed font
var local embed.FS

func GenTableOfContentsPdf(fontSize float64, rowsLimit int) error {
    pdf := gopdf.GoPdf{}
    pdf.Start(gopdf.Config{ PageSize: *gopdf.PageSizeA4})

    f, err := local.Open(path.Join("font","ipaexm.ttf"))
    if err != nil {
        return err
    }
    err = pdf.AddTTFFontByReader("ipaexm", f)
    if err != nil {
        return err
    }
    pdf.AddPage()

    pdfWidth := gopdf.PageSizeA4.W

    pdf.Br(120)
    err = genMokuzi(&pdf, pdfWidth, 24)
    if err != nil {
        return err
    }

    pdf.Br(30)
    pdf.SetX(0.1*pdfWidth)

    err = genTitles(&pdf, pdfWidth, fontSize, rowsLimit)
    if err != nil {
        return err
    }

    err = pdf.WritePdf(filepath.Join( tabeleOfContentsDir, tabeleOfContentsFile))
    if err != nil {
        return err
    }
    return nil
}
