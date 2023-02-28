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


func CheckUserUid(uidUserJwt string, uidUser string) bool {
    _, isAuthorized := VerifyRole(uidUserJwt, "")
    if isAuthorized == true || uidUserJwt == uidUser {
        return true
    }
    return false
}


func CheckRoleUserStore(uidUser string, uidStore string) bool {
    userRole, isAuthorized := VerifyRole(uidUser, "EMPLOYEE_ROLE")

    if isAuthorized != true {
        return false
    }
    if userRole != "ADMIN_ROLE" {
        checkUserStore := CheckUserStore(uidUser, uidStore)
        if checkUserStore != true {
            return false
        }
    }
    return true
}

func CheckUserStore(uidUser string, uidStore string) bool {
    client, _ := db.DbConnection()
    defer client.Close()
    var userStore model.UserStore

    query, _ := client.Prepare(db.REQ_GET_USER_STORE)

    err := query.QueryRow(uidUser, uidStore).Scan(&userStore.UidUser, &userStore.UidStore)
    if err != nil {
        return false
    }
    return true
}






