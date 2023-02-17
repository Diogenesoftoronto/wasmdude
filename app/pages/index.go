package pages

type Index struct {
}

func (c *Index) Title() string { return "Vugu: A modern UI library for Go+WebAssembly" }
func (c *Index) ShortTitle() string { return "Vugu Home" }
func (c *Index) MetaDescription() string { return "Pure Go. Targets WebAssembly (and/or server). Most modern browsers supported. Experimental, for now. Really cool." }
