package data

import (
	"fmt"
	"reflect"
	"time"

	"github.com/go-faker/faker/v4"
)

const (
	ValidForTag = "validFor"
)

func CustomFaker() {
	_ = faker.AddProvider(ValidForTag, func(v reflect.Value) (interface{}, error) {
		var (
			startDateStr string
			endDateStr   string
		)
		timeStamp := fmt.Sprintf("%s %s", faker.BaseDateFormat, faker.TimeFormat)

		t1 := time.Unix(faker.RandomUnixTime(), 0)
		t2 := time.Unix(faker.RandomUnixTime(), 0)

		if t1.Before(t2) {
			startDateStr = t1.Format(timeStamp)
			endDateStr = t2.Format(timeStamp)
		} else {
			startDateStr = t2.Format(timeStamp)
			endDateStr = t1.Format(timeStamp)
		}
		return ValidFor{StartDateTime: startDateStr, EndDateTime: endDateStr}, nil
	})
}
