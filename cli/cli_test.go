package cli_test

import (
	"os"
	"testing"

	"github.com/zeeraw/riksbank"
	"github.com/zeeraw/riksbank/cli"
	"github.com/zeeraw/riksbank/swea/mock"
)

const (
	bin = "riksbank"
)

var (
	tool cli.Tool
)

func TestMain(m *testing.M) {
	mck := mock.New()
	tool = cli.Tool{
		Riksbank: riksbank.New(riksbank.Config{
			SweaClient: mck,
		}),
	}
	os.Exit(m.Run())
}

func Test_Run(t *testing.T) {
	cases := []struct {
		name   string
		args   []string
		errors bool
	}{
		{
			name:   "(none)",
			args:   []string{},
			errors: true,
		},
		{
			name: "(only bin)",
			args: []string{
				bin,
			},
			errors: false,
		},
		{
			name: "rates",
			args: []string{
				bin,
				"rates",
			},
			errors: true,
		},
		{
			name: "rates with one series",
			args: []string{
				bin,
				"rates",
				"-s",
				"SEKNOKPMI",
			},
			errors: false,
		},
		{
			name: "series",
			args: []string{
				bin,
				"series",
			},
			errors: false,
		},
		{
			name: "groups",
			args: []string{
				bin,
				"groups",
			},
			errors: false,
		},
		{
			name: "days",
			args: []string{
				bin,
				"days",
			},
			errors: false,
		},
		{
			name: "exchange",
			args: []string{
				bin,
				"exchange",
			},
			errors: false,
		},
		{
			name: "exchange rates",
			args: []string{
				bin,
				"exchange",
				"rates",
			},
			errors: true,
		},
		{
			name: "exchange rates with currency pair",
			args: []string{
				bin,
				"exchange",
				"rates",
				"-c",
				"SEK/NOK",
			},
			errors: false,
		},
		{
			name: "exchange currencies",
			args: []string{
				bin,
				"exchange",
				"currencies",
			},
			errors: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := tool.Run(c.args)
			if c.errors {
				if err == nil {
					t.Errorf("command should error but didn't")
				}
			} else {
				if err != nil {
					t.Errorf("should not have error: %v", err)
				}
			}
		})
	}
}
