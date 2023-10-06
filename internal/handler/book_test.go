package handler

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/murat96k/kitaptar.kz/api"
	"github.com/murat96k/kitaptar.kz/internal/entity"
	mock_service "github.com/murat96k/kitaptar.kz/internal/service/mock"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_createBook(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, book *api.BookRequest)
	userID := "e79e360e-cb68-40a1-911e-a8a75068ef79"

	testTable := []struct {
		name                 string
		inputBody            string
		inputBook            *api.BookRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"annotation":"test_annotation", "name":"test_name", "genre":"test_genre", "image_path":"test_image_path"}`,
			inputBook: &api.BookRequest{
				Annotation: "test_annotation",
				Name:       "test_name",
				Genre:      "test_genre",
				ImagePath:  "test_image_path",
				FilePathId: uuid.Nil,
			},
			mockBehavior: func(s *mock_service.MockService, book *api.BookRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().CreateBook(gomock.Any(), book).Return(nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"author_id":"00000000-0000-0000-0000-000000000000","annotation":"test_annotation","name":"test_name","genre":"test_genre","image_path":"test_image_path","file_path_id":"00000000-0000-0000-0000-000000000000"}`,
		},
		{
			name:      "Wrong input (Missing annotation)",
			inputBody: `{"name":"test_name", "genre":"test_genre", "image_path":"test_image_path"}`,
			inputBook: &api.BookRequest{
				Name:       "test_name",
				Genre:      "test_genre",
				ImagePath:  "test_image_path",
				FilePathId: uuid.Nil,
			},
			mockBehavior: func(s *mock_service.MockService, book *api.BookRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().CreateBook(gomock.Any(), book).Return(nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"author_id":"00000000-0000-0000-0000-000000000000","annotation":"","name":"test_name","genre":"test_genre","image_path":"test_image_path","file_path_id":"00000000-0000-0000-0000-000000000000"}`,
		},
		{
			name:      "Service error",
			inputBody: `{"annotation":"test_annotation", "name":"test_name", "genre":"test_genre", "image_path":"test_image_path"}`,
			inputBook: &api.BookRequest{
				Annotation: "test_annotation",
				Name:       "test_name",
				Genre:      "test_genre",
				ImagePath:  "test_image_path",
				FilePathId: uuid.Nil,
			},
			mockBehavior: func(s *mock_service.MockService, book *api.BookRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().CreateBook(gomock.Any(), book).Return(errors.New("something went wrong"))
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
			testCase.mockBehavior(mockService, testCase.inputBook)

			mockHandler := New(mockService)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("POST", "/book/create", bytes.NewBufferString(testCase.inputBody))
			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func TestHandler_updateBook(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, bookID string, book *api.BookRequest)
	userID := "e79e360e-cb68-40a1-911e-a8a75068ef79"

	testTable := []struct {
		name                 string
		inputBody            string
		inputBookID          string
		inputBook            *api.BookRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK (Update all data)",
			inputBody:   `{"annotation":"test_annotation", "name":"test_name", "genre":"test_genre", "image_path":"test_image_path"}`,
			inputBookID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			inputBook: &api.BookRequest{
				Annotation: "test_annotation",
				Name:       "test_name",
				Genre:      "test_genre",
				ImagePath:  "test_image_path",
				FilePathId: uuid.Nil,
			},
			mockBehavior: func(s *mock_service.MockService, bookID string, book *api.BookRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().UpdateBook(gomock.Any(), bookID, book).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"author_id":"00000000-0000-0000-0000-000000000000","annotation":"test_annotation","name":"test_name","genre":"test_genre","image_path":"test_image_path","file_path_id":"00000000-0000-0000-0000-000000000000"}`,
		},
		{
			name:        "Input with missing annotation",
			inputBody:   `{"name":"test_name", "genre":"test_genre", "image_path":"test_image_path"}`,
			inputBookID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			inputBook: &api.BookRequest{
				Name:       "test_name",
				Genre:      "test_genre",
				ImagePath:  "test_image_path",
				FilePathId: uuid.Nil,
			},
			mockBehavior: func(s *mock_service.MockService, bookID string, book *api.BookRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().UpdateBook(gomock.Any(), bookID, book).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"author_id":"00000000-0000-0000-0000-000000000000","annotation":"","name":"test_name","genre":"test_genre","image_path":"test_image_path","file_path_id":"00000000-0000-0000-0000-000000000000"}`,
		},
		{
			name:        "Service error",
			inputBody:   `{"annotation":"test_annotation", "name":"test_name", "genre":"test_genre", "image_path":"test_image_path"}`,
			inputBookID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			inputBook: &api.BookRequest{
				Annotation: "test_annotation",
				Name:       "test_name",
				Genre:      "test_genre",
				ImagePath:  "test_image_path",
				FilePathId: uuid.Nil,
			},
			mockBehavior: func(s *mock_service.MockService, bookID string, book *api.BookRequest) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().UpdateBook(gomock.Any(), bookID, book).Return(errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
		{
			name:        "Empty book id param",
			inputBody:   `{"annotation":"test_annotation", "name":"test_name", "genre":"test_genre", "image_path":"test_image_path"}`,
			inputBookID: "",
			inputBook: &api.BookRequest{
				Annotation: "test_annotation",
				Name:       "test_name",
				Genre:      "test_genre",
				ImagePath:  "test_image_path",
				FilePathId: uuid.Nil,
			},
			mockBehavior: func(s *mock_service.MockService, bookID string, book *api.BookRequest) {

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
			testCase.mockBehavior(mockService, testCase.inputBookID, testCase.inputBook)

			mockHandler := New(mockService)

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("PUT", fmt.Sprintf("/book/update/%s", testCase.inputBookID), bytes.NewBufferString(testCase.inputBody))
			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func TestHandler_getBookById(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, bookID string)
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
			mockBehavior: func(s *mock_service.MockService, bookID string) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().GetBookById(gomock.Any(), bookID).Return(&entity.Book{
					AuthorId:   uuid.Nil,
					Annotation: "test_annotation",
					Name:       "test_name",
					Genre:      "test_genre",
					ImagePath:  "test_image_path",
					FilePathId: uuid.Nil,
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"book_id":"00000000-0000-0000-0000-000000000000","author_id":"00000000-0000-0000-0000-000000000000","Author":{"author_id":"00000000-0000-0000-0000-000000000000","firstname":"","lastname":"","image_path":"","about_author":"","CreatedAt":"0001-01-01T00:00:00Z"},"annotation":"test_annotation","name":"test_name","genre":"test_genre","image_path":"test_image_path","file_path_id":"00000000-0000-0000-0000-000000000000","FilePath":{"file_path_id":"00000000-0000-0000-0000-000000000000","mobi":"","fb2":"","epub":"","docx":"","CreatedAt":"0001-01-01T00:00:00Z"},"CreatedAt":"0001-01-01T00:00:00Z"}`,
		},
		{
			name:    "Empty book id",
			inputID: "",
			mockBehavior: func(s *mock_service.MockService, bookID string) {
			},
			expectedStatusCode:   404,
			expectedResponseBody: `404 page not found`,
		},
		{
			name:    "Service error",
			inputID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			mockBehavior: func(s *mock_service.MockService, bookID string) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().GetBookById(gomock.Any(), bookID).Return(nil, errors.New("something went wrong"))

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
			request := httptest.NewRequest("GET", strings.TrimSpace(fmt.Sprintf("/book/%s", testCase.inputID)), bytes.NewBufferString(testCase.inputID))
			fmt.Println("Request URL: ", request.URL)
			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func TestHandler_deleteBook(t *testing.T) {
	type mockBehavior = func(s *mock_service.MockService, bookID string)
	userID := "e79e360e-cb68-40a1-911e-a8a75068ef79"

	testTable := []struct {
		name                 string
		inputID              string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:    "Delete book by id (OK)",
			inputID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			mockBehavior: func(s *mock_service.MockService, bookID string) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().DeleteBook(gomock.Any(), bookID).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"Book deleted"}`,
		},
		{
			name:    "Empty book id",
			inputID: "",
			mockBehavior: func(s *mock_service.MockService, bookID string) {

			},
			expectedStatusCode:   404,
			expectedResponseBody: `404 page not found`,
		},
		{
			name:    "Service error",
			inputID: "9634d0d0-ba9f-4516-9459-83eb58ebdb86",
			mockBehavior: func(s *mock_service.MockService, bookID string) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().DeleteBook(gomock.Any(), bookID).Return(errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
		{
			name:    "Request by not authorizing user",
			inputID: "lalksdmvklasndvklsaklv",
			mockBehavior: func(s *mock_service.MockService, bookID string) {
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
			request := httptest.NewRequest("DELETE", fmt.Sprintf("/book/delete/%s", testCase.inputID), bytes.NewBufferString(testCase.inputID))
			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}

func TestHandler_getAllBooks(t *testing.T) {
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
				s.EXPECT().GetAllBooks(gomock.Any()).Return([]entity.Book{{
					AuthorId:   uuid.Nil,
					Annotation: "test_annotation",
					Name:       "test_name",
					Genre:      "test_genre",
					ImagePath:  "test_image_path",
					FilePathId: uuid.Nil,
				}}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `[{"book_id":"00000000-0000-0000-0000-000000000000","author_id":"00000000-0000-0000-0000-000000000000","Author":{"author_id":"00000000-0000-0000-0000-000000000000","firstname":"","lastname":"","image_path":"","about_author":"","CreatedAt":"0001-01-01T00:00:00Z"},"annotation":"test_annotation","name":"test_name","genre":"test_genre","image_path":"test_image_path","file_path_id":"00000000-0000-0000-0000-000000000000","FilePath":{"file_path_id":"00000000-0000-0000-0000-000000000000","mobi":"","fb2":"","epub":"","docx":"","CreatedAt":"0001-01-01T00:00:00Z"},"CreatedAt":"0001-01-01T00:00:00Z"}]`,
		},
		{
			name: "Service error",
			mockBehavior: func(s *mock_service.MockService) {
				s.EXPECT().VerifyToken("token").Return(userID, nil)
				s.EXPECT().GetAllBooks(gomock.Any()).Return([]entity.Book{}, errors.New("something went wrong"))
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
			request := httptest.NewRequest("GET", "/book/all", nil)

			request.Header.Set("Authorization", "Bearer token")

			mockHandler.InitRouter().ServeHTTP(recorder, request)

			require.Equal(t, testCase.expectedStatusCode, recorder.Code)
			require.Equal(t, testCase.expectedResponseBody, recorder.Body.String())

		})
	}
}
