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
type EventRequest struct {
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
// @Security BearerAuth
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
// @Param body body EventRequest true "Event payload"
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

	var event EventRequest
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



// GetOneEvent godoc
// @Summary Get a single event
// @Description Retrieve details of a specific event by its ID
// @Tags Event
// @Produce json
// @Param id path string true "Event ID"
// @Success 200 {object} EventResponse "Event details retrieved successfully"
// @Failure 400 {object} utils.ErrorResponseSchema "Invalid event ID"
// @Failure 404 {object} utils.ErrorResponseSchema "No event found"
// @Failure 500 {object} utils.ErrorResponseSchema "Internal server error"
// @Security BearerAuth
// @Router /events/{id} [get]
func (r *route)GetOneEvent(c *gin.Context){
	var event EventResponse
	eventId, ok := c.Params.Get("id")
	if !ok {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid event id",
		)
		
		return
	}
	ctx := c.Request.Context()
	row := r.Models.Event.GetOneEventById(ctx,eventId)
	err := row.Scan(&event.Id, &event.Name, &event.CreatorId, &event.CreatedAt)
	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to get event",
		)
		return
	}
	
	if err == pgx.ErrNoRows {
		utils.ErrorResponse(
			c,
			http.StatusNotFound,
			"no event found",
		)
		return
	}
	
	utils.SuccessResponse(
		c,
		http.StatusOK,
		"Event details",
		event,
	)
}



// DeleteEvent godoc
// @Summary Delete an event
// @Description Delete an existing event by its ID
// @Tags Event
// @Produce json
// @Param id path string true "Event ID"
// @Success 204 {object} utils.SuccessResponseSchema "Event deleted successfully"
// @Failure 400 {object} utils.ErrorResponseSchema "Invalid event ID"
// @Failure 500 {object} utils.ErrorResponseSchema "Failed to delete event"
// @Security BearerAuth
// @Router /events/{id} [delete]
func (r *route)DeleteEvent(c *gin.Context){
	eventId, ok := c.Params.Get("id")
	if !ok {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid event id",
		)
		
		return
	}
	ctx := c.Request.Context()
	value, err := r.Models.Event.DeleteEvent(ctx,eventId)
	if err != nil || !value.Delete() {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to delete event",
		)
		return
	}
	
	utils.SuccessResponse(
		c,
		http.StatusNoContent,
		"event deleted successfully",
	)
}




// UpdateEvent godoc
// @Summary Update an event
// @Description Update the name of an existing event by its ID
// @Tags Event
// @Accept json
// @Produce json
// @Param id path string true "Event ID"
// @Param body body EventRequest true "Update payload"
// @Success 200 {object} EventResponse "Event updated successfully"
// @Failure 400 {object} utils.ErrorResponseSchema "Invalid event ID or request body"
// @Failure 404 {object} utils.ErrorResponseSchema "Event not found"
// @Failure 500 {object} utils.ErrorResponseSchema "Internal server error"
// @Security BearerAuth
// @Router /events/{id} [put]
func (r *route)UpdateEvent(c *gin.Context){
	var event EventResponse
	eventId, ok := c.Params.Get("id")
	if !ok {
		utils.ErrorResponse(
			c,
			http.StatusBadRequest,
			"invalid event id",
		)
		
		return
	}
	
	var updateRequest EventRequest
	
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusBadGateway,
			"invalid request body",
		)
		return
	}
	
	ctx := c.Request.Context()
	row := r.Models.Event.UpdateEvent(ctx, updateRequest.Name, eventId)
	err :=  row.Scan(&event.Id,&event.Name,&event.CreatorId,&event.CreatedAt); 
	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to update event",
		)
		return 
	}
	
	if err == pgx.ErrNoRows {
		utils.ErrorResponse(
			c,
			http.StatusNotFound,
			"no event found",
		)
		return
	}
	
	utils.SuccessResponse(
		c,
		http.StatusOK,
		"event updated successfully",
		event,
	)
}
