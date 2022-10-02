package middleware

// import (
// 	"context"
// 	"net/http"

// 	"github.com/justinas/nosurf"
// 	"github.com/volatiletech/authboss/v3"
// )

// // layoutData is passing pointers to pointers be able to edit the current pointer
// // to the request. This is still safe as it still creates a new request and doesn't
// // modify the old one, it just modifies what we're pointing to in our methods so
// // we're able to skip returning an *http.Request everywhere
// func layoutData(w http.ResponseWriter, r **http.Request) authboss.HTMLData {
// 	currentUserName := ""
// 	userInter, err := ab.LoadCurrentUser(r)
// 	if userInter != nil && err == nil {
// 		currentUserName = userInter.(*User).Name
// 	}

// 	return authboss.HTMLData{
// 		"loggedin":          userInter != nil,
// 		"current_user_name": currentUserName,
// 		"csrf_token":        nosurf.Token(*r),
// 		"flash_success":     authboss.FlashSuccess(w, *r),
// 		"flash_error":       authboss.FlashError(w, *r),
// 	}
// }

// func dataInjector(handler http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		data := layoutData(w, &r)
// 		r = r.WithContext(context.WithValue(r.Context(), authboss.CTXKeyData, data))
// 		handler.ServeHTTP(w, r)
// 	})
// }
