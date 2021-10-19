package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"goTestApi/mocks"
	"goTestApi/repository"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetItemReturns404ForMissingItem(t *testing.T) {
	itemNumber := 404

	// expected error response
	j, _ := json.Marshal(repository.NewHTTPError(sql.ErrNoRows, http.StatusNotFound, repository.ENOTFOUND))
	s := fmt.Sprintf(string(j) + "\n")

	// mock database calls
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	store := mocks.NewMockItemsI(mockCtrl)
	store.EXPECT().GetById(itemNumber).Times(1).Return(nil, sql.ErrNoRows)

	// use Echo context for http request
	e := echo.New()
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/", nil)
	c := e.NewContext(request, recorder)
	c.SetPath("/v1/items")
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("404")
	require.NoError(t, err)

	// get our Handler
	h := NewHandler(store)

	// assertions for getItem()
	// assert status code is 404
	if assert.NoError(t, h.getItem(c)) {
		assert.Equal(t, http.StatusNotFound, recorder.Code)
		assert.Equal(t, s, recorder.Body.String())
	}

}
