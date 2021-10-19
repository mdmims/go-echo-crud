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
	j, _ := json.Marshal(repository.NotFound())
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

	// test call and verify HTTPError returned
	err = h.getItem(c)
	httpError, exists := err.(*echo.HTTPError)
	if exists {
		assert.Equal(t, http.StatusNotFound, httpError.Code)
	}
	respData, _ := json.Marshal(httpError.Message)
	respS := fmt.Sprintf(string(respData) + "\n")
	assert.Equal(t, s, respS)
}
