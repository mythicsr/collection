package collection

import (
	"github.com/clipperhouse/typewriter"
)

// a convenience for passing values into templates; in MVC it'd be called a view model
type model struct {
	Type      typewriter.Type
	SliceName string
	// these templates only ever happen to use one type parameter
	TypeParameter typewriter.Type
	typewriter.TagValue
}

var templates = typewriter.TemplateSlice{
	collection,

	aggregateT, //
	all,        //
	any,        //
	average,
	averageT,
	count,
	distinct,
	distinctBy,
	each,
	first,
	groupByT,
	keyByT,
	max,
	maxT,
	maxBy,
	min,
	minT,
	minBy,
	selectT,
	single,
	sum,
	sumT,
	where,

	sort,
	isSorted,
	sortDesc,
	isSortedDesc,

	sortBy,
	isSortedBy,
	sortByDesc,
	isSortedByDesc,

	sortImplementation,
	sortInterface,

	shuffle,

	//paginate 分页用
	//chunk
	//append
	//findOne //O(1) map
	//New  //顺便初始化map
}
