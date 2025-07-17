package main

// ActionType represents possible actions for grades/news
// Example values: "upgrade", "downgrade", "reiterate"
type ActionType string

const (
	ActionUpgrade   ActionType = "upgrade"
	ActionDowngrade ActionType = "downgrade"
	ActionReiterate ActionType = "reiterate"
)

// GradeType represents possible grades for ratings
// Example values: "Buy", "Hold", "Sell", "Outperform", "Underperform"
type GradeType string

const (
	GradeBuy          GradeType = "Buy"
	GradeHold         GradeType = "Hold"
	GradeSell         GradeType = "Sell"
	GradeOutperform   GradeType = "Outperform"
	GradeUnderperform GradeType = "Underperform"
)

// PeriodType represents reporting periods
// Example values: "Q1", "Q2", "Q3", "Q4", "FY", "annual", "quarter"
type PeriodType string

const (
	PeriodQ1      PeriodType = "Q1"
	PeriodQ2      PeriodType = "Q2"
	PeriodQ3      PeriodType = "Q3"
	PeriodQ4      PeriodType = "Q4"
	PeriodFY      PeriodType = "FY"
	PeriodAnnual  PeriodType = "annual"
	PeriodQuarter PeriodType = "quarter"
)

// ExchangeType represents stock/crypto/forex exchanges
// Example values: "NYSE", "NASDAQ", "AMEX", "OTC", "CME", "BINANCE"
type ExchangeType string

const (
	ExchangeNYSE    ExchangeType = "NYSE"
	ExchangeNASDAQ  ExchangeType = "NASDAQ"
	ExchangeAMEX    ExchangeType = "AMEX"
	ExchangeOTC     ExchangeType = "OTC"
	ExchangeCME     ExchangeType = "CME"
	ExchangeBINANCE ExchangeType = "BINANCE"
)

// SectorType represents company sectors
// Example values: "Technology", "Healthcare", "Financial Services", etc.
type SectorType string

const (
	SectorTechnology         SectorType = "Technology"
	SectorHealthcare         SectorType = "Healthcare"
	SectorFinancialServices  SectorType = "Financial Services"
	SectorConsumerCyclical   SectorType = "Consumer Cyclical"
	SectorConsumerDefensive  SectorType = "Consumer Defensive"
	SectorIndustrials        SectorType = "Industrials"
	SectorUtilities          SectorType = "Utilities"
	SectorBasicMaterials     SectorType = "Basic Materials"
	SectorEnergy             SectorType = "Energy"
	SectorRealEstate         SectorType = "Real Estate"
	SectorCommunication      SectorType = "Communication Services"
)

// CountryType represents countries
// Example values: "US", "CA", "GB", "DE", "JP"
type CountryType string

const (
	CountryUS CountryType = "US"
	CountryCA CountryType = "CA"
	CountryGB CountryType = "GB"
	CountryDE CountryType = "DE"
	CountryJP CountryType = "JP"
)