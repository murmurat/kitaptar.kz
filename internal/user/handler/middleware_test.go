package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mock_service "github.com/murat96k/kitaptar.kz/internal/user/service/mock"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"testing"
)

func TestHandler_authMiddleware(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer test_token",
			token:       "test_token",
			mockBehavior: func(s *mock_service.MockService, token string) {
				s.EXPECT().VerifyToken(token).Return("1", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "1",
		},
		{
			name:                 "Invalid Header Name",
			headerName:           "",
			headerValue:          "Bearer test_token",
			token:                "test_token",
			mockBehavior:         func(s *mock_service.MockService, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"authorization header is not set"}`,
		},
		{
			name:                 "Invalid header value",
			headerName:           "Authorization",
			headerValue:          "Bearr test_token",
			token:                "test_token",
			mockBehavior:         func(s *mock_service.MockService, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid header value"}`,
		},
		{
			name:                 "Empty token",
			headerName:           "Authorization",
			headerValue:          "Bearer ",
			token:                "test_token",
			mockBehavior:         func(s *mock_service.MockService, token string) {},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"empty token"}`,
		},
		{
			name:        "Parse error",
			headerName:  "Authorization",
			headerValue: "Bearer test_token",
			token:       "test_token",
			mockBehavior: func(s *mock_service.MockService, token string) {
				s.EXPECT().VerifyToken(token).Return("0", errors.New("invalid token"))
			},
			expectedStatusCode:   401,
			expectedResponseBody: `{"message":"invalid token"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			mockService := mock_service.NewMockService(controller)
			testCase.mockBehavior(mockService, testCase.token)

			mockHandler := New(mockService)

			// Init Endpoint
			router := gin.New()
			router.GET("/identity", mockHandler.authMiddleware(), func(c *gin.Context) {
				id, _ := c.Get(authUserID)
				c.String(200, "%s", id)
			})

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/identity", nil)
			request.Header.Set(testCase.headerName, testCase.headerValue)

			router.ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}
