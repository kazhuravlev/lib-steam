package meta

// Lang внутренний идентфиикатор языка
type Lang uint

const (
	// LangEnglish один из языков. Далее идет перечисление
	LangEnglish Lang = iota
	LangBulgarian
	LangCzech
	LangDanish
	LangDutch
	LangFinnish
	LangFrench
	LangGerman
	LangGreek
	LangHungarian
	LangItalian
	LangJapanese
	LangKoreana
	LangNorwegian
	LangPolish
	LangPortuguese
	LangBrazilian
	LangRomanian
	LangRussian
	LangSchinese
	LangSpanish
	LangSwedish
	LangTchinese
	LangThai
	LangTurkish
	LangUkrainian
)

var (
	// Langs маппинг языков на строки
	Langs = map[Lang]string{
		LangEnglish:    "english",
		LangBulgarian:  "bulgarian",
		LangCzech:      "czech",
		LangDanish:     "danish",
		LangDutch:      "dutch",
		LangFinnish:    "finnish",
		LangFrench:     "french",
		LangGerman:     "german",
		LangGreek:      "greek",
		LangHungarian:  "hungarian",
		LangItalian:    "italian",
		LangJapanese:   "japanese",
		LangKoreana:    "koreana",
		LangNorwegian:  "norwegian",
		LangPolish:     "polish",
		LangPortuguese: "portuguese",
		LangBrazilian:  "brazilian",
		LangRomanian:   "romanian",
		LangRussian:    "russian",
		LangSchinese:   "schinese",
		LangSpanish:    "spanish",
		LangSwedish:    "swedish",
		LangTchinese:   "tchinese",
		LangThai:       "thai",
		LangTurkish:    "turkish",
		LangUkrainian:  "ukrainian",
	}
	// ReverseLangs обратный маппинг
	ReverseLangs = map[string]Lang{
		"english":    LangEnglish,
		"bulgarian":  LangBulgarian,
		"czech":      LangCzech,
		"danish":     LangDanish,
		"dutch":      LangDutch,
		"finnish":    LangFinnish,
		"french":     LangFrench,
		"german":     LangGerman,
		"greek":      LangGreek,
		"hungarian":  LangHungarian,
		"italian":    LangItalian,
		"japanese":   LangJapanese,
		"koreana":    LangKoreana,
		"norwegian":  LangNorwegian,
		"polish":     LangPolish,
		"portuguese": LangPortuguese,
		"brazilian":  LangBrazilian,
		"romanian":   LangRomanian,
		"russian":    LangRussian,
		"schinese":   LangSchinese,
		"spanish":    LangSpanish,
		"swedish":    LangSwedish,
		"tchinese":   LangTchinese,
		"thai":       LangThai,
		"turkish":    LangTurkish,
		"ukrainian":  LangUkrainian,
	}
)
