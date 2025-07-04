package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/minulhasanrokan/students-api/internal/types"
	"github.com/minulhasanrokan/students-api/internal/utils/response"
)

func New() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		slog.Info("creating a student")

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {

			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("request Body is empty")))
			return
		}

		if err != nil {

			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err := validator.New().Struct(student); err != nil {

			validateError := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateError))

			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})

		w.Write([]byte("welcome api"))
	}
}
