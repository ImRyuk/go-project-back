package db

const REQ_CREATE_STORE = `
	INSERT INTO store
	(uid_store, name, post_code, address, city, type_store)
	Values (?,?,?,?,?,?);`

const REQ_GET_USER = `
	SELECT * from user
	WHERE email=? and password=?`

const REQ_GET_USER_BY_MAIL = `
	SELECT * from user
	WHERE email=?`

const REQ_CREATE_USER = `
	INSERT INTO user
	(uid_user, first_name, last_name, email, password, ROLE)
	Values (?,?, ?, ?, ?, ?);`

const REQ_GET_USER_ROLE = `
	SELECT uid_user, ROLE
	from user WHERE uid_user=?`

const REQ_LOGIN_USER = `
	SELECT uid_user, email from user
	WHERE email=? and password=?`

const REQ_CREATE_SERVICE = `
	INSERT INTO service
	(uid_service, name, store_uid, duration, price)
	Values (?,?,?,?,?);`

const REQ_GET_USER_STORE = `
	SELECT user_uid, store_uid
	FROM user_has_store
	WHERE user_uid =? and store_uid =?;`

const REQ_CREATE_USER_STORE = `
	INSERT INTO user_has_store (user_uid, store_uid)
	Values(?,?);`

const REQ_CREATE_APPOINTMENT = `
	INSERT INTO appointment
	(uid_appointment, datetime_start, datetime_end, user_uid, service_uid)
	Values(?,?,?,?,?);`

const REQ_GET_APPOINTMENTS_USER = `
	SELECT service.name as service, service.duration,
	service.price, appointment.datetime_start, appointment.datetime_end,
	store.name as store, store.city,
	store.address, store.post_code,
	store.type_store FROM appointment
	JOIN service ON ( appointment.service_uid = service.uid_service)
	JOIN store ON ( service.store_uid = store.uid_store)
	WHERE appointment.user_uid=?`

const REQ_GET_APPOINTMENTS_STORE = `
	SELECT service.name as service,
	user.first_name, user.last_name,
	user.email, appointment.datetime_start, appointment.datetime_end
	FROM appointment
	JOIN service ON (appointment.service_uid = service.uid_service)
	JOIN user ON (appointment.user_uid = user.uid_user)
	WHERE service.store_uid=?`

const REQ_GET_PROFILE_USER = `
	SELECT uid_user as uid, first_name, last_name,
	email, ROLE
	FROM user
	WHERE user.uid_user=?;`

const REQ_GET_STORE = `
	SELECT uid_store, name, type_store,
	city, post_code, address
	FROM store
	WHERE store.uid_store=?;`

const REQ_GET_STORES_BY_FILTER = `
    SELECT uid_store, name, type_store,
    city, post_code, address
    FROM store
    WHERE 1=1
`
const REQ_GET_STORES = `
	SELECT uid_store, name,
	type_store, city, post_code,
	address FROM store`

const REQ_GET_SERVICE_BY_UID =`
	SELECT uid_service, name, duration, price, store_uid
	FROM service WHERE uid_service=?`

const REQ_GET_SERVICE_BY_STORE =`
	SELECT uid_service, name, duration, price, store_uid
	FROM service WHERE store_uid=?`

const REQ_CHECK_BOOKING_EXISTS = `
	SELECT uid_appointment
	FROM appointment
	JOIN service ON (appointment.service_uid = service.uid_service)
	WHERE datetime_start<=? and datetime_end>=?
	and service.store_uid=?
	or datetime_start>=? and datetime_start<=?
	and service.store_uid=?`


const REQ_GET_STORES_BY_USER = `
    SELECT  user.uid_user as user_uid, store.uid_store as store_uid, store.name, store.type_store,
    store.city, store.post_code, store.address
    FROM user_has_store
    JOIN user ON (user_has_store.user_uid = user.uid_user)
    JOIN store ON (user_has_store.store_uid = store.uid_store)
    WHERE user_has_store.user_uid=?
`

