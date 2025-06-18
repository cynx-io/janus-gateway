package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"net/http"
	"reflect"
)

// autoInjectRoutes("/hermes.HermesUserService", publicRouter, h)
func autoInjectRoutes(basePath string, router *mux.Router, client interface{}) {
	clientVal := reflect.ValueOf(client)
	clientType := reflect.TypeOf(client)

	for i := 0; i < clientType.NumMethod(); i++ {
		method := clientType.Method(i)

		// Get method name, like "CheckUsername"
		route := basePath + "/" + method.Name

		// Create handler
		handler := func(method reflect.Method) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				// 1. Make a new zero value of the request type (*CheckUsernameRequest)
				if method.Type.NumIn() < 2 {
					http.Error(w, "Invalid method signature", http.StatusInternalServerError)
					return
				}

				reqType := method.Type.In(1).Elem() // *CheckUsernameRequest → CheckUsernameRequest
				req := reflect.New(reqType).Interface()

				// 2. Decode JSON into request
				if err := json.NewDecoder(r.Body).Decode(req); err != nil {
					http.Error(w, "Invalid request", http.StatusBadRequest)
					return
				}

				// 3. Call the gRPC method: client.CheckUsername(ctx, req)
				results := method.Func.Call([]reflect.Value{
					clientVal,
					reflect.ValueOf(r.Context()),
					reflect.ValueOf(req),
					reflect.ValueOf([]grpc.CallOption{}), // variadic opts
				})

				// 4. Handle response and error
				resp := results[0].Interface()
				errInterface := results[1].Interface()
				if errInterface != nil {
					err := errInterface.(error)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				err := HandleResponse(w, resp.(proto.Message))
				if err != nil {
					http.Error(w, "Failed to handle response", http.StatusInternalServerError)
					return
				}
			}
		}(method)

		// Register handler
		router.HandleFunc(route, handler).Methods("POST")
	}
}
