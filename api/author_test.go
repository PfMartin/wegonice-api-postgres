package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/PfMartin/wegonice-api/db/mock"
	db "github.com/PfMartin/wegonice-api/db/sqlc"
	"github.com/PfMartin/wegonice-api/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func randomAuthor(t *testing.T, userEmail string) (author db.Author) {
	return db.Author{
		AuthorName:  util.RandomString(6),
		Website:     util.RandomString(6),
		Instagram:   util.RandomString(6),
		Youtube:     util.RandomString(6),
		UserCreated: userEmail,
	}
}

func TestCreateAuthorAPI(t *testing.T) {
	user, _ := randomUser(t)
	author := randomAuthor(t, user.Email)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"author_name": author.AuthorName,
				"website":     author.Website,
				"instagram":   author.Instagram,
				"youtube":     author.Youtube,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateAuthorParams{
					AuthorName:  author.AuthorName,
					Website:     author.Website,
					Instagram:   author.Instagram,
					Youtube:     author.Youtube,
					UserCreated: author.UserCreated,
				}

				store.EXPECT().CreateAuthor(gomock.Any(), arg).Times(1).Return(author, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/authors"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			addAuthorization(t, request, server.tokenMaker, authorizationTypeBearer, user.Email, time.Minute)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
