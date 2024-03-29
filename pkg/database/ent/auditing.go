// Code generated by enthistory, DO NOT EDIT.
// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/attributehistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/attributetypehistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/certificationhistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/certificationtemplatehistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/companyhistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatchhistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/itembatchtoitembatchhistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/userhistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent/userstocompanyhistory"
	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/enthistory"
)

type Change struct {
	FieldName string
	Old       any
	New       any
}

func NewChange(fieldName string, old, new any) Change {
	return Change{
		FieldName: fieldName,
		Old:       old,
		New:       new,
	}
}

type HistoryDiff[T any] struct {
	Old     *T
	New     *T
	Changes []Change
}

var (
	MismatchedRefError    = errors.New("cannot take diff of histories with different Refs")
	IdenticalHistoryError = errors.New("cannot take diff of identical history")
)

func (ah *AttributeHistory) changes(new *AttributeHistory) []Change {
	var changes []Change
	if ah.Key != new.Key {
		changes = append(changes, NewChange(attributehistory.FieldKey, ah.Key, new.Key))
	}
	if ah.Value != new.Value {
		changes = append(changes, NewChange(attributehistory.FieldValue, ah.Value, new.Value))
	}
	if ah.CertUUID != new.CertUUID {
		changes = append(changes, NewChange(attributehistory.FieldCertUUID, ah.CertUUID, new.CertUUID))
	}
	if ah.AttributeTypeUUID != new.AttributeTypeUUID {
		changes = append(changes, NewChange(attributehistory.FieldAttributeTypeUUID, ah.AttributeTypeUUID, new.AttributeTypeUUID))
	}
	return changes
}

func (ah *AttributeHistory) Diff(history *AttributeHistory) (*HistoryDiff[AttributeHistory], error) {
	if ah.Ref != history.Ref {
		return nil, MismatchedRefError
	}

	ahUnix, historyUnix := ah.HistoryTime.Unix(), history.HistoryTime.Unix()
	ahOlder := ahUnix < historyUnix || (ahUnix == historyUnix && ah.ID.Time() < history.ID.Time())
	historyOlder := ahUnix > historyUnix || (ahUnix == historyUnix && ah.ID.Time() > history.ID.Time())

	if ahOlder {
		return &HistoryDiff[AttributeHistory]{
			Old:     ah,
			New:     history,
			Changes: ah.changes(history),
		}, nil
	} else if historyOlder {
		return &HistoryDiff[AttributeHistory]{
			Old:     history,
			New:     ah,
			Changes: history.changes(ah),
		}, nil
	}
	return nil, IdenticalHistoryError
}

func (ath *AttributeTypeHistory) changes(new *AttributeTypeHistory) []Change {
	var changes []Change
	if ath.Key != new.Key {
		changes = append(changes, NewChange(attributetypehistory.FieldKey, ath.Key, new.Key))
	}
	if ath.CompanyUUID != new.CompanyUUID {
		changes = append(changes, NewChange(attributetypehistory.FieldCompanyUUID, ath.CompanyUUID, new.CompanyUUID))
	}
	return changes
}

func (ath *AttributeTypeHistory) Diff(history *AttributeTypeHistory) (*HistoryDiff[AttributeTypeHistory], error) {
	if ath.Ref != history.Ref {
		return nil, MismatchedRefError
	}

	athUnix, historyUnix := ath.HistoryTime.Unix(), history.HistoryTime.Unix()
	athOlder := athUnix < historyUnix || (athUnix == historyUnix && ath.ID.Time() < history.ID.Time())
	historyOlder := athUnix > historyUnix || (athUnix == historyUnix && ath.ID.Time() > history.ID.Time())

	if athOlder {
		return &HistoryDiff[AttributeTypeHistory]{
			Old:     ath,
			New:     history,
			Changes: ath.changes(history),
		}, nil
	} else if historyOlder {
		return &HistoryDiff[AttributeTypeHistory]{
			Old:     history,
			New:     ath,
			Changes: history.changes(ath),
		}, nil
	}
	return nil, IdenticalHistoryError
}

