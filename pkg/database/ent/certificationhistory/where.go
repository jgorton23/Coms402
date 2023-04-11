// Code generated by ent, DO NOT EDIT.

package certificationhistory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/predicate"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/enthistory"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldLTE(FieldID, id))
}

// HistoryTime applies equality check predicate on the "history_time" field. It's identical to HistoryTimeEQ.
func HistoryTime(v time.Time) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldHistoryTime, v))
}

// Ref applies equality check predicate on the "ref" field. It's identical to RefEQ.
func Ref(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldRef, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldUpdatedBy, v))
}

// PrimaryAttribute applies equality check predicate on the "primaryAttribute" field. It's identical to PrimaryAttributeEQ.
func PrimaryAttribute(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldPrimaryAttribute, v))
}

// CompanyUUID applies equality check predicate on the "companyUUID" field. It's identical to CompanyUUIDEQ.
func CompanyUUID(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldCompanyUUID, v))
}

// ItemBatchUUID applies equality check predicate on the "itemBatchUUID" field. It's identical to ItemBatchUUIDEQ.
func ItemBatchUUID(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldItemBatchUUID, v))
}

// ImageUUID applies equality check predicate on the "imageUUID" field. It's identical to ImageUUIDEQ.
func ImageUUID(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldImageUUID, v))
}

// HistoryTimeEQ applies the EQ predicate on the "history_time" field.
func HistoryTimeEQ(v time.Time) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldHistoryTime, v))
}

// HistoryTimeNEQ applies the NEQ predicate on the "history_time" field.
func HistoryTimeNEQ(v time.Time) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNEQ(FieldHistoryTime, v))
}

// HistoryTimeIn applies the In predicate on the "history_time" field.
func HistoryTimeIn(vs ...time.Time) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldIn(FieldHistoryTime, vs...))
}

// HistoryTimeNotIn applies the NotIn predicate on the "history_time" field.
func HistoryTimeNotIn(vs ...time.Time) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNotIn(FieldHistoryTime, vs...))
}

// HistoryTimeGT applies the GT predicate on the "history_time" field.
func HistoryTimeGT(v time.Time) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldGT(FieldHistoryTime, v))
}

// HistoryTimeGTE applies the GTE predicate on the "history_time" field.
func HistoryTimeGTE(v time.Time) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldGTE(FieldHistoryTime, v))
}

// HistoryTimeLT applies the LT predicate on the "history_time" field.
func HistoryTimeLT(v time.Time) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldLT(FieldHistoryTime, v))
}

// HistoryTimeLTE applies the LTE predicate on the "history_time" field.
func HistoryTimeLTE(v time.Time) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldLTE(FieldHistoryTime, v))
}

// RefEQ applies the EQ predicate on the "ref" field.
func RefEQ(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldRef, v))
}

// RefNEQ applies the NEQ predicate on the "ref" field.
func RefNEQ(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNEQ(FieldRef, v))
}

// RefIn applies the In predicate on the "ref" field.
func RefIn(vs ...uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldIn(FieldRef, vs...))
}

// RefNotIn applies the NotIn predicate on the "ref" field.
func RefNotIn(vs ...uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNotIn(FieldRef, vs...))
}

// RefGT applies the GT predicate on the "ref" field.
func RefGT(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldGT(FieldRef, v))
}

// RefGTE applies the GTE predicate on the "ref" field.
func RefGTE(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldGTE(FieldRef, v))
}

// RefLT applies the LT predicate on the "ref" field.
func RefLT(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldLT(FieldRef, v))
}

// RefLTE applies the LTE predicate on the "ref" field.
func RefLTE(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldLTE(FieldRef, v))
}

// RefIsNil applies the IsNil predicate on the "ref" field.
func RefIsNil() predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldIsNull(FieldRef))
}

// RefNotNil applies the NotNil predicate on the "ref" field.
func RefNotNil() predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNotNull(FieldRef))
}

