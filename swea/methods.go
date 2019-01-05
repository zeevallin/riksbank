package swea

import (
	"context"
	"strconv"
	"strings"

	"cloud.google.com/go/civil"
	"github.com/zeeraw/riksbank/swea/responses"
)

// GetCalendarDays returns the business days between two dates
func (s *Swea) GetCalendarDays(ctx context.Context, req *GetCalendarDaysRequest) (*GetCalendarDaysResponse, error) {
	body, err := build(tmpl("get_calendar_days"), req)
	if err != nil {
		return nil, err
	}
	env := &responses.GetCalendarDaysResponseEnvelope{}
	err = s.call(ctx, body, env)
	if err != nil {
		return nil, err
	}
	res := &GetCalendarDaysResponse{
		From: req.From,
		To:   req.To,
		Days: make([]Day, len(env.Body.GetCalendarDaysResponse.Return)),
	}
	for idx, r := range env.Body.GetCalendarDaysResponse.Return {
		date, err := civil.ParseDate(r.Caldate.Text)
		if err != nil {
			return nil, err
		}
		week, err := strconv.Atoi(r.Week.Text)
		if err != nil {
			return nil, err
		}
		weekYear, err := strconv.Atoi(r.Weekyear.Text)
		if err != nil {
			return nil, err
		}
		res.Days[idx] = Day{
			Date:      date,
			Week:      week,
			WeekYear:  weekYear,
			IsBankDay: isTrue(r.Bankday.Text),
		}
	}
	return res, nil
}

// GetAllCrossNames returns the series names for exhcnage rates to SEK
func (s *Swea) GetAllCrossNames(ctx context.Context, req *GetAllCrossNamesRequest) (*GetAllCrossNamesResponse, error) {
	body, err := build(tmpl("get_all_cross_names"), req)
	if err != nil {
		return nil, err
	}
	env := &responses.GetAllCrossNamesResponseEnvelope{}
	err = s.call(ctx, body, env)
	res := &GetAllCrossNamesResponse{
		Language: req.Language,
		Series:   make([]SeriesInfo, len(env.Body.GetAllCrossNamesResponse.Return)),
	}
	for idx, r := range env.Body.GetAllCrossNamesResponse.Return {
		res.Series[idx] = SeriesInfo{
			ID:          r.Seriesid.Text,
			Name:        r.Seriesname.Text,
			Description: r.Seriesdescription.Text,
		}
	}
	return res, nil
}

// GetCrossRates returns the exchange rates for series
func (s *Swea) GetCrossRates(ctx context.Context, req *GetCrossRatesRequest) (*GetCrossRatesResponse, error) {
	body, err := build(tmpl("get_cross_rates"), req)
	if err != nil {
		return nil, err
	}
	env := &responses.GetCrossRatesResponse{}
	err = s.call(ctx, body, env)
	res := &GetCrossRatesResponse{
		CrossPairs:      req.CrossPairs,
		From:            req.From,
		To:              req.To,
		Language:        req.Language,
		AggregateMethod: req.AggregateMethod,
	}

	var crossRates = []CrossRate{}
	for _, s := range env.Body.GetCrossRatesResponse.Return.Groups.Series {
		for _, rr := range s.Resultrows {
			date, err := civil.ParseDate(strings.TrimSpace(rr.Date))
			if err != nil {
				return nil, err
			}
			var period string
			ptx := strings.TrimSpace(rr.Period.Text)
			if ptx != "" {
				period = ptx
			} else {
				period = date.String()
			}
			var value string
			switch req.AggregateMethod {
			case Weekly, Monthly, Quarterly, Yearly:
				value = strings.TrimSpace(rr.Average.Text)
			default:
				value = strings.TrimSpace(rr.Value)
			}
			cr := CrossRate{
				Base:    ParseCurrency(strings.TrimSpace(s.Seriesid1)),
				Counter: ParseCurrency(strings.TrimSpace(s.Seriesid2)),
				Date:    date,
				Value:   value,
				Period:  period,
			}
			crossRates = append(crossRates, cr)
		}
	}
	res.CrossRates = crossRates
	return res, nil
}
