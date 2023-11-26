package httputils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/csyezheng/memcard/pkg/logging"
	"github.com/hashicorp/go-multierror"
	"net/http"
	"strconv"
)

// jsonErrTmpl is the template to use when returning a JSON error. It is
// rendered using Printf, not json.Encode, so values must be escaped by the
// caller.
const jsonErrTmpl = `{"error":"%s"}`

// jsonOKResp is the return value for empty data responses.
const jsonOKResp = `{"ok":true}`

type singleError struct {
	Error string `json:"error,omitempty"`
}

type multiError struct {
	Errors []string `json:"errors,omitempty"`
}

// JsonResponse as an HTTP response function that consumes data to be serialized to JSON.
//
// If the provided data is nil and the response code is a 200, the result will
// be `{"ok":true}`. If the code is not a 200, the response will be of the
// format `{"error":"<val>"}` where val is the JSON-escaped http.StatusText for
// the provided code.
//
// If serialized fails, a generic 500 JSON response is returned. In dev mode, the
// error is included in the payload. If flushing the buffer to the response
// fails, an error is logged, but no recovery is attempted.
func JsonResponse(w http.ResponseWriter, code int, data interface{}) {
	// Avoid marshaling nil data.
	if data == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)

		logger := logging.NewLoggerFromEnv()
		logger.Info(strconv.FormatInt(int64(code), 10))

		// Return an OK response.
		if code >= 200 && code < 300 {
			fmt.Fprint(w, jsonOKResp)
			return
		}

		fmt.Fprintf(w, jsonErrTmpl, http.StatusText(code))
		return
	}

	// Special-case handle multi-error.
	if typ, ok := data.(*multierror.Error); ok {
		errs := typ.WrappedErrors()
		msgs := make([]string, 0, len(errs))
		for _, err := range errs {
			msgs = append(msgs, err.Error())
		}
		data = &multiError{Errors: msgs}
	}

	// If the provided value was an error, marshall accordingly.
	if typ, ok := data.(error); ok {
		data = &singleError{Error: typ.Error()}
	}
	b := bytes.NewBuffer(make([]byte, 0, 1024))
	// Render into the renderer
	if err := json.NewEncoder(b).Encode(data); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, jsonErrTmpl, http.StatusText(http.StatusInternalServerError))
		return
	}

	// Rendering worked, flush to the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = b.WriteTo(w)
}
