package co

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type stdTestCase struct {
	in              string
	out             string
	expectedErrText string
}

type stdCoFn func(string) (string, error)

func stdTestRunner(t *testing.T, testCases []stdTestCase, f stdCoFn) {

	for _, tt := range testCases {
		t.Run(tt.in, func(t *testing.T) {

			actualOut, actualErr := f(tt.in)

			assert.Equal(t, tt.out, actualOut)
			if tt.expectedErrText == "" {
				assert.NoError(t, actualErr)
			} else {
				assert.Error(t, actualErr)
				if actualErr != nil {
					assert.Equal(t, tt.expectedErrText, actualErr.Error())
				}
			}
		})
	}
}

func TestPatp(t *testing.T) {

	var testCases = []stdTestCase{
		{
			in:  "0",
			out: "~zod",
		},
		{
			in:  "255",
			out: "~fes",
		},
		{
			in:  "256",
			out: "~marzod",
		},
		{
			in:  "65535",
			out: "~fipfes",
		},
		{
			in:  "65536",
			out: "~dapnep-ronmyl",
		},
		{
			in:  "14287616",
			out: "~rosmur-hobrem",
		},
		{
			in:  "14287617",
			out: "~sallus-nodlut",
		},
		{
			in:  "14287618",
			out: "~marder-mopdur",
		},
		{
			in:  "14287619",
			out: "~laphec-savted",
		},
		{
			in:  "4294967295",
			out: "~dostec-risfen",
		},
		{
			in:  "4294967296",
			out: "~doznec-dozzod-dozzod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid integer string: abcdefg",
		},
	}

	stdTestRunner(t, testCases, Patp)
}

func TestPatq(t *testing.T) {

	var testCases = []stdTestCase{
		{
			in:  "0",
			out: "~zod",
		},
		{
			in:  "255",
			out: "~fes",
		},
		{
			in:  "256",
			out: "~marzod",
		},
		{
			in:  "65535",
			out: "~fipfes",
		},
		{
			in:  "65536",
			out: "~doznec-dozzod",
		},
		{
			in:  "14287616",
			out: "~dozler-wanzod",
		},
		{
			in:  "14287617",
			out: "~dozler-wannec",
		},
		{
			in:  "14287618",
			out: "~dozler-wanbud",
		},
		{
			in:  "14287619",
			out: "~dozler-wanwes",
		},
		{
			in:  "4294967295",
			out: "~fipfes-fipfes",
		},
		{
			in:  "4294967296",
			out: "~doznec-dozzod-dozzod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid integer string: abcdefg",
		},
	}

	stdTestRunner(t, testCases, Patq)
}

func TestClan(t *testing.T) {

	var testCases = []stdTestCase{
		{
			in:  "~zod",
			out: galaxy,
		},
		{
			in:  "~fes",
			out: galaxy,
		},
		{
			in:  "~marzod",
			out: star,
		},
		{
			in:  "~fipfes",
			out: star,
		},
		{
			in:  "~dapnep-ronmyl",
			out: planet,
		},
		{
			in:  "~rosmur-hobrem",
			out: planet,
		},
		{
			in:  "~sallus-nodlut",
			out: planet,
		},
		{
			in:  "~marder-mopdur",
			out: planet,
		},
		{
			in:  "~laphec-savted",
			out: planet,
		},
		{
			in:  "~dostec-risfen",
			out: planet,
		},
		{
			in:  "~divrul-dalred-samhec-sidrex",
			out: moon,
		},
		{
			in:  "~dotmec-niblyd-tocdys-ravryg--panper-hilsug-nidnev-marzod",
			out: comet,
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid @p: abcdefg",
		},
	}

	stdTestRunner(t, testCases, Clan)
}

func TestSein(t *testing.T) {

	var testCases = []stdTestCase{
		{
			in:  "~zod",
			out: "~zod",
		},
		{
			in:  "~fes",
			out: "~fes",
		},
		{
			in:  "~marzod",
			out: "~zod",
		},
		{
			in:  "~fipfes",
			out: "~fes",
		},
		{
			in:  "~dapnep-ronmyl",
			out: "~zod",
		},
		{
			in:  "~rosmur-hobrem",
			out: "~wanzod",
		},
		{
			in:  "~sallus-nodlut",
			out: "~wannec",
		},
		{
			in:  "~marder-mopdur",
			out: "~wanbud",
		},
		{
			in:  "~laphec-savted",
			out: "~wanwes",
		},
		{
			in:  "~dostec-risfen",
			out: "~fipfes",
		},
		{
			in:  "~divrul-dalred-samhec-sidrex",
			out: "~samhec-sidrex",
		},
		{
			in:  "~dotmec-niblyd-tocdys-ravryg--panper-hilsug-nidnev-marzod",
			out: "~zod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid @p: abcdefg",
		},
	}

	stdTestRunner(t, testCases, Sein)
}

func TestPatp2Dec(t *testing.T) {

	var testCases = []stdTestCase{
		{
			out: "0",
			in:  "~zod",
		},
		{
			out: "255",
			in:  "~fes",
		},
		{
			out: "256",
			in:  "~marzod",
		},
		{
			out: "65535",
			in:  "~fipfes",
		},
		{
			out: "65536",
			in:  "~dapnep-ronmyl",
		},
		{
			out: "14287616",
			in:  "~rosmur-hobrem",
		},
		{
			out: "14287617",
			in:  "~sallus-nodlut",
		},
		{
			out: "14287618",
			in:  "~marder-mopdur",
		},
		{
			out: "14287619",
			in:  "~laphec-savted",
		},
		{
			out: "4294967295",
			in:  "~dostec-risfen",
		},
		{
			out: "4294967296",
			in:  "~doznec-dozzod-dozzod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid @p: abcdefg",
		},
	}

	stdTestRunner(t, testCases, Patp2Dec)
}

