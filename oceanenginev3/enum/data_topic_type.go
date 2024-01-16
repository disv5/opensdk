package enum


type DataTopicType string

const (
	BASIC_DATA DataTopicType = "BASIC_DATA"
	QUERY_DATA DataTopicType = "QUERY_DATA"
	BIDWORD_DATA DataTopicType = "BIDWORD_DATA"
	MATERIAL_DATA DataTopicType = "MATERIAL_DATA"
)



type OrderByType string

const (
	ASC OrderByType = "ASC"
	DESC OrderByType = "DESC"
)
