package main

import (
	"log"
    "github.com/kienn-HCl/report-binding"
)

func main () {
    // err := genCsv()
    // if err != nil {
    //     log.Fatal(err)
    // }

    reports, err := reportbinding.NewReportDatas()
    if err != nil {
        log.Fatal(err)
    }

    //for _, pdf := range reports {
    //    fmt.Println(pdf.Filename)
    //}
    reportsNum := len(*reports)

    err = reports.UniteReport()
    if err != nil {
        log.Fatal(err)
    }

    err = reports.AddPagenum("./UniteReport/uniteReport.pdf", 3 + (reportsNum-1)/10)
    if err != nil {
        log.Fatal(err)
    }

    // pagenum, _ := api.PageCountFile("Unite/united.pdf")
    // println(pagenum)
}
