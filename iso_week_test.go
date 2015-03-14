package isoweek

import (
	"testing"
	"time"

	"bitbucket.org/splice/api/models"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	for _, day := range []int{13, 14, 15, 16, 17, 18, 19} {
		iw := New(time.Date(2015, 4, day, 0, 0, 0, 0, time.UTC))

		assert.Equal(t, time.Date(2015, 4, 13, 0, 0, 0, 0, time.UTC), iw.LowerBound, "lower bound: april %d", day)
		assert.Equal(t, time.Date(2015, 4, 19, 23, 59, 59, 999999999, time.UTC), iw.UpperBound, "upper bound: april %d", day)
	}
}

func TestISOWeekEquals(t *testing.T) {
	var dbt *models.DbTime
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
