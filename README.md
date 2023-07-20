# httputils

*httputils* is a library for testing `http.Handler` in unit tests without starting the actual web service.


```shell
go get github.com/jxsl13/httputils@latest
```


## Example test

```go

func TestRoundtrip(t *testing.T) {

	for _, ct := range pickle.ContentTypes {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if !strings.EqualFold(r.Method, http.MethodPost) {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			w.Header().Add("Content-Type", ct)
			w.WriteHeader(http.StatusOK)
			_, err := io.Copy(w, r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		})

		h := httputils.New(mux)

		var (
			body   = "peter porker"
			result = ""
		)
		err := h.Post("/", body, &result, ct)
		require.NoError(t, err)
		require.Equal(t, body, result)

	}
}
```
