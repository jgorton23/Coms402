// Code generated by ent, DO NOT EDIT.

package attribute

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldLTE(FieldID, id))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldKey, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldValue, v))
}

// CertUUID applies equality check predicate on the "certUUID" field. It's identical to CertUUIDEQ.
func CertUUID(v uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldCertUUID, v))
}

// AttributeTypeUUID applies equality check predicate on the "attributeTypeUUID" field. It's identical to AttributeTypeUUIDEQ.
func AttributeTypeUUID(v uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldAttributeTypeUUID, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.Attribute {
	return predicate.Attribute(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.Attribute {
	return predicate.Attribute(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldContainsFold(FieldKey, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.Attribute {
	return predicate.Attribute(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.Attribute {
	return predicate.Attribute(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldContainsFold(FieldValue, v))
}

// CertUUIDEQ applies the EQ predicate on the "certUUID" field.
func CertUUIDEQ(v uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldCertUUID, v))
}

// CertUUIDNEQ applies the NEQ predicate on the "certUUID" field.
func CertUUIDNEQ(v uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldNEQ(FieldCertUUID, v))
}

// CertUUIDIn applies the In predicate on the "certUUID" field.
func CertUUIDIn(vs ...uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldIn(FieldCertUUID, vs...))
}

// CertUUIDNotIn applies the NotIn predicate on the "certUUID" field.
func CertUUIDNotIn(vs ...uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldNotIn(FieldCertUUID, vs...))
}

// AttributeTypeUUIDEQ applies the EQ predicate on the "attributeTypeUUID" field.
func AttributeTypeUUIDEQ(v uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldAttributeTypeUUID, v))
}

// AttributeTypeUUIDNEQ applies the NEQ predicate on the "attributeTypeUUID" field.
func AttributeTypeUUIDNEQ(v uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldNEQ(FieldAttributeTypeUUID, v))
}

// AttributeTypeUUIDIn applies the In predicate on the "attributeTypeUUID" field.
func AttributeTypeUUIDIn(vs ...uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldIn(FieldAttributeTypeUUID, vs...))
}

// AttributeTypeUUIDNotIn applies the NotIn predicate on the "attributeTypeUUID" field.
func AttributeTypeUUIDNotIn(vs ...uuid.UUID) predicate.Attribute {
	return predicate.Attribute(sql.FieldNotIn(FieldAttributeTypeUUID, vs...))
}

// HasCertification applies the HasEdge predicate on the "certification" edge.
func HasCertification() predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CertificationTable, CertificationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCertificationWith applies the HasEdge predicate on the "certification" edge with a given conditions (other predicates).
func HasCertificationWith(preds ...predicate.Certification) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CertificationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CertificationTable, CertificationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAttributeType applies the HasEdge predicate on the "attributeType" edge.
func HasAttributeType() predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AttributeTypeTable, AttributeTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttributeTypeWith applies the HasEdge predicate on the "attributeType" edge with a given conditions (other predicates).
func HasAttributeTypeWith(preds ...predicate.AttributeType) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttributeTypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AttributeTypeTable, AttributeTypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Attribute) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Attribute) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
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
func Not(p predicate.Attribute) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		p(s.Not())
	})
}