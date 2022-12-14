package cart

import (
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
)

type CoreCart struct {
	ID           uint
	UserID       uint
	ProductID    uint
	Name         string
	Descriptions string
	Photo        string
	Price        int
}

type CoreCommunity struct {
	ID           uint
	Title        string
	Descriptions string
	Logo         string
	Members      int64
	Role         string
}

type CoreHistory struct {
	ID              uint
	Carts           []int
	Street          string
	City            string
	Province        string
	Type_Payment    string
	Status_Payment  string
	Virtual_Account string
	OrderID         string
	Gross           string
	Created_at      time.Time
}

type TransactionCart struct {
	ID            uint
	TransactionID uint
	CartID        uint
	Price         uint
}
type CoreProductResponse struct {
	ID    uint
	Name  string
	Photo string
	Price int
	Buyer string
}

type CorePayment struct {
	ID      uint
	UserID  uint
	OrderID string
	Groos   int
}

type DataInterface interface {
	InsertIntoCart(userid, productid int) (string, error)
	GetCommunity(userid, communityid int) (CoreCommunity, string, error)
	SelectCartList(userid, communityid int) ([]CoreCart, string, error)
	DeleteFromCart(userid, cartid int) (string, error)
	InsertIntoTransaction(core CoreHistory) (int, string, error)
	GetTotalTransaction(trasacid int) (int, string, error)
	DeleteCart(core CoreHistory) (string, error)
	CheckStock([]int) ([]int, string, error)
	UpdateStock([]int) (string, error)
	UpdateHistory(core CoreHistory, userid int) (string, error)
	GetUserRole(Userid, communityid int) (string, error)
	SelectCommunity(communityid int) (CoreCommunity, string, error)
	ListHistoryProduct(communityid int) ([]CoreProductResponse, string, error)
	CheckCommunity(prodid int) (int, string, error)
	CheckMember(userid, communityid int) (int, string, error)
}

type UsecaseInterface interface {
	AddToCart(userid, productid int) (string, error)
	GetCartList(userid, communityid int) (CoreCommunity, []CoreCart, string, error)
	DeleteFromCart(userid, cartid int) (string, error)
	InsertIntoTransaction(core CoreHistory) (int, int, string, error)
	GetCharge(orderid int, gross int, payment, table string) (coreapi.ChargeReq, string, error)
	ChargeRequest(transfer coreapi.ChargeReq, typename string) (coreapi.ChargeReq, string, error)
	UpdateHistory(core CoreHistory, userid int) (string, error)
	GetCommunityHistory(userid, communityid int) (CoreCommunity, []CoreProductResponse, string, error)
}
