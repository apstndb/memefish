package ast

import "testing"

func TestOptionsIntegerFieldBases(t *testing.T) {
	for _, tt := range []struct {
		value     string
		base      int
		want      int64
		wantError bool
	}{
		{"0x10", 16, 16, false},
		{"0X10", 16, 16, false},
		{"+0x10", 16, 16, false},
		{"-0x10", 16, -16, false},
		{"010", 10, 10, false},
		{"+010", 10, 10, false},
		{"-010", 10, -10, false},
		{"010", 0, 10, false},
		{"0x7fffffffffffffff", 16, 9223372036854775807, false},
		{"-0x8000000000000000", 16, -9223372036854775808, false},
		{"0x8000000000000000", 16, 0, true},
		{"0x", 16, 0, true},
	} {
		t.Run(tt.value, func(t *testing.T) {
			o := &Options{Records: []*OptionsDef{{
				Name:  &Ident{Name: "x"},
				Value: &IntLiteral{Value: tt.value, Base: tt.base},
			}}}
			got, err := o.IntegerField("x")
			if (err != nil) != tt.wantError {
				t.Fatalf("error = %v, want error = %v", err, tt.wantError)
			}
			if !tt.wantError && (got == nil || *got != tt.want) {
				t.Errorf("value = %v, want %d", got, tt.want)
			}
		})
	}
}
