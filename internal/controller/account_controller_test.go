package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/sreesanthv/micro-finance-service/internal/domain"
	"github.com/sreesanthv/micro-finance-service/internal/domain/mocks"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccount(t *testing.T) {
	input := domain.Account{
		Email:    "srv@test.local",
		Password: "helloworld",
		Name:     "Sreesanth",
	}

	accountService := &mocks.AccountService{}
	accountService.On("GetByEmail", mock.Anything, "srv@test.local").Return(&input, nil)
	accountService.On("Create", mock.Anything, mock.Anything).Return(1, nil)
	accountService.On("Get", mock.Anything, 1).Return(&input, nil)

	logger := &mocks.Logger{}
	logger.On("Errorf", mock.Anything, mock.Anything)
	rd := &mocks.Redis{}
	accountController := NewAccontController(accountService, logger, rd)

	Convey("Should create account", t, func() {
		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(input)
		if err != nil {
			log.Fatal(err)
		}

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/account", &buf)
		accountController.Create(w, req)

		So(w.Result().StatusCode, ShouldEqual, http.StatusCreated)
		decoder := json.NewDecoder(w.Body)
		out := domain.Account{}
		decoder.Decode(&out)
		So(out.Email, ShouldEqual, input.Email)
	})

	Convey("Should delete account", t, func() {
		accountService.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(nil)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("id", "1")

		var buf bytes.Buffer
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/account/1", &buf)
		req = req.WithContext(context.WithValue(req.Context(), "accountID", 1))
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
		accountController.Delete(w, req)
		So(w.Result().StatusCode, ShouldEqual, http.StatusCreated)
	})
}
