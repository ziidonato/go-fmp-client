package go_fmp

// Period represents the period type for financial statements
type Period string

const (
	PeriodQ1      Period = "Q1"
	PeriodQ2      Period = "Q2"
	PeriodQ3      Period = "Q3"
	PeriodQ4      Period = "Q4"
	PeriodFY      Period = "FY"
	PeriodAnnual  Period = "annual"
	PeriodQuarter Period = "quarter"
)

// FormType represents the form type for SEC filings
type FormType string

const (
	FormType10K   FormType = "10-K"
	FormType10Q   FormType = "10-Q"
	FormType8K    FormType = "8-K"
	FormType13F   FormType = "13F"
	FormType13G   FormType = "13G"
	FormType13D   FormType = "13D"
	FormType4     FormType = "4"
	FormTypeDEF14A FormType = "DEF 14A"
	FormTypeS1    FormType = "S-1"
	FormType424B4 FormType = "424B4"
)

// TransactionType represents the type of insider trading transaction
type TransactionType string

const (
	TransactionTypeBuy         TransactionType = "P-Purchase"
	TransactionTypeSell        TransactionType = "S-Sale"
	TransactionTypeGift        TransactionType = "G-Gift"
	TransactionTypeOption      TransactionType = "M-Option Exercise"
	TransactionTypeConversion  TransactionType = "C-Conversion"
	TransactionTypeAward       TransactionType = "A-Award"
	TransactionTypeDisposition TransactionType = "D-Disposition"
)

// GradeAction represents the action taken on a stock grade
type GradeAction string

const (
	GradeActionUpgrade   GradeAction = "upgrade"
	GradeActionDowngrade GradeAction = "downgrade"
	GradeActionMaintain  GradeAction = "maintain"
	GradeActionInitiate  GradeAction = "initiate"
	GradeActionReiterate GradeAction = "reiterate"
)

// Timeframe represents the timeframe for technical indicators
type Timeframe string

const (
	Timeframe1Min   Timeframe = "1min"
	Timeframe5Min   Timeframe = "5min"
	Timeframe15Min  Timeframe = "15min"
	Timeframe30Min  Timeframe = "30min"
	Timeframe1Hour  Timeframe = "1hour"
	Timeframe4Hour  Timeframe = "4hour"
	TimeframeDaily  Timeframe = "daily"
	TimeframeWeekly Timeframe = "weekly"
	TimeframeMonthly Timeframe = "monthly"
)

// MarketSituation represents the market situation in COT analysis
type MarketSituation string

const (
	MarketSituationBullish MarketSituation = "bullish"
	MarketSituationBearish MarketSituation = "bearish"
	MarketSituationNeutral MarketSituation = "neutral"
)

// MarketSentiment represents the market sentiment
type MarketSentiment string

const (
	MarketSentimentBullish     MarketSentiment = "bullish"
	MarketSentimentBearish     MarketSentiment = "bearish"
	MarketSentimentNeutral     MarketSentiment = "neutral"
	MarketSentimentMixed       MarketSentiment = "mixed"
	MarketSentimentUncertain   MarketSentiment = "uncertain"
)

// GradeConsensus represents the consensus rating
type GradeConsensus string

const (
	GradeConsensusStrongBuy  GradeConsensus = "Strong Buy"
	GradeConsensusBuy        GradeConsensus = "Buy"
	GradeConsensusHold       GradeConsensus = "Hold"
	GradeConsensusSell       GradeConsensus = "Sell"
	GradeConsensusStrongSell GradeConsensus = "Strong Sell"
)

// Exchange represents stock exchanges
type Exchange string

const (
	ExchangeNASDAQ Exchange = "NASDAQ"
	ExchangeNYSE   Exchange = "NYSE"
	ExchangeAMEX   Exchange = "AMEX"
	ExchangeLSE    Exchange = "LSE"
	ExchangeTSX    Exchange = "TSX"
	ExchangeEURONEXT Exchange = "EURONEXT"
	ExchangeJPX    Exchange = "JPX"
	ExchangeHKEX   Exchange = "HKEX"
	ExchangeSSE    Exchange = "SSE"
	ExchangeSZSE   Exchange = "SZSE"
	ExchangeASX    Exchange = "ASX"
	ExchangeNSE    Exchange = "NSE"
	ExchangeBSE    Exchange = "BSE"
	ExchangeKRX    Exchange = "KRX"
)

// Currency represents currency codes
type Currency string

const (
	CurrencyUSD Currency = "USD"
	CurrencyEUR Currency = "EUR"
	CurrencyGBP Currency = "GBP"
	CurrencyJPY Currency = "JPY"
	CurrencyCNY Currency = "CNY"
	CurrencyCAD Currency = "CAD"
	CurrencyAUD Currency = "AUD"
	CurrencyCHF Currency = "CHF"
	CurrencyHKD Currency = "HKD"
	CurrencySGD Currency = "SGD"
	CurrencyINR Currency = "INR"
	CurrencyKRW Currency = "KRW"
)

