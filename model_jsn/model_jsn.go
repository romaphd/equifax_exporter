package model_jsn

type Model_jsn []struct {
	B1 []struct {
		CustSid     string `json:"cust_sid"`
		FirstName   string `json:"first_name"`
		FatherName  string `json:"father_name"`
		SurnameName string `json:"surname_name"`
	} `json:"b1"`
	B2 []struct {
		CustSid     string `json:"cust_sid"`
		FirstName   string `json:"first_name"`
		FatherName  string `json:"father_name"`
		SurnameName string `json:"surname_name"`
	} `json:"b2"`
	B3 []struct {
		BirthDt            string `json:"birth_dt"`
		BirthPlaceTxt      string `json:"birth_place_txt"`
		BirthCntryOksmCode string `json:"birth_cntry_oksm_code"`
	} `json:"b3"`
	B4 []struct {
		Num                    string `json:"num"`
		PlanDt                 string `json:"plan_dt"`
		SerNum                 string `json:"ser_num"`
		IssueDt                string `json:"issue_dt"`
		IssueTxt               string `json:"issue_txt"`
		IssueCode              string `json:"issue_code"`
		CtlLoading             int    `json:"ctl_loading"`
		DocKindCode            string `json:"doc_kind_code"`
		OtherDocName           string `json:"other_doc_name"`
		CntryOksmDigitalCode   string `json:"cntry_oksm_digital_code"`
		RegAddrrecretCntryName string `json:"reg_addrrecret_cntry_name"`
	} `json:"b4"`
	B5 []struct {
		PrevNum                    string `json:"prev_num"`
		PrevPlanDt                 string `json:"prev_plan_dt"`
		PrevSerNum                 string `json:"prev_ser_num"`
		PrevIssueDt                string `json:"prev_issue_dt"`
		PrevIssueTxt               string `json:"prev_issue_txt"`
		PrevIssueCode              string `json:"prev_issue_code"`
		PrevDocKindCode            string `json:"prev_doc_kind_code"`
		PrevOtherDocName           string `json:"prev_other_doc_name"`
		PrevCntryOksmDigitalCode   string `json:"prev_cntry_oksm_digital_code"`
		PrevRegAddrrecretCntryName string `json:"prev_reg_addrrecret_cntry_name"`
	} `json:"b5"`
	B6 []struct {
		InnNum      string `json:"inn_num"`
		RegNum      string `json:"reg_num"`
		InnTypeCode string `json:"inn_type_code"`
	} `json:"b6"`
	B7 []struct {
		SnilsNum string `json:"snils_num"`
	} `json:"b7"`
	B8 []struct {
		RegDt                    string `json:"reg_dt"`
		RegOrgCode               string `json:"reg_org_code"`
		RegOrgName               string `json:"reg_org_name"`
		AddrRegCode              string `json:"addr_reg_code"`
		RegOwnershipTxt          string `json:"reg_ownership_txt"`
		RegLoctyOkatoCode        string `json:"reg_locty_okato_code"`
		RegOtherLoctyName        string `json:"reg_other_locty_name"`
		RegAddrrecretIdxNum      string `json:"reg_addrrecret_idx_num"`
		CntryOksmDigitalCode     string `json:"cntry_oksm_digital_code"`
		RegAddrrecretBulkNum     string `json:"reg_addrrecret_bulk_num"`
		RegAddrrecretFiasNum     string `json:"reg_addrrecret_fias_num"`
		RegAddrrecretFlatNum     string `json:"reg_addrrecret_flat_num"`
		RegAddrrecretHouseNum    string `json:"reg_addrrecret_house_num"`
		RegAddrrecretCntryName   string `json:"reg_addrrecret_cntry_name"`
		RegAddrrecretStreetName  string `json:"reg_addrrecret_street_name"`
		RegAddrrecretBuildingNum string `json:"reg_addrrecret_building_num"`
	} `json:"b8"`
	B9 []struct {
		FactOwnershipTxt          string `json:"fact_ownership_txt"`
		FactLoctyOkatoCode        string `json:"fact_locty_okato_code"`
		FactOtherLoctyName        string `json:"fact_other_locty_name"`
		CntryOksmDigitalCode      string `json:"cntry_oksm_digital_code"`
		FactAddrrecretIdxNum      string `json:"fact_addrrecret_idx_num"`
		FactAddrrecretBulkNum     string `json:"fact_addrrecret_bulk_num"`
		FactAddrrecretFiasNum     string `json:"fact_addrrecret_fias_num"`
		FactAddrrecretFlatNum     string `json:"fact_addrrecret_flat_num"`
		FactAddrrecretHouseNum    string `json:"fact_addrrecret_house_num"`
		FactAddrrecretCntryName   string `json:"fact_addrrecret_cntry_name"`
		FactAddrrecretStreetName  string `json:"fact_addrrecret_street_name"`
		FactAddrrecretBuildingNum string `json:"fact_addrrecret_building_num"`
	} `json:"b9"`
	Rn  int    `json:"rn"`
	B11 string `json:"b11"`
	B12 string `json:"b12"`
	B18 []struct {
		ReqstNum           int64  `json:"reqst_num"`
		DateBeginDt        string `json:"date_begin_dt"`
		LiabSrcFlag        int    `json:"liab_src_flag"`
		CredTypeCode       string `json:"cred_type_code"`
		DealTypeCode       string `json:"deal_type_code"`
		LiabSubjFlag       int    `json:"liab_subj_flag"`
		RefCtlSrcID        string `json:"ref_ctl_src_id"`
		NewCreditFlag      int    `json:"new_credit_flag"`
		CredProgramSid     string `json:"cred_program_sid"`
		CreditPurpCode     string `json:"credit_purp_code"`
		CreditTypeCode     string `json:"credit_type_code"`
		LiabSubjEndDt      string `json:"liab_subj_end_dt"`
		RefAgrCredSid      string `json:"ref_agr_cred_sid"`
		CredProgramName    string `json:"cred_program_name"`
		PotrebCreditFlag   int    `json:"potreb_credit_flag"`
		UsePaymentCardFlag int    `json:"use_payment_card_flag"`
	} `json:"b18"`
	B19 []struct {
		CrncyCode        string  `json:"crncy_code"`
		LiabCcyAmt       float64 `json:"liab_ccy_amt"`
		CollatLiabCcyAmt float64 `json:"collat_liab_ccy_amt"`
	} `json:"b19"`
	B20 []struct {
		CogrntrCnt      int `json:"cogrntr_cnt"`
		SolidaryDebtCnt int `json:"solidary_debt_cnt"`
	} `json:"b20"`
	B21 string `json:"b21"`
	B22 []struct {
		FullCostCredCalcDt    string  `json:"full_cost_cred_calc_dt"`
		FullCostCredCcyAmt    float64 `json:"full_cost_cred_ccy_amt"`
		FullCostCredIntrstQty float64 `json:"full_cost_cred_intrst_qty"`
	} `json:"b22"`
	B23 string `json:"b23"`
	B24 []struct {
		FinSubIssueDt string `json:"fin_sub_issue_dt"`
	} `json:"b24"`
	B25 []struct {
		AccPayFlag         int     `json:"acc_pay_flag"`
		AccPayCcyAmt       float64 `json:"acc_pay_ccy_amt"`
		NotPrefPeriodFlag  int     `json:"not_pref_period_flag"`
		OtherAccPayCcyAmt  float64 `json:"other_acc_pay_ccy_amt"`
		PrncpAccPayCcyAmt  float64 `json:"prncp_acc_pay_ccy_amt"`
		BeginlAccPayCcyAmt float64 `json:"beginl_acc_pay_ccy_amt"`
		IntrstAccPayCcyAmt float64 `json:"intrst_acc_pay_ccy_amt"`
	} `json:"b25"`
	B26 []struct {
		NonOvrCcyAmt       float64 `json:"non_ovr_ccy_amt"`
		IssueNonOvrDt      string  `json:"issue_non_ovr_dt"`
		NonOvrOtherCcyAmt  float64 `json:"non_ovr_other_ccy_amt"`
		NonOvrPrncpCcyAmt  float64 `json:"non_ovr_prncp_ccy_amt"`
		NonOvrIntrstCcyAmt float64 `json:"non_ovr_intrst_ccy_amt"`
	} `json:"b26"`
	B27 []struct {
		OvrCcyAmt                float64 `json:"ovr_ccy_amt"`
		IssueOvrDt               string  `json:"issue_ovr_dt"`
		OvrOtherCcyAmt           float64 `json:"ovr_other_ccy_amt"`
		OvrPrncpCcyAmt           float64 `json:"ovr_prncp_ccy_amt"`
		OvrIntrstCcyAmt          float64 `json:"ovr_intrst_ccy_amt"`
		LastMissingPymntPrncpDt  string  `json:"last_missing_pymnt_prncp_dt"`
		LastMissingPymntIntrstDt string  `json:"last_missing_pymnt_intrst_dt"`
	} `json:"b27"`
	B28 []struct {
		LastPymntDt           string  `json:"last_pymnt_dt"`
		LengthOvrCnt          int     `json:"length_ovr_cnt"`
		AllPymntCcyAmt        float64 `json:"all_pymnt_ccy_amt"`
		LastPymntCcyAmt       float64 `json:"last_pymnt_ccy_amt"`
		CredPymntAmtCode      string  `json:"cred_pymnt_amt_code"`
		CredPymntTermCode     string  `json:"cred_pymnt_term_code"`
		AllPymntOtherCcyAmt   float64 `json:"all_pymnt_other_ccy_amt"`
		AllPymntPrncpCcyAmt   float64 `json:"all_pymnt_prncp_ccy_amt"`
		AllPymntIntrstCcyAmt  float64 `json:"all_pymnt_intrst_ccy_amt"`
		LastPymntOtherCcyAmt  float64 `json:"last_pymnt_other_ccy_amt"`
		LastPymntPrncpCcyAmt  float64 `json:"last_pymnt_prncp_ccy_amt"`
		LastPymntIntrstCcyAmt float64 `json:"last_pymnt_intrst_ccy_amt"`
	} `json:"b28"`
	B29 []struct {
		AvgMonPymntCcyAmt int    `json:"avg_mon_pymnt_ccy_amt"`
		CalcAvgMonPymntDt string `json:"calc_avg_mon_pymnt_dt"`
	} `json:"b29"`
	B32 string `json:"b32"`
	B33 string `json:"b33"`
	B34 string `json:"b34"`
	B35 string `json:"b35"`
	B36 string `json:"b36"`
	B37 string `json:"b37"`
	B38 string `json:"b38"`
	B39 string `json:"b39"`
	B45 string `json:"b45"`
	B51 string `json:"b51"`
	B52 string `json:"b52"`
	B54 string `json:"b54"`
	B55 []struct {
		ReqstDt             string  `json:"reqst_dt"`
		SrcCode             string  `json:"src_code"`
		ReqstCode           string  `json:"reqst_code"`
		ReqstTypeCode       string  `json:"reqst_type_code"`
		ReqstAprvdEndDt     string  `json:"reqst_aprvd_end_dt"`
		ReqstCredCcyAmt     float64 `json:"reqst_cred_ccy_amt"`
		ReqstAprvdStartDt   string  `json:"reqst_aprvd_start_dt"`
		ReqstCredCrncyCode  string  `json:"reqst_cred_crncy_code"`
		ReqstDeactivationDt string  `json:"reqst_deactivation_dt"`
	} `json:"b55"`
	B56 []struct {
		LiabEndFlag    int    `json:"liab_end_flag"`
		CreditTypeCode string `json:"credit_type_code"`
		Ovr90DaysFlag  int    `json:"ovr_90_days_flag"`
		FinSubjIssueDt string `json:"fin_subj_issue_dt"`
	} `json:"b56"`
	B57 []struct {
		RjctDt         string `json:"rjct_dt"`
		ReqstNum       int64  `json:"reqst_num"`
		ReasonRjctCode string `json:"reason_rjct_code"`
	} `json:"b57"`
	Evt []struct {
		ID           string `json:"id"`
		Rn           int    `json:"rn"`
		EvtDt        string `json:"evt_dt"`
		EntSid       int64  `json:"ent_sid"`
		CustSid      string `json:"cust_sid"`
		EvtCode      string `json:"evt_code"`
		CtlLoading   int    `json:"ctl_loading"`
		EntCtlSrcID  string `json:"ent_ctl_src_id"`
		CustCtlSrcID string `json:"cust_ctl_src_id"`
	} `json:"evt"`
	B10E []struct {
		EmailContTxt string `json:"email_cont_txt"`
	} `json:"b10e"`
	B10P []struct {
		PhoneNumTxt  string `json:"phone_num_txt"`
		PhoneContNum string `json:"phone_cont_num"`
	} `json:"b10p"`
	B551 []struct {
		Rn                  int     `json:"rn"`
		OldID               string  `json:"old_id"`
		ReqstDt             string  `json:"reqst_dt"`
		SrcCode             string  `json:"src_code"`
		ReqstNum            int     `json:"reqst_num"`
		ReqstCode           string  `json:"reqst_code"`
		LiabUniqSid         string  `json:"liab_uniq_sid"`
		CustRoleCode        string  `json:"cust_role_code"`
		ReqstTypeCode       string  `json:"reqst_type_code"`
		ReqstAprvdEndDt     string  `json:"reqst_aprvd_end_dt"`
		ReqstCredCcyAmt     float64 `json:"reqst_cred_ccy_amt"`
		ReqstAprvdStartDt   string  `json:"reqst_aprvd_start_dt"`
		ReqstCredCrncyCode  string  `json:"reqst_cred_crncy_code"`
		ReqstDeactivationDt string  `json:"reqst_deactivation_dt"`
	} `json:"b551"`
	Fltr []struct {
		SendToOkb     int `json:"send_to_okb"`
		SendToNbki    int `json:"send_to_nbki"`
		SendToEquifax int `json:"send_to_equifax"`
	} `json:"fltr"`
	Link []struct {
		AgrApplSid         int64  `json:"agr_appl_sid"`
		LiabUniqSid        string `json:"liab_uniq_sid"`
		CustRoleCode       string `json:"cust_role_code"`
		EntityTypeTxt      string `json:"entity_type_txt"`
		CustPartygroupCode string `json:"cust_partygroup_code"`
	} `json:"link"`
	Oldid []struct {
		OldID      string `json:"old_id"`
		CtlSrcID   string `json:"ctl_src_id"`
		AgrCredSid int64  `json:"agr_cred_sid"`
	} `json:"oldid"`
	EvtCode    string `json:"evt_code"`
	ActionType string `json:"action_type"`
}
