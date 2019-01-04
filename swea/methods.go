package swea

import (
	"context"
	"strconv"

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

// GetAllCrossNames returns the business days between two dates
func (s *Swea) GetAllCrossNames(ctx context.Context, req *GetAllCrossNamesRequest) (*GetAllCrossNamesResponse, error) {
	body, err := build(tmpl("get_all_cross_names"), req)
	if err != nil {
		return nil, err
	}
	env := &responses.GetAllCrossNamesResponseEnvelope{}
	err = s.call(ctx, body, env)
	res := &GetAllCrossNamesResponse{
		Language: req.Language,
		Series:   make([]Series, len(env.Body.GetAllCrossNamesResponse.Return)),
	}
	for idx, r := range env.Body.GetAllCrossNamesResponse.Return {
		res.Series[idx] = Series{
			ID:          r.Seriesid.Text,
			Name:        r.Seriesname.Text,
			Description: r.Seriesdescription.Text,
		}
	}
	return res, nil
}
