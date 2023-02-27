package permissions
import (
	"go_app/db"
	"go_app/model"
)


func VerifyRole(uidUser string, role string) bool {
	var user model.User
	client, _ := db.DbConnection()
	defer client.Close()
    err := client.QueryRow(db.REQ_GET_ROLE_USER , uidUser).Scan(&user.Email, &user.ROLE)

    if err != nil {
        return false
    }
    if user.ROLE != role {
    	return false
    }
    return true
}

// func VerifyUserIsAdmin(uidUser string) bool {
//     var user model.User
//     defer client.Close()
//     err := client.QueryRow(db.REQ_GET_ROLE_USER , uidUser).Scan(&user.Email, &user.ROLE, &user.isAdmin)

//     if err != nil {
//         return false
//     }
//     if user.isAdmin!= true {
//         return false
//     }
// }


func CheckUserUid(uidUserJwt string, uidUser string) bool {
    if uidUserJwt == uidUser {
        return true
    }
    return false
}


func CheckRoleUserStore(uidUser string, uidStore string) bool {
    checkRole := VerifyRole(uidUser, "ROLE_EMPLOYER")
    if checkRole != true {
        return false
    }
    checkUserStore := CheckUserStore(uidUser, uidStore)
    if checkUserStore != true {
        return false
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


