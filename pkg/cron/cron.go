package cron

import (
	"log"
	"time"
)

type Task struct {
	Name     string
	Schedule Schedule
	Action   func() error
}

type Schedule struct {
	IsDate  bool
	Day     int
	Hours   int
	Minuts  int
	Seconds int
}

type Cron struct {
	tasks []*Task
}

func NewCron() *Cron {
	return &Cron{}
}

func (c *Cron) AddTask(task Task) {
	c.tasks = append(c.tasks, &task)
}

func (c *Cron) Start() {
	for _, task := range c.tasks {
		go c.runTask(task)
	}
}

func (c *Cron) runTask(task *Task) {
	for {
		nextRun := task.Schedule.getTime()
		now := time.Now()
		duration := nextRun.Sub(now)

		if duration < 0 {
			nextRun = task.Schedule.getTime()
			duration = nextRun.Sub(now)
		}

		time.Sleep(duration)

		err := task.Action()
		if err != nil {
			log.Printf("Action %s is failed with error: %s", task.Name, err)
		}

		nextRun = task.Schedule.getTime()
	}
}

func (s *Schedule) calcTime() time.Duration {
	return (time.Hour*24*time.Duration(s.Day) + time.Hour*time.Duration(s.Hours) + time.Minute*time.Duration(s.Minuts) + time.Second*time.Duration(s.Seconds))
}

func (s *Schedule) getTime() time.Time {
	if s.IsDate {
		year, month, day := time.Now().AddDate(0, 0, 1+s.Day).Date()
		loc := time.Now().Location()
		date := time.Date(year, month, day, s.Hours, s.Minuts, s.Seconds, 0, loc)
		return date
	}
	return time.Now().Add(s.calcTime())
}