func (ch *CertificationHistory) changes(new *CertificationHistory) []Change {
	var changes []Change
	if ch.PrimaryAttribute != new.PrimaryAttribute {
		changes = append(changes, NewChange(certificationhistory.FieldPrimaryAttribute, ch.PrimaryAttribute, new.PrimaryAttribute))
	}
	if ch.CompanyUUID != new.CompanyUUID {
		changes = append(changes, NewChange(certificationhistory.FieldCompanyUUID, ch.CompanyUUID, new.CompanyUUID))
	}
	if ch.ItemBatchUUID != new.ItemBatchUUID {
		changes = append(changes, NewChange(certificationhistory.FieldItemBatchUUID, ch.ItemBatchUUID, new.ItemBatchUUID))
	}
	if ch.ImageUUID != new.ImageUUID {
		changes = append(changes, NewChange(certificationhistory.FieldImageUUID, ch.ImageUUID, new.ImageUUID))
	}
	return changes
}

func (ch *CertificationHistory) Diff(history *CertificationHistory) (*HistoryDiff[CertificationHistory], error) {
	if ch.Ref != history.Ref {
		return nil, MismatchedRefError
	}

	chUnix, historyUnix := ch.HistoryTime.Unix(), history.HistoryTime.Unix()
	chOlder := chUnix < historyUnix || (chUnix == historyUnix && ch.ID.Time() < history.ID.Time())
	historyOlder := chUnix > historyUnix || (chUnix == historyUnix && ch.ID.Time() > history.ID.Time())

	if chOlder {
		return &HistoryDiff[CertificationHistory]{
			Old:     ch,
			New:     history,
			Changes: ch.changes(history),
		}, nil
	} else if historyOlder {
		return &HistoryDiff[CertificationHistory]{
			Old:     history,
			New:     ch,
			Changes: history.changes(ch),
		}, nil
	}
	return nil, IdenticalHistoryError
}

func (cth *CertificationTemplateHistory) changes(new *CertificationTemplateHistory) []Change {
	var changes []Change
	if cth.Description != new.Description {
		changes = append(changes, NewChange(certificationtemplatehistory.FieldDescription, cth.Description, new.Description))
	}
	if cth.CompanyUUID != new.CompanyUUID {
		changes = append(changes, NewChange(certificationtemplatehistory.FieldCompanyUUID, cth.CompanyUUID, new.CompanyUUID))
	}
	return changes
}

func (cth *CertificationTemplateHistory) Diff(history *CertificationTemplateHistory) (*HistoryDiff[CertificationTemplateHistory], error) {
	if cth.Ref != history.Ref {
		return nil, MismatchedRefError
	}

	cthUnix, historyUnix := cth.HistoryTime.Unix(), history.HistoryTime.Unix()
	cthOlder := cthUnix < historyUnix || (cthUnix == historyUnix && cth.ID.Time() < history.ID.Time())
	historyOlder := cthUnix > historyUnix || (cthUnix == historyUnix && cth.ID.Time() > history.ID.Time())

	if cthOlder {
		return &HistoryDiff[CertificationTemplateHistory]{
			Old:     cth,
			New:     history,
			Changes: cth.changes(history),
		}, nil
	} else if historyOlder {
		return &HistoryDiff[CertificationTemplateHistory]{
			Old:     history,
			New:     cth,
			Changes: history.changes(cth),
		}, nil
	}
	return nil, IdenticalHistoryError
}

func (ch *CompanyHistory) changes(new *CompanyHistory) []Change {
	var changes []Change
	if ch.Name != new.Name {
		changes = append(changes, NewChange(companyhistory.FieldName, ch.Name, new.Name))
	}
	return changes
}

