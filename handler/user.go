package handler

// type omit bool

// // ListUsers all users
// func ListUsers(w http.ResponseWriter, r *http.Request) {
// 	render.JSON(w, r, model.ListUsers())
// }

// // GetUser returns user
// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	login := chi.URLParam(r, "username")

// 	u, err := model.GetUser(login)
// 	if err != nil {
// 		render.JSON(w, r, model.Message{Msg: err.Error()})
// 		log.Println(err)
// 		return
// 	}

// 	render.JSON(w, r, &struct {
// 		*model.User
// 		ID       string `json:"id"`
// 		Password omit   `json:"password,omitempty"`
// 	}{
// 		User: &u,
// 		ID:   login,
// 	})
// }
