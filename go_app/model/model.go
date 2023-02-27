package model

import (
	"time"
)

type User struct {
	Uid string`json:"uidUser"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
	ROLE string `json:"ROLE"`
}

type Store struct {
	Uid string `json:"uidStore"`
	Name string `json:"name"`
	PostCode int16 `json:"postCode"`
	Address string `json:"address"`
	City string `json:"city"`
	TypeStore string `json:"typeStore"`
}

type UserStore struct {
	UidUser string `json:"uidUser"`
	UidStore string `json:"uidStore"`
}

type Service struct {
	Uid string `json:"uidService"`
	Duration float32 `json:"duration"`
	Price float32 `json:"price"`
	Name string `json:"name"`
	UidStore string `json:"uidStore"`
}

type Appointment struct {
	Uid string `json:"uidAppointment"`
	DatetimeStart time.Time `json:"DatetimeStart"`
	UidService string `json:"uidService"`
    UidUser string `json:"uidUser"`
}

type CreateStore struct {
	UidUser string `json:"uidUser"`
	StoreName string `json:"storeName"`
	StoreType string `json:"StoreType"`
	PostCode int16 `json:"postCode"`
	Address string `json:"address"`
	City string `json:"city"`
	TypeStore string `json:"typeStore"`
}

type CreateService struct {
	UidUser string `json:"uidUser"`
	ServiceName string `json:"Servicename"`
	Duration float32 `json:"duration"`
	Price float32 `json:"price"`
	UidStore string `json:"uidStore"`
}

type AppointmentsUser struct {
	UidUser string `json:"uidUser"`
	ServiceName string `json:"Servicename"`
	Duration float32 `json:"Duration"`
	Price float32 `json:"price"`
	DatetimeStart string`json:"DatetimeStart"`
	StoreName string `json:"storeName"`
	City string `json:"city"`
	Address string `json:"address"`
	PostCode int16 `json:"postCode"`
	TypeStore string `json:"typeStore"`
}


type AppointmentsStore struct {
	UidUser string `json:"uidUser"`
	UidStore string `json:"uidStore"`
	ServiceName string `json:"Servicename"`
	DatetimeStart string`json:"DatetimeStart"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email string `json:"email"`
}


type ProfileUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email string `json:"email"`
	ServiceName string `json:"Servicename"`
	DatetimeStart string`json:"DatetimeStart"`
	Duration float32 `json:"Duration"`
	Price float32 `json:"price"`
	StoreName string `json:"storeName"`
	City string `json:"city"`
	Address string `json:"address"`
	PostCode int16 `json:"postCode"`
	TypeStore string `json:"typeStore"`
}





