package handlers

import (
	"github.com/afirthes/ws-quiz/internal/errors"
	"github.com/afirthes/ws-quiz/template"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
)

type RestHandlers struct {
	log          *zap.SugaredLogger
	errorHandler *errors.ErrorHandler
}

func NewRestHandlers(log *zap.SugaredLogger, eh *errors.ErrorHandler) *RestHandlers {
	return &RestHandlers{
		log:          log,
		errorHandler: eh,
	}
}

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (rh *RestHandlers) Home(w http.ResponseWriter, r *http.Request) {

	quizes := []template.Quiz{
		{
			Title: "Викторина на знание английских слов",
			Id:    "1",
		},
		{
			Title: "Викторина на знание столиц",
			Id:    "2",
		},
		{
			Title: "Викторина на общеобразовательные знания (общая эрудиция)",
			Id:    "3",
		},
	}

	err := template.RootLayout("Quizzer page",
		template.TwoInRow(
			template.QuizList(quizes),
			template.UserSetup()),
		template.TwoInRow(
			template.UsersList(),
			template.Question()),
	).Render(r.Context(), w)

	if err != nil {
		rh.log.Error("Error rendering template", zap.Error(err))
		rh.errorHandler.InternalServerError(w, r, errors.ErrorRenderingTemplate)
	}
}

func (rh *RestHandlers) LoadImageMain(w http.ResponseWriter, r *http.Request) {

	err := template.LoadImageMain().Render(r.Context(), w)

	if err != nil {
		rh.log.Error("Error rendering template", zap.Error(err))
		rh.errorHandler.InternalServerError(w, r, errors.ErrorRenderingTemplate)
	}
}
