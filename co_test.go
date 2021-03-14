package urbitgob

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatp(t *testing.T) {

	var testCases = []struct {
		num          string
		expectedName string
		expectedErr  string
	}{
		{
			num:          "0",
			expectedName: "~zod",
		},
		{
			num:          "255",
			expectedName: "~fes",
		},
		{
			num:          "256",
			expectedName: "~marzod",
		},
		{
			num:          "65535",
			expectedName: "~fipfes",
		},
		{
			num:          "65536",
			expectedName: "~dapnep-ronmyl",
		},
		{
			num:          "14287616",
			expectedName: "~rosmur-hobrem",
		},
		{
			num:          "14287617",
			expectedName: "~sallus-nodlut",
		},
		{
			num:          "14287618",
			expectedName: "~marder-mopdur",
		},
		{
			num:          "14287619",
			expectedName: "~laphec-savted",
		},
		{
			num:          "4294967295",
			expectedName: "~dostec-risfen",
		},
		{
			num:          "4294967296",
			expectedName: "~doznec-dozzod-dozzod",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.num, func(t *testing.T) {

			actualName, actualErr := Patp(tt.num)

			assert.Equal(t, tt.expectedName, actualName)
			if tt.expectedErr == "" {
				assert.NoError(t, actualErr)
			} else {
				assert.Error(t, actualErr)
				if actualErr != nil {
					assert.Equal(t, tt.expectedErr, actualErr.Error())
				}
			}
		})
	}
}
