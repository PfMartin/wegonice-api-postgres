package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/PfMartin/wegonice-api/db/mock"
	db "github.com/PfMartin/wegonice-api/db/sqlc"
	"github.com/PfMartin/wegonice-api/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func randomAuthor(t *testing.T, userEmail string) (author db.Author) {
	return db.Author{
		ID:          1,
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
				requireBodyMatchAuthor(t, recorder.Body, author)
			},
		},
		{
			name: "ForeignKeyViolation",
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

				store.EXPECT().CreateAuthor(gomock.Any(), arg).Times(1).Return(db.Author{}, &pq.Error{Code: "23503"})
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name: "UniqueViolation",
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

				store.EXPECT().CreateAuthor(gomock.Any(), arg).Times(1).Return(db.Author{}, &pq.Error{Code: "23505"})
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name: "UniqueViolation",
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

				store.EXPECT().CreateAuthor(gomock.Any(), arg).Times(1).Return(db.Author{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{
				"website":   author.Website,
				"instagram": author.Instagram,
				"youtube":   author.Youtube,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateAuthor(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
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

func TestGetAuthorAPI(t *testing.T) {
	user, _ := randomUser(t)
	author := randomAuthor(t, user.Email)

	testCases := []struct {
		name          string
		url           string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			url:  fmt.Sprintf("/authors/%d", author.ID),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAuthor(gomock.Any(), author.ID).Times(1).Return(author, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAuthor(t, recorder.Body, author)
			},
		},
		{
			name: "InvalidURI",
			url:  "/authors",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAuthor(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			url:  "/authors/0",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAuthor(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "NotFound",
			url:  fmt.Sprintf("/authors/%d", author.ID),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAuthor(gomock.Any(), author.ID).Times(1).Return(db.Author{}, sql.ErrNoRows)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InternalError",
			url:  fmt.Sprintf("/authors/%d", author.ID),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAuthor(gomock.Any(), author.ID).Times(1).Return(db.Author{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "UniqueViolation",
			url:  fmt.Sprintf("/authors/%d", author.ID),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetAuthor(gomock.Any(), author.ID).Times(1).Return(db.Author{}, &pq.Error{Code: "23503"})
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
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

			request, err := http.NewRequest(http.MethodGet, tc.url, bytes.NewBuffer([]byte{}))
			require.NoError(t, err)

			addAuthorization(t, request, server.tokenMaker, authorizationTypeBearer, user.Email, time.Minute)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestDeleteAuthorAPI(t *testing.T) {
	user, _ := randomUser(t)
	author := randomAuthor(t, user.Email)

	testCases := []struct {
		name          string
		url           string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			url:  fmt.Sprintf("/authors/%d", author.ID),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().DeleteAuthorById(gomock.Any(), author.ID).Times(1).Return(author, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAuthor(t, recorder.Body, author)
			},
		},
		{
			name: "InvalidURI",
			url:  "/authors/wrong",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().DeleteAuthorById(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			url:  "/authors/0",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().DeleteAuthorById(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalError",
			url:  fmt.Sprintf("/authors/%d", author.ID),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().DeleteAuthorById(gomock.Any(), author.ID).Times(1).Return(db.Author{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
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

			request, err := http.NewRequest(http.MethodDelete, tc.url, bytes.NewBuffer([]byte{}))
			require.NoError(t, err)

			addAuthorization(t, request, server.tokenMaker, authorizationTypeBearer, user.Email, time.Minute)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestUpdateAuthorAPI(t *testing.T) {
	user, _ := randomUser(t)
	author := randomAuthor(t, user.Email)

	updateAuthorParams := db.UpdateAuthorByIdParams{
		ID:         author.ID,
		AuthorName: util.RandomString(6),
		Website:    util.RandomString(10),
		Instagram:  util.RandomString(10),
		Youtube:    util.RandomString(10),
	}

	updatedAuthor := db.Author{
		ID:         author.ID,
		AuthorName: updateAuthorParams.AuthorName,
		Website:    updateAuthorParams.Website,
		Instagram:  updateAuthorParams.Instagram,
		Youtube:    updateAuthorParams.Youtube,
	}

	testCases := []struct {
		name          string
		url           string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"author_name": updateAuthorParams.AuthorName,
				"website":     updateAuthorParams.Website,
				"instagram":   updateAuthorParams.Instagram,
				"youtube":     updateAuthorParams.Youtube,
			},
			url: fmt.Sprintf("/authors/%d", author.ID),
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().UpdateAuthorById(gomock.Any(), updateAuthorParams).Times(1).Return(updatedAuthor, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAuthor(t, recorder.Body, updatedAuthor)
			},
		},
		{
			name: "InvalidBody",
			body: gin.H{
				"website":   updateAuthorParams.Website,
				"instagram": updateAuthorParams.Instagram,
				"youtube":   updateAuthorParams.Youtube,
			},
			url: fmt.Sprintf("/authors/%d", author.ID),
			buildStubs: func(store *mockdb.MockStore) {

				store.EXPECT().UpdateAuthorById(gomock.Any(), updateAuthorParams).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidID",
			body: gin.H{
				"website":   updateAuthorParams.Website,
				"instagram": updateAuthorParams.Instagram,
				"youtube":   updateAuthorParams.Youtube,
			},
			url: "/authors/0",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().UpdateAuthorById(gomock.Any(), updateAuthorParams).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"author_name": updateAuthorParams.AuthorName,
				"website":     updateAuthorParams.Website,
				"instagram":   updateAuthorParams.Instagram,
				"youtube":     updateAuthorParams.Youtube,
			},
			url: fmt.Sprintf("/authors/%d", author.ID),
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().UpdateAuthorById(gomock.Any(), updateAuthorParams).Times(1).Return(db.Author{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
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

			request, err := http.NewRequest(http.MethodPut, tc.url, bytes.NewReader(data))
			require.NoError(t, err)

			addAuthorization(t, request, server.tokenMaker, authorizationTypeBearer, user.Email, time.Minute)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func TestListAuthorsAPI(t *testing.T) {

}

func requireBodyMatchAuthor(t *testing.T, body *bytes.Buffer, author db.Author) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAuthor db.Author
	err = json.Unmarshal(data, &gotAuthor)
	require.NoError(t, err)
	require.Equal(t, author.AuthorName, gotAuthor.AuthorName)
	require.Equal(t, author.Website, gotAuthor.Website)
	require.Equal(t, author.Instagram, gotAuthor.Instagram)
	require.Equal(t, author.Youtube, gotAuthor.Youtube)
	require.Equal(t, author.UserCreated, gotAuthor.UserCreated)
}
