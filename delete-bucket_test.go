package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteBucketHandler(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		s3                 S3Client
		expectedStatusCode int
		expectedBody       string
	}{
		"success": {
			s3:                 &S3ClientMock{},
			expectedStatusCode: http.StatusNoContent,
			expectedBody:       "",
		},
		"s3 error": {
			s3: &S3ClientMock{
				Err: errors.New("mocked S3 error"),
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedBody:       "error removing bucket\n",
		},
	}

	for _, tc := range tests {
		req, err := http.NewRequest("DELETE", "/api/buckets/bucketName", nil)
		assert.NoError(err)

		rr := httptest.NewRecorder()
		handler := DeleteBucketHandler(tc.s3)

		handler.ServeHTTP(rr, req)

		assert.Equal(tc.expectedStatusCode, rr.Code)
		assert.Equal(tc.expectedBody, rr.Body.String())
	}
}