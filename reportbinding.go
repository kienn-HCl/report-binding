package reportbinding

type ReportData struct {
    PageCount   int    `csv:"PageCount"`
    Author      string `csv:"Author"`
    Title       string `csv:"Title"`
    Filename    string `csv:"Filename"`
}

type ReportDatas []ReportData

