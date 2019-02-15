package product

type Product struct {
	CSCode             string
	Status             string
	Name               string
	Category           string
	Comment            string
	Description        string
	Size               int
	CasePack           int
	CostPerOunce       float64
	NewRetail          float64
	CurrentRetail      float64
	WarehouseInventory int
	WarehouseOnOrder   int
	Stores             []store
}

type store struct {
	ID      string
	Name    string
	Qty     int
	Address string
	City    string
	Phone   string
}

func (p *Product) QtyInStores() int {
	count := 0
	for _, s := range p.Stores {
		count += s.Qty
	}
	return count
}

func (p *Product) StatusName() string {
	switch p.Status {
	case "1":
		return "General Distribution"
	case "D":
		return "Discontinued General Item"
	case "S":
		return "Special Order"
	case "L":
		return "Regular Limited Item"
	case "X":
		return "Limited Discontinued"
	case "N":
		return "Unavailable General Item"
	case "A":
		return "Limited Allocated Product"
	case "U":
		return "Unavailable Limited Item"
	case "T":
		return "Trial"
	}
	return ""
}

func (p *Product) GeneralCategory() string {
	if len(p.Category) < 1 {
		return ""
	}
	switch p.Category[:1] {
	case "A":
		return "SPIRITS"
	case "C":
		return "LIQUEURS - SCHNAPPS - PREMIXED"
	case "E":
		return "WINE - DESSERT, APERITIF, & LATE HARVEST"
	case "G":
		return "WINE - FRUIT, CIDER, AND SAKE"
	case "I":
		return "WINE - CHAMPAGNE & SPARKLING"
	case "K":
		return "WINE - WHITE TABLE"
	case "L":
		return "LIMITED TIME OFFER"
	case "M":
		return "WINE - BLUSH & ROSE"
	case "P":
		return "WINE - RED TABLE"
	case "R":
		return "FLAVORED MALT BEVERAGES - R"
	case "T":
		return "BEER"
	case "Y":
		return "GIFT & SPECIAL ORDER - Y"
	}
	return ""
}

