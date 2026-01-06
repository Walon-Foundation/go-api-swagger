package routes

import (
	"net/http"
	"time"

	"github.com/Walon-Foundation/go-gin-doc/cmd/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// EventResponse represents a single event returned by GetEvent
type EventResponse struct {
	Id        string    `json:"id" example:"evt_12345"`
	Name      string    `json:"name" example:"Hackathon"`
	CreatorId string    `json:"creator_id" example:"usr_12345"`
	CreatedAt time.Time `json:"created_at" example:"2026-01-06T15:00:00Z"`
}

// CreateEventRequest represents the payload for creating a new event
type CreateEventRequest struct {
	Name string `json:"name" binding:"required,min=2" example:"Hackathon"`
}

// ===================== GetEvent =====================

// GetEvent godoc
// @Summary Get all events
// @Description Retrieve all events from the system
// @Tags Event
// @Produce json
// @Success 200 {object} []EventResponse "All events retrieved"
// @Success 200 {object} utils.SuccessResponseSchema "No events yet"
// @Failure 500 {object} utils.ErrorResponseSchema "Internal server error"
// @Router /events [get]
func (r *route) GetEvent(c *gin.Context) {
	var events []EventResponse
	ctx := c.Request.Context()

	rows, err := r.Models.Event.GetEvent(ctx)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	if err == pgx.ErrNoRows {
		utils.SuccessResponse(c, http.StatusOK, "No events yet")
		return
	}

	for rows.Next() {
		var eventSingle EventResponse
		if err := rows.Scan(&eventSingle.Id, &eventSingle.Name, &eventSingle.CreatorId, &eventSingle.CreatedAt); err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get events")
			return
		}
		events = append(events, eventSingle)
	}

	rows.Close()

	utils.SuccessResponse(c, http.StatusOK, "All events", events)
}



// CreateEvent godoc
// @Summary Create a new event
// @Description Create a new event in the system
// @Tags Event
// @Accept json
// @Produce json
// @Param body body CreateEventRequest true "Event payload"
// @Success 201 {object} utils.SuccessResponseSchema "Event created successfully"
// @Failure 400 {object} utils.ErrorResponseSchema "Invalid request body"
// @Failure 401 {object} utils.ErrorResponseSchema "Unauthorized user"
// @Failure 409 {object} utils.ErrorResponseSchema "Event already exists"
// @Failure 500 {object} utils.ErrorResponseSchema "Internal server error"
// @Security BearerAuth
// @Router /events [post]
func (r *route) CreateEvent(c *gin.Context) {
	userId, ok := c.Get("user")
	if !ok {
		utils.ErrorResponse(c, http.StatusUnauthorized, "invalid user")
		return
	}

	var event CreateEventRequest
	if err := c.ShouldBindJSON(&event); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	// check if the event with that name and creator id already exists
	var tempName string
	ctx := c.Request.Context()

	checkRow := r.Models.Event.GetEventByNameAndCreatorId(ctx, event.Name, userId.(string))
	if err := checkRow.Scan(&tempName); err != pgx.ErrNoRows {
		utils.ErrorResponse(c, http.StatusConflict, "Event already exists")
		return
	}

	id, _ := gonanoid.New(20)

	row := r.Models.Event.CreateEvent(ctx, id, event.Name, userId.(string))
	var newEventId string
	if err := row.Scan(&newEventId); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "failed to create event")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "event created", newEventId)
}
