package go_fmp

// Period represents financial reporting periods
type Period string

const (
	PeriodAnnual   Period = "annual"
	PeriodQuarter  Period = "quarter"
	PeriodQ1       Period = "Q1"
	PeriodQ2       Period = "Q2"
	PeriodQ3       Period = "Q3"
	PeriodQ4       Period = "Q4"
	PeriodFY       Period = "FY"
)

// TransactionType represents insider trading transaction types
type TransactionType string

const (
	TransactionTypeBuy          TransactionType = "P-Purchase"
	TransactionTypeSale         TransactionType = "S-Sale"
	TransactionTypeGift         TransactionType = "G-Gift"
	TransactionTypeExercise     TransactionType = "M-Exercise"
	TransactionTypeConversion   TransactionType = "C-Conversion"
)

// FormType represents SEC form types
type FormType string

const (
	FormType10K    FormType = "10-K"
	FormType10Q    FormType = "10-Q"
	FormType8K     FormType = "8-K"
	FormType13F    FormType = "13F"
	FormType4      FormType = "4"
	FormTypeDEF14A FormType = "DEF 14A"
)

// SharesType represents types of shares
type SharesType string

const (
	SharesTypeCommon    SharesType = "COM"
	SharesTypeCall      SharesType = "CALL"
	SharesTypePut       SharesType = "PUT"
	SharesTypePreferred SharesType = "PRF"
)

// AssetClass represents asset classes
type AssetClass string

const (
	AssetClassEquity      AssetClass = "Equity"
	AssetClassFixedIncome AssetClass = "Fixed Income"
	AssetClassCommodity   AssetClass = "Commodity"
	AssetClassAlternative AssetClass = "Alternative"
	AssetClassCurrency    AssetClass = "Currency"
)

// Exchange represents stock exchanges
type Exchange string

const (
	ExchangeNASDAQ Exchange = "NASDAQ"
	ExchangeNYSE   Exchange = "NYSE"
	ExchangeAMEX   Exchange = "AMEX"
	ExchangeOTC    Exchange = "OTC"
	ExchangeLSE    Exchange = "LSE"
	ExchangeTSE    Exchange = "TSE"
)

// Interval represents time intervals for charts
type Interval string

const (
	Interval1Min   Interval = "1min"
	Interval5Min   Interval = "5min"
	Interval15Min  Interval = "15min"
	Interval30Min  Interval = "30min"
	Interval1Hour  Interval = "1hour"
	Interval4Hour  Interval = "4hour"
	IntervalDaily  Interval = "daily"
)

// Rating represents stock ratings
type Rating string

const (
	RatingStrongBuy  Rating = "Strong Buy"
	RatingBuy        Rating = "Buy"
	RatingHold       Rating = "Hold"
	RatingSell       Rating = "Sell"
	RatingStrongSell Rating = "Strong Sell"
)

// AcquiredOrDisposed represents whether shares were acquired or disposed
type AcquiredOrDisposed string

const (
	Acquired AcquiredOrDisposed = "A"
	Disposed AcquiredOrDisposed = "D"
)

// TypeOfOwner represents types of insider owners
type TypeOfOwner string

const (
	TypeOfOwnerDirector        TypeOfOwner = "director"
	TypeOfOwnerOfficer         TypeOfOwner = "officer"
	TypeOfOwnerTenPercentOwner TypeOfOwner = "10% owner"
	TypeOfOwnerOther          TypeOfOwner = "other"
)