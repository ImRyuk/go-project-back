package permissions
import (
	"go_app/db"
	"go_app/model"
)


func VerifyRole(uidUser string, role string) (string,bool){
	var user model.User
	client, _ := db.DbConnection()
	defer client.Close()
    err := client.QueryRow(db.REQ_GET_USER_ROLE , uidUser).Scan(&user.Email, &user.ROLE)
    if err != nil {
        return "", false
    }
    if user.ROLE == role {
        return "EMPLOYEE_ROLE", true
    }
    if user.ROLE == "ADMIN_ROLE" {
        return "ADMIN_ROLE", true
    }

    return "", false
}


func CheckUserUid(userUidJwt string, userUid string) bool {
    _, isAuthorized := VerifyRole(userUidJwt, "")
    if isAuthorized == true || userUidJwt == userUid {
        return true
    }
    return false
}


func CheckRoleUserStore(userUid string, storeUid string) bool {
    userRole, isAuthorized := VerifyRole(userUid, "EMPLOYEE_ROLE")

    if isAuthorized != true {
        return false
    }
    if userRole != "ADMIN_ROLE" {
        checkUserStore := CheckUserStore(userUid, storeUid)
        if checkUserStore != true {
            return false
        }
    }
    return true
}

func CheckUserStore(userUid string, storeUid string) bool {
    client, _ := db.DbConnection()
    defer client.Close()
    var userStore model.UserStore
    query, _ := client.Prepare(db.REQ_GET_USER_STORE)

    err := query.QueryRow(userUid, storeUid).Scan(&userStore.UserUid, &userStore.StoreUid)
    if err != nil {
        return false
    }
    return true
}






