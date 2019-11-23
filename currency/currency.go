package currency

type Currency int

const (
	USD     Currency = 249
	GBP     Currency = 179
	EUR     Currency = 220
	RUB     Currency = 17000
	BRL     Currency = 899
	RMB     Currency = 0
	Unknown Currency = 0
	JPY     Currency = 28800
	NOK     Currency = 2170
	IDR     Currency = 3359900
	MYR     Currency = 1000
	PHP     Currency = 11800
	SGD     Currency = 349
	THB     Currency = 8900
	VND     Currency = 5650000
	KRW     Currency = 289000
	TRY     Currency = 730
	UAH     Currency = 3000
	MXN     Currency = 4300
	CAD     Currency = 325
	AUD     Currency = 325
	NZD     Currency = 369
	PLN     Currency = 0
	CHF     Currency = 255
	AED     Currency = 925
	CLP     Currency = 167000
	CNY     Currency = 1600
	COP     Currency = 765000
	PEN     Currency = 880
	SAR     Currency = 935
	TWD     Currency = 8100
	HKD     Currency = 1930
	ZAR     Currency = 3800
	INR     Currency = 17200
)

var (
	Map = map[Currency]int{
		USD: 249,
		GBP: 179,
		EUR: 220,
		RUB: 17000,
		BRL: 899,
		//RMB:     0,
		Unknown: 0,
		JPY:     28800,
		NOK:     2170,
		IDR:     3359900,
		MYR:     1000,
		PHP:     11800,
		SGD:     349,
		THB:     8900,
		VND:     5650000,
		KRW:     289000,
		TRY:     730,
		UAH:     3000,
		MXN:     4300,
		//CAD:     325,
		//AUD:     325,
		NZD: 369,
		//PLN:     0,
		CHF: 255,
		AED: 925,
		CLP: 167000,
		CNY: 1600,
		COP: 765000,
		PEN: 880,
		SAR: 935,
		TWD: 8100,
		HKD: 1930,
		ZAR: 3800,
		INR: 17200,
	}

	ReverseMap = func() map[int]Currency {
		m := make(map[int]Currency, len(Map))
		for k, v := range Map {
			m[v] = k
		}

		return m
	}()
)
