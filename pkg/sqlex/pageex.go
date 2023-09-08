package sqlex

import (
	"strconv"
)

type OrderItem struct {
	Column string `json:"column"`
	Asc    bool   `json:"asc"`
}

type Page struct {
	PageNo      int64       `json:"page_no,optional"`
	PageSize    int64       `json:"page_size,optional"`
	StartTime   int64       `json:"start_time,optional"`
	EndTime     int64       `json:"end_time,optional"`
	SortBy      []OrderItem `json:"sort_by,optional"`
	IgnoreTotal bool        `json:"need_total,optional"`
}

func (p *Page) OrderBy() string {
	size := len(p.SortBy)
	if size == 0 {
		return "order by id desc"
	}

	order := "order by "
	for i, v := range p.SortBy {
		order = order + v.Column + " "
		if !v.Asc {
			order = order + " desc "
		}
		if size-1 == i {

		} else {
			order = order + ","
		}
	}
	return order
}

// OrderByFieldMap 映射字段名
func (p *Page) OrderByFieldMap(fieldMap map[string]string) string {
	size := len(p.SortBy)
	if size == 0 {
		field := "id"
		if newField, ok := fieldMap[field]; ok {
			field = newField
		}
		return "order by " + field + " desc"
	}

	order := "order by "
	for i, v := range p.SortBy {
		field := v.Column
		if newField, ok := fieldMap[field]; ok {
			field = newField
		}
		order = order + field + " "
		if !v.Asc {
			order = order + " desc "
		}
		if size-1 != i {
			order = order + ","
		}
	}
	return order
}

func (p *Page) PageLimit() string {
	if p.PageNo == 0 {
		p.PageNo = 1
	}

	if p.PageSize == 0 {
		p.PageSize = 20
	}
	return "limit " + strconv.FormatInt((p.PageNo-1)*p.PageSize, 10) + "," + strconv.FormatInt(p.PageSize, 10)
}

func (p *Page) PageTimeRange(filed string) string {
	where := " where 1=1 "
	if p.StartTime > 0 {
		where += " and " + filed + " >= " + strconv.Itoa(int(p.StartTime))
	}
	if p.EndTime > 0 {
		where += " and " + filed + " < " + strconv.Itoa(int(p.EndTime))
	}
	return where
}
