package utils

import (
	"github.com/go-gota/gota/dataframe"
)

func DataFramePaginate(data interface{}, page int, limit int, sortOf string, sortBy string) []map[string]interface{} {
	var indexes []int
	df := dataframe.LoadStructs(data)
	if sortBy != "" {
		if sortOf == "desc" {
			df = df.Arrange(dataframe.RevSort(sortBy))
		} else {
			df = df.Arrange(dataframe.Sort(sortBy))
		}
	}
	startIndex := page * limit
	endIndex := (page + 1) * limit
	if len(df.Maps()) < endIndex {
		endIndex = len(df.Maps())
	}
	for startIndex < endIndex {
		indexes = append(indexes, startIndex)
		startIndex++
	}
	df = df.Subset(indexes)
	return df.Maps()
}
func DataFrameLeftJoin(data1 interface{}, data2 interface{}, key string) []map[string]interface{} {
	df1 := dataframe.LoadStructs(data1)
	df2 := dataframe.LoadStructs(data2)
	dfJoin := df1.LeftJoin(df2, key)
	return dfJoin.Maps()
}
func DataFrameInnerJoin(data1 interface{}, data2 interface{}, key string) []map[string]interface{} {
	df1 := dataframe.LoadStructs(data1)
	df2 := dataframe.LoadStructs(data2)
	dfJoin := df1.InnerJoin(df2, key)
	return dfJoin.Maps()
}
