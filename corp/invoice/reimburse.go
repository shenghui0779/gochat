package invoice

import (
	"encoding/json"

	"github.com/chenghonour/gochat/urls"
	"github.com/chenghonour/gochat/wx"
)

// ReimburseStatus 发票报销状态
type ReimburseStatus string

const (
	InvioceReimburseInit    ReimburseStatus = "INVOICE_REIMBURSE_INIT"    // 发票初始状态，未锁定
	InvioceReimburseLock    ReimburseStatus = "INVOICE_REIMBURSE_LOCK"    // 发票已锁定，无法重复提交报销
	InvioceReimburseClosure ReimburseStatus = "INVOICE_REIMBURSE_CLOSURE" // 发票已核销，从用户卡包中移除
)

type InvoiceUserInfo struct {
	Fee                   int                   `json:"fee"`
	Title                 string                `json:"title"`
	BillingTime           int64                 `json:"billing_time"`
	BillingNO             string                `json:"billing_no"`
	BillingCode           string                `json:"billing_code"`
	Tax                   int                   `json:"tax"`
	FeeWithoutTax         int                   `json:"fee_without_tax"`
	Detail                string                `json:"detail"`
	PdfURL                string                `json:"pdf_url"`
	TriPdfUrl             string                `json:"tri_pdf_url"`
	CheckCode             string                `json:"check_code"`
	BuyerNumber           string                `json:"buyer_number"`
	BuyerAddressAndPhone  string                `json:"buyer_address_and_phone"`
	BuyerBankAccount      string                `json:"buyer_bank_account"`
	SellerNumber          string                `json:"seller_number"`
	SellerAddressAndPhone string                `json:"seller_address_and_phone"`
	SellerBankAccount     string                `json:"seller_bank_account"`
	Remarks               string                `json:"remarks"`
	Cashier               string                `json:"cashier"`
	Maker                 string                `json:"maker"`
	ReimburseStatus       ReimburseStatus       `json:"reimburse_status"`
	OrderID               string                `json:"order_id"`
	Info                  []*InvoiceProductInfo `json:"info"`
}

type InvoiceProductInfo struct {
	Name  string `json:"name"`
	Num   int    `json:"num"`
	Unit  string `json:"unit"`
	Fee   int    `json:"fee"`
	Price int    `json:"price"`
}

type ParamsInvoice struct {
	CardID      string `json:"card_id"`
	EncryptCode string `json:"encrypt_code"`
}

type ResultInvoiceInfo struct {
	CardID    string           `json:"card_id"`
	BeginTime int64            `json:"begin_time"`
	EndTime   int64            `json:"end_time"`
	OpenID    string           `json:"openid"`
	Type      string           `json:"type"`
	Payee     string           `json:"payee"`
	Detail    string           `json:"detail"`
	UserInfo  *InvoiceUserInfo `json:"user_info"`
}

// GetInvoiceInfo 查询电子发票
func GetInvoiceInfo(cardID, encryptCode string, result *ResultInvoiceInfo) wx.Action {
	params := &ParamsInvoice{
		CardID:      cardID,
		EncryptCode: encryptCode,
	}

	return wx.NewPostAction(urls.CorpInvoiceGetInfo,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsInvoiceStatusUpdate struct {
	CardID          string          `json:"card_id"`
	EncryptCode     string          `json:"encrypt_code"`
	ReimburseStatus ReimburseStatus `json:"reimburse_status"`
}

// UpdateInvoiceStatus 更新发票状态
func UpdateInvoiceStatus(cardID, encryptCode string, status ReimburseStatus) wx.Action {
	params := &ParamsInvoiceStatusUpdate{
		CardID:          cardID,
		EncryptCode:     encryptCode,
		ReimburseStatus: status,
	}

	return wx.NewPostAction(urls.CorpInvoiceUpdateStatus,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}

type ParamsInvoiceBatchInfo struct {
	ItemList []*ParamsInvoice `json:"item_list"`
}

type ResultInvoiceBatchInfo struct {
	ItemList []*ResultInvoiceInfo `json:"item_list"`
}

// BatchGetInvoiceInfo 批量查询电子发票
func BatchGetInvoiceInfo(invoices []*ParamsInvoice, result *ResultInvoiceBatchInfo) wx.Action {
	params := &ParamsInvoiceBatchInfo{
		ItemList: invoices,
	}

	return wx.NewPostAction(urls.CorpInvoiceBatchGetInfo,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

type ParamsInvoiceStatusBatchUpdate struct {
	OpenID          string           `json:"openid"`
	ReimburseStatus ReimburseStatus  `json:"reimburse_status"`
	InvoiceList     []*ParamsInvoice `json:"invoice_list"`
}

// BatchUpdateInvoiceStatus 批量更新发票状态
func BatchUpdateInvoiceStatus(openID string, status ReimburseStatus, invoices ...*ParamsInvoice) wx.Action {
	params := &ParamsInvoiceStatusBatchUpdate{
		OpenID:          openID,
		ReimburseStatus: status,
		InvoiceList:     invoices,
	}

	return wx.NewPostAction(urls.CorpInvoiceBatchUpdateStatus,
		wx.WithBody(func() ([]byte, error) {
			return wx.MarshalNoEscapeHTML(params)
		}),
	)
}
