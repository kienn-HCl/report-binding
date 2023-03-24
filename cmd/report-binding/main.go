package main

import (
	"os"

	"github.com/kienn-HCl/report-binding"
	"github.com/urfave/cli/v2"
)

func initDirectory(c *cli.Context) error {
    if c.Bool("csv") {
        return reportbinding.GenCsv()
    }
    return reportbinding.InitReportBinding()
}

func bindReport(c *cli.Context) error {
    reports, err := reportbinding.NewReportDatas()
    if err != nil {
        return err
    }

    err = reports.UniteReport()
    if err != nil {
        return err
    }

    reportsNum := len(*reports)

    err = reports.AddPagenum("./UniteReport/uniteReport.pdf", 4 + (reportsNum-1)/10)
    if err != nil {
        return err
    }
    return nil
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
                    Aliases: []string{"s"},
                    Usage: "only output csv file",
                    Value: false,
                },
            },
            Action: initDirectory,
        },
        {
            Name: "bindReport",
            Aliases: []string{"b"},
            Usage: "bind report",
            Action: bindReport,
        },
    }

    app.Name = "reportbinding"
    app.Run(os.Args)
}