// OperationEQ applies the EQ predicate on the "operation" field.
func OperationEQ(v enthistory.OpType) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldOperation, v))
}

// OperationNEQ applies the NEQ predicate on the "operation" field.
func OperationNEQ(v enthistory.OpType) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNEQ(FieldOperation, v))
}

// OperationIn applies the In predicate on the "operation" field.
func OperationIn(vs ...enthistory.OpType) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldIn(FieldOperation, vs...))
}

// OperationNotIn applies the NotIn predicate on the "operation" field.
func OperationNotIn(vs ...enthistory.OpType) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNotIn(FieldOperation, vs...))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// PrimaryAttributeEQ applies the EQ predicate on the "primaryAttribute" field.
func PrimaryAttributeEQ(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldPrimaryAttribute, v))
}

// PrimaryAttributeNEQ applies the NEQ predicate on the "primaryAttribute" field.
func PrimaryAttributeNEQ(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNEQ(FieldPrimaryAttribute, v))
}

// PrimaryAttributeIn applies the In predicate on the "primaryAttribute" field.
func PrimaryAttributeIn(vs ...string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldIn(FieldPrimaryAttribute, vs...))
}

// PrimaryAttributeNotIn applies the NotIn predicate on the "primaryAttribute" field.
func PrimaryAttributeNotIn(vs ...string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNotIn(FieldPrimaryAttribute, vs...))
}

// PrimaryAttributeGT applies the GT predicate on the "primaryAttribute" field.
func PrimaryAttributeGT(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldGT(FieldPrimaryAttribute, v))
}

// PrimaryAttributeGTE applies the GTE predicate on the "primaryAttribute" field.
func PrimaryAttributeGTE(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldGTE(FieldPrimaryAttribute, v))
}

// PrimaryAttributeLT applies the LT predicate on the "primaryAttribute" field.
func PrimaryAttributeLT(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldLT(FieldPrimaryAttribute, v))
}

// PrimaryAttributeLTE applies the LTE predicate on the "primaryAttribute" field.
func PrimaryAttributeLTE(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldLTE(FieldPrimaryAttribute, v))
}

// PrimaryAttributeContains applies the Contains predicate on the "primaryAttribute" field.
func PrimaryAttributeContains(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldContains(FieldPrimaryAttribute, v))
}

// PrimaryAttributeHasPrefix applies the HasPrefix predicate on the "primaryAttribute" field.
func PrimaryAttributeHasPrefix(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldHasPrefix(FieldPrimaryAttribute, v))
}

// PrimaryAttributeHasSuffix applies the HasSuffix predicate on the "primaryAttribute" field.
func PrimaryAttributeHasSuffix(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldHasSuffix(FieldPrimaryAttribute, v))
}

// PrimaryAttributeEqualFold applies the EqualFold predicate on the "primaryAttribute" field.
func PrimaryAttributeEqualFold(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEqualFold(FieldPrimaryAttribute, v))
}

// PrimaryAttributeContainsFold applies the ContainsFold predicate on the "primaryAttribute" field.
func PrimaryAttributeContainsFold(v string) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldContainsFold(FieldPrimaryAttribute, v))
}

// CompanyUUIDEQ applies the EQ predicate on the "companyUUID" field.
func CompanyUUIDEQ(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldCompanyUUID, v))
}

// CompanyUUIDNEQ applies the NEQ predicate on the "companyUUID" field.
func CompanyUUIDNEQ(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNEQ(FieldCompanyUUID, v))
}

// CompanyUUIDIn applies the In predicate on the "companyUUID" field.
func CompanyUUIDIn(vs ...uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldIn(FieldCompanyUUID, vs...))
}

// CompanyUUIDNotIn applies the NotIn predicate on the "companyUUID" field.
func CompanyUUIDNotIn(vs ...uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNotIn(FieldCompanyUUID, vs...))
}

