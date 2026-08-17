package csvjson

import "testing"

func TestProbeHeaderOnlyCSVUsesEmptyArrays(t *testing.T) {
	columns, rows, err := CSVToJSON("name,age\n", Options{Delimiter: ',', Header: true})
	if err != nil {
		t.Fatalf("CSVToJSON returned error: %v", err)
	}
	if columns == nil || rows == nil {
		t.Fatalf("header-only CSV must return non-nil empty arrays: columns=%#v rows=%#v", columns, rows)
	}
	if len(columns) != 2 || len(rows) != 0 {
		t.Fatalf("unexpected result: columns=%#v rows=%#v", columns, rows)
	}
}