func (ch *CompanyHistory) Diff(history *CompanyHistory) (*HistoryDiff[CompanyHistory], error) {
	if ch.Ref != history.Ref {
		return nil, MismatchedRefError
	}

	chUnix, historyUnix := ch.HistoryTime.Unix(), history.HistoryTime.Unix()
	chOlder := chUnix < historyUnix || (chUnix == historyUnix && ch.ID.Time() < history.ID.Time())
	historyOlder := chUnix > historyUnix || (chUnix == historyUnix && ch.ID.Time() > history.ID.Time())

	if chOlder {
		return &HistoryDiff[CompanyHistory]{
			Old:     ch,
			New:     history,
			Changes: ch.changes(history),
		}, nil
	} else if historyOlder {
		return &HistoryDiff[CompanyHistory]{
			Old:     history,
			New:     ch,
			Changes: history.changes(ch),
		}, nil
	}
	return nil, IdenticalHistoryError
}

func (ibh *ItemBatchHistory) changes(new *ItemBatchHistory) []Change {
	var changes []Change
	if ibh.ItemNumber != new.ItemNumber {
		changes = append(changes, NewChange(itembatchhistory.FieldItemNumber, ibh.ItemNumber, new.ItemNumber))
	}
	if ibh.Description != new.Description {
		changes = append(changes, NewChange(itembatchhistory.FieldDescription, ibh.Description, new.Description))
	}
	if ibh.CompanyUUID != new.CompanyUUID {
		changes = append(changes, NewChange(itembatchhistory.FieldCompanyUUID, ibh.CompanyUUID, new.CompanyUUID))
	}
	return changes
}

func (ibh *ItemBatchHistory) Diff(history *ItemBatchHistory) (*HistoryDiff[ItemBatchHistory], error) {
	if ibh.Ref != history.Ref {
		return nil, MismatchedRefError
	}

	ibhUnix, historyUnix := ibh.HistoryTime.Unix(), history.HistoryTime.Unix()
	ibhOlder := ibhUnix < historyUnix || (ibhUnix == historyUnix && ibh.ID.Time() < history.ID.Time())
	historyOlder := ibhUnix > historyUnix || (ibhUnix == historyUnix && ibh.ID.Time() > history.ID.Time())

	if ibhOlder {
		return &HistoryDiff[ItemBatchHistory]{
			Old:     ibh,
			New:     history,
			Changes: ibh.changes(history),
		}, nil
	} else if historyOlder {
		return &HistoryDiff[ItemBatchHistory]{
			Old:     history,
			New:     ibh,
			Changes: history.changes(ibh),
		}, nil
	}
	return nil, IdenticalHistoryError
}

func (ibtibh *ItemBatchToItemBatchHistory) changes(new *ItemBatchToItemBatchHistory) []Change {
	var changes []Change
	if ibtibh.ChildUUID != new.ChildUUID {
		changes = append(changes, NewChange(itembatchtoitembatchhistory.FieldChildUUID, ibtibh.ChildUUID, new.ChildUUID))
	}
	if ibtibh.ParentUUID != new.ParentUUID {
		changes = append(changes, NewChange(itembatchtoitembatchhistory.FieldParentUUID, ibtibh.ParentUUID, new.ParentUUID))
	}
	return changes
}

func (ibtibh *ItemBatchToItemBatchHistory) Diff(history *ItemBatchToItemBatchHistory) (*HistoryDiff[ItemBatchToItemBatchHistory], error) {
	if ibtibh.Ref != history.Ref {
		return nil, MismatchedRefError
	}

	ibtibhUnix, historyUnix := ibtibh.HistoryTime.Unix(), history.HistoryTime.Unix()
	ibtibhOlder := ibtibhUnix < historyUnix || (ibtibhUnix == historyUnix && ibtibh.ID.Time() < history.ID.Time())
	historyOlder := ibtibhUnix > historyUnix || (ibtibhUnix == historyUnix && ibtibh.ID.Time() > history.ID.Time())

	if ibtibhOlder {
		return &HistoryDiff[ItemBatchToItemBatchHistory]{
			Old:     ibtibh,
			New:     history,
			Changes: ibtibh.changes(history),
		}, nil
	} else if historyOlder {
		return &HistoryDiff[ItemBatchToItemBatchHistory]{
			Old:     history,
			New:     ibtibh,
			Changes: history.changes(ibtibh),
		}, nil
	}
	return nil, IdenticalHistoryError
}