// CompanyUUIDGT applies the GT predicate on the "companyUUID" field.
func CompanyUUIDGT(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldGT(FieldCompanyUUID, v))
}

// CompanyUUIDGTE applies the GTE predicate on the "companyUUID" field.
func CompanyUUIDGTE(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldGTE(FieldCompanyUUID, v))
}

// CompanyUUIDLT applies the LT predicate on the "companyUUID" field.
func CompanyUUIDLT(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldLT(FieldCompanyUUID, v))
}

// CompanyUUIDLTE applies the LTE predicate on the "companyUUID" field.
func CompanyUUIDLTE(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldLTE(FieldCompanyUUID, v))
}

// ItemBatchUUIDEQ applies the EQ predicate on the "itemBatchUUID" field.
func ItemBatchUUIDEQ(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldItemBatchUUID, v))
}

// ItemBatchUUIDNEQ applies the NEQ predicate on the "itemBatchUUID" field.
func ItemBatchUUIDNEQ(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNEQ(FieldItemBatchUUID, v))
}

// ItemBatchUUIDIn applies the In predicate on the "itemBatchUUID" field.
func ItemBatchUUIDIn(vs ...uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldIn(FieldItemBatchUUID, vs...))
}

// ItemBatchUUIDNotIn applies the NotIn predicate on the "itemBatchUUID" field.
func ItemBatchUUIDNotIn(vs ...uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNotIn(FieldItemBatchUUID, vs...))
}

// ItemBatchUUIDGT applies the GT predicate on the "itemBatchUUID" field.
func ItemBatchUUIDGT(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldGT(FieldItemBatchUUID, v))
}

// ItemBatchUUIDGTE applies the GTE predicate on the "itemBatchUUID" field.
func ItemBatchUUIDGTE(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldGTE(FieldItemBatchUUID, v))
}

// ItemBatchUUIDLT applies the LT predicate on the "itemBatchUUID" field.
func ItemBatchUUIDLT(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldLT(FieldItemBatchUUID, v))
}

// ItemBatchUUIDLTE applies the LTE predicate on the "itemBatchUUID" field.
func ItemBatchUUIDLTE(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldLTE(FieldItemBatchUUID, v))
}

// ImageUUIDEQ applies the EQ predicate on the "imageUUID" field.
func ImageUUIDEQ(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldEQ(FieldImageUUID, v))
}

// ImageUUIDNEQ applies the NEQ predicate on the "imageUUID" field.
func ImageUUIDNEQ(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNEQ(FieldImageUUID, v))
}

// ImageUUIDIn applies the In predicate on the "imageUUID" field.
func ImageUUIDIn(vs ...uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldIn(FieldImageUUID, vs...))
}

// ImageUUIDNotIn applies the NotIn predicate on the "imageUUID" field.
func ImageUUIDNotIn(vs ...uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldNotIn(FieldImageUUID, vs...))
}

// ImageUUIDGT applies the GT predicate on the "imageUUID" field.
func ImageUUIDGT(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldGT(FieldImageUUID, v))
}

// ImageUUIDGTE applies the GTE predicate on the "imageUUID" field.
func ImageUUIDGTE(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldGTE(FieldImageUUID, v))
}

// ImageUUIDLT applies the LT predicate on the "imageUUID" field.
func ImageUUIDLT(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldLT(FieldImageUUID, v))
}

// ImageUUIDLTE applies the LTE predicate on the "imageUUID" field.
func ImageUUIDLTE(v uuid.UUID) predicate.CertificationHistory {
	return predicate.CertificationHistory(sql.FieldLTE(FieldImageUUID, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CertificationHistory) predicate.CertificationHistory {
	return predicate.CertificationHistory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CertificationHistory) predicate.CertificationHistory {
	return predicate.CertificationHistory(func(s *sql.Selector) {
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
func Not(p predicate.CertificationHistory) predicate.CertificationHistory {
	return predicate.CertificationHistory(func(s *sql.Selector) {
		p(s.Not())
	})
}