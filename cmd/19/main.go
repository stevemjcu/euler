package main

import (
	"fmt"
	"strconv"
)

var daysInMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var daysOfWeek = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

type Date struct {
	weekday, day int
	month, year  int
}

func (d Date) Next() Date {
	d.weekday = (d.weekday + 1) % 7
	leap := 0
	if d.month == 1 && d.IsLeapYear() {
		leap = 1
	}
	d.day = (d.day + 1) % (daysInMonth[d.month] + leap)
	if d.day == 0 {
		d.month = (d.month + 1) % 12
		if d.month == 0 {
			d.year++
		}
	}
	return d
}

func (d Date) IsLeapYear() bool {
	return d.year%4 == 0 && (d.year%100 != 0 || d.year%400 == 0)
}

func (d Date) String() string {
	s := daysOfWeek[d.weekday] + " "
	s += strconv.Itoa(d.day+1) + "-"
	s += strconv.Itoa(d.month+1) + "-"
	s += strconv.Itoa(d.year)
	return s
}

func main() {
	count, d := 0, Date{1, 0, 0, 1900}
	for i := 0; d.year < 2001; i++ {
		if d.year > 1900 && d.day == 0 && d.weekday == 0 {
			fmt.Println(d)
			count++
		}
		d = d.Next()
	}
	fmt.Println(count)
}
