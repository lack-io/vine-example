// Copyright 2021 lack
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cron

import "time"

type Crontab interface {
	// Return the next activation time, later than the given time.
	// Next is invoked initially, and then each time the job is run.
	Next(time.Time) time.Time
}

// ConstantDelaySchedule represents a simple recurring duty cycle, e.g. "Every 5 minutes".
// It does not support jobs more frequent than once a second.
type ConstantDelayCron struct {
	Delay time.Duration
}

// Every returns a crontab Schedule that activates once every duration.
// Delays of less than a second are not supported (will round up to 1 second).
// Any fields less than a Second are truncated.
func Every(duration time.Duration) ConstantDelayCron {
	if duration < time.Second {
		duration = time.Second
	}
	return ConstantDelayCron{
		Delay: duration - time.Duration(duration.Nanoseconds())%time.Second,
	}
}

// Next returns the next time this should be run.
// This rounds so that the next activation time will be on the second.
func (cron ConstantDelayCron) Next(t time.Time) time.Time {
	return t.Add(cron.Delay - time.Duration(t.Nanosecond())*time.Nanosecond)
}

type ConstantLastCron struct {
	Delay time.Duration
}

func LastDay(duration time.Duration) ConstantLastCron {
	if duration < time.Second {
		duration = time.Second
	}

	return ConstantLastCron{
		Delay: duration - time.Duration(duration.Nanoseconds())%time.Second,
	}
}

func (cron ConstantLastCron) Next(t time.Time) time.Time {

	year, mon, day := t.AddDate(0, 1, -t.Day()).Date()

	return time.Date(year, mon, day, 0, 0, 0, 0, t.Location()).Add(cron.Delay)
}
