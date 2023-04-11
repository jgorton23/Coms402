// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"git.las.iastate.edu/SeniorDesignComS/2023spr/online-certificate-repo/backend/pkg/database/ent"
)

// The AttributeFunc type is an adapter to allow the use of ordinary
// function as Attribute mutator.
type AttributeFunc func(context.Context, *ent.AttributeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AttributeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AttributeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AttributeMutation", m)
}

// The AttributeHistoryFunc type is an adapter to allow the use of ordinary
// function as AttributeHistory mutator.
type AttributeHistoryFunc func(context.Context, *ent.AttributeHistoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AttributeHistoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AttributeHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AttributeHistoryMutation", m)
}

// The AttributeTypeFunc type is an adapter to allow the use of ordinary
// function as AttributeType mutator.
type AttributeTypeFunc func(context.Context, *ent.AttributeTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AttributeTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AttributeTypeMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AttributeTypeMutation", m)
}

// The AttributeTypeHistoryFunc type is an adapter to allow the use of ordinary
// function as AttributeTypeHistory mutator.
type AttributeTypeHistoryFunc func(context.Context, *ent.AttributeTypeHistoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AttributeTypeHistoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AttributeTypeHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AttributeTypeHistoryMutation", m)
}

// The AuthorizationPolicyFunc type is an adapter to allow the use of ordinary
// function as AuthorizationPolicy mutator.
type AuthorizationPolicyFunc func(context.Context, *ent.AuthorizationPolicyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AuthorizationPolicyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.AuthorizationPolicyMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AuthorizationPolicyMutation", m)
}

// The CertificationFunc type is an adapter to allow the use of ordinary
// function as Certification mutator.
type CertificationFunc func(context.Context, *ent.CertificationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CertificationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CertificationMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CertificationMutation", m)
}

// The CertificationHistoryFunc type is an adapter to allow the use of ordinary
// function as CertificationHistory mutator.
type CertificationHistoryFunc func(context.Context, *ent.CertificationHistoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CertificationHistoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CertificationHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CertificationHistoryMutation", m)
}

// The CertificationTemplateFunc type is an adapter to allow the use of ordinary
// function as CertificationTemplate mutator.
type CertificationTemplateFunc func(context.Context, *ent.CertificationTemplateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CertificationTemplateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CertificationTemplateMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CertificationTemplateMutation", m)
}

// The CertificationTemplateHistoryFunc type is an adapter to allow the use of ordinary
// function as CertificationTemplateHistory mutator.
type CertificationTemplateHistoryFunc func(context.Context, *ent.CertificationTemplateHistoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CertificationTemplateHistoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CertificationTemplateHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CertificationTemplateHistoryMutation", m)
}

// The CompanyFunc type is an adapter to allow the use of ordinary
// function as Company mutator.
type CompanyFunc func(context.Context, *ent.CompanyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CompanyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CompanyMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CompanyMutation", m)
}

// The CompanyHistoryFunc type is an adapter to allow the use of ordinary
// function as CompanyHistory mutator.
type CompanyHistoryFunc func(context.Context, *ent.CompanyHistoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CompanyHistoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.CompanyHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CompanyHistoryMutation", m)
}

// The ItemBatchFunc type is an adapter to allow the use of ordinary
// function as ItemBatch mutator.
type ItemBatchFunc func(context.Context, *ent.ItemBatchMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ItemBatchFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ItemBatchMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ItemBatchMutation", m)
}

// The ItemBatchHistoryFunc type is an adapter to allow the use of ordinary
// function as ItemBatchHistory mutator.
type ItemBatchHistoryFunc func(context.Context, *ent.ItemBatchHistoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ItemBatchHistoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ItemBatchHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ItemBatchHistoryMutation", m)
}

// The ItemBatchToItemBatchFunc type is an adapter to allow the use of ordinary
// function as ItemBatchToItemBatch mutator.
type ItemBatchToItemBatchFunc func(context.Context, *ent.ItemBatchToItemBatchMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ItemBatchToItemBatchFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ItemBatchToItemBatchMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ItemBatchToItemBatchMutation", m)
}

// The ItemBatchToItemBatchHistoryFunc type is an adapter to allow the use of ordinary
// function as ItemBatchToItemBatchHistory mutator.
type ItemBatchToItemBatchHistoryFunc func(context.Context, *ent.ItemBatchToItemBatchHistoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f ItemBatchToItemBatchHistoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.ItemBatchToItemBatchHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.ItemBatchToItemBatchHistoryMutation", m)
}

// The SessionFunc type is an adapter to allow the use of ordinary
// function as Session mutator.
type SessionFunc func(context.Context, *ent.SessionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SessionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.SessionMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SessionMutation", m)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
}

// The UserHistoryFunc type is an adapter to allow the use of ordinary
// function as UserHistory mutator.
type UserHistoryFunc func(context.Context, *ent.UserHistoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserHistoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UserHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserHistoryMutation", m)
}

// The UsersToCompanyFunc type is an adapter to allow the use of ordinary
// function as UsersToCompany mutator.
type UsersToCompanyFunc func(context.Context, *ent.UsersToCompanyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UsersToCompanyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UsersToCompanyMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UsersToCompanyMutation", m)
}

// The UsersToCompanyHistoryFunc type is an adapter to allow the use of ordinary
// function as UsersToCompanyHistory mutator.
type UsersToCompanyHistoryFunc func(context.Context, *ent.UsersToCompanyHistoryMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UsersToCompanyHistoryFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(*ent.UsersToCompanyHistoryMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UsersToCompanyHistoryMutation", m)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
