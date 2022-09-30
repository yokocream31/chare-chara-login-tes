package main

// import ( 
// 	"net/http"
// 	"testing"
// 	"strings"
// 	"net/http/httptest"
// 	"bytes"
// )

// func TestMyHandler(t *testing.T) {
//     reqBody := bytes.NewBufferString("request body")
//     req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", reqBody)

//     got := httptest.NewRecorder()
//     handler(got, req)

//     if got.Code != http.StatusOK {
//         t.Errorf("want OK, but %d", got.Code)
//     }
//     wantBody := "Hello"
//     if got := got.Body.String(); !strings.Contains(got, wantBody) {
//         t.Errorf("get %s : response body does not contain %s", got, wantBody)
//     }
// }