// Code generated by ent, DO NOT EDIT.

package builder

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the builder type in the database.
	Label = "builder"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURI holds the string denoting the uri field in the database.
	FieldURI = "uri"
	// EdgeSlsaAttestations holds the string denoting the slsa_attestations edge name in mutations.
	EdgeSlsaAttestations = "slsa_attestations"
	// Table holds the table name of the builder in the database.
	Table = "builders"
	// SlsaAttestationsTable is the table that holds the slsa_attestations relation/edge.
	SlsaAttestationsTable = "slsa_attestations"
	// SlsaAttestationsInverseTable is the table name for the SLSAAttestation entity.
	// It exists in this package in order to avoid circular dependency with the "slsaattestation" package.
	SlsaAttestationsInverseTable = "slsa_attestations"
	// SlsaAttestationsColumn is the table column denoting the slsa_attestations relation/edge.
	SlsaAttestationsColumn = "built_by_id"
)

// Columns holds all SQL columns for builder fields.
var Columns = []string{
	FieldID,
	FieldURI,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Builder queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByURI orders the results by the uri field.
func ByURI(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURI, opts...).ToFunc()
}

// BySlsaAttestationsCount orders the results by slsa_attestations count.
func BySlsaAttestationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSlsaAttestationsStep(), opts...)
	}
}

// BySlsaAttestations orders the results by slsa_attestations terms.
func BySlsaAttestations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSlsaAttestationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSlsaAttestationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SlsaAttestationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, SlsaAttestationsTable, SlsaAttestationsColumn),
	)
}