package main

import (
	"os"
    "log"

	"github.com/kienn-HCl/report-binding"
	"github.com/urfave/cli/v2"
)

func initDirectory(c *cli.Context) error {
    if c.Bool("csv") {
        return reportbinding.GenCsv()
    }
    return reportbinding.InitReportBinding()
}

func uniteReport(c *cli.Context) error {
    reports, err := reportbinding.NewReportDatas()
    if err != nil {
        return err
    }
    return reports.UniteReport()
}

func addPagenum(c *cli.Context) error {
    return reportbinding.AddPagenum(c.Path("filepath"), c.Int("startPagenum"))
}

func genTableOfContents(c *cli.Context) error {
    return reportbinding.GenTableOfContentsPdf(c.Float64("fontsize"), c.Int("rows"))
}

func bindReport(c *cli.Context) error {
    return reportbinding.BindReport()
}

func main () {
    app := cli.NewApp()

    app.Commands = []*cli.Command{
        {
            Name: "init",
            Aliases: []string{"i"},
            Usage: "initialize directory",
            Flags: []cli.Flag{
                &cli.BoolFlag{
                    Name: "csv",
                    Aliases: []string{"c"},
                    Usage: "only output csv file",
                    Value: false,
                },
            },
            Action: initDirectory,
        },
        {
            Name: "uniteReport",
            Aliases: []string{"u"},
            Usage: "unite report and output \"./UnitedReport\" directory",
            Action: uniteReport,
        },
        {
            Name: "addPagenum",
            Aliases: []string{"a"},
            Usage: "add pagenum",
            Flags: []cli.Flag{
                &cli.PathFlag{
                    Name: "filepath",
                    Aliases: []string{"f"},
                    Usage: "File to which you want to add pagenum",
                    Value: "./UnitedReport/unitedReport.pdf",
                },
                &cli.IntFlag{
                    Name: "startPagenum",
                    Aliases: []string{"s"},
                    Usage: "the number of pagenum you want to start with",
                    Value: 1,
                },
            },
            Action: addPagenum,
        },
        {
            Name: "genTableOfContents",
            Aliases: []string{"t"},
            Usage: "make tableOfContents.pdf and output \"./TableOfContents\" directory",
            Flags: []cli.Flag{
                &cli.Float64Flag{
                    Name: "fontsize",
                    Aliases: []string{"f"},
                    Usage: "fontsize",
                    Value: 15,
                },
                &cli.IntFlag{
                    Name: "rows",
                    Aliases: []string{"r"},
                    Usage: "rows number per page",
                    Value: 12,
                },
            },
            Action: genTableOfContents,
        },
        {
            Name: "bind report",
            Aliases: []string{"b"},
            Usage: "bind report and output \"./BoundReport\" directory",
            Action: bindReport,
        },
    }

    app.Name = "reportbinding"
    app.Usage = "report binding utils"
    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
