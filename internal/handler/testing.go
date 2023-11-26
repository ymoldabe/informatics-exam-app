package handler

import (
	"Zhantasov/internal/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *HandlerType) Testing(c *gin.Context) {
	grades, err := strconv.Atoi(c.Param("grades"))
	if err != nil {
		h.NewError(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	questions, err := h.service.GetQuestions(grades)
	if err != nil || questions == nil {
		h.NewError(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	c.HTML(http.StatusOK, "testing.html", gin.H{
		"Grades":    grades,
		"Questions": questions,
	})
}

func (h *HandlerType) TestingResult(c *gin.Context) {

	grades, err := strconv.Atoi(c.Param("grades"))
	if err != nil {
		h.NewError(c, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	questions, err := h.service.GetQuestions(grades)
	if err != nil || questions == nil {
		h.NewError(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	result := []models.Result{}
	resLen := 0

	for key, value := range questions {
		res := models.Result{}
		name := "answers-" + strconv.Itoa(key)
		fmt.Println("answers-0", name)
		answer := c.PostForm(name)
		ans, err := strconv.Atoi(value.Answer)
		if err != nil {
			h.NewError(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
		if answer == value.Answer {
			res.Res = true
			resLen += 1
		} else {

			res.Res = false
		}
		res.Answer = value.Options[ans]
		res.Number = value.Number
		res.Query = value.Query
		result = append(result, res)
	}
	lenght := len(questions)

	c.HTML(http.StatusOK, "result.html", gin.H{
		"Lenght": lenght,
		"ResLen": resLen,
		"Grades": grades,
		"Result": result,
	})
}
