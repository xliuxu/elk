// Code generated by entc, DO NOT EDIT.

package http

import (
	json "encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	easyjson "github.com/mailru/easyjson"
	"github.com/masseelch/elk/internal/integration/pets/ent"
	"github.com/masseelch/elk/internal/integration/pets/ent/category"
	"github.com/masseelch/elk/internal/integration/pets/ent/owner"
	"github.com/masseelch/elk/internal/integration/pets/ent/pet"
	"github.com/masseelch/render"
	"go.uber.org/zap"
)

// Payload of a ent.Category create request.
type CategoryCreateRequest struct {
	Name *string `json:"name"`
	Pets []int   `json:"pets"`
}

// Create creates a new ent.Category and stores it in the database.
func (h CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d CategoryCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		render.BadRequest(w, r, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Category.Create()
	if d.Name != nil {
		b.SetName(*d.Name)
	}
	if d.Pets != nil {
		b.AddPetIDs(d.Pets...)
	}
	e, err := b.Save(r.Context())
	if err != nil {
		l.Error("error saving category", zap.Error(err))
		render.InternalServerError(w, r, nil)
		return
	}
	// Reload entry.
	q := h.client.Category.Query().Where(category.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Int("id", e.ID), zap.Error(err))
			render.NotFound(w, r, msg)
		default:
			l.Error("error fetching category from db", zap.Int("id", e.ID), zap.Error(err))
			render.InternalServerError(w, r, nil)
		}
		return
	}
	l.Info("category rendered", zap.Int("id", e.ID))
	easyjson.MarshalToHTTPResponseWriter(NewCategoryCreateResponse(e), w)
}

// Payload of a ent.Owner create request.
type OwnerCreateRequest struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
	Pets []int   `json:"pets"`
}

// Create creates a new ent.Owner and stores it in the database.
func (h OwnerHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d OwnerCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		render.BadRequest(w, r, "invalid json string")
		return
	}
	// Save the data.
	b := h.client.Owner.Create()
	if d.Name != nil {
		b.SetName(*d.Name)
	}
	if d.Age != nil {
		b.SetAge(*d.Age)
	}
	if d.Pets != nil {
		b.AddPetIDs(d.Pets...)
	}
	e, err := b.Save(r.Context())
	if err != nil {
		l.Error("error saving owner", zap.Error(err))
		render.InternalServerError(w, r, nil)
		return
	}
	// Reload entry.
	q := h.client.Owner.Query().Where(owner.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Int("id", e.ID), zap.Error(err))
			render.NotFound(w, r, msg)
		default:
			l.Error("error fetching owner from db", zap.Int("id", e.ID), zap.Error(err))
			render.InternalServerError(w, r, nil)
		}
		return
	}
	l.Info("owner rendered", zap.Int("id", e.ID))
	easyjson.MarshalToHTTPResponseWriter(NewOwnerCreateResponse(e), w)
}

// Payload of a ent.Pet create request.
type PetCreateRequest struct {
	Name     *string `json:"name" validate:"required"`
	Age      *int    `json:"age" validate:"required,gt=0"`
	Category []int   `json:"category"`
	Owner    *int    `json:"owner" validate:"required"`
	Friends  []int   `json:"friends"`
}

// Create creates a new ent.Pet and stores it in the database.
func (h PetHandler) Create(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Create"))
	// Get the post data.
	var d PetCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		l.Error("error decoding json", zap.Error(err))
		render.BadRequest(w, r, "invalid json string")
		return
	}
	// Validate the data.
	if err := h.validator.Struct(d); err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			l.Error("error validating request data", zap.Error(err))
			render.InternalServerError(w, r, nil)
			return
		}
		l.Info("validation failed", zap.Error(err))
		render.BadRequest(w, r, err)
		return
	}
	// Save the data.
	b := h.client.Pet.Create()
	if d.Name != nil {
		b.SetName(*d.Name)
	}
	if d.Age != nil {
		b.SetAge(*d.Age)
	}
	if d.Category != nil {
		b.AddCategoryIDs(d.Category...)
	}
	if d.Owner != nil {
		b.SetOwnerID(*d.Owner)
	}
	if d.Friends != nil {
		b.AddFriendIDs(d.Friends...)
	}
	e, err := b.Save(r.Context())
	if err != nil {
		l.Error("error saving pet", zap.Error(err))
		render.InternalServerError(w, r, nil)
		return
	}
	// Reload entry.
	q := h.client.Pet.Query().Where(pet.ID(e.ID))
	e, err = q.Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Int("id", e.ID), zap.Error(err))
			render.NotFound(w, r, msg)
		default:
			l.Error("error fetching pet from db", zap.Int("id", e.ID), zap.Error(err))
			render.InternalServerError(w, r, nil)
		}
		return
	}
	l.Info("pet rendered", zap.Int("id", e.ID))
	easyjson.MarshalToHTTPResponseWriter(NewPetCreateResponse(e), w)
}
