// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/masseelch/elk/internal/integration/pets/ent/playgroup"
)

// PlayGroup is the model entity for the PlayGroup schema.
type PlayGroup struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Weekday holds the value of the "weekday" field.
	Weekday playgroup.Weekday `json:"weekday,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlayGroupQuery when eager-loading is set.
	Edges PlayGroupEdges `json:"edges"`
}

// PlayGroupEdges holds the relations/edges for other nodes in the graph.
type PlayGroupEdges struct {
	// Participants holds the value of the participants edge.
	Participants []*Pet `json:"participants,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ParticipantsOrErr returns the Participants value or an error if the edge
// was not loaded in eager-loading.
func (e PlayGroupEdges) ParticipantsOrErr() ([]*Pet, error) {
	if e.loadedTypes[0] {
		return e.Participants, nil
	}
	return nil, &NotLoadedError{edge: "participants"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PlayGroup) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case playgroup.FieldID:
			values[i] = new(sql.NullInt64)
		case playgroup.FieldTitle, playgroup.FieldDescription, playgroup.FieldWeekday:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PlayGroup", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PlayGroup fields.
func (pg *PlayGroup) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case playgroup.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pg.ID = int(value.Int64)
		case playgroup.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pg.Title = value.String
			}
		case playgroup.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				pg.Description = value.String
			}
		case playgroup.FieldWeekday:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field weekday", values[i])
			} else if value.Valid {
				pg.Weekday = playgroup.Weekday(value.String)
			}
		}
	}
	return nil
}

// QueryParticipants queries the "participants" edge of the PlayGroup entity.
func (pg *PlayGroup) QueryParticipants() *PetQuery {
	return (&PlayGroupClient{config: pg.config}).QueryParticipants(pg)
}

// Update returns a builder for updating this PlayGroup.
// Note that you need to call PlayGroup.Unwrap() before calling this method if this PlayGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (pg *PlayGroup) Update() *PlayGroupUpdateOne {
	return (&PlayGroupClient{config: pg.config}).UpdateOne(pg)
}

// Unwrap unwraps the PlayGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pg *PlayGroup) Unwrap() *PlayGroup {
	tx, ok := pg.config.driver.(*txDriver)
	if !ok {
		panic("ent: PlayGroup is not a transactional entity")
	}
	pg.config.driver = tx.drv
	return pg
}

// String implements the fmt.Stringer.
func (pg *PlayGroup) String() string {
	var builder strings.Builder
	builder.WriteString("PlayGroup(")
	builder.WriteString(fmt.Sprintf("id=%v", pg.ID))
	builder.WriteString(", title=")
	builder.WriteString(pg.Title)
	builder.WriteString(", description=")
	builder.WriteString(pg.Description)
	builder.WriteString(", weekday=")
	builder.WriteString(fmt.Sprintf("%v", pg.Weekday))
	builder.WriteByte(')')
	return builder.String()
}

// PlayGroups is a parsable slice of PlayGroup.
type PlayGroups []*PlayGroup

func (pg PlayGroups) config(cfg config) {
	for _i := range pg {
		pg[_i].config = cfg
	}
}
