// Code generated by ent, DO NOT EDIT.

package itembatch

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldLTE(FieldID, id))
}

// ItemNumber applies equality check predicate on the "itemNumber" field. It's identical to ItemNumberEQ.
func ItemNumber(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldEQ(FieldItemNumber, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldEQ(FieldDescription, v))
}

// CompanyUUID applies equality check predicate on the "companyUUID" field. It's identical to CompanyUUIDEQ.
func CompanyUUID(v uuid.UUID) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldEQ(FieldCompanyUUID, v))
}

// ItemNumberEQ applies the EQ predicate on the "itemNumber" field.
func ItemNumberEQ(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldEQ(FieldItemNumber, v))
}

// ItemNumberNEQ applies the NEQ predicate on the "itemNumber" field.
func ItemNumberNEQ(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldNEQ(FieldItemNumber, v))
}

// ItemNumberIn applies the In predicate on the "itemNumber" field.
func ItemNumberIn(vs ...string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldIn(FieldItemNumber, vs...))
}

// ItemNumberNotIn applies the NotIn predicate on the "itemNumber" field.
func ItemNumberNotIn(vs ...string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldNotIn(FieldItemNumber, vs...))
}

// ItemNumberGT applies the GT predicate on the "itemNumber" field.
func ItemNumberGT(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldGT(FieldItemNumber, v))
}

// ItemNumberGTE applies the GTE predicate on the "itemNumber" field.
func ItemNumberGTE(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldGTE(FieldItemNumber, v))
}

// ItemNumberLT applies the LT predicate on the "itemNumber" field.
func ItemNumberLT(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldLT(FieldItemNumber, v))
}

// ItemNumberLTE applies the LTE predicate on the "itemNumber" field.
func ItemNumberLTE(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldLTE(FieldItemNumber, v))
}

// ItemNumberContains applies the Contains predicate on the "itemNumber" field.
func ItemNumberContains(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldContains(FieldItemNumber, v))
}

// ItemNumberHasPrefix applies the HasPrefix predicate on the "itemNumber" field.
func ItemNumberHasPrefix(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldHasPrefix(FieldItemNumber, v))
}

// ItemNumberHasSuffix applies the HasSuffix predicate on the "itemNumber" field.
func ItemNumberHasSuffix(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldHasSuffix(FieldItemNumber, v))
}

// ItemNumberEqualFold applies the EqualFold predicate on the "itemNumber" field.
func ItemNumberEqualFold(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldEqualFold(FieldItemNumber, v))
}

// ItemNumberContainsFold applies the ContainsFold predicate on the "itemNumber" field.
func ItemNumberContainsFold(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldContainsFold(FieldItemNumber, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldContainsFold(FieldDescription, v))
}

// CompanyUUIDEQ applies the EQ predicate on the "companyUUID" field.
func CompanyUUIDEQ(v uuid.UUID) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldEQ(FieldCompanyUUID, v))
}

// CompanyUUIDNEQ applies the NEQ predicate on the "companyUUID" field.
func CompanyUUIDNEQ(v uuid.UUID) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldNEQ(FieldCompanyUUID, v))
}

// CompanyUUIDIn applies the In predicate on the "companyUUID" field.
func CompanyUUIDIn(vs ...uuid.UUID) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldIn(FieldCompanyUUID, vs...))
}

// CompanyUUIDNotIn applies the NotIn predicate on the "companyUUID" field.
func CompanyUUIDNotIn(vs ...uuid.UUID) predicate.ItemBatch {
	return predicate.ItemBatch(sql.FieldNotIn(FieldCompanyUUID, vs...))
}

// HasCompany applies the HasEdge predicate on the "company" edge.
func HasCompany() predicate.ItemBatch {
	return predicate.ItemBatch(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCompanyWith applies the HasEdge predicate on the "company" edge with a given conditions (other predicates).
func HasCompanyWith(preds ...predicate.Company) predicate.ItemBatch {
	return predicate.ItemBatch(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CompanyInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CompanyTable, CompanyColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ItemBatch) predicate.ItemBatch {
	return predicate.ItemBatch(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ItemBatch) predicate.ItemBatch {
	return predicate.ItemBatch(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ItemBatch) predicate.ItemBatch {
	return predicate.ItemBatch(func(s *sql.Selector) {
		p(s.Not())
	})
}
