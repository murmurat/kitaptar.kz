package handler

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/kitaptar/entity"
	mock_service "github.com/murat96k/kitaptar.kz/internal/kitaptar/service/mock"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_createFilePath(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, filePath *api.FilePathRequest)
	userID := "e79e360e-cb68-40a1-911e-a8a75068ef79"
	filePathID := "7d02919a-12b7-44a2-9382-8ad8664076ca"

	testTable := []struct {
		name                 string
		inputBody            string
		inputFilePath        *api.FilePathRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"mobi":"test", "fb2":"test", "epub":"test", "docx":"test"}`,
			inputFilePath: &api.FilePathRequest{
				Mobi: "test",
				Fb2:  "test",
				Epub: "test",
				Docx: "test",
			},
			mockBehavior: func(s *mock_service.MockService, filePath *api.FilePathRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().CreateFilePath(gomock.Any(), filePath).Return(filePathID, nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"message":"7d02919a-12b7-44a2-9382-8ad8664076ca"}`,
		},
		{
			name:      "Wrong input (Missing mobi)",
			inputBody: `{"fb2":"test", "epub":"test", "docx":"test"}`,
			inputFilePath: &api.FilePathRequest{
				Fb2:  "test",
				Epub: "test",
				Docx: "test",
			},
			mockBehavior: func(s *mock_service.MockService, filePath *api.FilePathRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().CreateFilePath(gomock.Any(), filePath).Return(filePathID, nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"message":"7d02919a-12b7-44a2-9382-8ad8664076ca"}`,
		},
		{
			name:      "Service error",
			inputBody: `{"mobi":"test", "fb2":"test", "epub":"test", "docx":"test"}`,
			inputFilePath: &api.FilePathRequest{
				Mobi: "test",
				Fb2:  "test",
				Epub: "test",
				Docx: "test",
			},
			mockBehavior: func(s *mock_service.MockService, filePath *api.FilePathRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().CreateFilePath(gomock.Any(), filePath).Return("", errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			mockService := mock_service.NewMockService(controller)
			testCase.mockBehavior(mockService, testCase.inputFilePath)

			mockHandler := New(mockService)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/file_paths", bytes.NewBufferString(testCase.inputBody))
			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func TestHandler_updateFilePath(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, filePathID string, filePath *api.FilePathRequest)
	userID := "e79e360e-cb68-40a1-911e-a8a75068ef79"

	testTable := []struct {
		name                 string
		inputBody            string
		inputFilePathID      string
		inputFilePath        *api.FilePathRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:            "OK (Update all data)",
			inputBody:       `{"mobi":"test", "fb2":"test", "epub":"test", "docx":"test"}`,
			inputFilePathID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			inputFilePath: &api.FilePathRequest{
				Mobi: "test",
				Fb2:  "test",
				Epub: "test",
				Docx: "test",
			},
			mockBehavior: func(s *mock_service.MockService, filePathID string, filePath *api.FilePathRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().UpdateFilePath(gomock.Any(), filePathID, filePath).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"mobi":"test","fb2":"test","epub":"test","docx":"test"}`,
		},
		{
			name:            "Input with missing mobi",
			inputBody:       `{"fb2":"test", "epub":"test", "docx":"test"}`,
			inputFilePathID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			inputFilePath: &api.FilePathRequest{
				Fb2:  "test",
				Epub: "test",
				Docx: "test",
			},
			mockBehavior: func(s *mock_service.MockService, filePathID string, filePath *api.FilePathRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().UpdateFilePath(gomock.Any(), filePathID, filePath).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"mobi":"","fb2":"test","epub":"test","docx":"test"}`,
		},
		{
			name:            "Service error",
			inputBody:       `{"mobi":"test", "fb2":"test", "epub":"test", "docx":"test"}`,
			inputFilePathID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			inputFilePath: &api.FilePathRequest{
				Mobi: "test",
				Fb2:  "test",
				Epub: "test",
				Docx: "test",
			},
			mockBehavior: func(s *mock_service.MockService, filePathID string, filePath *api.FilePathRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().UpdateFilePath(gomock.Any(), filePathID, filePath).Return(errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
		{
			name:            "Empty file path id param",
			inputBody:       `{"mobi":"test", "fb2":"test", "epub":"test", "docx":"test"}`,
			inputFilePathID: "",
			inputFilePath: &api.FilePathRequest{
				Mobi: "test",
				Fb2:  "test",
				Epub: "test",
				Docx: "test",
			},
			mockBehavior: func(s *mock_service.MockService, filePathID string, filePath *api.FilePathRequest) {
			},
			expectedStatusCode:   404,
			expectedResponseBody: `404 page not found`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			mockService := mock_service.NewMockService(controller)
			testCase.mockBehavior(mockService, testCase.inputFilePathID, testCase.inputFilePath)

			mockHandler := New(mockService)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("PUT", fmt.Sprintf("/file_paths/%s", testCase.inputFilePathID), bytes.NewBufferString(testCase.inputBody))
			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func TestHandler_getFilePathById(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, filePathID string)
	userID := "e79e360e-cb68-40a1-911e-a8a75068ef79"

	testTable := []struct {
		name                 string
		inputID              string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:    "OK",
			inputID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			mockBehavior: func(s *mock_service.MockService, filePathID string) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().GetFilePathById(gomock.Any(), filePathID).Return(&entity.FilePath{
					Mobi: "test",
					Fb2:  "test",
					Epub: "test",
					Docx: "test",
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"file_path_id":"00000000-0000-0000-0000-000000000000","mobi":"test","fb2":"test","epub":"test","docx":"test","CreatedAt":"0001-01-01T00:00:00Z"}`,
		},
		{
			name:    "Service error",
			inputID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			mockBehavior: func(s *mock_service.MockService, filePathID string) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().GetFilePathById(gomock.Any(), filePathID).Return(nil, errors.New("something went wrong"))

			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			mockService := mock_service.NewMockService(controller)
			testCase.mockBehavior(mockService, testCase.inputID)

			mockHandler := New(mockService)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", strings.TrimSpace(fmt.Sprintf("/file_paths/%s", testCase.inputID)), bytes.NewBufferString(testCase.inputID))
			fmt.Println("Request URL: ", request.URL)
			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func TestHandler_deleteFilePath(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, filePathID string)
	userID := "e79e360e-cb68-40a1-911e-a8a75068ef79"

	testTable := []struct {
		name                 string
		inputID              string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:    "Delete file path by id (OK)",
			inputID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			mockBehavior: func(s *mock_service.MockService, filePathID string) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().DeleteFilePath(gomock.Any(), filePathID).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"File path deleted"}`,
		},
		{
			name:    "Empty file path id",
			inputID: "",
			mockBehavior: func(s *mock_service.MockService, filePathID string) {

			},
			expectedStatusCode:   404,
			expectedResponseBody: `404 page not found`,
		},
		{
			name:    "Service error",
			inputID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			mockBehavior: func(s *mock_service.MockService, filePathID string) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().DeleteFilePath(gomock.Any(), filePathID).Return(errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
		{
			name:    "Request by not authorizing user",
			inputID: "lalksdmvklasndvklsaklv",
			mockBehavior: func(s *mock_service.MockService, filePathID string) {
				s.EXPECT().VerifyToken("token").Return("", errors.New("error"))

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
			testCase.mockBehavior(mockService, testCase.inputID)

			mockHandler := New(mockService)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("DELETE", fmt.Sprintf("/file_paths/%s", testCase.inputID), bytes.NewBufferString(testCase.inputID))
			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func TestHandler_getAllFilePaths(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService)
	userID := "e79e360e-cb68-40a1-911e-a8a75068ef79"

	testTable := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockService) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().GetAllFilePaths(gomock.Any()).Return([]entity.FilePath{{
					Mobi: "test",
					Fb2:  "test",
					Epub: "test",
					Docx: "test",
				}}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `[{"file_path_id":"00000000-0000-0000-0000-000000000000","mobi":"test","fb2":"test","epub":"test","docx":"test","CreatedAt":"0001-01-01T00:00:00Z"}]`,
		},
		{
			name: "Service error",
			mockBehavior: func(s *mock_service.MockService) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().GetAllFilePaths(gomock.Any()).Return([]entity.FilePath{}, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
		{
			name: "Request by not authorizing user",
			mockBehavior: func(s *mock_service.MockService) {
				s.EXPECT().VerifyToken("token").Return("", errors.New("error"))
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
			testCase.mockBehavior(mockService)

			mockHandler := New(mockService)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", "/file_paths", nil)

			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}
