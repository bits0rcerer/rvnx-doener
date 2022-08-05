// Code generated by ent, DO NOT EDIT.

package scorerating

import (
	"rvnx_doener_service/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v float64) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScore), v))
	})
}

// Anonymous applies equality check predicate on the "anonymous" field. It's identical to AnonymousEQ.
func Anonymous(v bool) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAnonymous), v))
	})
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreated), v))
	})
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.ScoreRating {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ScoreRating(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreated), v...))
	})
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.ScoreRating {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ScoreRating(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreated), v...))
	})
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreated), v))
	})
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreated), v))
	})
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreated), v))
	})
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreated), v))
	})
}

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v float64) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScore), v))
	})
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v float64) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldScore), v))
	})
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...float64) predicate.ScoreRating {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ScoreRating(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldScore), v...))
	})
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...float64) predicate.ScoreRating {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ScoreRating(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldScore), v...))
	})
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v float64) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldScore), v))
	})
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v float64) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldScore), v))
	})
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v float64) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldScore), v))
	})
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v float64) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldScore), v))
	})
}

// AnonymousEQ applies the EQ predicate on the "anonymous" field.
func AnonymousEQ(v bool) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAnonymous), v))
	})
}

// AnonymousNEQ applies the NEQ predicate on the "anonymous" field.
func AnonymousNEQ(v bool) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAnonymous), v))
	})
}

// HasShop applies the HasEdge predicate on the "shop" edge.
func HasShop() predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ShopTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ShopTable, ShopColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShopWith applies the HasEdge predicate on the "shop" edge with a given conditions (other predicates).
func HasShopWith(preds ...predicate.KebabShop) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ShopInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ShopTable, ShopColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.TwitchUser) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ScoreRating) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ScoreRating) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
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
func Not(p predicate.ScoreRating) predicate.ScoreRating {
	return predicate.ScoreRating(func(s *sql.Selector) {
		p(s.Not())
	})
}
