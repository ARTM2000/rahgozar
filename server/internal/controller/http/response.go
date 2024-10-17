package http

import "github.com/gofiber/fiber/v2"

// ------------- commons -------------

type ResponseData struct {
	Data    map[string]interface{}
	Message string
	IsError bool
}

type Final struct {
	TrackId string                 `json:"track_id"`
	Error   bool                   `json:"error"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func FormatResponse(c *fiber.Ctx, resData ResponseData) Final {
	message := ""
	if resData.Message != "" {
		message = resData.Message
	}

	data := map[string]interface{}{}
	if resData.Data != nil {
		data = resData.Data
	}

	trackId := c.GetRespHeader(fiber.HeaderXRequestID)

	if trackId == "" {
		panic("track id should be defined in format response function")
	}

	return Final{
		TrackId: trackId,
		Error:   resData.IsError,
		Message: message,
		Data:    data,
	}
}

// -----------------------------------
