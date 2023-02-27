package db

const REQ_CREATE_STORE = `
	INSERT INTO store
	(uid_store, name, post_code, address, city, type_store)
	Values (?,?,?,?,?,?);`

const REQ_GET_USER = `
	SELECT * from user
	WHERE email=? and password=?`

const REQ_CREATE_USER = `
	INSERT INTO user
	(uid_user, first_name, last_name, email, password, ROLE)
	Values (?,?, ?, ?, ?, ?);`

const REQ_GET_ROLE_USER = `
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
	(uid_appointment, datetime_start, user_uid, service_uid)
	Values(?,?,?,?);`

const REQ_GET_APPOINTMENTS_USER = `
	SELECT service.name as service, service.duration,
	service.price, appointment.datetime_start,
	store.name as store, store.city,
	store.address, store.post_code,
	store.type_store FROM appointment
	JOIN service ON ( appointment.service_uid = service.uid_service)
	JOIN store ON ( service.store_uid = store.uid_store)
	WHERE appointment.user_uid=?`

const REQ_GET_APPOINTMENTS_STORE = `
	SELECT service.name as service,
	user.first_name, user.last_name,
	user.email, appointment.datetime_start
	FROM appointment
	JOIN service ON (appointment.service_uid = service.uid_service)
	JOIN user ON (appointment.user_uid = user.uid_user)
	WHERE service.store_uid=?`

const REQ_GET_PROFILE_USER = `
	SELECT uid_user as uid, first_name, last_name,
	email, ROLE
	FROM user
	WHERE user.uid_user=?`






