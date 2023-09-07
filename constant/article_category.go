package constant

const (
	ARTICLE_CATEGORY_TYPE_PRIMARY   = 1
	ARTICLE_CATEGORY_TYPE_SECONDARY = 2
)

var ARTICLE_CATEGORY_TYPE = map[int]string{
	ARTICLE_CATEGORY_TYPE_PRIMARY:   "Primary",
	ARTICLE_CATEGORY_TYPE_SECONDARY: "Secondary",
}
