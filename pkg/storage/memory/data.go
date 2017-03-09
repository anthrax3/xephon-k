package memory

import (
	"sort"
	"sync"

	"github.com/xephonhq/xephon-k/pkg/common"
)

var initPointsLength = 100

type SeriesID string

// TODO: should be able to allow double etc later
type Data map[SeriesID]IntSeriesStore

type IntSeriesStore struct {
	mu     sync.RWMutex
	series common.IntSeries
	length int
}

func NewIntSeriesStore() *IntSeriesStore {
	series := common.IntSeries{}
	series.Points = make([]common.IntPoint, initPointsLength)
	return &IntSeriesStore{series: series, length: 0}
}

func (store *IntSeriesStore) WriteSeries(newSeries common.IntSeries) {
	store.mu.Lock()
	defer store.mu.Unlock()

	// merge the old series with new series
	// TODO: efficient merge sort
	// TODO: actually we can remove duplicate when merge by comparing with previous point

	// store.series should already be sorted, so we only sort the newSeries
	sort.Sort(common.ByTime(newSeries.Points))
	i := 0
	j := 0
	k := 0
	// NOTE: we can't use len(store.series.Points) because there might be duplicate
	oldLength := store.length
	newLength := len(newSeries.Points)
	points := make([]common.IntPoint, oldLength+newLength)
	for i < oldLength && j < newLength {
		if store.series.Points[i].TimeNano < newSeries.Points[j].TimeNano {
			points[k] = store.series.Points[i]
			i++
		} else if store.series.Points[i].TimeNano == newSeries.Points[j].TimeNano {
			// if there is duplicate, overwrite the old point with new point
			points[k] = store.series.Points[j]
			i++
			j++
		} else {
			points[k] = store.series.Points[j]
			j++
		}
		k++
	}
	store.length = k

	// copy what is left, should only have one array left
	// https://github.com/golang/go/wiki/SliceTricks
	if i < oldLength {
		points = append(points, store.series.Points[i:]...)
		// TODO: will there be +1 problem?
		store.length = k + oldLength - i
	}
	if j < newLength {
		points = append(points, newSeries.Points[j:]...)
		store.length = k + newLength - i
	}

	// TODO: maybe using pool is a good idea since there are a lot of merge when inserting series
	store.series.Points = points
}
