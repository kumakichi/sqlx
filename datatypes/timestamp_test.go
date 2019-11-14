package datatypes

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	DATE_LAYOUT_S1 = "2006-01-02"
	TIME_LAYOUT_S1 = "2006-01-02 15:04:05"
)

func TestTimestamp(t *testing.T) {
	tt := require.New(t)

	t0, _ := time.Parse(time.RFC3339, "2017-03-27T23:58:59+08:00")
	dt := Timestamp(t0)
	tt.Equal("2017-03-27T23:58:59+08:00", dt.String())
	tt.Equal("2017-03-27T23:58:59+08:00", dt.Format(time.RFC3339))
	tt.Equal(int64(1490630339), dt.Unix())
	tt.Equal(TimestampUnixZero.Unix(), int64(0))
	tt.Equal(TimestampUnixZero.IsZero(), true)
	tt.Equal("1970-01-01T08:00:00+08:00", TimestampUnixZero.String())

	{
		input := "1970-01-01"
		r, err := ParseTimestampFromStringWithLayout(input, DATE_LAYOUT_S1)
		tt.Nil(err)
		input = "1970-01-01 08:00:00"
		r, err = ParseTimestampFromStringWithLayout(input, TIME_LAYOUT_S1)
		tt.Nil(err)
		tt.Equal(TimestampZero.IsZero(), r.IsZero())

		input = "1970-01-01 08:00:00"
		_, err = ParseTimestampFromStringWithLayout(input, DATE_LAYOUT_S1)
		tt.NotNil(err)
		input = "1970-01-01"
		_, err = ParseTimestampFromStringWithLayout(input, TIME_LAYOUT_S1)
		tt.NotNil(err)
	}

	{
		dateString, err := dt.MarshalText()
		tt.NoError(err)
		tt.Equal("2017-03-27T23:58:59+08:00", string(dateString))

		dt2 := TimestampZero
		tt.True(dt2.IsZero())
		err = dt2.UnmarshalText(dateString)
		tt.NoError(err)
		tt.Equal(dt, dt2)
		tt.False(dt2.IsZero())
	}
	{
		value, err := dt.Value()
		tt.NoError(err)
		tt.Equal(int64(1490630339), value.(int64))

		dt2 := TimestampZero
		tt.True(dt2.IsZero())
		err = dt2.Scan(value)
		tt.NoError(err)
		tt.Equal(dt.In(CST), dt2.In(CST))
		tt.False(dt2.IsZero())
	}
	{
		dt3 := TimestampZero
		err := dt3.UnmarshalText([]byte(""))
		tt.NoError(err)
	}
}