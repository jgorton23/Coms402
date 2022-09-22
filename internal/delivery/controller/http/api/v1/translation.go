package v1

// import (
// 	"net/http"

// 	"github.com/go-chi/chi/v5"

// 	"iseage/bank/internal/entity"
// 	"iseage/bank/internal/usecase"
// 	"iseage/bank/pkg/logger"
// )

// type translationRoutes struct {
// 	t usecase.Translation
// 	l logger.Interface
// }

// func newTranslationRoutes(handler chi.Router, t usecase.Translation, l logger.Interface) {
// 	routes := &translationRoutes{t, l}

// 	handler.Route("translation", func(r chi.Router) {
// 		r.Get("/history", routes.history)
// 		r.Post("/do-translate", routes.doTranslate)
// 	})

// }

// type historyResponse struct {
// 	History []entity.Translation `json:"history"`
// }

// // @Summary     Show history
// // @Description Show all translation history
// // @ID          history
// // @Tags  	    translation
// // @Accept      json
// // @Produce     json
// // @Success     200 {object} historyResponse
// // @Failure     500 {object} response
// // @Router      /translation/history [get]
// func (tr *translationRoutes) history(w http.ResponseWriter, r *http.Request) {
// 	translations, err := tr.t.History(r.Context())
// 	if err != nil {
// 		tr.l.Error(err, "http - v1 - history")
// 		jsonResponse(w, http.StatusInternalServerError, response{Error: "database problems"})
// 		return
// 	}

// 	jsonResponse(w, http.StatusOK, historyResponse{translations})
// }

// type doTranslateRequest struct {
// 	Source      string `json:"source"       binding:"required"  example:"auto"`
// 	Destination string `json:"destination"  binding:"required"  example:"en"`
// 	Original    string `json:"original"     binding:"required"  example:"текст для перевода"`
// }

// // @Summary     Translate
// // @Description Translate a text
// // @ID          do-translate
// // @Tags  	    translation
// // @Accept      json
// // @Produce     json
// // @Param       request body doTranslateRequest true "Set up translation"
// // @Success     200 {object} entity.Translation
// // @Failure     400 {object} response
// // @Failure     500 {object} response
// // @Router      /translation/do-translate [post]
// func (tr *translationRoutes) doTranslate(w http.ResponseWriter, r *http.Request) {
// 	var request doTranslateRequest
// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		r.l.Error(err, "http - v1 - doTranslate")
// 		errorResponse(c, http.StatusBadRequest, "invalid request body")

// 		return
// 	}

// 	translation, err := r.t.Translate(
// 		c.Request.Context(),
// 		entity.Translation{
// 			Source:      request.Source,
// 			Destination: request.Destination,
// 			Original:    request.Original,
// 		},
// 	)
// 	if err != nil {
// 		r.l.Error(err, "http - v1 - doTranslate")
// 		errorResponse(c, http.StatusInternalServerError, "translation service problems")

// 		return
// 	}

// 	c.JSON(http.StatusOK, translation)
// }
