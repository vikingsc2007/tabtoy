package printer

type PrinterContext struct {
	outFile string
	p       Printer
	name    string
}

func (self *PrinterContext) Start(g *Globals) bool {

	log.Infof("[%s] %s\n", self.name, self.outFile)

	bf := self.p.Run(g, self.outFile)

	if bf == nil {
		return false
	}

	var allok bool = true 

	for k, v := range bf {
		// fmt.Println(k, v)
		log.Infof("----(%s)\n", k)
		if v == nil {
			return false
		}
	
		allok =  (v.WriteFile(k) == nil) && allok
	}

	return allok
}

type Printer interface {
	Run(g *Globals,outFile string) map[string]*Stream
}

var printerByExt = make(map[string]Printer)

func RegisterPrinter(ext string, p Printer) {

	if _, ok := printerByExt[ext]; ok {
		panic("duplicate printer")
	}

	printerByExt[ext] = p
}
