package controller

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/sreesanthv/micro-finance-service/internal/domain"
)

type AccountController struct {
	Controller
	accountService domain.AccountService
	log            domain.Logger
	redis          domain.Redis
}

func NewAccontController(actService domain.AccountService, l domain.Logger, r domain.Redis) *AccountController {
	return &AccountController{
		accountService: actService,
		log:            l,
		redis:          r,
	}
}

func (a *AccountController) Create(w http.ResponseWriter, r *http.Request) {
	req := domain.Account{}
	err := a.Decode(r, &req)
	if err != nil {
		a.log.Errorf("Error decoding account details. err: %s", err)
		a.SendMessage(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if req.Email == "" || req.Password == "" {
		a.SendMessage(w, http.StatusBadRequest, "Email and password are mandatory")
		return
	}

	act, _ := a.accountService.GetByEmail(r.Context(), req.Email)
	if act != nil && act.ID != 0 {
		a.SendMessage(w, http.StatusBadRequest, "Account with Email already exists")
		return
	}

	req.Password, err = a.hash(req.Password)
	if err != nil {
		a.log.Errorf("Error in hashing password: %s", err)
		a.SendMessage(w, http.StatusInternalServerError, "")
		return
	}

	id, err := a.accountService.Create(r.Context(), &req)
	if err != nil {
		a.SendMessage(w, http.StatusInternalServerError, "")
		return
	}

	act, err = a.accountService.Get(r.Context(), id)
	if err != nil {
		a.SendMessage(w, http.StatusInternalServerError, "")
		return
	}
	act.Password = ""

	a.SendResponse(w, http.StatusCreated, act)
}

func (a *AccountController) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		a.SendMessage(w, http.StatusBadRequest, "")
		return
	}

	if id != a.GetAccountId(r) {
		a.SendMessage(w, http.StatusUnauthorized, "")
		return
	}

	act, err := a.accountService.Get(r.Context(), id)
	if err != nil {
		a.SendMessage(w, http.StatusNotFound, "Account doesn't exist")
		return
	}
	act.Password = ""
	a.SendResponse(w, http.StatusCreated, act)
}

func (a *AccountController) Update(w http.ResponseWriter, r *http.Request) {
	req := domain.Account{}
	err := a.Decode(r, &req)
	if err != nil {
		a.log.Errorf("Error decoding account details. err: %s", err)
		a.SendMessage(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		a.SendMessage(w, http.StatusBadRequest, "")
		return
	}

	if id != a.GetAccountId(r) {
		a.SendMessage(w, http.StatusUnauthorized, "")
		return
	}

	act, err := a.accountService.Get(r.Context(), id)
	if err != nil {
		a.SendMessage(w, http.StatusNotFound, "Account doesn't exist")
		return
	}

	if req.Email != "" {
		act.Email = req.Email
	}
	if req.Password != "" {
		act.Password, _ = a.hash(req.Password)
	}
	if req.Address != "" {
		act.Address = req.Address
	}
	if req.Email != "" {
		act.Email = req.Email
	}
	if req.Location != "" {
		act.Location = req.Location
	}
	if req.Name != "" {
		act.Name = req.Name
	}
	if req.Nationality != "" {
		act.Nationality = req.Nationality
	}
	if req.Pan != "" {
		act.Pan = req.Pan
	}
	if req.Phone != "" {
		act.Phone = req.Phone
	}
	if req.Sex != "" {
		act.Sex = req.Sex
	}

	a.accountService.Update(r.Context(), act)
	act.Password = ""
	a.SendResponse(w, http.StatusCreated, act)
}

func (a *AccountController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		a.SendMessage(w, http.StatusBadRequest, "")
		return
	}

	if id != a.GetAccountId(r) {
		a.SendMessage(w, http.StatusUnauthorized, "")
		return
	}

	act, err := a.accountService.Get(r.Context(), id)
	if err != nil {
		a.SendMessage(w, http.StatusNotFound, "Account doesn't exist")
		return
	}

	a.accountService.Delete(r.Context(), act.ID)
	a.SendMessage(w, http.StatusCreated, "Account deleted")
}

func (a *AccountController) Login(w http.ResponseWriter, r *http.Request) {
	req := domain.Account{}
	err := a.Decode(r, &req)
	if err != nil {
		a.log.Errorf("Error decoding login details. err: %s", err)
		a.SendMessage(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	act, err := a.accountService.GetByEmail(r.Context(), req.Email)
	valid := a.compareHash(act.Password, req.Password)
	if !valid {
		a.SendMessage(w, http.StatusUnauthorized, "Invalid credentials")
	}

	bytes := make([]byte, 32)
	rand.Read(bytes)
	token := hex.EncodeToString(bytes)

	a.redis.Set(token, act.ID, time.Hour*24)
	a.SendResponse(w, http.StatusOK, map[string]string{
		"token": token,
	})
}

func (a *AccountController) Logout(w http.ResponseWriter, r *http.Request) {
	a.redis.Del(a.GetToken(r))
}

type TransactionRequest struct {
	Amount    float64
	Narration string
}

func (a *AccountController) Debit(w http.ResponseWriter, r *http.Request) {
	req := TransactionRequest{}
	err := a.Decode(r, &req)
	if err != nil {
		a.log.Errorf("Error decoding transaction details. err: %s", err)
		a.SendMessage(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if req.Amount == 0 {
		a.SendMessage(w, http.StatusBadRequest, "Amount should greater than zero")
	}

	tr := &domain.Transaction{
		AccountId: a.GetAccountId(r),
		Debit:     req.Amount,
		Narration: req.Narration,
	}
	a.accountService.SaveTransaction(r.Context(), tr)
}

func (a *AccountController) Credit(w http.ResponseWriter, r *http.Request) {
	req := TransactionRequest{}
	err := a.Decode(r, &req)
	if err != nil {
		a.log.Errorf("Error decoding transaction details. err: %s", err)
		a.SendMessage(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	if req.Amount == 0 {
		a.SendMessage(w, http.StatusBadRequest, "Amount should greater than zero")
	}

	tr := &domain.Transaction{
		AccountId: a.GetAccountId(r),
		Credit:    req.Amount,
		Narration: req.Narration,
	}
	a.accountService.SaveTransaction(r.Context(), tr)
}

func (a *AccountController) GetAccountId(r *http.Request) int {
	return r.Context().Value("accountID").(int)
}

func (a *AccountController) GetToken(r *http.Request) string {
	return r.Context().Value("token").(string)
}
