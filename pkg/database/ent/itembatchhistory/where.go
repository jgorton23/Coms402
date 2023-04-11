// Code generated by ent, DO NOT EDIT.

package itembatchhistory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/enthistory"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldLTE(FieldID, id))
}

// HistoryTime applies equality check predicate on the "history_time" field. It's identical to HistoryTimeEQ.
func HistoryTime(v time.Time) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEQ(FieldHistoryTime, v))
}

// Ref applies equality check predicate on the "ref" field. It's identical to RefEQ.
func Ref(v uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEQ(FieldRef, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEQ(FieldUpdatedBy, v))
}

// ItemNumber applies equality check predicate on the "itemNumber" field. It's identical to ItemNumberEQ.
func ItemNumber(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEQ(FieldItemNumber, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEQ(FieldDescription, v))
}

// CompanyUUID applies equality check predicate on the "companyUUID" field. It's identical to CompanyUUIDEQ.
func CompanyUUID(v uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEQ(FieldCompanyUUID, v))
}

// HistoryTimeEQ applies the EQ predicate on the "history_time" field.
func HistoryTimeEQ(v time.Time) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEQ(FieldHistoryTime, v))
}

// HistoryTimeNEQ applies the NEQ predicate on the "history_time" field.
func HistoryTimeNEQ(v time.Time) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNEQ(FieldHistoryTime, v))
}

// HistoryTimeIn applies the In predicate on the "history_time" field.
func HistoryTimeIn(vs ...time.Time) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldIn(FieldHistoryTime, vs...))
}

// HistoryTimeNotIn applies the NotIn predicate on the "history_time" field.
func HistoryTimeNotIn(vs ...time.Time) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNotIn(FieldHistoryTime, vs...))
}

// HistoryTimeGT applies the GT predicate on the "history_time" field.
func HistoryTimeGT(v time.Time) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldGT(FieldHistoryTime, v))
}

// HistoryTimeGTE applies the GTE predicate on the "history_time" field.
func HistoryTimeGTE(v time.Time) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldGTE(FieldHistoryTime, v))
}

// HistoryTimeLT applies the LT predicate on the "history_time" field.
func HistoryTimeLT(v time.Time) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldLT(FieldHistoryTime, v))
}

// HistoryTimeLTE applies the LTE predicate on the "history_time" field.
func HistoryTimeLTE(v time.Time) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldLTE(FieldHistoryTime, v))
}

// RefEQ applies the EQ predicate on the "ref" field.
func RefEQ(v uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEQ(FieldRef, v))
}

// RefNEQ applies the NEQ predicate on the "ref" field.
func RefNEQ(v uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNEQ(FieldRef, v))
}

// RefIn applies the In predicate on the "ref" field.
func RefIn(vs ...uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldIn(FieldRef, vs...))
}

// RefNotIn applies the NotIn predicate on the "ref" field.
func RefNotIn(vs ...uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNotIn(FieldRef, vs...))
}

// RefGT applies the GT predicate on the "ref" field.
func RefGT(v uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldGT(FieldRef, v))
}

// RefGTE applies the GTE predicate on the "ref" field.
func RefGTE(v uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldGTE(FieldRef, v))
}

// RefLT applies the LT predicate on the "ref" field.
func RefLT(v uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldLT(FieldRef, v))
}

// RefLTE applies the LTE predicate on the "ref" field.
func RefLTE(v uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldLTE(FieldRef, v))
}

// RefIsNil applies the IsNil predicate on the "ref" field.
func RefIsNil() predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldIsNull(FieldRef))
}

// RefNotNil applies the NotNil predicate on the "ref" field.
func RefNotNil() predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNotNull(FieldRef))
}

// OperationEQ applies the EQ predicate on the "operation" field.
func OperationEQ(v enthistory.OpType) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEQ(FieldOperation, v))
}

// OperationNEQ applies the NEQ predicate on the "operation" field.
func OperationNEQ(v enthistory.OpType) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNEQ(FieldOperation, v))
}

// OperationIn applies the In predicate on the "operation" field.
func OperationIn(vs ...enthistory.OpType) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldIn(FieldOperation, vs...))
}