func (p *Product) CategoryDescription() string {
	switch p.Category {
	case "ADF":
		return "VODKA - BASIC"
	case "ADM":
		return "VODKA - FLAVORED"
	case "ADW":
		return "VODKA - DOMESTIC"
	case "ADX":
		return "VODKA - IMPORTED"
	case "AHG":
		return "GIN - DRY"
	case "AHM":
		return "GIN - FLAVORED"
	case "AHW":
		return "GIN - DOMESTIC"
	case "AHX":
		return "GIN - IMPORTED"
	case "ALE":
		return "RUM - LIGHT"
	case "ALJ":
		return "RUM - DARK"
	case "ALP":
		return "RUM - SPICED AND FLAVORED"
	case "ALU":
		return "RUM - MISC"
	case "ALW":
		return "RUM - DOMESTIC"
	case "ALX":
		return "RUM - IMPORTED"
	case "APD":
		return "TEQUILA - SILVER"
	case "APF":
		return "TEQUILA - GOLD"
	case "APJ":
		return "TEQUILA - REPOSADO & ANEJO"
	case "APM":
		return "TEQUILA - FLAVORED"
	case "APQ":
		return "MEZCAL"
	case "ATB":
		return "BRANDY - BASIC"
	case "ATG":
		return "BRANDY - COGNAC & ARMAGNAC"
	case "ATJ":
		return "BRANDY - CALVADOS & APPLEJACK"
	case "ATM":
		return "BRANDY - SPICED & FLAVORED"
	case "ATR":
		return "BRANDY - GRAPPA & EAUX DE VIE"
	case "ATU":
		return "BRANDY - MISC"
	case "ATW":
		return "BRANDY - DOMESTIC"
	case "ATX":
		return "BRANDY - IMPORTED"
	case "AWB":
		return "WHISKEY - CANADIAN"
	case "AWE":
		return "WHISKEY - BLENDED"
	case "AWH":
		return "WHISKEY - BOURBON & TENNESSEE"
	case "AWK":
		return "WHISKEY - BOURBON SINGLE BARRE"
	case "AWN":
		return "WHISKEY - IRISH"
	case "AWR":
		return "WHISKY - SCOTCH BLENDED"
	case "AWS":
		return "WHISKY - SCOTCH SINGLE MALT"
	case "AWT":
		return "WHISKEY - FLAVORED"
	case "AWU":
		return "WHISKEY - MISC"
	case "AWW":
		return "WHISKEY - DOMESTIC"
	case "AWX":
		return "WHISKEY - IMPORTED"
	case "CH":
		return "LIQUEURS"
	case "CHB":
		return "LIQUEURS - TRIPLE SEC & ORANGE"
	case "CHE":
		return "LIQUEURS - FRUIT"
	case "CHH":
		return "LIQUEURS - AMARETTO"
	case "CHK":
		return "LIQUEURS - COFFEE"
	case "CHN":
		return "LIQUEURS - CHOCOLATE"
	case "CHQ":
		return "LIQUEURS - ANISE & SAMBUCA"
	case "CHS":
		return "LIQUEURS - BITTERS & HERBAL"
	case "CHU":
		return "LIQUEURS - MISC"
	case "CHV":
		return "LIQUEURS - CREAM"
	case "CHW":
		return "LIQUEURS - DOMESTIC"
	case "CHX":
		return "LIQUEURS - IMPORTED"
	case "CPC":
		return "SCHNAPPS - PEPPERMINT"
	case "CPG":
		return "SCHNAPPS - FRUIT"
	case "CPK":
		return "SCHNAPPS - HOT & SPICED"
	case "CPP":
		return "SCHNAPPS - BUTTERSCOTCH"
	case "CPS":
		return "SCHNAPPS - GOLD"
	case "CPU":
		return "SCHNAPPS - MISC"
	case "CPW":
		return "SCHNAPPS - DOMESTIC"
	case "CPX":
		return "SCHNAPPS - IMPORTED"
	case "CWF":
		return "PREMIXED - MARGARITA"
	case "CWL":
		return "PREMIXED - LONG ISLAND TEA"
	case "CWR":
		return "PREMIXED - MUDSLIDE"
	case "CWU":
		return "PREMIXED - MISC"
	case "EDD":
		return "VERMOUTH - DRY"
	case "EDP":
		return "VERMOUTH - SWEET"
	case "EDU":
		return "VERMOUTH - MISC, HERBAL & FLAV"
	case "EDW":
		return "VERMOUTH - DOMESTIC"
	case "EDX":
		return "VERMOUTH - IMPORTED"
	case "EHD":
		return "SHERRY - DRY - FINO"
	case "EHJ":
		return "SHERRY - MEDIUM - AMONTILLADO"
	case "EHP":
		return "SHERRY - CREAM"
	case "EHU":
		return "SHERRY - MISC"
	case "EHW":
		return "SHERRY - DOMESTIC"
	case "EHX":
		return "SHERRY - IMPORTED"
	case "ELD":
		return "MARSALA - DRY"
	case "ELP":
		return "MARSALA - SWEET"
	case "ELW":
		return "MARSALA - DOMESTIC"
	case "ELX":
		return "MARSALA - IMPORTED"
	case "EPD":
		return "MADIERA - DRY- RAINWATER, SERC"
	case "EPP":
		return "MADEIRA - SWEET - BUAL & MALMS"
	case "EPW":
		return "MADEIRA - DOMESTIC"
	case "EPX":
		return "MADEIRA - IMPORTED"
	case "ESC":
		return "PORT - RUBY"
	case "ESG":
		return "PORT - TAWNY"
	case "ESK":
		return "PORT - VINTAGE CHARACTER"
	case "ESN":
		return "PORT - LATE BOTTLED"
	case "ESR":
		return "PORT - VINTAGE & QUINTA"
	case "ESU":
		return "PORT - MISC"
	case "ESW":
		return "PORT - DOMESTIC"
	case "ESX":
		return "PORT - IMPORTED"
	case "EVF":
		return "LATE HARVEST - MUSCAT"
	case "EVL":
		return "LATE HARVEST - RIESLING, BA, T"
	case "EVR":
		return "LATE HARVEST - SAUV, SEMILLON,"
	case "EVU":
		return "LATE HARVEST - WHITE MISC"
	case "EVV":
		return "LATE HARVEST - RED MISC"
	case "GFE":
		return "WINE - FRUIT"
	case "GFL":
		return "WINE - COOLER"
	case "GFS":
		return "WINE - SANGRIA"
	case "GLC":
		return "CIDER - BASIC"
	case "GLM":
		return "CIDER - FLAVORED"
	case "GPW":
		return "SAKE - DOMESTIC"
	case "GPX":
		return "SAKE - IMPORTED"
	case "IHF":
		return "SPARKLING WINE - COLD DUCK & P"
	case "IHI":
		return "SPARKLING WINE - SPUMANTE & DE"
	case "IHL":
		return "SPARKLING WINE - EXTRA DRY"
	case "IHP":
		return "SPARKLING WINE - BRUT & BLANC"
	case "IHS":
		return "SPARKLING WINE - BLANC DE NOIR"
	case "IHU":
		return "SPARKLING WINE - MISC"
	case "IPD":
		return "SPARKLING WINE - DOMESTIC"
	case "IPH":
		return "SPARKLING WINE - FRENCH & CHAM"
	case "IPI":
		return "SPARKLING WINE - PROSECCO"
	case "IPL":
		return "SPARKLING WINE - ITALIAN"
	case "IPP":
		return "SPARKLING WINE - SPANISH CAVA"
	case "IPX":
		return "SPARKLING WINE - IMPORTED MISC"
	case "KFE":
		return "WHITE GENERIC - RHINE"
	case "KFJ":
		return "WHITE GENERIC - CHABLIS"
	case "KFP":
		return "WHITE GENERIC - TABLE & PROPRI"
	case "KKB":
		return "WHITE VARIETAL - CHARDONNAY"
	case "KKD":
		return "WHITE VARIETAL - SAUVIGNON BLA"
	case "KKF":
		return "WHITE VARIETAL - PINOT BLANC"
	case "KKH":
		return "WHITE VARIETAL - PINOT GRIS"
	case "KKJ":
		return "WHITE VARIETAL - VIOGNIER"
	case "KKK":
		return "WHITE VARIETAL - MARSANNE & RO"
	case "KKL":
		return "WHITE VARIETAL - SEMILLON"
	case "KKN":
		return "WHITE VARIETAL - GEWURZTRAMINE"
	case "KKP":
		return "WHITE VARIETAL - CHENIN BLANC"
	case "KKR":
		return "WHITE VARIETAL - RIESLING"
	case "KKT":
		return "WHITE VARIETAL - MUSCAT & MOSC"
	case "KKU":
		return "WHITE VARIETAL - MISC"
	case "KRA":
		return "FRENCH WHITE - BURGUNDY"
	case "KRB":
		return "FRENCH WHITE - BORDEAUX & LOIR"
	case "KRC":
		return "FRENCH WHITE - ALSATIAN"
	case "KRD":
		return "FRENCH WHITE - MISC"
	case "KRE":
		return "FRENCH WHITE - VARIETAL"
	case "KRG":
		return "GERMAN WHITE - RHEIN"
	case "KRH":
		return "GERMAN WHITE - MOSEL"
	case "KRI":
		return "GERMAN WHITE - MISC & VARIETAL"
	case "KRJ":
		return "AUSTRIAN WHITE"
	case "KRK":
		return "ITALIAN WHITE - PINOT GRIGIO"
	case "KRL":
		return "ITALIAN WHITE - MISC DOC"
	case "KRM":
		return "ITALIAN WHITE - VARIETAL"
	case "KRN":
		return "SPANISH WHITE"
	case "KRO":
		return "PORTUGESE WHITE"
	case "KRP":
		return "GREEK WHITE"
	case "KRQ":
		return "SOUTH AFRICAN WHITE"
	case "KRR":
		return "CHILEAN WHITE"
	case "KRS":
		return "ARGENTINE WHITE"
	case "KRT":
		return "AUSTRALIAN WHITE"
	case "KRV":
		return "NEW ZEALAND WHITE"
	case "KRY":
		return "IMPORTED WHITE - MISC"
	case "LTO":
		return "WINE OFFER"
	case "MHB":
		return "BLUSH WINE - GENERIC"
	case "MHJ":
		return "BLUSH WINE - WHITE ZINFANDEL"
	case "MHQ":
		return "BLUSH WINE - WHITE GRENACHE"
	case "MHU":
		return "BLUSH WINE - MISC"
	case "MHW":
		return "BLUSH WINE - DOMESTIC"
	case "MPB":
		return "ROSE WINE - GENERIC"
	case "MPQ":
		return "ROSE WINE - VARIETAL & VIN GRI"
	case "MPU":
		return "ROSE WINE - MISC"
	case "MPW":
		return "ROSE WINE - DOMESTIC"
	case "MPX":
		return "ROSE WINE - IMPORTED"
	case "PFD":
		return "RED GENERIC - BURGUNDY"
	case "PFP":
		return "RED GENERIC - RED TABLE & PROP"
	case "PLB":
		return "RED VARIETAL - CABERNET"
	case "PLC":
		return "RED VARIETAL - CARMENERE"
	case "PLD":
		return "RED VARIETAL - MERITAGE"
	case "PLE":
		return "RED VARIETAL - MALBEC"
	case "PLF":
		return "RED VARIETAL - MERLOT"
	case "PLG":
		return "RED VARIETAL - CAB FRANC, VERD"
	case "PLH":
		return "RED VARIETAL - ZINFANDEL"
	case "PLJ":
		return "RED VARIETAL - SYRAH, PETITE S"
	case "PLL":
		return "RED VARIETAL - GRENACHE"
	case "PLN":
		return "RED VARIETAL - MOURVEDRE"
	case "PLP":
		return "RED VARIETAL - BARBERA"
	case "PLQ":
		return "RED VARIETAL - SANGIOVESE, NEB"
	case "PLR":
		return "RED VARIETAL - PINOT NOIR"
	case "PLS":
		return "RED VARIETAL - GAMAY & BEAUJOL"
	case "PLU":
		return "RED VARIETAL - MISC"
	case "PRA":
		return "FRENCH RED - BORDEAUX"
	case "PRB":
		return "FRENCH RED - BURGUNDY"
	case "PRC":
		return "FRENCH RED - BEAUJOLAIS"
	case "PRD":
		return "FRENCH RED - RHONE"
	case "PRE":
		return "FRENCH RED - MISC AOC"
	case "PRF":
		return "FRENCH RED - VARIETAL"
	case "PRG":
		return "GERMAN & AUSTRIAN RED"
	case "PRH":
		return "ITALIAN RED - CHIANTI"
	case "PRI":
		return "ITALIAN RED - TUSCANY"
	case "PRJ":
		return "ITALIAN RED - PIEDMONT"
	case "PRK":
		return "ITALIAN RED - MISC DOC"
	case "PRL":
		return "ITALIAN RED - VARIETAL"
	case "PRM":
		return "ITALIAN RED - SICILY"
	case "PRN":
		return "SPANISH RED - RIOJA"
	case "PRP":
		return "SPANISH RED - MISC"
	case "PRQ":
		return "PORTUGESE RED"
	case "PRR":
		return "GREEK RED"
	case "PRS":
		return "SOUTH AFRICAN RED"
	case "PRT":
		return "CHILEAN RED"
	case "PRU":
		return "ARGENTINE RED - MALBEC"
	case "PRV":
		return "ARGENTINE RED - OTHER"
	case "PRW":
		return "NEW ZEALAND RED"
	case "PRY":
		return "AUSTRALIAN RED"
	case "PRZ":
		return "IMPORTED RED - MISC"
	case "RCP":
		return "FLAVORED MALT BEVERAGES"
	case "TCA":
		return "BEER - GLUTEN FREE/UP TO 20PPM"
	case "TCE":
		return "BEER - WHEAT & HEFEWEIZEN"
	case "TCJ":
		return "BEER - ICE & MALT LIQUOR"
	case "TCP":
		return "BEER - FRUIT & FLAVORED"
	case "THE":
		return "LAGER - PILSENER"
	case "THJ":
		return "LAGER - OKTOBERFEST"
	case "THP":
		return "LAGER - HEAVY & DARK"
	case "THU":
		return "LAGER - MISC & SEASONAL"
	case "TNC":
		return "ALE - PALE"
	case "TNF":
		return "ALE - BITTER & ESB"
	case "TNI":
		return "ALE - WHITE & BELGIAN STYLE"
	case "TNL":
		return "ALE - RED, AMBER, BROWN"
	case "TNP":
		return "ALE - HEAVY, PORTER, STOUT"
	case "TNU":
		return "ALE - MISC & SEASONAL"
	case "TVB":
		return "DOMESTIC BEER"
	case "TVD":
		return "IMPORTED BEER - CANADA"
	case "TVF":
		return "IMPORTED BEER - MEXICO"
	case "TVH":
		return "IMPORTED BEER - GERMANY, AUSTR"
	case "TVJ":
		return "IMPORTED BEER - CZECH REPUBLIC"
	case "TVL":
		return "IMPORTED BEER - ITALY"
	case "TVN":
		return "IMPORTED BEER - BELGIUM & HOLL"
	case "TVP":
		return "IMPORTED BEER - UNITED KINGDOM"
	case "TVR":
		return "IMPORTED BEER - EAST & SOUTH A"
	case "TVT":
		return "IMPORTED BEER - AUSTRALIA & NE"
	case "TVU":
		return "IMPORTED BEER - MISC"
	case "YGA":
		return "GIFT SETS - SPIRITS"
	case "YGC":
		return "GIFT SETS - LIQUEURS"
	case "YGE":
		return "GIFT SETS - WINES"
	}
	return ""
}
