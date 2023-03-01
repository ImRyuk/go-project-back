package model

import (
	"time"
)

type User struct {
	UserUid string`json:"userUid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
	ROLE string `json:"ROLE"`
}

type Store struct {
	StoreUid string `json:"storeUid"`
	Name string `json:"name"`
	PostCode int32 `json:"postCode"`
	Address string `json:"address"`
	City string `json:"city"`
	StoreType string `json:"storeType"`
}

type UserStore struct {
	UserUid string `json:"userUid"`
	StoreUid string `json:"storeUid"`
}

type Service struct {
	ServiceUid string `json:"serviceUid"`
	Duration int32 `json:"duration"`
	Price float32 `json:"price"`
	Name string `json:"name"`
	StoreUid string `json:"storeUid"`
}

type Appointment struct {
	AppointmentUid string `json:"appointmentUid"`
	DatetimeStart time.Time `json:"datetimeStart"`
	DatetimeEnd time.Time `json:"datetimeEnd"`
	ServiceUid string `json:"serviceUid"`
    UserUid string `json:"userUid"`
}

type CreateService struct {
	UidUser string `json:"uidUser"`
	ServiceName string `json:"Servicename"`
	Duration float32 `json:"duration"`
	Price float32 `json:"price"`
	UidStore string `json:"uidStore"`
}

type AppointmentsUser struct {
	UserUid string `json: userUid"`
	ServiceName string `json:"serviceName"`
	Duration float32 `json:"duration"`
	Price float32 `json:"price"`
	DatetimeStart string`json:"datetimeStart"`
	StoreName string `json:"storeName"`
	City string `json:"city"`
	Address string `json:"address"`
	PostCode int32 `json:"postCode"`
	StoreType string `json:"storeType"`
}


type AppointmentsStore struct {
	UserUid string `json:"userUid"`
	StoreUid string `json:"storeuid"`
	ServiceName string `json:"serviceName"`
	DatetimeStart string`json:"datetimeStart"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email string `json:"email"`
}






