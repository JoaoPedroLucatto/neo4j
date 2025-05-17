package neo4j

import (
	"context"
	"neo4j-etl/internal/entity"

	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (r *Neo4j) LoadCountries(c *entity.Country) error {
	ctx := context.Background()
	session := r.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MERGE (country:Country {iso3: $iso3})
			SET country.id = $id, country.name = $name
		`
		_, err := tx.Run(ctx, query, map[string]interface{}{
			"id":   uuid.NewString(),
			"iso3": c.Iso3,
			"name": c.Name,
		})

		return nil, err
	})

	return err
}

func (r *Neo4j) LoadVaccines(vaccine *entity.Vaccine) error {
	ctx := context.Background()
	session := r.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MERGE (vaccine:Vaccine {name: $name})
			ON CREATE SET vaccine.id = $id
			WITH vaccine
			MATCH (country:Country {iso3: $iso3})
			MERGE (country)-[:USES]->(vaccine)
		`
		_, err := tx.Run(ctx, query, map[string]interface{}{
			"id":   uuid.NewString(),
			"iso3": vaccine.CountryIso,
			"name": vaccine.Name,
		})

		return nil, err
	})

	return err
}

func (r *Neo4j) LoadCovidCases(cc *entity.CovidCase) error {
	ctx := context.Background()
	session := r.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (country:Country {iso3: $iso3})
			MERGE (case:CovidCase {countryIso: $iso3, date: $date})
			SET case.id = $id, case.totalCases = $cases, case.totalDeaths = $deaths
			MERGE (country)-[:HAS_CASE]->(case)
		`
		_, err := tx.Run(ctx, query, map[string]interface{}{
			"id":     uuid.NewString(),
			"iso3":   cc.CountryIso,
			"date":   cc.Date,
			"cases":  cc.TotalCases,
			"deaths": cc.TotalDeaths,
		})

		return nil, err
	})

	return err
}

func (r *Neo4j) LoadVaccinationStats(vs *entity.VaccinationStats) error {
	ctx := context.Background()
	session := r.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (c:Country {iso3: $iso3})
			MERGE (vs:VaccinationStats {countryIso: $iso3, date: $date})
			SET vs.id = $id, vs.totalVaccinated = $totalVaccinated
			MERGE (c)-[:VACCINATED_ON]->(vs)
		`
		_, err := tx.Run(ctx, query, map[string]interface{}{
			"id":              uuid.NewString(),
			"iso3":            vs.CountryIso,
			"date":            vs.Date,
			"totalVaccinated": vs.TotalVaccinated,
		})

		return nil, err
	})

	return err
}

func (r *Neo4j) LoadVaccineApprovals(va *entity.VaccineApproval) error {
	ctx := context.Background()
	session := r.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MERGE (v:Vaccine {name: $vaccine})
			SET v.id = $vaccineId
	
			MERGE (va:VaccineApproval {id: $id})
			SET va.vaccine = $vaccine, va.date = date($date)
	
			MERGE (v)-[:APPROVED_ON]->(va)
	
			MERGE (c:Country {iso3: $iso3})
			SET c.name = $countryName
	
			MERGE (c)-[:USES]->(v)
		`
		_, err := tx.Run(ctx, query, map[string]interface{}{
			"id":          uuid.NewString(),
			"vaccineId":   va.Vaccine,
			"iso3":        va.CountryIso,
			"vaccine":     va.Vaccine,
			"countryName": va.CountryName,
			"date":        va.Date,
		})

		if err != nil {
			return nil, err
		}

		return nil, err
	})

	return err
}
