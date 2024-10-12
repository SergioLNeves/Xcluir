package handler

import (
	"github.com/SergioLNeves/Xcluir/domain"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"log/slog"
	"net/http"
)

type TweetHandler struct {
	tweetService domain.TweetService
}

func NewTweetHandler(tweetService domain.TweetService) *TweetHandler {
	return &TweetHandler{tweetService: tweetService}
}

func (h *TweetHandler) DeleteTweetsFromFile(echoContext echo.Context) error {
	logger := slog.With(
		slog.String("handler", "TweetHandler"),
		slog.String("func", "DeleteTweetsFromFile"),
	)

	file, err := echoContext.FormFile("file")
	if err != nil {
		logger.Error("error: File Not Found", slog.Any("err", err))
		return echoContext.JSON(http.StatusBadRequest, domain.NewAPIError(http.StatusBadRequest, err.Error()))
	}

	src, err := file.Open()
	if err != nil {
		logger.Error("error: Cannot Open File", slog.Any("err", err))
		return echoContext.JSON(http.StatusInternalServerError, domain.NewAPIError(http.StatusInternalServerError, err.Error()))
	}
	defer src.Close()

	fileRead, err := ioutil.ReadAll(src)
	if err != nil {
		logger.Error("error: Cannot Read File", slog.Any("err", err))
		return echoContext.JSON(http.StatusInternalServerError, domain.NewAPIError(http.StatusInternalServerError, err.Error()))
	}

	tempFile, err := ioutil.TempFile("", "*.json")
	if err != nil {
		logger.Error("error: Cannot Create Temp File", slog.Any("err", err))
		return echoContext.JSON(http.StatusInternalServerError, domain.NewAPIError(http.StatusInternalServerError, err.Error()))
	}
	defer tempFile.Close()

	_, err = tempFile.Write(fileRead)
	if err != nil {
		logger.Error("error: Cannot Write File", slog.Any("err", err))
		return echoContext.JSON(http.StatusInternalServerError, domain.NewAPIError(http.StatusInternalServerError, err.Error()))
	}

	err = h.tweetService.DeleteTweetsFromFile(tempFile.Name())
	if err != nil {
		logger.Error("error: Cannot Delete File", slog.Any("err", err))
		return echoContext.JSON(http.StatusInternalServerError, domain.NewAPIError(http.StatusInternalServerError, err.Error()))
	}

	return echoContext.JSON(http.StatusOK, tempFile.Name())

}

func (h *TweetHandler) DeleteTweetsFromPatch(echoContext echo.Context) error {
	logger := slog.With(
		slog.String("handler", "TweetHandler"),
		slog.String("func", "DeleteTweetsFromPatch"))

	filePatch := echoContext.Param("filepath")

	err := h.tweetService.DeleteTweetsFromFile(filePatch)
	if err != nil {
		logger.Error("error: Cannot Delete File", slog.Any("err", err))
		return echoContext.JSON(http.StatusInternalServerError, domain.NewAPIError(http.StatusInternalServerError, err.Error()))
	}

	return echoContext.JSON(http.StatusOK, "ok")

}
