// Code generated by ent, DO NOT EDIT.

package kebabshop

import (
	"rvnx_doener_service/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OsmID applies equality check predicate on the "osm_id" field. It's identical to OsmIDEQ.
func OsmID(v int) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOsmID), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// Lat applies equality check predicate on the "lat" field. It's identical to LatEQ.
func Lat(v float64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLat), v))
	})
}

// Lng applies equality check predicate on the "lng" field. It's identical to LngEQ.
func Lng(v float64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLng), v))
	})
}

// Visible applies equality check predicate on the "visible" field. It's identical to VisibleEQ.
func Visible(v bool) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVisible), v))
	})
}

// OsmIDEQ applies the EQ predicate on the "osm_id" field.
func OsmIDEQ(v int) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOsmID), v))
	})
}

// OsmIDNEQ applies the NEQ predicate on the "osm_id" field.
func OsmIDNEQ(v int) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOsmID), v))
	})
}

// OsmIDIn applies the In predicate on the "osm_id" field.
func OsmIDIn(vs ...int) predicate.KebabShop {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KebabShop(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOsmID), v...))
	})
}

// OsmIDNotIn applies the NotIn predicate on the "osm_id" field.
func OsmIDNotIn(vs ...int) predicate.KebabShop {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KebabShop(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOsmID), v...))
	})
}

// OsmIDGT applies the GT predicate on the "osm_id" field.
func OsmIDGT(v int) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOsmID), v))
	})
}

// OsmIDGTE applies the GTE predicate on the "osm_id" field.
func OsmIDGTE(v int) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOsmID), v))
	})
}

// OsmIDLT applies the LT predicate on the "osm_id" field.
func OsmIDLT(v int) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOsmID), v))
	})
}

// OsmIDLTE applies the LTE predicate on the "osm_id" field.
func OsmIDLTE(v int) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOsmID), v))
	})
}

// OsmIDIsNil applies the IsNil predicate on the "osm_id" field.
func OsmIDIsNil() predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOsmID)))
	})
}

// OsmIDNotNil applies the NotNil predicate on the "osm_id" field.
func OsmIDNotNil() predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOsmID)))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.KebabShop {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KebabShop(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.KebabShop {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KebabShop(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreated), v))
	})
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.KebabShop {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KebabShop(func(s *sql.Selector) {
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
func CreatedNotIn(vs ...time.Time) predicate.KebabShop {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KebabShop(func(s *sql.Selector) {
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
func CreatedGT(v time.Time) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreated), v))
	})
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreated), v))
	})
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreated), v))
	})
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreated), v))
	})
}

// LatEQ applies the EQ predicate on the "lat" field.
func LatEQ(v float64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLat), v))
	})
}

// LatNEQ applies the NEQ predicate on the "lat" field.
func LatNEQ(v float64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLat), v))
	})
}

// LatIn applies the In predicate on the "lat" field.
func LatIn(vs ...float64) predicate.KebabShop {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KebabShop(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLat), v...))
	})
}

// LatNotIn applies the NotIn predicate on the "lat" field.
func LatNotIn(vs ...float64) predicate.KebabShop {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KebabShop(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLat), v...))
	})
}

// LatGT applies the GT predicate on the "lat" field.
func LatGT(v float64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLat), v))
	})
}

// LatGTE applies the GTE predicate on the "lat" field.
func LatGTE(v float64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLat), v))
	})
}

// LatLT applies the LT predicate on the "lat" field.
func LatLT(v float64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLat), v))
	})
}

// LatLTE applies the LTE predicate on the "lat" field.
func LatLTE(v float64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLat), v))
	})
}

// LngEQ applies the EQ predicate on the "lng" field.
func LngEQ(v float64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLng), v))
	})
}

// LngNEQ applies the NEQ predicate on the "lng" field.
func LngNEQ(v float64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLng), v))
	})
}

// LngIn applies the In predicate on the "lng" field.
func LngIn(vs ...float64) predicate.KebabShop {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KebabShop(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLng), v...))
	})
}

// LngNotIn applies the NotIn predicate on the "lng" field.
func LngNotIn(vs ...float64) predicate.KebabShop {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.KebabShop(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLng), v...))
	})
}

// LngGT applies the GT predicate on the "lng" field.
func LngGT(v float64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLng), v))
	})
}

// LngGTE applies the GTE predicate on the "lng" field.
func LngGTE(v float64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLng), v))
	})
}

// LngLT applies the LT predicate on the "lng" field.
func LngLT(v float64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLng), v))
	})
}

// LngLTE applies the LTE predicate on the "lng" field.
func LngLTE(v float64) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLng), v))
	})
}

// VisibleEQ applies the EQ predicate on the "visible" field.
func VisibleEQ(v bool) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVisible), v))
	})
}

// VisibleNEQ applies the NEQ predicate on the "visible" field.
func VisibleNEQ(v bool) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVisible), v))
	})
}

// HasUserScores applies the HasEdge predicate on the "user_scores" edge.
func HasUserScores() predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserScoresTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserScoresTable, UserScoresColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserScoresWith applies the HasEdge predicate on the "user_scores" edge with a given conditions (other predicates).
func HasUserScoresWith(preds ...predicate.ScoreRating) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserScoresInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserScoresTable, UserScoresColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserPrices applies the HasEdge predicate on the "user_prices" edge.
func HasUserPrices() predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserPricesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserPricesTable, UserPricesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserPricesWith applies the HasEdge predicate on the "user_prices" edge with a given conditions (other predicates).
func HasUserPricesWith(preds ...predicate.ShopPrice) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserPricesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserPricesTable, UserPricesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserOpinions applies the HasEdge predicate on the "user_opinions" edge.
func HasUserOpinions() predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserOpinionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserOpinionsTable, UserOpinionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserOpinionsWith applies the HasEdge predicate on the "user_opinions" edge with a given conditions (other predicates).
func HasUserOpinionsWith(preds ...predicate.UserOpinion) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserOpinionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserOpinionsTable, UserOpinionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.KebabShop) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.KebabShop) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
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
func Not(p predicate.KebabShop) predicate.KebabShop {
	return predicate.KebabShop(func(s *sql.Selector) {
		p(s.Not())
	})
}
