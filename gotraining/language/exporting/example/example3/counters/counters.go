package counters

type alertCounters int64

// New creates and returns values of the unexported type alertCounter.
func New(v int64) alertCounters {
	return alertCounters(v)
}
