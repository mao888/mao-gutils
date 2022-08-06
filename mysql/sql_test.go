package mysql_test

import (
	"go-zero-demo/internal/utils"
	"testing"
)

func TestWarpSqlValues(t *testing.T) {
	values := []string{
		"1", "2", "3", "4", "5",
	}
	ret := utils.WarpSqlValues(values)
	if ret != "'1', '2', '3', '4', '5'" {
		t.Error()
	}
}

func TestWarpBracketedSqlValues(t *testing.T) {
	values := []string{
		"1", "2", "3", "4", "5",
	}
	ret := utils.WarpBracketedSqlValues(values)
	if ret != "('1', '2', '3', '4', '5')" {
		t.Error()
	}
}
