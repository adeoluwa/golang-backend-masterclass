package api

import (
	mockdb "backend_masterclass/db/mock"
	db "backend_masterclass/db/sqlc"
	"backend_masterclass/util"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T){
	account := randomAccount()

	testCases := []struct{
		name string
		accountID int64
		buildStuds func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			accountID: account.ID,
			buildStuds: func(store *mockdb.MockStore){
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)
			},
		   checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
			require.Equal(t, http.StatusOK, recorder.Code)
	        requireBodyMatchAccount(t, recorder.Body, account)
		   },
		},
		{
			name: "NotFound",
			accountID: account.ID,
			buildStuds: func(store *mockdb.MockStore){
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(db.Account{}, sql.ErrNoRows)
			},
		    checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
			require.Equal(t, http.StatusNotFound, recorder.Code)
		   },
		},
		{
			name: "InternalError",
			accountID: account.ID,
			buildStuds: func(store *mockdb.MockStore){
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(db.Account{}, sql.ErrConnDone)
			},
		    checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
			require.Equal(t, http.StatusInternalServerError, recorder.Code)
		   },
		},
		{
			name: "InvalidID",
			accountID: 0,
			buildStuds: func(store *mockdb.MockStore){
				store.EXPECT().GetAccount(gomock.Any(), gomock.Any()).Times(0)
			},
		    checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
			require.Equal(t, http.StatusBadRequest, recorder.Code)
		   },
		},
	}

	for i := range testCases{

		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T){

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
		
			store := mockdb.NewMockStore(ctrl)
			tc.buildStuds(store)
		
			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()
		
			url := fmt.Sprintf("/account/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
		
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestCreateAccountAPI(t *testing.T){
	account := randomAccount()

	testCases := []struct {
		name string
		body gin.H
		buildStuds func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"Owner":account.Owner,
				"Currency":account.Currency,
				"Balance":account.Balance,
			},
			buildStuds: func(store *mockdb.MockStore){
				arg := db.CreateAccountParams{
					Owner: account.Owner,
					Currency: account.Currency,
					Balance: 0,
				}
				store.EXPECT().CreateAccount(gomock.Any(), gomock.Eq(arg)).Times(1).Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"Owner":account.Owner,
				"Currency":account.Currency,
				"Balance":account.Balance,
			},
			buildStuds: func(store *mockdb.MockStore){
				arg := db.CreateAccountParams{
					Owner: account.Owner,
					Currency: account.Currency,
					Balance: 0,
				}
				store.EXPECT().CreateAccount(gomock.Any(), gomock.Eq(arg)).Times(1).Return(db.Account{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
				// requireBodyMatchAccount(t, recorder.Body, account)
			},
		},
		{
			name: "InvalidCurrency",
			body: gin.H{
				"currency":"invalid",
			},
			buildStuds: func(store *mockdb.MockStore){
				store.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			} ,
		},
	}

	for i := range testCases{

		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T){

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
		
			store := mockdb.NewMockStore(ctrl)
			tc.buildStuds(store)
		
			// start test server and send request
			server := NewServer(store)
			// recorder := httptest.NewRecorder()
		
			url := "/account"
			bodyBytes, err := json.Marshal(tc.body)
			require.NoError(t, err)
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyBytes))
			require.NoError(t, err)
		
			recorder := httptest.NewRecorder()
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}



func randomAccount() db.Account{
	return db.Account{
		ID: util.RandomInt(1, 1000),
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T,body *bytes.Buffer, account db.Account){
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}