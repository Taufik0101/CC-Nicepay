package dto

type CreateTransactionNicePay struct {
	Amount     int    `json:"amount"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostNumber int    `json:"postNumber"`
	Country    string `json:"country"`
	CartData   struct {
		Count int `json:"count"`
		Item  []struct {
			ImgURL      string `json:"img_url"`
			GoodsName   string `json:"goods_name"`
			GoodsDetail string `json:"goods_detail"`
			GoodsAmt    int    `json:"goods_amt"`
		} `json:"item"`
	} `json:"cartData"`
}

type ResponseRegistrationCCNicePay struct {
	Timestamp   string `json:"timeStamp"`
	ResultCD    string `json:"resultCd"`
	ResultMSG   string `json:"resultMsg"`
	TxId        string `json:"tXid"`
	ReferenceNo string `json:"referenceNo"`
	PayMethod   string `json:"payMethod"`
	AMT         string `json:"amt"`
	TransDt     string `json:"transDt"`
	TransTm     string `json:"transTm"`
}

type RequestStatusInquiry struct {
	Timestamp   string `json:"timeStamp"`
	TxId        string `json:"tXid"`
	ReferenceNo string `json:"referenceNo"`
	AMT         string `json:"amt"`
}

type ResponseStatusInquiry struct {
	Currency    string `json:"currency"`
	ResultCD    string `json:"resultCd"`    // check
	ResultMSG   string `json:"resultMsg"`   // check
	TxId        string `json:"tXid"`        //check
	ReferenceNo string `json:"referenceNo"` //check
	PayMethod   string `json:"payMethod"`   //check
	AMT         string `json:"amt"`         //check
	InstmntMon  string `json:"instmntMon"`
	InstmntType string `json:"instmntType"`
	GoodsNm     string `json:"goodsNm"`
	BillingNm   string `json:"billingNm"`
	ReqDT       string `json:"reqDt"`
	ReqTm       string `json:"reqTm"`
	Status      string `json:"status"`
}

type RequestPaymentCCNicePay struct {
	Timestamp    string `json:"timeStamp"`
	TxId         string `json:"tXid"`
	ReferenceNo  string `json:"referenceNo"`
	AMT          string `json:"amt"`
	CardNo       string `json:"cardNo"`
	CardExpYyMm  string `json:"cardExpYymm"`
	CardCVV      string `json:"cardCvv"`
	CardHolderNm string `json:"cardHolderNm"`
	CallBackURL  string `json:"callBackUrl"`
}
