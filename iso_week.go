package isoweek

import "time"

// ISOWeek represents the ISO 8601 Week number with the corresponding year.
// It provides timestamps representing the boundaries of the week.
type ISOWeek struct {
	Year       int
	Week       int
	UpperBound time.Time
	LowerBound time.Time
}

// New returns a new ISOWeek.
func New(t time.Time) *ISOWeek {
	yr, wk := t.ISOWeek()

	mon := mondayOf(t)
	sun := mon.AddDate(0, 0, 7).Add(-1 * time.Nanosecond)
	return &ISOWeek{
		Year:       yr,
		Week:       wk,
		LowerBound: mon,
		UpperBound: sun,
	}
}

// FromWeekNbr is useful to retrieve lower and uppper bounds of a week number (using UTC time).
func FromWeekNbr(yr, wk int) *ISOWeek {
	lower := time.Date(yr, 0, 0, 0, 0, 0, 0, time.UTC)
	isoYear, isoWeek := lower.ISOWeek()

	//  back to Monday
	for lower.Weekday() != time.Monday {
		lower = lower.AddDate(0, 0, -1)
		isoYear, isoWeek = lower.ISOWeek()
	}

	// forward to the first day of the first week
	for isoYear < yr {
		lower = lower.AddDate(0, 0, 7)
		isoYear, isoWeek = lower.ISOWeek()
	}

	// forward to the first day of the given week
	for isoWeek < wk {
		lower = lower.AddDate(0, 0, 7)
		isoYear, isoWeek = lower.ISOWeek()
	}
	upper := lower.AddDate(0, 0, 7)
	upper = upper.Add(-1)
	return &ISOWeek{
		Year:       yr,
		Week:       wk,
		LowerBound: lower,
		UpperBound: upper,
	}
}

// Equals checks whether or not two iso weeks have the same value (year + week).
func (iw *ISOWeek) Equals(other *ISOWeek) bool {
	if iw == nil || other == nil {
		return false
	}

	return iw.Week == other.Week && iw.Year == other.Year
}

func mondayOf(t time.Time) time.Time {
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)

	if t.Weekday() == time.Sunday {
		return t.AddDate(0, 0, -6)
	}
	return t.AddDate(0, 0, int(1-t.Weekday()))
}
