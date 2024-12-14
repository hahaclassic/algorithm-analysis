package tableoutput

import (
	"os"

	"github.com/jedib0t/go-pretty/table"
)

func PrintTable(style table.Style, headers []string, rows [][]interface{}) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	tableHeader := make(table.Row, 0, len(headers))
	for i := range headers {
		tableHeader = append(tableHeader, headers[i])
	}
	t.AppendHeader(tableHeader)

	for _, row := range rows {
		t.AppendRow(row)
	}
	t.SetStyle(style)
	t.Render()
}
