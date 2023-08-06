package dbscan

import (
	"fmt"
	"sort"
	"time"
)

type ClusterablePoint interface {
	GetPoint() []float64
	String() string
	GetID() int64
	GetDate() time.Time
}

type IDPoint struct {
	ID    int64
	Point []float64
	Date  time.Time
}

func NewIDPoint(ID int64, point []float64, date time.Time) *IDPoint {
	return &IDPoint{
		ID:    ID,
		Point: point,
		Date:  date,
	}
}

func (self *IDPoint) String() string {
	return fmt.Sprintf("\"%d\": %v", self.ID, self.Point)
}

func (self *IDPoint) GetPoint() []float64 {
	return self.Point
}

func (self *IDPoint) GetID() int64 {
	return self.ID
}

func (self *IDPoint) GetDate() time.Time {
	return self.Date
}

func (self *IDPoint) Copy() *IDPoint {
	var p = new(IDPoint)
	p.ID = self.ID
	p.Point = self.Point
	p.Date = self.Date
	copy(p.Point, self.Point)
	return p
}

// ClusterablePointSlice Slice attaches the methods of Interface to []float64, sorting in increasing order.
type ClusterablePointSlice struct {
	Data          []ClusterablePoint
	SortDimension int
}

func (self ClusterablePointSlice) Len() int { return len(self.Data) }
func (self ClusterablePointSlice) Less(i, j int) bool {
	return self.Data[i].GetPoint()[self.SortDimension] < self.Data[j].GetPoint()[self.SortDimension]
}
func (self ClusterablePointSlice) Swap(i, j int) {
	self.Data[i], self.Data[j] = self.Data[j], self.Data[i]
}

// Sort is a convenience method.
func (self ClusterablePointSlice) Sort() { sort.Sort(self) }

func NamedPointToClusterablePoint(in []*IDPoint) (out []ClusterablePoint) {
	out = make([]ClusterablePoint, len(in))
	for i, v := range in {
		out[i] = v
	}
	return
}