func TestPatq2Dec(t *testing.T) {

	var testCases = []stdTestCase{
		{
			out: "0",
			in:  "~zod",
		},
		{
			out: "255",
			in:  "~fes",
		},
		{
			out: "256",
			in:  "~marzod",
		},
		{
			out: "65535",
			in:  "~fipfes",
		},
		{
			out: "65536",
			in:  "~doznec-dozzod",
		},
		{
			out: "14287616",
			in:  "~dozler-wanzod",
		},
		{
			out: "14287617",
			in:  "~dozler-wannec",
		},
		{
			out: "14287618",
			in:  "~dozler-wanbud",
		},
		{
			out: "14287619",
			in:  "~dozler-wanwes",
		},
		{
			out: "4294967295",
			in:  "~fipfes-fipfes",
		},
		{
			out: "4294967296",
			in:  "~doznec-dozzod-dozzod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid @q: abcdefg",
		},
	}

	stdTestRunner(t, testCases, Patq2Dec)
}

func TestPatp2Hex(t *testing.T) {

	var testCases = []stdTestCase{
		{
			out: "00",
			in:  "~zod",
		},
		{
			out: "ff",
			in:  "~fes",
		},
		{
			out: "0100",
			in:  "~marzod",
		},
		{
			out: "ffff",
			in:  "~fipfes",
		},
		{
			out: "010000",
			in:  "~dapnep-ronmyl",
		},
		{
			out: "da0300",
			in:  "~rosmur-hobrem",
		},
		{
			out: "da0301",
			in:  "~sallus-nodlut",
		},
		{
			out: "da0302",
			in:  "~marder-mopdur",
		},
		{
			out: "da0303",
			in:  "~laphec-savted",
		},
		{
			out: "ffffffff",
			in:  "~dostec-risfen",
		},
		{
			out: "0100000000",
			in:  "~doznec-dozzod-dozzod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid @p: abcdefg",
		},
	}

	stdTestRunner(t, testCases, Patp2Hex)
}

func TestPatq2Hex(t *testing.T) {

	var testCases = []stdTestCase{
		{
			out: "00",
			in:  "~zod",
		},
		{
			out: "ff",
			in:  "~fes",
		},
		{
			out: "0100",
			in:  "~marzod",
		},
		{
			out: "ffff",
			in:  "~fipfes",
		},
		{
			out: "00010000",
			in:  "~doznec-dozzod",
		},
		// TODO: Look into these leading zeroes; does the same in JS implementation
		{
			out: "00da0300",
			in:  "~dozler-wanzod",
		},
		{
			out: "00da0301",
			in:  "~dozler-wannec",
		},
		{
			out: "00da0302",
			in:  "~dozler-wanbud",
		},
		{
			out: "00da0303",
			in:  "~dozler-wanwes",
		},
		{
			out: "ffffffff",
			in:  "~fipfes-fipfes",
		},
		{
			out: "000100000000",
			in:  "~doznec-dozzod-dozzod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid @q: abcdefg",
		},
	}

	stdTestRunner(t, testCases, Patq2Hex)
}

func TestHex2Patp(t *testing.T) {

	var testCases = []stdTestCase{
		{
			in:  "00",
			out: "~zod",
		},
		{
			in:  "ff",
			out: "~fes",
		},
		{
			in:  "0100",
			out: "~marzod",
		},
		{
			in:  "ffff",
			out: "~fipfes",
		},
		{
			in:  "010000",
			out: "~dapnep-ronmyl",
		},
		{
			in:  "da0300",
			out: "~rosmur-hobrem",
		},
		{
			in:  "da0301",
			out: "~sallus-nodlut",
		},
		{
			in:  "da0302",
			out: "~marder-mopdur",
		},
		{
			in:  "da0303",
			out: "~laphec-savted",
		},
		{
			in:  "ffffffff",
			out: "~dostec-risfen",
		},
		{
			in:  "0100000000",
			out: "~doznec-dozzod-dozzod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid hexadecimal string: abcdefg",
		},
	}

	stdTestRunner(t, testCases, Hex2Patp)
}

func TestHex2Patq(t *testing.T) {

	var testCases = []stdTestCase{
		{
			in:  "00",
			out: "~zod",
		},
		{
			in:  "ff",
			out: "~fes",
		},
		{
			in:  "0100",
			out: "~marzod",
		},
		{
			in:  "ffff",
			out: "~fipfes",
		},
		{
			in:  "00010000",
			out: "~doznec-dozzod",
		},
		{
			in:  "00da0300",
			out: "~dozler-wanzod",
		},
		{
			in:  "00da0301",
			out: "~dozler-wannec",
		},
		{
			in:  "00da0302",
			out: "~dozler-wanbud",
		},
		{
			in:  "00da0303",
			out: "~dozler-wanwes",
		},
		{
			in:  "ffffffff",
			out: "~fipfes-fipfes",
		},
		{
			in:  "000100000000",
			out: "~doznec-dozzod-dozzod",
		},
		{
			in:              "abcdefg",
			expectedErrText: "invalid hexadecimal string: abcdefg",
		},
	}

	stdTestRunner(t, testCases, Hex2Patq)
}
