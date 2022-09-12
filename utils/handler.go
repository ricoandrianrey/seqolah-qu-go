package utils

import (
	"encoding/json"
	"net/http"
	"seqolah-qu/types"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

func HandlerWrapper(h types.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		params := func(key string) string {
			value := chi.URLParam(r, key)
			return value
		}
		body := func(data interface{}) {
			json.NewDecoder(r.Body).Decode(data)
		}
		query := func(data interface{}) {
			decoder.Decode(data, r.URL.Query())
		}

		Request := types.Request{
			Params: params,
			Body:   body,
			Query:  query,
		}

		response, _ := h(Request, ctx)

		b, _ := json.Marshal(response)
		w.Write(b)
	}
}
