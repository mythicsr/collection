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

	reduceT, //
	isAll,   //
	isAny,   //
	//average,
	averageT, //num
	count,
	//distinct,
	distinctBy, //
	forEach,
	findOne,
	groupByT,
	keyByT,
	//max,
	//maxT,
	maxBy,
	//min,
	//minT,
	minBy,
	selectT,
	single,
	//sum,
	sumT,
	find,

	//sort,
	//isSorted,
	//sortDesc,
	//isSortedDesc,

	sortBy,
	//sortByDesc,
	isSortedBy,
	//isSortedByDesc,

	shuffle,
	chunk,
	deepCopy,

	includes,
	toSlice,

	toInterfaces,

	lpush,
	lpop,
	rpush,
	rpop,
	removeByIndex,
	remove,

	paginate,

	sortImplementation,
	sortInterface,
}
