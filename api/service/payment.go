package service

import (
	"CC-Nicepay/api/dbtemporary"
	"CC-Nicepay/api/dto"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	numberSet = "0123456789"
)

// ReferenceNo --> ExternalID in Xendit
// MerchantToken --> SHA256( timestamp + IONPAYTEST + ReferenceNo + amount +

func RandomReferenceNo() string {
	transactionID := "ORD"
	for i := 0; i < 4; i++ {
		random := rand.Intn(len(numberSet))
		transactionID = transactionID + string(numberSet[random])
	}

	return transactionID
}
func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}
func RegistrationCCNicePay(input dto.CreateTransactionNicePay) (*dto.ResponseRegistrationCCNicePay, error) {
	var nicePayMerchantID = "IONPAYTEST"
	var nicePayMerchantKey = "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A=="
	var startDB = dbtemporary.DB
	URL, _ := url.Parse("https://dev.nicepay.co.id/nicepay/direct/v2/registration")
	//URL, _ := url.Parse("https://staging.nicepay.co.id/nicepay/direct/v2/registration")

	ReferenceNo := RandomReferenceNo()

	cartJSON, _ := json.Marshal(input.CartData)
	amount := strconv.Itoa(input.Amount)
	timeNow := time.Now()
	timestamp := timeNow.Format("20060102150405")
	timeDate := timeNow.Format(time.RFC3339)
	timeTomorrow := timeNow.AddDate(0, 0, 1).Format("20060102")

	stringMerchantToken := timestamp + nicePayMerchantID + ReferenceNo + amount + nicePayMerchantKey

	hash := NewSHA256([]byte(stringMerchantToken))
	signature := hex.EncodeToString(hash)

	postBody, _ := json.Marshal(map[string]interface{}{
		"timeStamp":      timestamp,
		"iMid":           nicePayMerchantID,
		"payMethod":      "01",
		"currency":       "IDR",
		"amt":            amount,
		"referenceNo":    ReferenceNo,
		"goodsNm":        "Tes CC",
		"billingNm":      input.Name,
		"billingPhone":   input.Phone,
		"billingEmail":   input.Email,
		"billingAddr":    input.Address,
		"billingCity":    input.City,
		"billingState":   input.State,
		"cartData":       cartJSON,
		"merFixAcctId":   4,
		"billingPostCd":  strconv.Itoa(input.PostNumber),
		"billingCountry": "Indonesia",
		"dbProcessUrl":   "https://merchant.com/api/dbProcessUrl/Notif",
		"merchantToken":  signature,
		"userIP":         "127.0.0.1",
		"userAgent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML,like Gecko) Chrome/60.0.3112.101 Safari/537.36",
		"instmntType":    "2",
		"instmntMon":     "1",
		"recurrOpt":      "2",
		"vacctValidDt":   timeTomorrow,
		"vacctValidTm":   "235959",
	})

	reqBody := ioutil.NopCloser(strings.NewReader(string(postBody)))

	req := &http.Request{
		Method: "POST",
		URL:    URL,
		Body:   reqBody,
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	var dataForCallbacks dto.ResponseRegistrationCCNicePay
	err = json.NewDecoder(resp.Body).Decode(&dataForCallbacks)
	if err != nil {
		return nil, err
	}

	dataForCallbacks.Timestamp = timestamp
	if dataForCallbacks.ResultCD != "0000" {
		return nil, errors.New(dataForCallbacks.ResultMSG)
	}

	startDB = append(startDB, dataForCallbacks)

	res, _ := json.Marshal(dataForCallbacks)

	if _, err := os.Stat("./log.json"); err == nil {
		f, err := os.OpenFile("./log.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		_, _ = f.WriteString("\n" + `"` + timeDate + `"` + "\n")

		_, _ = f.WriteString(`"Request Registration CC Nice Pay"` + "\n")

		_, _ = f.Write(postBody)

		_, _ = f.WriteString("\n")

		_, _ = f.WriteString(`"Response Registration CC Nice Pay"` + "\n")

		_, _ = f.Write(res)

		_, _ = f.WriteString("\n\n")
	} else {
		f, _ := os.Create("./log.json")

		defer f.Close()

		_, _ = f.WriteString(`"` + timeDate + `"` + "\n")

		_, _ = f.WriteString(`"Request Registration CC Nice Pay"` + "\n")

		_, _ = f.Write(postBody)

		_, _ = f.WriteString("\n")

		_, _ = f.WriteString(`"Response Registration CC Nice Pay"` + "\n")

		_, _ = f.Write(res)

		_, _ = f.WriteString("\n\n")
	}

	return &dataForCallbacks, nil
}

func CheckStatusInquiry(input dto.RequestStatusInquiry) (*dto.ResponseStatusInquiry, error) {
	var nicePayMerchantID = "IONPAYTEST"
	var nicePayMerchantKey = "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A=="
	URL, _ := url.Parse("https://dev.nicepay.co.id/nicepay/direct/v2/inquiry")

	ReferenceNo := input.ReferenceNo
	amount := input.AMT
	timestamp := input.Timestamp
	TxID := input.TxId

	timeDate := time.Now().Format(time.RFC3339)

	stringMerchantToken := timestamp + nicePayMerchantID + ReferenceNo + amount + nicePayMerchantKey

	hash := NewSHA256([]byte(stringMerchantToken))
	signature := hex.EncodeToString(hash)

	postBody, _ := json.Marshal(map[string]interface{}{
		"timeStamp":     timestamp,
		"iMid":          nicePayMerchantID,
		"tXid":          TxID,
		"referenceNo":   ReferenceNo,
		"amt":           amount,
		"merchantToken": signature,
	})

	reqBody := ioutil.NopCloser(strings.NewReader(string(postBody)))

	req := &http.Request{
		Method: "POST",
		URL:    URL,
		Body:   reqBody,
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	var dataForCallbacks dto.ResponseStatusInquiry
	err = json.NewDecoder(resp.Body).Decode(&dataForCallbacks)
	if err != nil {
		return nil, err
	}

	if dataForCallbacks.ResultCD != "0000" {
		return nil, errors.New(dataForCallbacks.ResultMSG)
	}

	res, _ := json.Marshal(dataForCallbacks)

	if _, err := os.Stat("./log.json"); err == nil {
		f, err := os.OpenFile("./log.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		_, _ = f.WriteString("\n" + `"` + timeDate + `"` + "\n")

		_, _ = f.WriteString(`"Request Inquiry CC Nice Pay"` + "\n")

		_, _ = f.Write(postBody)

		_, _ = f.WriteString("\n")

		_, _ = f.WriteString(`"Response Inquiry CC Nice Pay"` + "\n")

		_, _ = f.Write(res)

		_, _ = f.WriteString("\n\n")
	} else {
		f, _ := os.Create("./log.json")

		defer f.Close()

		_, _ = f.WriteString(`"` + timeDate + `"` + "\n")

		_, _ = f.WriteString(`"Request Inquiry CC Nice Pay"` + "\n")

		_, _ = f.Write(postBody)

		_, _ = f.WriteString("\n")

		_, _ = f.WriteString(`"Response Inquiry CC Nice Pay"` + "\n")

		_, _ = f.Write(res)

		_, _ = f.WriteString("\n\n")
	}

	return &dataForCallbacks, nil
}

func PaymentCCNicePay(input dto.RequestPaymentCCNicePay) (bool, error) {
	var nicePayMerchantID = "IONPAYTEST"
	var nicePayMerchantKey = "33F49GnCMS1mFYlGXisbUDzVf2ATWCl9k3R++d5hDd3Frmuos/XLx8XhXpe+LDYAbpGKZYSwtlyyLOtS/8aD7A=="

	ReferenceNo := input.ReferenceNo
	amount := input.AMT
	timestamp := input.Timestamp
	TxID := input.TxId

	timeDate := time.Now().Format(time.RFC3339)

	stringMerchantToken := timestamp + nicePayMerchantID + ReferenceNo + amount + nicePayMerchantKey

	hash := NewSHA256([]byte(stringMerchantToken))
	signature := hex.EncodeToString(hash)

	urlParams := "timeStamp=" + timestamp +
		"&tXid=" + TxID +
		"&merchantToken=" + signature +
		"&cardNo=" + input.CardNo +
		"&cardExpYymm=" + input.CardExpYyMm +
		"&cardCvv=" + input.CardCVV +
		"&cardHolderNm=" + input.CardHolderNm +
		"&callBackUrl=" + input.CallBackURL

	URL, _ := url.Parse("https://dev.nicepay.co.id/nicepay/direct/v2/payment?" + urlParams)

	req := &http.Request{
		Method: "POST",
		URL:    URL,
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		return false, nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return false, nil
	}
	sb := string(body)

	i := strings.Index(sb, "SUCCESS")

	if i == -1 {
		return false, nil
	}

	if _, err := os.Stat("./log.json"); err == nil {
		f, err := os.OpenFile("./log.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		_, _ = f.WriteString("\n" + `"` + timeDate + `"` + "\n")

		_, _ = f.WriteString(`"Request Payment CC Nice Pay"` + "\n")

		_, _ = f.WriteString(urlParams)

		_, _ = f.WriteString("\n")

		_, _ = f.WriteString(`"Response Payment CC Nice Pay"` + "\n")

		_, _ = f.WriteString("SUCCESS")

		_, _ = f.WriteString("\n\n")
	} else {
		f, _ := os.Create("./log.json")

		defer f.Close()

		_, _ = f.WriteString(`"` + timeDate + `"` + "\n")

		_, _ = f.WriteString(`"Request Payment CC Nice Pay"` + "\n")

		_, _ = f.WriteString(urlParams)

		_, _ = f.WriteString("\n")

		_, _ = f.WriteString(`"Response Payment CC Nice Pay"` + "\n")

		_, _ = f.WriteString("SUCCESS")

		_, _ = f.WriteString("\n\n")
	}

	return true, nil
}
