package controller

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

type Middleware struct {
	controller *AccountController
}

func NewMiddleware(ctrl *AccountController) *Middleware {
	return &Middleware{
		controller: ctrl,
	}
}

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", 1)
		ctx := r.Context()

		id := m.controller.redis.Get(token).Val()
		if id == "" {
			m.controller.SendMessage(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		actId, err := strconv.Atoi(id)
		if err != nil {
			m.controller.SendMessage(w, http.StatusInternalServerError, "")
			return
		}

		act, err := m.controller.accountService.Get(ctx, actId)
		if err != nil {
			m.controller.SendMessage(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		ctx = context.WithValue(ctx, "accountID", act.ID)
		ctx = context.WithValue(ctx, "token", token)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
