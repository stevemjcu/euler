package main

import (
	"fmt"
	"strconv"
	"time"
)

var daysPerMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

type Date struct {
	weekday, day int
	month, year  int
}

func (d Date) Weekday() int { return d.weekday + 1 }
func (d Date) Day() int     { return d.day + 1 }
func (d Date) Month() int   { return d.month + 1 }
func (d Date) Year() int    { return d.year }

func (d Date) Next() Date {
	d.weekday = (d.weekday + 1) % 7
	leap := 0
	if d.month == 1 && d.IsLeapYear() {
		leap = 1
	}
	d.day = (d.day + 1) % (daysPerMonth[d.month] + leap)
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
	s := strconv.Itoa(d.Day()) + "-"
	s += strconv.Itoa(d.Month()) + "-"
	s += strconv.Itoa(d.Year())
	return s
}

func solveCustom() int {
	count, d := 0, Date{1, 0, 0, 1900}
	for i := 0; d.year < 2001; i++ {
		if d.Year() > 1900 && d.Day() == 1 && d.Weekday() == 1 {
			fmt.Println(d)
			count++
		}
		d = d.Next()
	}
	return count
}

func solveBuiltIn() int {
	count, d := 0, time.Date(1901, 1, 1, 0, 0, 0, 0, time.UTC)
	for d.Year() < 2001 {
		if d.Day() == 1 && d.Weekday() == time.Sunday {
			fmt.Println(d)
			count++
		}
		d = d.Add(time.Hour * 24)
	}
	return count
}

func main() {
	fmt.Println(solveCustom())
}
