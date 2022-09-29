package usecase

import "capstone/happyApp/features/event"

type usecaseEvent struct {
	eventData event.DataInterface
}

func New(data event.DataInterface) event.UsecaseInterface {
	return &usecaseEvent{
		eventData: data,
	}
}

func (usecase *usecaseEvent) PostEvent(data event.EventCore, id int) int {

	if data.Description == "" || data.Location == "" || data.Title == "" {
		return -3
	}

	row := usecase.eventData.InsertEvent(data, id)
	return row

}

func (usecase *usecaseEvent) GetEvent(search string) ([]event.Response, error) {

	data, err := usecase.eventData.SelectEvent(search)
	if err != nil {
		return nil, err
	}

	return data, nil
}