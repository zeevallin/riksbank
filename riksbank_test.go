package riksbank_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/zeeraw/riksbank"
	"github.com/zeeraw/riksbank/currency"
	"github.com/zeeraw/riksbank/swea"
	"github.com/zeeraw/riksbank/swea/mock"
)

var (
	api *riksbank.Riksbank
	mck *mock.API

	crossRateInfo = []swea.CrossRateInfo{
		swea.CrossRateInfo{
			Base:    "SEK",
			Counter: "NOK",
			Average: "44.0",
			Value:   "0.88",
			Date:    today(),
			Period:  fmtDate(today()),
		},
	}

	ratesInfo = []swea.RateInfo{
		swea.RateInfo{
			GroupID:    "1",
			GroupName:  "Group 1",
			SeriesID:   "SERIES1",
			SeriesName: "Series 1",
			Value:      "0.0",
			Average:    "1.1",
			Min:        "2.2",
			Max:        "3.3",
			Ultimo:     "4.4",
			Date:       today(),
			Period:     fmtDate(today()),
		},
		swea.RateInfo{
			GroupID:    "2",
			GroupName:  "Group 2",
			SeriesID:   "SERIES2",
			SeriesName: "Series 2",
			Value:      "5.5",
			Average:    "6.6",
			Min:        "7.7",
			Max:        "8.8",
			Ultimo:     "9.9",
			Date:       today(),
			Period:     fmtDate(today()),
		},
	}
)

func TestMain(m *testing.M) {
	mck = mock.New()
	api = riksbank.New(riksbank.Config{
		Client: mck,
	})
	os.Exit(m.Run())
}

func Test_ExchangeRates(t *testing.T) {
	cases := []struct {
		name            string
		aggregateMethod riksbank.AggregateMethod
		expectedValue   float64
	}{
		{
			name:            "when aggregate method is daily",
			aggregateMethod: riksbank.Daily,
			expectedValue:   float64(0.88),
		},
		{
			name:            "when aggregate method is weekly",
			aggregateMethod: riksbank.Yearly,
			expectedValue:   float64(44.00),
		},
		{
			name:            "when aggregate method is monthly",
			aggregateMethod: riksbank.Monthly,
			expectedValue:   float64(44.00),
		},
		{
			name:            "when aggregate method is quarterly",
			aggregateMethod: riksbank.Quarterly,
			expectedValue:   float64(44.00),
		},
		{
			name:            "when aggregate method is yearly",
			aggregateMethod: riksbank.Yearly,
			expectedValue:   float64(44.00),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mck.GetCrossRatesResponse.CrossRates = crossRateInfo
			req := &riksbank.ExchangeRatesRequest{
				CurrencyPairs:   []currency.Pair{},
				AggregateMethod: c.aggregateMethod,
				From:            sevenDaysAgo(),
				To:              today(),
			}
			res, err := api.ExchangeRates(context.Background(), req)
			if err != nil {
				t.Errorf("request should not error: %v", err)
			}
			if len(res.ExchangeRates) != len(mck.GetCrossRatesResponse.CrossRates) {
				t.Errorf("exchange rates (%d) are not the same length as cross rates (%d)", len(res.ExchangeRates), len(mck.GetCrossRatesResponse.CrossRates))
			}
			er := res.ExchangeRates[0]
			if er.Value == nil {
				t.Errorf("exchange rate value should not be nil")
			}
			if c.expectedValue != *er.Value {
				t.Errorf("the exchange rate value was not %f, was %f", c.expectedValue, *er.Value)
			}
		})
	}
}

func Test_Rates(t *testing.T) {
	cases := []struct {
		name            string
		aggregateMethod riksbank.AggregateMethod
		analysisMethod  riksbank.AnalysisMethod
		expectedValues  []float64
	}{
		{
			name:            "when analysis is real and aggregate daily",
			analysisMethod:  riksbank.Real,
			aggregateMethod: riksbank.Daily,
			expectedValues:  []float64{0.0, 5.5},
		},
		{
			name:            "when analysis is mean and aggregate weekly",
			analysisMethod:  riksbank.Mean,
			aggregateMethod: riksbank.Weekly,
			expectedValues:  []float64{1.1, 6.6},
		},
		{
			name:            "when analysis is min and aggregate weekly",
			analysisMethod:  riksbank.Min,
			aggregateMethod: riksbank.Weekly,
			expectedValues:  []float64{2.2, 7.7},
		},
		{
			name:            "when analysis is max and aggregate weekly",
			analysisMethod:  riksbank.Max,
			aggregateMethod: riksbank.Weekly,
			expectedValues:  []float64{3.3, 8.8},
		},
		{
			name:            "when analysis is ultimo and aggregate weekly",
			analysisMethod:  riksbank.Ultimo,
			aggregateMethod: riksbank.Weekly,
			expectedValues:  []float64{4.4, 9.9},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mck.GetInterestAndExchangeRatesResponse.Rates = ratesInfo
			req := &riksbank.RatesRequest{
				Series:          []string{},
				AnalysisMethod:  c.analysisMethod,
				AggregateMethod: c.aggregateMethod,
				From:            sevenDaysAgo(),
				To:              today(),
			}
			res, err := api.Rates(context.Background(), req)
			if err != nil {
				t.Errorf("request should not error: %v", err)
			}
			if len(res.Rates) != len(mck.GetInterestAndExchangeRatesResponse.Rates) {
				t.Errorf("rates (%d) are not the same length as series (%d)", len(res.Rates), len(mck.GetInterestAndExchangeRatesResponse.Rates))
			}
			for idx, expected := range c.expectedValues {
				actual := res.Rates[idx].Value
				if actual == nil {
					t.Errorf("rate at index %d should not have nil value", idx)
				}
				if *actual != expected {
					t.Errorf("rates value was %f expected: %f", *actual, expected)
				}
			}
		})
	}
}

