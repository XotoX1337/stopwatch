package stopwatch_test

import (
	"testing"
	"time"

	"github.com/XotoX1337/stopwatch"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDuration(t *testing.T) {
	assert := assert.New(t)

	sectionName := genSectionName()

	stopwatch := stopwatch.New()
	stopwatch.Start(sectionName)

	startTime := time.Now()
	time.Sleep(250 * time.Millisecond)
	stopwatch.Stop(sectionName)
	endTime := startTime.Add(250 * time.Millisecond)

	section, _ := stopwatch.Get(sectionName)
	assert.WithinDuration(endTime, section.End(), 50*time.Millisecond)
}

func TestStartAlreadyExists(t *testing.T) {
	assert := assert.New(t)

	sectionName := genSectionName()

	stopwatch := stopwatch.New()
	_, err := stopwatch.Start(sectionName)
	assert.NoError(err)
	_, err = stopwatch.Start(sectionName)
	assert.Error(err)
}

func TestStopDoesNotExist(t *testing.T) {
	assert := assert.New(t)

	sectionName := genSectionName()

	stopwatch := stopwatch.New()
	_, err := stopwatch.Stop(sectionName)
	assert.Error(err)
}

func TestMultipleSections(t *testing.T) {
	assert := assert.New(t)

	firstSectionName := genSectionName()
	secondSectionName := genSectionName()

	timeOutOne := 250 * time.Millisecond
	timeOutTwo := 125 * time.Millisecond

	expectedDurationOne := timeOutOne
	expectedDurationTwo := timeOutTwo + timeOutOne

	stopwatch := stopwatch.New()
	firstSection, err := stopwatch.Start(firstSectionName)
	assert.NoError(err)
	startTime := time.Now()

	time.Sleep(timeOutOne)

	seconSection, err := stopwatch.Start(secondSectionName)
	assert.NoError(err)

	_, err = stopwatch.Stop(firstSectionName)
	endTimeOne := startTime.Add(expectedDurationOne)
	assert.NoError(err)

	time.Sleep(timeOutTwo)
	_, err = stopwatch.Stop(secondSectionName)
	endTimeTwo := startTime.Add(expectedDurationTwo)
	assert.NoError(err)

	//check if firstSection.duration equals ~250ms
	assert.WithinDuration(endTimeOne, firstSection.End(), 50*time.Millisecond)
	//check if secondSection.duration equals ~375ms
	assert.WithinDuration(endTimeTwo, seconSection.End(), 50*time.Millisecond)
}

func genSectionName() string {
	return uuid.New().String()
}

//ToDo:test lap