// Country represents country codes
type Country string

const (
	CountryUS Country = "US"
	CountryCA Country = "CA"
	CountryGB Country = "GB"
	CountryDE Country = "DE"
	CountryFR Country = "FR"
	CountryJP Country = "JP"
	CountryCN Country = "CN"
	CountryAU Country = "AU"
	CountryIN Country = "IN"
	CountryKR Country = "KR"
	CountryBR Country = "BR"
	CountryMX Country = "MX"
)

// Sector represents business sectors
type Sector string

const (
	SectorTechnology          Sector = "Technology"
	SectorHealthcare          Sector = "Healthcare"
	SectorFinancials          Sector = "Financials"
	SectorConsumerDiscretionary Sector = "Consumer Discretionary"
	SectorConsumerStaples     Sector = "Consumer Staples"
	SectorIndustrials         Sector = "Industrials"
	SectorEnergy              Sector = "Energy"
	SectorMaterials           Sector = "Materials"
	SectorRealEstate          Sector = "Real Estate"
	SectorUtilities           Sector = "Utilities"
	SectorCommunicationServices Sector = "Communication Services"
)

// AssetClass represents the asset class for ETFs
type AssetClass string

const (
	AssetClassEquity      AssetClass = "Equity"
	AssetClassFixedIncome AssetClass = "Fixed Income"
	AssetClassCommodity   AssetClass = "Commodity"
	AssetClassCurrency    AssetClass = "Currency"
	AssetClassAlternative AssetClass = "Alternative"
	AssetClassMixed       AssetClass = "Mixed"
)

// EconomicImpact represents the impact level of economic events
type EconomicImpact string

const (
	EconomicImpactHigh   EconomicImpact = "High"
	EconomicImpactMedium EconomicImpact = "Medium"
	EconomicImpactLow    EconomicImpact = "Low"
)

// Units represents measurement units
type Units string

const (
	UnitsPercent Units = "percent"
	UnitsDollars Units = "dollars"
	UnitsShares  Units = "shares"
	UnitsRatio   Units = "ratio"
	UnitsIndex   Units = "index"
	UnitsPoints  Units = "points"
)

// SecurityType represents the type of security
type SecurityType string

const (
	SecurityTypeCommonStock    SecurityType = "Common Stock"
	SecurityTypePreferredStock SecurityType = "Preferred Stock"
	SecurityTypeBond           SecurityType = "Bond"
	SecurityTypeOption         SecurityType = "Option"
	SecurityTypeWarrant        SecurityType = "Warrant"
	SecurityTypeETF            SecurityType = "ETF"
	SecurityTypeMutualFund     SecurityType = "Mutual Fund"
)

// ReportedCurrency represents the currency used in financial reports
type ReportedCurrency string

const (
	ReportedCurrencyUSD ReportedCurrency = "USD"
	ReportedCurrencyEUR ReportedCurrency = "EUR"
	ReportedCurrencyGBP ReportedCurrency = "GBP"
	ReportedCurrencyJPY ReportedCurrency = "JPY"
	ReportedCurrencyCNY ReportedCurrency = "CNY"
	ReportedCurrencyCAD ReportedCurrency = "CAD"
	ReportedCurrencyAUD ReportedCurrency = "AUD"
)

// ESGRiskRating represents ESG risk rating levels
type ESGRiskRating string

const (
	ESGRiskRatingNegligible ESGRiskRating = "Negligible"
	ESGRiskRatingLow        ESGRiskRating = "Low"
	ESGRiskRatingMedium     ESGRiskRating = "Medium"
	ESGRiskRatingHigh       ESGRiskRating = "High"
	ESGRiskRatingSevere     ESGRiskRating = "Severe"
)

// IndustryRank represents industry ranking categories
type IndustryRank string

const (
	IndustryRankTop10    IndustryRank = "Top 10%"
	IndustryRankTop25    IndustryRank = "Top 25%"
	IndustryRankTop50    IndustryRank = "Top 50%"
	IndustryRankBottom50 IndustryRank = "Bottom 50%"
	IndustryRankBottom25 IndustryRank = "Bottom 25%"
)

// PayoffProfile represents the payoff profile for fund disclosures
type PayoffProfile string

const (
	PayoffProfileLong  PayoffProfile = "Long"
	PayoffProfileShort PayoffProfile = "Short"
)

// AssetCategory represents asset categories for fund disclosures
type AssetCategory string

const (
	AssetCategoryEquity      AssetCategory = "Equity"
	AssetCategoryDebt        AssetCategory = "Debt"
	AssetCategoryDerivative  AssetCategory = "Derivative"
	AssetCategoryCommodity   AssetCategory = "Commodity"
	AssetCategoryCash        AssetCategory = "Cash"
)

// FairValueLevel represents fair value hierarchy levels
type FairValueLevel string

const (
	FairValueLevel1 FairValueLevel = "Level 1"
	FairValueLevel2 FairValueLevel = "Level 2"
	FairValueLevel3 FairValueLevel = "Level 3"
)