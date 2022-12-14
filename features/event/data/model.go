package data

import (
	"capstone/happyApp/features/event"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title       string
	Description string
	Price       uint64
	Date        time.Time
	Location    string
	CommunityID uint
	Community   Community
	Event       []JoinEvent
}

type JoinEvent struct {
	gorm.Model
	UserID           uint
	EventID          uint
	Order_id         string
	Type_payment     string
	Payment_method   string
	Status_payment   string
	Midtrans_virtual string
	GrossAmount      string
}

type User struct {
	gorm.Model
}

type Community struct {
	gorm.Model
	Logo  string
	Title string
}

type JoinCommunity struct {
	gorm.Model
	UserID      uint
	CommunityID uint
	Role        string
}

type temp struct {
	ID          uint
	Logo        string
	Title       string
	Description string
	Count       int64
}

type tempDetail struct {
	ID            uint
	Title         string
	Description   string
	Penyelenggara string
	Date          time.Time
	Price         uint64
	Location      string
}

type tempRespon struct {
	ID           uint
	Logo         string
	Title        string
	Descriptions string
	Date         time.Time
	Price        int64
}

func fromCore(data event.EventCore) Event {
	return Event{
		Title:       data.Title,
		Description: data.Description,
		Price:       data.Price,
		Date:        data.Date,
		Location:    data.Location,
		CommunityID: data.CommunityID,
	}
}

func EventList(data []tempRespon) []event.Response {

	var dataRespon []event.Response
	for _, v := range data {
		dataRespon = append(dataRespon, event.Response{
			ID:           v.ID,
			Logo:         v.Logo,
			Title:        v.Title,
			Descriptions: v.Descriptions,
			Date:         v.Date,
			Price:        v.Price,
		})
	}

	return dataRespon
}

func EventListComu(data []event.Response, dataComu temp, role string) event.CommunityEvent {
	return event.CommunityEvent{
		ID:          dataComu.ID,
		Role:        role,
		Logo:        dataComu.Logo,
		Title:       dataComu.Title,
		Description: dataComu.Description,
		Count:       dataComu.Count,
		Event:       data,
	}
}

func EventDetails(data tempDetail, role string) event.EventDetail {
	return event.EventDetail{
		ID:            data.ID,
		Title:         data.Title,
		Description:   data.Description,
		Status:        role,
		Penyelenggara: data.Penyelenggara,
		Date:          data.Date,
		Price:         data.Price,
		Location:      data.Location,
	}
}

func toModelJoinEvent(data event.JoinEventCore) JoinEvent {

	return JoinEvent{
		UserID:           data.UserID,
		EventID:          data.EventID,
		Type_payment:     data.Type_payment,
		Payment_method:   data.Payment_method,
		Order_id:         data.Order_id,
		Status_payment:   data.Status_payment,
		Midtrans_virtual: data.Midtrans_virtual,
		GrossAmount:      data.GrossAmount,
	}
}