func Test_Days(t *testing.T) {
	cases := []struct {
		name     string
		swea     []swea.DayInfo
		expected riksbank.Days
	}{
		{
			name: "when the days contain a wide range of different type of days",
			swea: []swea.DayInfo{
				swea.DayInfo{
					Date:      today(),
					IsBankDay: false,
					Week:      1,
					WeekYear:  today().Year(),
				},
			},
			expected: riksbank.Days{
				riksbank.Day{
					Date:      today(),
					Week:      1,
					Year:      today().Year(),
					IsBankDay: false,
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mck.GetCalendarDaysResponse.Days = c.swea
			req := &riksbank.DaysRequest{
				From: sevenDaysAgo(),
				To:   today(),
			}
			res, err := api.Days(context.Background(), req)
			if err != nil {
				t.Errorf("request should not error: %v", err)
			}
			for idx, day := range c.expected {
				if res.Days[idx] != day {
					t.Errorf("the day %v is not as expected: %v", res.Days[idx], day)
				}
			}
		})
	}
}

func Test_Groups(t *testing.T) {
	cases := []struct {
		name     string
		swea     []swea.GroupInfo
		expected riksbank.Groups
	}{
		{
			name: "when there are several groups",
			swea: []swea.GroupInfo{
				swea.GroupInfo{
					ID:          "1",
					ParentID:    "",
					Name:        "Group 1",
					Description: "This is a group called Group 1",
				},
				swea.GroupInfo{
					ID:          "2",
					ParentID:    "1",
					Name:        "Group 2",
					Description: "This is a group called Group 2",
				},
			},
			expected: riksbank.Groups{
				riksbank.Group{
					ID:          "1",
					ParentID:    "",
					Name:        "Group 1",
					Description: "This is a group called Group 1",
				},
				riksbank.Group{
					ID:          "2",
					ParentID:    "1",
					Name:        "Group 2",
					Description: "This is a group called Group 2",
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			mck.GetInterestAndExchangeGroupNamesResponse.Groups = c.swea
			res, err := api.Groups(context.Background(), &riksbank.GroupsRequest{})
			if err != nil {
				t.Errorf("request should not error: %v", err)
			}
			for idx, day := range c.expected {
				if res.Groups[idx] != day {
					t.Errorf("the day %v is not as expected: %v", res.Groups[idx], day)
				}
			}
		})
	}
}

func Test_Series(t *testing.T) {
	var (
		sweaG1 = swea.GroupInfo{
			ID:          "1",
			ParentID:    "",
			Name:        "Group 1",
			Description: "This is a group called Group 1",
		}
		sweaG2 = swea.GroupInfo{
			ID:          "2",
			ParentID:    "1",
			Name:        "Group 2",
			Description: "This is a group called Group 2",
		}
		sweaS1 = swea.SeriesInfo{
			ID:              "SERIES1",
			Name:            "Series 1",
			Description:     "Description",
			LongDescription: "Long description",
			Source:          "Source",
			GroupID:         "1",
			From:            sevenDaysAgo(),
			To:              today(),
		}
	)
	cases := []struct {
		name       string
		groups     []string
		sweaGroups []swea.GroupInfo
		sweaSeries []swea.SeriesInfo
		expected   riksbank.SeriesGroups
	}{
		{
			name:   "when no groups are specified",
			groups: []string{},
			sweaGroups: []swea.GroupInfo{
				sweaG1,
				sweaG2,
			},
			sweaSeries: []swea.SeriesInfo{
				sweaS1,
			},
			expected: riksbank.SeriesGroups{
				{
					Group: riksbank.Group{
						ID:          "1",
						ParentID:    "",
						Name:        "Group 1",
						Description: "This is a group called Group 1",
					},
				},
				{
					Group: riksbank.Group{
						ID:          "2",
						ParentID:    "1",
						Name:        "Group 2",
						Description: "This is a group called Group 2",
					},
				},
			},
		},
		{
			name:   "when no groups are specified",
			groups: []string{"2"},
			sweaGroups: []swea.GroupInfo{
				sweaG1,
				sweaG2,
			},
			sweaSeries: []swea.SeriesInfo{
				sweaS1,
			},
			expected: riksbank.SeriesGroups{
				{
					Group: riksbank.Group{
						ID:          "2",
						ParentID:    "1",
						Name:        "Group 2",
						Description: "This is a group called Group 2",
					},
				},
			},
		},
	}
	for _, c := range cases {
		mck.GetInterestAndExchangeGroupNamesResponse.Groups = c.sweaGroups
		mck.GetInterestAndExchangeNamesResponse.Series = c.sweaSeries
		res, err := api.Series(context.Background(), &riksbank.SeriesRequest{
			Groups: c.groups,
		})
		if err != nil {
			t.Errorf("request should not error: %v", err)
		}
		if len(res.Groups) != len(c.expected) {
			t.Errorf("response groups (%d) are not the same length expected (%d)", len(res.Groups), len(c.expected))
		}
	}
}

func today() time.Time {
	return date(time.Now())
}

func sevenDaysAgo() time.Time {
	return date(time.Now().AddDate(0, 0, -7))
}

func date(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

func fmtDate(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())
}
