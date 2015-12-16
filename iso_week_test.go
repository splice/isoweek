package isoweek

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	nullTime = "0000-00-00 00:00:00"
	// Time layout used to process DB time values.
	DbtimeLayout = "2006-01-02 15:04:05"
)

func TestNew(t *testing.T) {
	for _, day := range []int{13, 14, 15, 16, 17, 18, 19} {
		iw := New(time.Date(2015, 4, day, 0, 0, 0, 0, time.UTC))

		assert.Equal(t, time.Date(2015, 4, 13, 0, 0, 0, 0, time.UTC), iw.LowerBound, "lower bound: april %d", day)
		assert.Equal(t, time.Date(2015, 4, 19, 23, 59, 59, 999999999, time.UTC), iw.UpperBound, "upper bound: april %d", day)
	}
}

func TestFromWeekNbr(t *testing.T) {
	testCases := []struct {
		yr    int
		wk    int
		lower time.Time
		upper time.Time
	}{
		{2015, 01,
			time.Date(2014, 12, 29, 0, 0, 0, 0, time.UTC),
			time.Date(2015, 1, 4, 23, 59, 59, 999999999, time.UTC),
		},
		{2015, 9,
			time.Date(2015, 2, 23, 0, 0, 0, 0, time.UTC),
			time.Date(2015, 3, 1, 23, 59, 59, 999999999, time.UTC),
		},
		{2015, 53,
			time.Date(2015, 12, 28, 0, 0, 0, 0, time.UTC),
			time.Date(2016, 1, 3, 23, 59, 59, 999999999, time.UTC),
		},
	}

	for i, tc := range testCases {
		t.Logf("test case %d\n", i)
		wk := FromWeekNbr(tc.yr, tc.wk)
		assert.Equal(t, tc.lower, wk.LowerBound, "lower bound of year: %d week: %d - %+v", tc.yr, tc.wk, wk)
		assert.Equal(t, tc.upper, wk.UpperBound, "upper bound of year: %d week: %d - %+v", tc.yr, tc.wk, wk)
	}
}

func TestISOWeekEquals(t *testing.T) {
	var dbt *DbTime
	iw0 := New(dbt.ToTime())
	iw1 := New(time.Date(2014, 1, 4, 0, 0, 0, 0, time.UTC))
	iw2 := New(time.Date(2014, 12, 31, 0, 0, 0, 0, time.UTC))
	iw3 := New(time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC))
	iw4 := New(time.Date(2015, 1, 5, 0, 0, 0, 0, time.UTC))

	// zero times are not equal to anything
	assert.False(t, iw0.Equals(iw1))
	assert.False(t, iw1.Equals(iw0))

	assert.False(t, iw1.Equals(iw2)) // year mismatch
	assert.True(t, iw2.Equals(iw3))
	assert.False(t, iw3.Equals(iw4)) // week mismatch
}

type DbTime struct {
	String string
	Time   time.Time
}

func (t *DbTime) ToTime() time.Time {
	if t == nil {
		return time.Time{}
	}
	if !t.Time.IsZero() {
		return t.Time
	}
	if t.String != "" && t.String != nullTime {
		ts, err := time.Parse(DbtimeLayout, t.String)
		if err != nil {
			return time.Time{}
		}
		t.Time = ts
		return t.Time
	}
	return time.Time{}
}

func (t *DbTime) ToString() string {
	if t.String == "" && !t.Time.IsZero() {
		t.String = t.Time.Format(DbtimeLayout)
	}
	return t.String
}
