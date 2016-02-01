package edb

import "sort"
import "sync"

type Hist struct {
  Max int
  Min int
  Count int
  Avg float64

  num_buckets int
  bucket_size int
  values []int
  avgs []float64

  outliers []int

  m *sync.Mutex
}

func (t *Table) NewHist(info *IntInfo) *Hist {

  buckets := 200 // resolution?
  h := &Hist{}

  h.num_buckets = buckets

  h.Max = int(info.Max)
  h.Min = int(info.Min)
  h.Avg = 0
  h.Count = 0

  h.outliers = make([]int, 0)
  
  h.m = &sync.Mutex{}


  size := info.Max - info.Min
  h.bucket_size = size / buckets
  if h.bucket_size == 0 {
    if (size < 100) {
      h.bucket_size = 1
      h.num_buckets = size
    } else {
      h.bucket_size = size / 100
      h.num_buckets = size / h.num_buckets
    }
  }

  h.num_buckets += 1


  h.values = make([]int, h.num_buckets + 1)
  h.avgs = make([]float64, h.num_buckets + 1)
  // we should use X buckets to start...
  return h
}

func (h *Hist) addValue(value int) {
  h.m.Lock()

  bucket_value := (value - h.Min) / h.bucket_size

  if bucket_value > h.num_buckets {
    h.outliers = append(h.outliers, value)
    h.m.Unlock()
    return
  }

  partial := h.avgs[bucket_value]

  if (value > h.Max) {
    h.Max = value
  } 

  if (value < h.Min) {
    h.Min = value
  }

  h.Count++
  h.Avg = h.Avg + (float64(value) - h.Avg) / float64(h.Count)

  // update counts
  count := h.values[bucket_value]
  count++
  h.values[bucket_value] = count

  // update bucket averages
  h.avgs[bucket_value] = partial + (float64(value) - partial) / float64(h.values[bucket_value])
  h.m.Unlock()
}

type ByVal []int
func (a ByVal) Len() int           { return len(a) }
func (a ByVal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVal) Less(i, j int) bool { return a[i] < a[j] }


func (h *Hist) getPercentiles() []int {	
  if h.Count == 0 {
    return make([]int, 0)
  }
  percentiles := make([]int, 101)
  keys := make([]int, 0)

  // unpack the bucket values!
  for k,_ := range h.values {
    keys = append(keys, k)
  }
  sort.Sort(ByVal(keys))

  percentiles[0] = h.Min
  count := 0
  prev_p := 0
  for _, k := range keys {
    key_count := h.values[k]
    count = count + key_count
    p := 100 * count / h.Count
    for ip := prev_p; ip < p; ip++ {
      percentiles[ip] = (k * h.bucket_size) + h.Min
    }
    percentiles[p] = k
    prev_p = p
  }



  return percentiles
}