func (uh *UserHistory) changes(new *UserHistory) []Change {
	var changes []Change
	if uh.Email != new.Email {
		changes = append(changes, NewChange(userhistory.FieldEmail, uh.Email, new.Email))
	}
	if !uh.CreatedAt.Equal(new.CreatedAt) {
		changes = append(changes, NewChange(userhistory.FieldCreatedAt, uh.CreatedAt, new.CreatedAt))
	}
	if !uh.UpdatedAt.Equal(new.UpdatedAt) {
		changes = append(changes, NewChange(userhistory.FieldUpdatedAt, uh.UpdatedAt, new.UpdatedAt))
	}
	if uh.PasswordHash != new.PasswordHash {
		changes = append(changes, NewChange(userhistory.FieldPasswordHash, uh.PasswordHash, new.PasswordHash))
	}
	if uh.AttemptCount != new.AttemptCount {
		changes = append(changes, NewChange(userhistory.FieldAttemptCount, uh.AttemptCount, new.AttemptCount))
	}
	if !uh.LastAttempt.Equal(new.LastAttempt) {
		changes = append(changes, NewChange(userhistory.FieldLastAttempt, uh.LastAttempt, new.LastAttempt))
	}
	if !uh.Locked.Equal(new.Locked) {
		changes = append(changes, NewChange(userhistory.FieldLocked, uh.Locked, new.Locked))
	}
	if uh.Role != new.Role {
		changes = append(changes, NewChange(userhistory.FieldRole, uh.Role, new.Role))
	}
	return changes
}

func (uh *UserHistory) Diff(history *UserHistory) (*HistoryDiff[UserHistory], error) {
	if uh.Ref != history.Ref {
		return nil, MismatchedRefError
	}

	uhUnix, historyUnix := uh.HistoryTime.Unix(), history.HistoryTime.Unix()
	uhOlder := uhUnix < historyUnix || (uhUnix == historyUnix && uh.ID.Time() < history.ID.Time())
	historyOlder := uhUnix > historyUnix || (uhUnix == historyUnix && uh.ID.Time() > history.ID.Time())

	if uhOlder {
		return &HistoryDiff[UserHistory]{
			Old:     uh,
			New:     history,
			Changes: uh.changes(history),
		}, nil
	} else if historyOlder {
		return &HistoryDiff[UserHistory]{
			Old:     history,
			New:     uh,
			Changes: history.changes(uh),
		}, nil
	}
	return nil, IdenticalHistoryError
}

func (utch *UsersToCompanyHistory) changes(new *UsersToCompanyHistory) []Change {
	var changes []Change
	if utch.CompanyUUID != new.CompanyUUID {
		changes = append(changes, NewChange(userstocompanyhistory.FieldCompanyUUID, utch.CompanyUUID, new.CompanyUUID))
	}
	if utch.UserUUID != new.UserUUID {
		changes = append(changes, NewChange(userstocompanyhistory.FieldUserUUID, utch.UserUUID, new.UserUUID))
	}
	if utch.RoleType != new.RoleType {
		changes = append(changes, NewChange(userstocompanyhistory.FieldRoleType, utch.RoleType, new.RoleType))
	}
	if utch.Approved != new.Approved {
		changes = append(changes, NewChange(userstocompanyhistory.FieldApproved, utch.Approved, new.Approved))
	}
	return changes
}

func (utch *UsersToCompanyHistory) Diff(history *UsersToCompanyHistory) (*HistoryDiff[UsersToCompanyHistory], error) {
	if utch.Ref != history.Ref {
		return nil, MismatchedRefError
	}

	utchUnix, historyUnix := utch.HistoryTime.Unix(), history.HistoryTime.Unix()
	utchOlder := utchUnix < historyUnix || (utchUnix == historyUnix && utch.ID.Time() < history.ID.Time())
	historyOlder := utchUnix > historyUnix || (utchUnix == historyUnix && utch.ID.Time() > history.ID.Time())

	if utchOlder {
		return &HistoryDiff[UsersToCompanyHistory]{
			Old:     utch,
			New:     history,
			Changes: utch.changes(history),
		}, nil
	} else if historyOlder {
		return &HistoryDiff[UsersToCompanyHistory]{
			Old:     history,
			New:     utch,
			Changes: history.changes(utch),
		}, nil
	}
	return nil, IdenticalHistoryError
}

