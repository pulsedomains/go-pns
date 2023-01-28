package pns

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNormaliseDomain(t *testing.T) {
	tests := []struct {
		input  string
		output string
		err    error
	}{
		{"", "", nil},
		{".", ".", nil},
		{"pls", "pls", nil},
		{"PLS", "pls", nil},
		{".pls", ".pls", nil},
		{".pls.", ".pls.", nil},
		{"wealdtech.pls", "wealdtech.pls", nil},
		{".wealdtech.pls", ".wealdtech.pls", nil},
		{"subdomain.wealdtech.pls", "subdomain.wealdtech.pls", nil},
		{"*.wealdtech.pls", "*.wealdtech.pls", nil},
		{"omg.thetoken.pls", "omg.thetoken.pls", nil},
		{"_underscore.thetoken.pls", "_underscore.thetoken.pls", nil},
		{"點看.pls", "點看.pls", nil},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := NormaliseDomain(tt.input)
			if tt.err != nil {
				if err == nil {
					t.Fatalf("missing expected error")
				}
				if tt.err.Error() != err.Error() {
					t.Errorf("unexpected error value %v", err)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error %v", err)
				}
				if tt.output != result {
					t.Errorf("%v => %v (expected %v)", tt.input, result, tt.output)
				}
			}
		})
	}
}

func TestNormaliseDomainStrict(t *testing.T) {
	tests := []struct {
		input  string
		output string
		err    error
	}{
		{"", "", nil},
		{".", ".", nil},
		{"pls", "pls", nil},
		{"PLS", "pls", nil},
		{".pls", ".pls", nil},
		{".pls.", ".pls.", nil},
		{"wealdtech.pls", "wealdtech.pls", nil},
		{".wealdtech.pls", ".wealdtech.pls", nil},
		{"subdomain.wealdtech.pls", "subdomain.wealdtech.pls", nil},
		{"*.wealdtech.pls", "*.wealdtech.pls", nil},
		{"omg.thetoken.pls", "omg.thetoken.pls", nil},
		{"_underscore.thetoken.pls", "", errors.New("idna: disallowed rune U+005F")},
		{"點看.pls", "點看.pls", nil},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := NormaliseDomainStrict(tt.input)
			if tt.err != nil {
				if err == nil {
					t.Fatalf("missing expected error")
				}
				if tt.err.Error() != err.Error() {
					t.Errorf("unexpected error value %v", err)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error %v", err)
				}
				if tt.output != result {
					t.Errorf("%v => %v (expected %v)", tt.input, result, tt.output)
				}
			}
		})
	}
}

func TestTld(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"", ""},
		{".", ""},
		{"pls", "pls"},
		{"PLS", "pls"},
		{".pls", "pls"},
		{"wealdtech.pls", "pls"},
		{".wealdtech.pls", "pls"},
		{"subdomain.wealdtech.pls", "pls"},
	}

	for _, tt := range tests {
		result := Tld(tt.input)
		if tt.output != result {
			t.Errorf("Failure: %v => %v (expected %v)\n", tt.input, result, tt.output)
		}
	}
}

func TestDomainPart(t *testing.T) {
	tests := []struct {
		input  string
		part   int
		output string
		err    bool
	}{
		{"", 1, "", false},
		{"", 2, "", true},
		{"", -1, "", false},
		{"", -2, "", true},
		{".", 1, "", false},
		{".", 2, "", false},
		{".", 3, "", true},
		{".", -1, "", false},
		{".", -2, "", false},
		{".", -3, "", true},
		{"PLS", 1, "pls", false},
		{"PLS", 2, "", true},
		{"PLS", -1, "pls", false},
		{"PLS", -2, "", true},
		{".PLS", 1, "", false},
		{".PLS", 2, "pls", false},
		{".PLS", 3, "", true},
		{".PLS", -1, "pls", false},
		{".PLS", -2, "", false},
		{".PLS", -3, "", true},
		{"wealdtech.pls", 1, "wealdtech", false},
		{"wealdtech.pls", 2, "pls", false},
		{"wealdtech.pls", 3, "", true},
		{"wealdtech.pls", -1, "pls", false},
		{"wealdtech.pls", -2, "wealdtech", false},
		{"wealdtech.pls", -3, "", true},
		{".wealdtech.pls", 1, "", false},
		{".wealdtech.pls", 2, "wealdtech", false},
		{".wealdtech.pls", 3, "pls", false},
		{".wealdtech.pls", 4, "", true},
		{".wealdtech.pls", -1, "pls", false},
		{".wealdtech.pls", -2, "wealdtech", false},
		{".wealdtech.pls", -3, "", false},
		{".wealdtech.pls", -4, "", true},
		{"subdomain.wealdtech.pls", 1, "subdomain", false},
		{"subdomain.wealdtech.pls", 2, "wealdtech", false},
		{"subdomain.wealdtech.pls", 3, "pls", false},
		{"subdomain.wealdtech.pls", 4, "", true},
		{"subdomain.wealdtech.pls", -1, "pls", false},
		{"subdomain.wealdtech.pls", -2, "wealdtech", false},
		{"subdomain.wealdtech.pls", -3, "subdomain", false},
		{"subdomain.wealdtech.pls", -4, "", true},
		{"a.b.c", 1, "a", false},
		{"a.b.c", 2, "b", false},
		{"a.b.c", 3, "c", false},
		{"a.b.c", 4, "", true},
		{"a.b.c", -1, "c", false},
		{"a.b.c", -2, "b", false},
		{"a.b.c", -3, "a", false},
		{"a.b.c", -4, "", true},
	}

	for _, tt := range tests {
		result, err := DomainPart(tt.input, tt.part)
		if err != nil && !tt.err {
			t.Errorf("Failure: %v, %v => error (unexpected)\n", tt.input, tt.part)
		}
		if err == nil && tt.err {
			t.Errorf("Failure: %v, %v => no error (unexpected)\n", tt.input, tt.part)
		}
		if tt.output != result {
			t.Errorf("Failure: %v, %v => %v (expected %v)\n", tt.input, tt.part, result, tt.output)
		}
	}
}

func TestUnqualifiedName(t *testing.T) {
	tests := []struct {
		domain string
		root   string
		name   string
		err    error
	}{
		{
			domain: "",
			root:   "",
			name:   "",
		},
		{
			domain: "wealdtech.pls",
			root:   "pls",
			name:   "wealdtech",
		},
	}

	for i, test := range tests {
		name, err := UnqualifiedName(test.domain, test.root)
		if test.err != nil {
			assert.Equal(t, test.err, err, fmt.Sprintf("Incorrect error at test %d", i))
		} else {
			require.Nil(t, err, fmt.Sprintf("Unexpected error at test %d", i))
			assert.Equal(t, test.name, name, fmt.Sprintf("Incorrect result at test %d", i))
		}
	}
}
