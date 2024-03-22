// Package stopwatch provides a simple way to profile code
package stopwatch

import (
	"fmt"
	"time"
)

type Stopwatch interface {
	Start(name string) (*section, error)
	Stop(name string) (*section, error)
	Get(name string) (*section, error)
}

type Section interface {
	End() time.Time
	Start() time.Time
	Duration() time.Duration
	Lap()
}

type Lap interface {
	Duration() time.Duration
}

type stopwatch struct {
	sections []*section
}

type section struct {
	name    string
	event   *event
	laps    []*lap
	lastLap *lap
}

type lap struct {
	event *event
}

type event struct {
	begin, end time.Time
	duration   time.Duration
}

/*
Type Stopwatch
*/

// Creates a new instancen of stopwatch
func New() *stopwatch {
	return &stopwatch{}
}

// Starts a timer for the given section inside the stopwatch.
//
// Throws an error if a section with the given name already exists.
func (s *stopwatch) Start(name string) (*section, error) {
	sctn := s.getSection(name)
	if sctn != nil {
		return nil, fmt.Errorf("%s already exists", name)
	}
	sctn = &section{
		name:  name,
		event: newEvent(),
	}
	s.sections = append(s.sections, sctn)
	return sctn, nil
}

// Stops the timer for the given section inside the stopwatch.
//
// Laps within the given sections will also be stopped.
//
// Throws an error if no section with the given name exists.
func (s *stopwatch) Stop(name string) (*section, error) {
	sctn := s.getSection(name)

	if sctn == nil {
		return nil, fmt.Errorf("%s does not exist", name)
	}
	sctn.event.stop()
	if sctn.lastLap != nil {
		sctn.lastLap.event.stop()
	}
	return sctn, nil
}

// Returns the section with the given name, if it exists.
func (s *stopwatch) Get(name string) (*section, error) {
	sctn := s.getSection(name)
	if sctn == nil {
		return &section{}, fmt.Errorf("%s does not exist", name)
	}
	return sctn, nil
}

func (s *stopwatch) getSection(name string) *section {
	for _, sctn := range s.sections {
		if sctn.name == name {
			return sctn
		}
	}

	return nil
}

/*
Type Section
*/

// Returns the duration
func (sctn *section) Duration() time.Duration {
	return sctn.event.duration
}

// Returns the start time
func (sctn *section) Start() time.Time {
	return sctn.event.begin
}

// Returns the end time
func (sctn *section) End() time.Time {
	return sctn.event.end
}

// Creates a new Lap and stopped the current one if it exists.
func (sctn *section) Lap() {
	lap := &lap{
		event: newEvent(),
	}
	if sctn.lastLap != nil {
		sctn.lastLap.event.stop()
	}
	sctn.lastLap = lap
	sctn.laps = append(sctn.laps, lap)
}

// Returns the laps
func (sctn *section) Laps() []*lap {
	return sctn.laps
}

/*
Type Lap
*/

// Returns the duration
func (l *lap) Duration() time.Duration {
	return l.event.duration
}

/*
Type Event
*/

func newEvent() *event {
	e := event{}
	e.start()
	return &e
}

func (e *event) start() {
	e.begin = time.Now()
}

func (e *event) stop() {
	e.end = time.Now()
	e.duration = e.end.Sub(e.begin)
}
