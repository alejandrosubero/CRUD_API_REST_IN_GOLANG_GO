package commons

import "net/http"

func SendResponse(write http.ResponseWriter, status int, data []byte) {
	write.Header().Set("Content-type", "application/json")
	write.WriteHeader(status)
	write.Write(data)
}

func SendError(write http.ResponseWriter, status int) {
	data := []byte(`{}`)
	write.Header().Set("Content-type", "application/json")
	write.WriteHeader(status)
	write.Write(data)

}