func (c Change) String(op enthistory.OpType) string {
	var newstr, oldstr string
	if c.New != nil {
		val, err := json.Marshal(c.New)
		if err != nil {
			newstr = fmt.Sprintf("%#v", c.New)
		} else {
			newstr = string(val)
		}
	}
	if c.Old != nil {
		val, err := json.Marshal(c.Old)
		if err != nil {
			oldstr = fmt.Sprintf("%#v", c.Old)
		} else {
			oldstr = string(val)
		}
	}
	switch op {
	case enthistory.OpTypeInsert:
		return fmt.Sprintf("%s: %#s", c.FieldName, newstr)
	case enthistory.OpTypeDelete:
		return fmt.Sprintf("%s: %#s", c.FieldName, oldstr)
	default:
		return fmt.Sprintf("%s: %#s -> %#s", c.FieldName, oldstr, newstr)
	}
}

func (c *Client) Audit(ctx context.Context) ([][]string, error) {
	records := [][]string{
		{"Table", "Ref Id", "History Time", "Operation", "Changes", "Updated By"},
	}
	var record [][]string
	var err error

	record, err = auditAttributeHistory(ctx, c.config)
	if err != nil {
		return nil, err
	}
	records = append(records, record...)

	record, err = auditAttributeTypeHistory(ctx, c.config)
	if err != nil {
		return nil, err
	}
	records = append(records, record...)

	record, err = auditCertificationHistory(ctx, c.config)
	if err != nil {
		return nil, err
	}
	records = append(records, record...)

	record, err = auditCertificationTemplateHistory(ctx, c.config)
	if err != nil {
		return nil, err
	}
	records = append(records, record...)

	record, err = auditCompanyHistory(ctx, c.config)
	if err != nil {
		return nil, err
	}
	records = append(records, record...)

	record, err = auditItemBatchHistory(ctx, c.config)
	if err != nil {
		return nil, err
	}
	records = append(records, record...)

	record, err = auditItemBatchToItemBatchHistory(ctx, c.config)
	if err != nil {
		return nil, err
	}
	records = append(records, record...)

	record, err = auditUserHistory(ctx, c.config)
	if err != nil {
		return nil, err
	}
	records = append(records, record...)

	record, err = auditUsersToCompanyHistory(ctx, c.config)
	if err != nil {
		return nil, err
	}
	records = append(records, record...)

	return records, nil
}

type record struct {
	Table       string
	RefId       uuid.UUID
	HistoryTime time.Time
	Operation   enthistory.OpType
	Changes     []Change
	UpdatedBy   *string
}

func (r *record) toRow() []string {
	row := make([]string, 6)

	row[0] = r.Table
	row[1] = fmt.Sprintf("%v", r.RefId)
	row[2] = r.HistoryTime.Format(time.ANSIC)
	row[3] = r.Operation.String()
	for i, change := range r.Changes {
		if i == 0 {
			row[4] = change.String(r.Operation)
			continue
		}
		row[4] = fmt.Sprintf("%s\n%s", row[4], change.String(r.Operation))
	}
	if r.UpdatedBy != nil {
		row[5] = fmt.Sprintf("%v", *r.UpdatedBy)
	}
	return row
}

type ref struct {
	Ref uuid.UUID
}

