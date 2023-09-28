package sabre

import (
	"os"
	"testing"

	"github.com/olekukonko/tablewriter"
	"github.com/stretchr/testify/assert"
)

func TestParseAmount(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want float64
	}{
		{"383 -> 38.3", "383", 3.83},
		{"394900 -> 3949.00", "394900", 3949.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := ParseAmount(tt.raw)
			assert.Equal(t, tt.want, rs)
		})
	}
}

func Test_parseDuration(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int32
	}{
		// TODO: Add test cases.
		{"Duration NDC", "PT07H40M", 460},
		{"Duration NDC", "PT3H6M", 186},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := ParseDuration(tt.arg)
			assert.Equal(t, tt.want, rs)
		})
	}
}

func TestParseExpiration(t *testing.T) {
	tests := []struct {
		name string
		exp  string
		want string
	}{
		{"Valid expiration", "2039-01", "0139"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rs := ParseExpiration(tt.exp)
			assert.Equal(t, tt.want, rs)
		})
	}
}

func printTable(title, footer []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(title)
	table.SetAutoMergeCells(false)
	table.SetAutoWrapText(false)
	table.SetBorder(true)  // Set Border to false
	table.AppendBulk(data) // Add Bulk Data
	table.SetFooter(footer)
	table.Render()
}
