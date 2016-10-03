package main

import (
	"fmt"
	"sort"
	"time"
)

type room struct {
	id       int
	capacity int
}

type event struct {
	roomId int
	start  time.Time
	end    time.Time
}

type result struct {
	room
	available bool
}

type schedule []event

func (s schedule) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s schedule) Len() int {
	return len(s)
}

func (s schedule) Less(i, j int) bool {
	if s[i].start.Before(s[j].start) {
		return true
	}
	if s[i].start == s[j].start {
		return s[i].end.Before(s[j].end)
	}
	return false
}

var (
	rooms []room = []room{
		room{1, 5},
		room{2, 10},
		room{3, 8},
		room{4, 20},
		room{5, 6},
		room{6, 4},
		room{7, 7},
		room{8, 12},
		room{9, 15},
		room{10, 50},
	}

	begin = time.Date(2016, time.August, 10, 9, 0, 0, 0, time.Local)

	sched []event = []event{
		event{1, begin, begin.Add(time.Hour)},
		event{2, begin.Add(time.Hour), begin.Add(2 * time.Hour)},
		event{3, begin.Add(3 * time.Hour), begin.Add(5 * time.Hour)},
		event{1, begin.Add(2 * time.Hour), begin.Add(4 * time.Hour)},
	}
)

func main() {
	end := begin.Add(2 * time.Hour)
	attendees := 7
	fmt.Printf("Event\nStart: %v\nEnd: %v\nAttendees: %d\n\n", begin, end, attendees)
	sort.Sort(schedule(sched))
	printSchedule(sched)
	printResults(availability(begin, end, attendees), attendees)
}

func availability(start, end time.Time, capacity int) []result {
	var results []result

	max := sort.Search(len(sched), func(i int) bool {
		e := sched[i]
		if e.start.After(end) {
			return true
		}
		if e.end.Before(start) {
			return true
		}
		return false
	})

	min := sort.Search(len(sched), func(i int) bool {
		t := sched[i].start
		return t == start || t.After(start)
	})

	// fmt.Printf("min: %d, max: %d\n", min, max)

	m := map[int]struct{}{}

	for _, v := range sched[min:max] {
		if _, ok := m[v.roomId]; ok {
			continue
		}
		r := rooms[v.roomId-1]
		results = append(results, result{r, false})
		m[v.roomId] = struct{}{}
	}

	for _, v := range rooms {
		if _, ok := m[v.id]; !ok {
			results = append(results, result{v, v.capacity >= capacity})
			m[v.id] = struct{}{}
		}
	}

	return results
}

func printResults(r []result, capacity int) {
	for i, v := range r {
		if i > 0 {
			fmt.Println()
		}
		status := "Available"
		if !v.available {
			if v.capacity < capacity {
				status = "Insufficient capacity"
			} else {
				status = "Booked"
			}
		}
		fmt.Printf("Room: %d\nCapacity: %d\nStatus: %s\n", v.id, v.capacity, status)
	}
}

func printSchedule(evts []event) {
	for i, e := range evts {
		if i > 0 {
			fmt.Println()
		}
		fmt.Printf("Start: %v\nEnd: %v\nRoom: %d\n\n", e.start, e.end, e.roomId)
	}
}
