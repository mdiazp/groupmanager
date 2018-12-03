package tests

import (
	"database/sql"
	"math/rand"
	"testing"

	"github.com/mdiazp/groupmanager/server/db/models"
)

func TestUserHandler(t *testing.T) {
	var e error
	db := GetDBHandler(t)

	// Add users
	n := 500
	users := make(map[string]models.User)
	for sz := 0; sz < n; sz++ {
		username := GetRandString(10)

		if _, ok := users[username]; ok {
			sz--
			continue
		}

		u := models.User{
			Username: username,
			Rol:      GetRandomUserRol(),
		}

		e = db.CreateUser(&u)
		if e != nil {
			t.Fatalf("Fail creating user: %s", e.Error())
		}
		users[username] = u
	}

	// Check Unique Username Constraint
	// Check Not Null Username Constraint
	// Check Not Null Rol Constraint
	for username := range users {
		u := models.User{}
		u.Username = username
		u.Rol = models.RolSuperadmin
		if rand.Int()%2 == 0 {
			u.Rol = models.RolGroupadmin
		}

		e = db.CreateUser(&u)
		if e == nil {
			t.Fatalf("Fail checking unique username constraints")
		} else {
			t.Logf("Unique username constraint database error: %s", e.Error())
		}

		u.Username = ""
		e = db.CreateUser(&u)
		if e == nil {
			t.Fatalf("Fail checking not null username constraints")
		} else {
			t.Logf("Not Null username constraint database error: %s", e.Error())
		}

		u.Username = "x"
		u.Rol = ""
		e = db.CreateUser(&u)
		if e == nil {
			t.Fatalf("Fail checking not null rol constraints")
		} else {
			t.Logf("Not Null rol constraint database error: %s", e.Error())
		}
	}

	// RetrieveUser and check good values
	for _, user := range users {
		u := models.User{}
		u.ID = user.ID

		e = db.RetrieveUser(&u)
		if e != nil {
			t.Fatalf("Fail retrieve users: %s", e.Error())
		}

		if u.ID != user.ID || u.Username != user.Username ||
			u.Rol != user.Rol {
			t.Fatal("Fail user has bad wrong values:", u, user)
		}
	}

	// UpdateUsers
	for username, user := range users {
		username += "Z"
		u := user
		u.Username = username

		users[username] = u
		e = db.UpdateUser(&u)
		if e != nil {
			t.Fatalf("Fail UpdateUser: %s", e.Error())
		}
	}

	// DeleteUsers
	for _, user := range users {
		u := user
		e = db.DeleteUser(&u)
		if e != nil {
			t.Fatalf("Fail Delete User: %s", e.Error())
		}
		n--

		// Count
		cnt, e := db.CountUsers(nil, nil, nil)
		if e != nil {
			t.Fatalf("Error CountUsers: %s", e.Error())
		}

		if cnt != n {
			t.Fatalf("Error deleting user, cnt != n, cnt=%d n=%d", cnt, n)
		}

		u = user
		e = db.RetrieveUser(&u)
		if e != nil && e != sql.ErrNoRows {
			t.Fatalf("Fail at RetrieveUser after delete it: %s", e.Error())
		}
	}
}