func auditAttributeHistory(ctx context.Context, config config) ([][]string, error) {
	var records = [][]string{}
	var refs []ref
	client := NewAttributeHistoryClient(config)
	err := client.Query().
		Unique(true).
		Order(Asc(attributehistory.FieldRef)).
		Select(attributehistory.FieldRef).
		Scan(ctx, &refs)

	if err != nil {
		return nil, err
	}
	for _, currRef := range refs {
		histories, err := client.Query().
			Where(attributehistory.Ref(currRef.Ref)).
			Order(Asc(attributehistory.FieldHistoryTime)).
			All(ctx)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(histories); i++ {
			curr := histories[i]
			record := record{
				Table:       "AttributeHistory",
				RefId:       curr.Ref,
				HistoryTime: curr.HistoryTime,
				Operation:   curr.Operation,
				UpdatedBy:   curr.UpdatedBy,
			}
			switch curr.Operation {
			case enthistory.OpTypeInsert:
				record.Changes = (&AttributeHistory{}).changes(curr)
			case enthistory.OpTypeDelete:
				record.Changes = curr.changes(&AttributeHistory{})
			default:
				if i == 0 {
					record.Changes = (&AttributeHistory{}).changes(curr)
				} else {
					record.Changes = histories[i-1].changes(curr)
				}
			}
			records = append(records, record.toRow())
		}
	}
	return records, nil
}
func auditAttributeTypeHistory(ctx context.Context, config config) ([][]string, error) {
	var records = [][]string{}
	var refs []ref
	client := NewAttributeTypeHistoryClient(config)
	err := client.Query().
		Unique(true).
		Order(Asc(attributetypehistory.FieldRef)).
		Select(attributetypehistory.FieldRef).
		Scan(ctx, &refs)

	if err != nil {
		return nil, err
	}
	for _, currRef := range refs {
		histories, err := client.Query().
			Where(attributetypehistory.Ref(currRef.Ref)).
			Order(Asc(attributetypehistory.FieldHistoryTime)).
			All(ctx)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(histories); i++ {
			curr := histories[i]
			record := record{
				Table:       "AttributeTypeHistory",
				RefId:       curr.Ref,
				HistoryTime: curr.HistoryTime,
				Operation:   curr.Operation,
				UpdatedBy:   curr.UpdatedBy,
			}
			switch curr.Operation {
			case enthistory.OpTypeInsert:
				record.Changes = (&AttributeTypeHistory{}).changes(curr)
			case enthistory.OpTypeDelete:
				record.Changes = curr.changes(&AttributeTypeHistory{})
			default:
				if i == 0 {
					record.Changes = (&AttributeTypeHistory{}).changes(curr)
				} else {
					record.Changes = histories[i-1].changes(curr)
				}
			}
			records = append(records, record.toRow())
		}
	}
	return records, nil
}
func auditCertificationHistory(ctx context.Context, config config) ([][]string, error) {
	var records = [][]string{}
	var refs []ref
	client := NewCertificationHistoryClient(config)
	err := client.Query().
		Unique(true).
		Order(Asc(certificationhistory.FieldRef)).
		Select(certificationhistory.FieldRef).
		Scan(ctx, &refs)

	if err != nil {
		return nil, err
	}
	for _, currRef := range refs {
		histories, err := client.Query().
			Where(certificationhistory.Ref(currRef.Ref)).
			Order(Asc(certificationhistory.FieldHistoryTime)).
			All(ctx)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(histories); i++ {
			curr := histories[i]
			record := record{
				Table:       "CertificationHistory",
				RefId:       curr.Ref,
				HistoryTime: curr.HistoryTime,
				Operation:   curr.Operation,
				UpdatedBy:   curr.UpdatedBy,
			}
			switch curr.Operation {
			case enthistory.OpTypeInsert:
				record.Changes = (&CertificationHistory{}).changes(curr)
			case enthistory.OpTypeDelete:
				record.Changes = curr.changes(&CertificationHistory{})
			default:
				if i == 0 {
					record.Changes = (&CertificationHistory{}).changes(curr)
				} else {
					record.Changes = histories[i-1].changes(curr)
				}
			}
			records = append(records, record.toRow())
		}
	}
	return records, nil
}
func auditCertificationTemplateHistory(ctx context.Context, config config) ([][]string, error) {
	var records = [][]string{}
	var refs []ref
	client := NewCertificationTemplateHistoryClient(config)
	err := client.Query().
		Unique(true).
		Order(Asc(certificationtemplatehistory.FieldRef)).
		Select(certificationtemplatehistory.FieldRef).
		Scan(ctx, &refs)

	if err != nil {
		return nil, err
	}
	for _, currRef := range refs {
		histories, err := client.Query().
			Where(certificationtemplatehistory.Ref(currRef.Ref)).
			Order(Asc(certificationtemplatehistory.FieldHistoryTime)).
			All(ctx)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(histories); i++ {
			curr := histories[i]
			record := record{
				Table:       "CertificationTemplateHistory",
				RefId:       curr.Ref,
				HistoryTime: curr.HistoryTime,
				Operation:   curr.Operation,
				UpdatedBy:   curr.UpdatedBy,
			}
			switch curr.Operation {
			case enthistory.OpTypeInsert:
				record.Changes = (&CertificationTemplateHistory{}).changes(curr)
			case enthistory.OpTypeDelete:
				record.Changes = curr.changes(&CertificationTemplateHistory{})
			default:
				if i == 0 {
					record.Changes = (&CertificationTemplateHistory{}).changes(curr)
				} else {
					record.Changes = histories[i-1].changes(curr)
				}
			}
			records = append(records, record.toRow())
		}
	}
	return records, nil
}
func auditCompanyHistory(ctx context.Context, config config) ([][]string, error) {
	var records = [][]string{}
	var refs []ref
	client := NewCompanyHistoryClient(config)
	err := client.Query().
		Unique(true).
		Order(Asc(companyhistory.FieldRef)).
		Select(companyhistory.FieldRef).
		Scan(ctx, &refs)

	if err != nil {
		return nil, err
	}
	for _, currRef := range refs {
		histories, err := client.Query().
			Where(companyhistory.Ref(currRef.Ref)).
			Order(Asc(companyhistory.FieldHistoryTime)).
			All(ctx)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(histories); i++ {
			curr := histories[i]
			record := record{
				Table:       "CompanyHistory",
				RefId:       curr.Ref,
				HistoryTime: curr.HistoryTime,
				Operation:   curr.Operation,
				UpdatedBy:   curr.UpdatedBy,
			}
			switch curr.Operation {
			case enthistory.OpTypeInsert:
				record.Changes = (&CompanyHistory{}).changes(curr)
			case enthistory.OpTypeDelete:
				record.Changes = curr.changes(&CompanyHistory{})
			default:
				if i == 0 {
					record.Changes = (&CompanyHistory{}).changes(curr)
				} else {
					record.Changes = histories[i-1].changes(curr)
				}
			}
			records = append(records, record.toRow())
		}
	}
	return records, nil
}
func auditItemBatchHistory(ctx context.Context, config config) ([][]string, error) {
	var records = [][]string{}
	var refs []ref
	client := NewItemBatchHistoryClient(config)
	err := client.Query().
		Unique(true).
		Order(Asc(itembatchhistory.FieldRef)).
		Select(itembatchhistory.FieldRef).
		Scan(ctx, &refs)

	if err != nil {
		return nil, err
	}
	for _, currRef := range refs {
		histories, err := client.Query().
			Where(itembatchhistory.Ref(currRef.Ref)).
			Order(Asc(itembatchhistory.FieldHistoryTime)).
			All(ctx)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(histories); i++ {
			curr := histories[i]
			record := record{
				Table:       "ItemBatchHistory",
				RefId:       curr.Ref,
				HistoryTime: curr.HistoryTime,
				Operation:   curr.Operation,
				UpdatedBy:   curr.UpdatedBy,
			}
			switch curr.Operation {
			case enthistory.OpTypeInsert:
				record.Changes = (&ItemBatchHistory{}).changes(curr)
			case enthistory.OpTypeDelete:
				record.Changes = curr.changes(&ItemBatchHistory{})
			default:
				if i == 0 {
					record.Changes = (&ItemBatchHistory{}).changes(curr)
				} else {
					record.Changes = histories[i-1].changes(curr)
				}
			}
			records = append(records, record.toRow())
		}
	}
	return records, nil
}
func auditItemBatchToItemBatchHistory(ctx context.Context, config config) ([][]string, error) {
	var records = [][]string{}
	var refs []ref
	client := NewItemBatchToItemBatchHistoryClient(config)
	err := client.Query().
		Unique(true).
		Order(Asc(itembatchtoitembatchhistory.FieldRef)).
		Select(itembatchtoitembatchhistory.FieldRef).
		Scan(ctx, &refs)

	if err != nil {
		return nil, err
	}
	for _, currRef := range refs {
		histories, err := client.Query().
			Where(itembatchtoitembatchhistory.Ref(currRef.Ref)).
			Order(Asc(itembatchtoitembatchhistory.FieldHistoryTime)).
			All(ctx)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(histories); i++ {
			curr := histories[i]
			record := record{
				Table:       "ItemBatchToItemBatchHistory",
				RefId:       curr.Ref,
				HistoryTime: curr.HistoryTime,
				Operation:   curr.Operation,
				UpdatedBy:   curr.UpdatedBy,
			}
			switch curr.Operation {
			case enthistory.OpTypeInsert:
				record.Changes = (&ItemBatchToItemBatchHistory{}).changes(curr)
			case enthistory.OpTypeDelete:
				record.Changes = curr.changes(&ItemBatchToItemBatchHistory{})
			default:
				if i == 0 {
					record.Changes = (&ItemBatchToItemBatchHistory{}).changes(curr)
				} else {
					record.Changes = histories[i-1].changes(curr)
				}
			}
			records = append(records, record.toRow())
		}
	}
	return records, nil
}
func auditUserHistory(ctx context.Context, config config) ([][]string, error) {
	var records = [][]string{}
	var refs []ref
	client := NewUserHistoryClient(config)
	err := client.Query().
		Unique(true).
		Order(Asc(userhistory.FieldRef)).
		Select(userhistory.FieldRef).
		Scan(ctx, &refs)

	if err != nil {
		return nil, err
	}
	for _, currRef := range refs {
		histories, err := client.Query().
			Where(userhistory.Ref(currRef.Ref)).
			Order(Asc(userhistory.FieldHistoryTime)).
			All(ctx)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(histories); i++ {
			curr := histories[i]
			record := record{
				Table:       "UserHistory",
				RefId:       curr.Ref,
				HistoryTime: curr.HistoryTime,
				Operation:   curr.Operation,
				UpdatedBy:   curr.UpdatedBy,
			}
			switch curr.Operation {
			case enthistory.OpTypeInsert:
				record.Changes = (&UserHistory{}).changes(curr)
			case enthistory.OpTypeDelete:
				record.Changes = curr.changes(&UserHistory{})
			default:
				if i == 0 {
					record.Changes = (&UserHistory{}).changes(curr)
				} else {
					record.Changes = histories[i-1].changes(curr)
				}
			}
			records = append(records, record.toRow())
		}
	}
	return records, nil
}
func auditUsersToCompanyHistory(ctx context.Context, config config) ([][]string, error) {
	var records = [][]string{}
	var refs []ref
	client := NewUsersToCompanyHistoryClient(config)
	err := client.Query().
		Unique(true).
		Order(Asc(userstocompanyhistory.FieldRef)).
		Select(userstocompanyhistory.FieldRef).
		Scan(ctx, &refs)

	if err != nil {
		return nil, err
	}
	for _, currRef := range refs {
		histories, err := client.Query().
			Where(userstocompanyhistory.Ref(currRef.Ref)).
			Order(Asc(userstocompanyhistory.FieldHistoryTime)).
			All(ctx)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(histories); i++ {
			curr := histories[i]
			record := record{
				Table:       "UsersToCompanyHistory",
				RefId:       curr.Ref,
				HistoryTime: curr.HistoryTime,
				Operation:   curr.Operation,
				UpdatedBy:   curr.UpdatedBy,
			}
			switch curr.Operation {
			case enthistory.OpTypeInsert:
				record.Changes = (&UsersToCompanyHistory{}).changes(curr)
			case enthistory.OpTypeDelete:
				record.Changes = curr.changes(&UsersToCompanyHistory{})
			default:
				if i == 0 {
					record.Changes = (&UsersToCompanyHistory{}).changes(curr)
				} else {
					record.Changes = histories[i-1].changes(curr)
				}
			}
			records = append(records, record.toRow())
		}
	}
	return records, nil
}