// OperationNotIn applies the NotIn predicate on the "operation" field.
func OperationNotIn(vs ...enthistory.OpType) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNotIn(FieldOperation, vs...))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// ItemNumberEQ applies the EQ predicate on the "itemNumber" field.
func ItemNumberEQ(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEQ(FieldItemNumber, v))
}

// ItemNumberNEQ applies the NEQ predicate on the "itemNumber" field.
func ItemNumberNEQ(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNEQ(FieldItemNumber, v))
}

// ItemNumberIn applies the In predicate on the "itemNumber" field.
func ItemNumberIn(vs ...string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldIn(FieldItemNumber, vs...))
}

// ItemNumberNotIn applies the NotIn predicate on the "itemNumber" field.
func ItemNumberNotIn(vs ...string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNotIn(FieldItemNumber, vs...))
}

// ItemNumberGT applies the GT predicate on the "itemNumber" field.
func ItemNumberGT(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldGT(FieldItemNumber, v))
}

// ItemNumberGTE applies the GTE predicate on the "itemNumber" field.
func ItemNumberGTE(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldGTE(FieldItemNumber, v))
}

// ItemNumberLT applies the LT predicate on the "itemNumber" field.
func ItemNumberLT(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldLT(FieldItemNumber, v))
}

// ItemNumberLTE applies the LTE predicate on the "itemNumber" field.
func ItemNumberLTE(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldLTE(FieldItemNumber, v))
}

// ItemNumberContains applies the Contains predicate on the "itemNumber" field.
func ItemNumberContains(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldContains(FieldItemNumber, v))
}

// ItemNumberHasPrefix applies the HasPrefix predicate on the "itemNumber" field.
func ItemNumberHasPrefix(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldHasPrefix(FieldItemNumber, v))
}

// ItemNumberHasSuffix applies the HasSuffix predicate on the "itemNumber" field.
func ItemNumberHasSuffix(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldHasSuffix(FieldItemNumber, v))
}

// ItemNumberEqualFold applies the EqualFold predicate on the "itemNumber" field.
func ItemNumberEqualFold(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEqualFold(FieldItemNumber, v))
}

// ItemNumberContainsFold applies the ContainsFold predicate on the "itemNumber" field.
func ItemNumberContainsFold(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldContainsFold(FieldItemNumber, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldContainsFold(FieldDescription, v))
}

// CompanyUUIDEQ applies the EQ predicate on the "companyUUID" field.
func CompanyUUIDEQ(v uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldEQ(FieldCompanyUUID, v))
}

// CompanyUUIDNEQ applies the NEQ predicate on the "companyUUID" field.
func CompanyUUIDNEQ(v uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNEQ(FieldCompanyUUID, v))
}

// CompanyUUIDIn applies the In predicate on the "companyUUID" field.
func CompanyUUIDIn(vs ...uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldIn(FieldCompanyUUID, vs...))
}

// CompanyUUIDNotIn applies the NotIn predicate on the "companyUUID" field.
func CompanyUUIDNotIn(vs ...uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldNotIn(FieldCompanyUUID, vs...))
}

// CompanyUUIDGT applies the GT predicate on the "companyUUID" field.
func CompanyUUIDGT(v uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldGT(FieldCompanyUUID, v))
}

// CompanyUUIDGTE applies the GTE predicate on the "companyUUID" field.
func CompanyUUIDGTE(v uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldGTE(FieldCompanyUUID, v))
}

// CompanyUUIDLT applies the LT predicate on the "companyUUID" field.
func CompanyUUIDLT(v uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldLT(FieldCompanyUUID, v))
}

// CompanyUUIDLTE applies the LTE predicate on the "companyUUID" field.
func CompanyUUIDLTE(v uuid.UUID) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(sql.FieldLTE(FieldCompanyUUID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ItemBatchHistory) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ItemBatchHistory) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(func(s *sql.Selector) {
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
func Not(p predicate.ItemBatchHistory) predicate.ItemBatchHistory {
	return predicate.ItemBatchHistory(func(s *sql.Selector) {
		p(s.Not())
	})
}
