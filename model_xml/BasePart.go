package model_xml

import (
	"encoding/xml"
)

type BasePart struct {
	XMLName    xml.Name `xml:"base_part,omitempty"`
	AddrReg    *AddrReg
	AddrFact   *AddrFact
	Contacts   *Contacts
	Ogrnip     *Ogrnip
	Incapacity *Incapacity
	Contract   *Contract
}
type AddrReg struct {
	XMLName           xml.Name `xml:"addr_reg,omitempty"`
	RegCode           string   `xml:"reg_code,omitempty"`
	Index             string   `xml:"index,omitempty"`
	Country           string   `xml:"country,omitempty"`
	CountryText       string   `xml:"country_text,omitempty"`
	Fias              string   `xml:"fias,omitempty"`
	Okato             string   `xml:"okato,omitempty"`
	OtherStatement    string   `xml:"other_statement,omitempty"`
	Street            string   `xml:"street,omitempty"`
	House             string   `xml:"house,omitempty"`
	Domain            string   `xml:"domain,omitempty"`
	Block             string   `xml:"block,omitempty"`
	Build             string   `xml:"build,omitempty"`
	Apartment         string   `xml:"apartment,omitempty"`
	RegDate           string   `xml:"reg_date,omitempty"`
	RegPlace          string   `xml:"reg_place,omitempty"`
	RegDepartmentCode string   `xml:"reg_department_code,omitempty"`
}
type AddrFact struct {
	XMLName        xml.Name `xml:"addr_fact,omitempty"`
	Index          string   `xml:"index,omitempty"`
	Country        string   `xml:"country,omitempty"`
	CountryText    string   `xml:"country_text,omitempty"`
	Fias           string   `xml:"fias,omitempty"`
	Okato          string   `xml:"okato,omitempty"`
	OtherStatement string   `xml:"other_statement,omitempty"`
	Street         string   `xml:"street,omitempty"`
	House          string   `xml:"house,omitempty"`
	Domain         string   `xml:"domain,omitempty"`
	Block          string   `xml:"block,omitempty"`
	Build          string   `xml:"build,omitempty"`
	Apartment      string   `xml:"apartment,omitempty"`
}
type Contacts struct {
	XMLName xml.Name `xml:"contacts,omitempty"`
	Phone   struct {
		XMLName xml.Name `xml:"phone,omitempty"`
		Number  string   `xml:"number,omitempty"`
		Comment string   `xml:"comment,omitempty"`
	}
	Email string `xml:"email,omitempty"`
}
type Ogrnip struct {
	XMLName xml.Name `xml:"ogrnip,omitempty"`
	No      string   `xml:"no,omitempty"`
	Date    string   `xml:"date,omitempty"`
}
type Incapacity struct {
	XMLName           xml.Name `xml:"incapacity,omitempty"`
	Code              string   `xml:"code,omitempty"`
	CourtDecisionDate string   `xml:"court_decision_date,omitempty"`
	CourtDecisionNo   string   `xml:"court_decision_no,omitempty"`
	CourtName         string   `xml:"court_name,omitempty"`
}

type Contract struct {
	XMLName             xml.Name `xml:"contract,omitempty"`
	Uid                 *Uid
	Deal                *Deal
	ContractAmount      *ContractAmount
	JointDebtors        *JointDebtors
	PaymentTerms        *PaymentTerms
	FullCost            *FullCost
	ContractChanges     []ContractChanges
	CredStartDebt       *CredStartDebt
	Debt                *Debt
	DebtCurrent         *DebtCurrent
	DebtOverdue         *DebtOverdue
	Payments            *Payments
	AveragePayment      *AveragePayment
	Collaterals         []Collaterals
	Guarantees          []Guarantees
	IndieGuarantees     []IndieGuarantees
	CollateralInsce     []CollateralInsce
	RepaymentCollateral []RepaymentCollateral
	GuaranteeReturn     []GuaranteeReturn
	ContractEnd         *ContractEnd
	Court               []Court
	StopLoad            *StopLoad
}

