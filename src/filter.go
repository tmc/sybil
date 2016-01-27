package edb

type Filter interface {
  Filter(*Record) bool;
}

type NoFilter struct { }
func (f NoFilter) Filter(r *Record) bool {
  return false;
}

type IntFilter struct {
  Field string
  FieldId int16
  Op string
  Value int

  table *Table
}

func (filter IntFilter) Filter(r *Record) bool {
  if r.Populated[filter.FieldId] == false {
    return true
  }

  field := r.Ints[filter.FieldId]
  switch filter.Op {
    case "gt":
      return int(field) < int(filter.Value)

    case "lt":
      return int(field) > int(filter.Value)

    default:

  }

  return true
}


func (t *Table) IntFilter(name string, op string, value int) IntFilter {
  intFilter := IntFilter{Field: name, FieldId: t.get_key_id(name), Op: op, Value: value}
  intFilter.table = t;

  return intFilter
  
}
