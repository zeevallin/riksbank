package swea

import (
	"strconv"

	"cloud.google.com/go/civil"
	"github.com/zeeraw/riksbank/api/swea/responses"
)

// GetCalendarDays returns the business days between two dates
func GetCalendarDays(req *GetCalendarDaysRequest) (res *GetCalendarDaysResponse, err error) {
	r, err := request(tmpl("get_calendar_days"), req)
	if err != nil {
		return
	}

	env := &responses.GetCalendarDaysResponseEnvelope{}
	err = call(r, env)
	if err != nil {
		return
	}

	res = &GetCalendarDaysResponse{
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

	return
}

// GetAllCrossNames returns the business days between two dates
func GetAllCrossNames(req *GetAllCrossNamesRequest) (res *GetAllCrossNamesResponse, err error) {
	r, err := request(tmpl("get_all_cross_names"), req)
	if err != nil {
		return
	}

	env := &responses.GetAllCrossNamesResponseEnvelope{}
	err = call(r, env)

	res = &GetAllCrossNamesResponse{
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

	return
}
