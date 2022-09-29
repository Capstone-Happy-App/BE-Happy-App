package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	userData "capstone/happyApp/features/user/data"
	userDelivery "capstone/happyApp/features/user/delivery"
	userUsecase "capstone/happyApp/features/user/usecase"

	loginData "capstone/happyApp/features/login/data"
	loginDelivery "capstone/happyApp/features/login/delivery"
	loginUsecase "capstone/happyApp/features/login/usecase"

	communityData "capstone/happyApp/features/community/data"
	communityDelivery "capstone/happyApp/features/community/delivery"
	communityUsecase "capstone/happyApp/features/community/usecase"

	eventData "capstone/happyApp/features/event/data"
	eventDelivery "capstone/happyApp/features/event/delivery"
	eventUsecase "capstone/happyApp/features/event/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {

	eventDataFactory := eventData.New(db)
	eventUsecaseFactory := eventUsecase.New(eventDataFactory)
	eventDelivery.New(e, eventUsecaseFactory)

	communityDataFactory := communityData.New(db)
	communityUsecaseFactory := communityUsecase.New(communityDataFactory)
	communityDelivery.New(e, communityUsecaseFactory)

	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	loginDataFactory := loginData.New(db)
	loginUsecaseFactory := loginUsecase.New(loginDataFactory)
	loginDelivery.New(e, loginUsecaseFactory)

}