type Uid struct {
	XMLName xml.Name `xml:"uid,omitempty"`
	ID      string   `xml:"id,omitempty"`
}
type Deal struct {
	XMLName             xml.Name `xml:"deal,omitempty"`
	Ratio               string   `xml:"ratio,omitempty"`
	Date                string   `xml:"date,omitempty"`
	Category            string   `xml:"category,omitempty"`
	Type                string   `xml:"type,omitempty"`
	Purpose             string   `xml:"purpose,omitempty"`
	SignCredit          string   `xml:"sign_credit,omitempty"`
	SignCreditCard      string   `xml:"sign_credit_card,omitempty"`
	SignRenovation      string   `xml:"sign_renovation,omitempty"`
	SignDealCashSource  string   `xml:"sign_deal_cash_source,omitempty"`
	SignDealCashSubject string   `xml:"sign_deal_cash_subject,omitempty"`
	EndDate             string   `xml:"end_date,omitempty"`
}
type ContractAmount struct {
	XMLName     xml.Name `xml:"contract_amount,omitempty"`
	Sum         string   `xml:"sum,omitempty"`
	Currency    string   `xml:"currency,omitempty"`
	SecuritySum string   `xml:"security_sum,omitempty"`
}
type JointDebtors struct {
	XMLName xml.Name `xml:"joint_debtors,omitempty"`
	Count   string   `xml:"count,omitempty"`
}
type PaymentTerms struct {
	XMLName               xml.Name `xml:"payment_terms,omitempty"`
	OpNextPayoutSum       string   `xml:"op_next_payout_sum,omitempty"`
	OpNextPayoutDate      string   `xml:"op_next_payout_date,omitempty"`
	PercentNextPayoutSum  string   `xml:"percent_next_payout_sum,omitempty"`
	PercentNextPayoutDate string   `xml:"percent_next_payout_date,omitempty"`
	Regularity            string   `xml:"regularity,omitempty"`
	MinSumPayCc           string   `xml:"min_sum_pay_cc,omitempty"`
	GraceDate             string   `xml:"grace_date,omitempty"`
	GraceEndDate          string   `xml:"grace_end_date,omitempty"`
	PercentEndDate        string   `xml:"percent_end_date,omitempty"`
}
type FullCost struct {
	XMLName xml.Name `xml:"full_cost,omitempty"`
	Percent string   `xml:"percent,omitempty"`
	Sum     string   `xml:"sum,omitempty"`
	Date    string   `xml:"date,omitempty"`
}
type ContractChanges struct {
	XMLName       xml.Name `xml:"contract_changes,omitempty"`
	Date          string   `xml:"date,omitempty"`
	Type          string   `xml:"type,omitempty"`
	SpecialType   string   `xml:"special_type,omitempty"`
	TypeText      string   `xml:"type_text,omitempty"`
	ApplyDate     string   `xml:"apply_date,omitempty"`
	EndDate       string   `xml:"end_date,omitempty"`
	FinishDate    string   `xml:"finish_date,omitempty"`
	Finish        string   `xml:"finish,omitempty"`
	CurrencyPrice string   `xml:"currency_price,omitempty"`
}
type CredStartDebt struct {
	XMLName xml.Name `xml:"cred_start_debt,omitempty"`
	Date    string   `xml:"date,omitempty"`
}
type Debt struct {
	XMLName                   xml.Name `xml:"basedebt_part,omitempty"`
	SignCalcLastPayout        string   `xml:"sign_calc_last_payout,omitempty"`
	CalcDate                  string   `xml:"calc_date,omitempty"`
	FirstSum                  string   `xml:"first_sum,omitempty"`
	Sum                       string   `xml:"sum,omitempty"`
	OpSum                     string   `xml:"op_sum,omitempty"`
	PercentSum                string   `xml:"percent_sum,omitempty"`
	OtherSum                  string   `xml:"other_sum,omitempty"`
	SignUnacceptedGracePeriod string   `xml:"sign_unaccepted_grace_period,omitempty"`
}
type DebtCurrent struct {
	XMLName    xml.Name `xml:"debt_current,omitempty"`
	Date       string   `xml:"date,omitempty"`
	Sum        string   `xml:"sum,omitempty"`
	OpSum      string   `xml:"op_sum,omitempty"`
	PercentSum string   `xml:"percent_sum,omitempty"`
	OtherSum   string   `xml:"other_sum,omitempty"`
}
type DebtOverdue struct {
	XMLName               xml.Name `xml:"debt_overdue,omitempty"`
	Date                  string   `xml:"date,omitempty"`
	Sum                   string   `xml:"sum,omitempty"`
	OpSum                 string   `xml:"op_sum,omitempty"`
	PercentSum            string   `xml:"percent_sum,omitempty"`
	OtherSum              string   `xml:"other_sum,omitempty"`
	OpMissPayoutDate      string   `xml:"op_miss_payout_date,omitempty"`
	PercentMissPayoutDate string   `xml:"percent_miss_payout_date,omitempty"`
}
type Payments struct {
	XMLName              xml.Name `xml:"payments,omitempty"`
	LastPayoutDate       string   `xml:"last_payout_date,omitempty"`
	LastPayoutSum        string   `xml:"last_payout_sum,omitempty"`
	LastPayoutOpSum      string   `xml:"last_payout_op_sum,omitempty"`
	LastPayoutPercentSum string   `xml:"last_payout_percent_sum,omitempty"`
	LastPayoutOtherSum   string   `xml:"last_payout_other_sum,omitempty"`
	PaidSum              string   `xml:"paid_sum,omitempty"`
	PaidOpSum            string   `xml:"paid_op_sum,omitempty"`
	PaidPercentSum       string   `xml:"paid_percent_sum,omitempty"`
	PaidOtherSum         string   `xml:"paid_other_sum,omitempty"`
	SizePaymentsType     string   `xml:"size_payments_type,omitempty"`
	PaymentsDeadlineType string   `xml:"payments_deadline_type,omitempty"`
	OverdueDay           string   `xml:"overdue_day,omitempty"`
}
type AveragePayment struct {
	XMLName xml.Name `xml:"average_payment,omitempty"`
	Sum     string   `xml:"sum,omitempty"`
	Date    string   `xml:"date,omitempty"`
}
type Collaterals struct {
	XMLName    xml.Name `xml:"collaterals,omitempty"`
	Collateral struct {
		XMLName        xml.Name `xml:"collateral,omitempty"`
		ItemType       string   `xml:"item_type,omitempty"`
		ID             string   `xml:"id,omitempty"`
		Date           string   `xml:"date,omitempty"`
		Sum            string   `xml:"sum,omitempty"`
		Currency       string   `xml:"currency,omitempty"`
		DateAssessment string   `xml:"date_assessment,omitempty"`
		ItemBurden     string   `xml:"item_burden,omitempty"`
		EndDate        string   `xml:"end_date,omitempty"`
		FactEndDate    string   `xml:"fact_end_date,omitempty"`
	}
}
type Guarantees struct {
	XMLName   xml.Name `xml:"guarantees,omitempty"`
	Guarantee struct {
		XMLName     xml.Name `xml:"guarantee,omitempty"`
		Uid         string   `xml:"uid,omitempty"`
		Sum         string   `xml:"sum,omitempty"`
		Currency    string   `xml:"currency,omitempty"`
		Date        string   `xml:"date,omitempty"`
		EndDate     string   `xml:"end_date,omitempty"`
		FactEndDate string   `xml:"fact_end_date,omitempty"`
		EndReason   string   `xml:"end_reason,omitempty"`
	}
}
type IndieGuarantees struct {
	XMLName        xml.Name `xml:"indie_guarantees,omitempty"`
	IndieGuarantee struct {
		XMLName     xml.Name `xml:"indie_guarantee,omitempty"`
		Uid         string   `xml:"uid,omitempty"`
		Sum         string   `xml:"sum,omitempty"`
		Currency    string   `xml:"currency,omitempty"`
		Date        string   `xml:"date,omitempty"`
		EndDate     string   `xml:"end_date,omitempty"`
		FactEndDate string   `xml:"fact_end_date,omitempty"`
		EndReason   string   `xml:"end_reason,omitempty"`
	}
}
type CollateralInsce struct {
	XMLName     xml.Name `xml:"collateral_insce,omitempty"`
	Limit       string   `xml:"limit,omitempty"`
	Currency    string   `xml:"currency,omitempty"`
	Franchise   string   `xml:"franchise,omitempty"`
	Date        string   `xml:"date,omitempty"`
	EndDate     string   `xml:"end_date,omitempty"`
	FactEndDate string   `xml:"fact_end_date,omitempty"`
	EndReason   string   `xml:"end_reason,omitempty"`
}
type RepaymentCollateral struct {
	XMLName xml.Name `xml:"repayment_collateral,omitempty"`
	Code    string   `xml:"code,omitempty"`
	Date    string   `xml:"date,omitempty"`
	Sum     string   `xml:"sum,omitempty"`
}
type GuaranteeReturn struct {
	XMLName       xml.Name `xml:"guarantee_return,omitempty"`
	ReturningSum  string   `xml:"returning_sum,omitempty"`
	PrincipalSum  string   `xml:"principal_sum,omitempty"`
	ReturnCorrect string   `xml:"return_correct,omitempty"`
}
type ContractEnd struct {
	XMLName xml.Name `xml:"contract_end,omitempty"`
	Reason  string   `xml:"reason,omitempty"`
	Date    string   `xml:"date,omitempty"`
}
type Court []struct {
	XMLName  xml.Name `xml:"court,omitempty"`
	CourtAct struct {
		XMLName    xml.Name `xml:"court_act,omitempty"`
		Date       string   `xml:"date,omitempty"`
		No         string   `xml:"no,omitempty"`
		Resolution string   `xml:"resolution,omitempty"`
		Accepted   string   `xml:"accepted,omitempty"`
	}
}

type StopLoad struct {
	XMLName xml.Name `xml:"stop_load,omitempty"`
	Code    string   `xml:"code,omitempty"`
	Date    string   `xml:"date,omitempty"`
}
