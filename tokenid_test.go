package pns

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeriveTokenId(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		input    string
		err      string
	}{
		{
			name:     "Valid PNS domain",
			expected: "79233663829379634837589865448569342784712482819484549289560981379859480642508",
			input:    "vitalik.pls",
		},
		{
			name:     "Invalid PNS domain",
			expected: "",
			input:    "foo.bar",
			err:      "unregistered name",
		},
		{
			name:     "Blank PNS domain",
			expected: "",
			input:    "",
			err:      "empty domain",
		},
	}
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/831a5442dc2e4536a9f8dee4ea1707a6")
	require.NoError(t, err)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := DeriveTokenID(client, test.input)
			if test.err != "" {
				require.EqualError(t, err, test.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.expected, actual)
			}
		})
	}
}
