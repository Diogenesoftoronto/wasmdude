package pages

type MyArticles struct{}

func (c *MyArticles) Title() string      { return "All the Articles and Blogs" }
func (c *MyArticles) ShortTitle() string { return "My Articles" }
func (c *MyArticles) MetaDescription() string {
	return "Various topics about cyberspace, VR, Philosophy, and being a black enby gopher"
}

//
