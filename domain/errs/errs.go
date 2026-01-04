package errs

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

const (
	ErrInvalidBody        = "4001"
	ErrInvalidCredentials = "4002"
	ErrUserNotFound       = "4003"
	ErrMissingParameter   = "4004"
	ErrRateNotFound       = "4005"
	ErrGroupAlreadyExists = "4006"

	ErrInternalServer  = "5000"
	ErrWhenSearchMovie = "5001"
	ErrTMDBConnection  = "5002"
	ErrTMDBAPIError    = "5003"
)

var errorMessages = map[string]string{
	ErrInvalidCredentials: "Credenciais inválidas",
	ErrUserNotFound:       "Usuário não encontrado",
	ErrInvalidBody:        "Dados inválidos",
	ErrMissingParameter:   "Parâmetro obrigatório ausente",
	ErrInternalServer:     "Ocorreu um erro inesperado",
	ErrWhenSearchMovie:    "Ocorreu um erro inesperado ao buscar filmes",
	ErrTMDBConnection:     "Falha na conexão com o serviço de filmes",
	ErrTMDBAPIError:       "Erro na API de filmes - verifique sua chave de acesso",
	ErrRateNotFound:       "Avaliação não encontrada",
	ErrGroupAlreadyExists: "Um grupo com esse nome já existe",
}

type DomainError struct {
	Code       string
	Message    string
	Err        error
	StackTrace string
}

func (e *DomainError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s", e.Code, e.Err.Error())
	}
	return e.Code
}

func (e *DomainError) GetDetails() map[string]interface{} {
	details := map[string]interface{}{
		"code":    e.Code,
		"message": e.Message,
	}

	if os.Getenv("ENV") == "dev" {
		if e.StackTrace != "" {
			details["stackTrace"] = e.getFilteredStackTrace()
		}
	}

	return details
}

func (e *DomainError) getFilteredStackTrace() []string {
	lines := strings.Split(e.StackTrace, "\n")
	var filtered []string

	for i, line := range lines {
		if strings.Contains(line, "crud/") {
			filtered = append(filtered, strings.TrimSpace(line))
			if i+1 < len(lines) {
				filtered = append(filtered, strings.TrimSpace(lines[i+1]))
			}
		}
	}

	return filtered
}

func New(code, message string, err error) *DomainError {
	return &DomainError{
		Code:       code,
		Message:    message,
		Err:        err,
		StackTrace: string(debug.Stack()),
	}
}

func NewWithCode(code string, err error) *DomainError {
	message := errorMessages[code]
	if message == "" {
		message = "Erro desconhecido"
	}

	return &DomainError{
		Code:       code,
		Message:    message,
		Err:        err,
		StackTrace: string(debug.Stack()),
	}
}
