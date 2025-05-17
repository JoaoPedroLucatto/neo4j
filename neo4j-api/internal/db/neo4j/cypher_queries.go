package neo4j

import (
	"context"
	"fmt"
	"neo4j-api/internal/entity"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

func (r *Neo4j) GetCovidSummary(iso3 string, date time.Time) (*map[string]interface{}, error) {
	ctx := context.Background()
	session := r.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (c:Country {iso3: $iso3})-[:HAS_CASE]->(cc:CovidCase)
			WHERE date(cc.date) = date($date)
			RETURN cc.totalCases AS cases, cc.totalDeaths AS deaths
		`

		res, err := tx.Run(ctx, query, map[string]interface{}{
			"iso3": iso3,
			"date": date,
		})
		if err != nil {
			return nil, err
		}

		if res.Next(ctx) {
			record := res.Record()
			return &map[string]interface{}{
				"cases":  record.Values[0],
				"deaths": record.Values[1],
			}, nil
		}

		return nil, err
	})
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	data := result.(*map[string]interface{})
	return data, nil
}

func (r *Neo4j) GetFirstDoseCount(iso3 string, date time.Time) (*map[string]interface{}, error) {
	ctx := context.Background()
	session := r.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (c:Country {iso3: $iso3})-[:VACCINATED_ON]->(vs:VaccinationStats)
			WHERE date(vs.date) = date($date)
			RETURN vs.totalVaccinated AS totalVaccinated
		`

		res, err := tx.Run(ctx, query, map[string]interface{}{
			"iso3": iso3,
			"date": date,
		})
		if err != nil {
			return nil, err
		}

		if res.Next(ctx) {
			record := res.Record()
			return &map[string]interface{}{
				"totalVaccinated": record.Values[0],
			}, nil
		}

		return nil, err
	})
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	data := result.(*map[string]interface{})
	return data, nil
}

func (r *Neo4j) GetVaccinesUsedByCountry(iso3 string) (*map[string]interface{}, error) {
	ctx := context.Background()
	session := r.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (c:Country {iso3: $iso3})-[:USES]->(v:Vaccine)
			RETURN v.name AS vaccineName

		`

		res, err := tx.Run(ctx, query, map[string]interface{}{
			"iso3": iso3,
		})
		if err != nil {
			return nil, err
		}

		var vaccines []string

		for res.Next(ctx) {
			record := res.Record()
			if name, ok := record.Values[0].(string); ok {
				vaccines = append(vaccines, name)
			}
		}

		if err = res.Err(); err != nil {
			return nil, err
		}

		return map[string]interface{}{
			"vaccines": vaccines,
		}, nil
	})
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	data := result.(map[string]interface{})

	return &data, nil
}

func (r *Neo4j) GetVaccineApprovalDates(vaccineName string) (*[]entity.VaccineApprovalOutPut, error) {
	ctx := context.Background()
	session := r.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (v:Vaccine)-[:APPROVED_ON]->(va:VaccineApproval)
			RETURN v.name AS vaccineName, va.date AS approvalDate
			ORDER BY approvalDate
		`

		res, err := tx.Run(ctx, query, nil)
		if err != nil {
			return nil, err
		}

		var approvals []entity.VaccineApprovalOutPut

		for res.Next(ctx) {
			record := res.Record()
			vaccineName, existVaccine := record.Get("vaccineName")
			fmt.Printf("tem vaccineName %s\n", vaccineName)
			fmt.Printf("O tipo de vacciName é: %T\n", vaccineName)
			if !existVaccine {
				fmt.Printf("tem vacina %t", existVaccine)
			}

			approvalDate, existDate := record.Get("approvalDate")
			fmt.Printf("tem approvalDate: %+v\n", approvalDate)
			fmt.Printf("O tipo de approvalDate é: %T\n", approvalDate)
			if !existDate {
				fmt.Printf("tem approvalDate %t", existDate)

			}

			approvalDateStr := approvalDate.(dbtype.Date).String()

			approvals = append(approvals, entity.VaccineApprovalOutPut{
				VaccineName:  vaccineName.(string),
				ApprovalDate: approvalDateStr,
			})
		}

		if err := res.Err(); err != nil {
			return nil, err
		}

		return approvals, nil
	})
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	approvals, ok := result.([]entity.VaccineApprovalOutPut)
	if !ok {
		return nil, fmt.Errorf("falha ao converter resultado")
	}

	return &approvals, nil
}

func (r *Neo4j) GetCountriesByVaccine(vaccineName string) (*[]entity.CountriesByVaccineOutPut, error) {
	var countries []entity.CountriesByVaccineOutPut
	ctx := context.Background()
	session := r.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MATCH (c:Country)-[:USES]->(v:Vaccine {name: $vaccineName})
			RETURN c.name AS countryName, c.iso3 AS iso3
			ORDER BY c.name
		`

		res, err := tx.Run(ctx, query, map[string]interface{}{
			"vaccineName": vaccineName,
		})
		if err != nil {
			return nil, err
		}

		for res.Next(ctx) {
			record := res.Record()
			name, ok1 := record.Values[0].(string)
			iso3, ok2 := record.Values[1].(string)
			if ok1 && ok2 {
				countries = append(countries, entity.CountriesByVaccineOutPut{
					CountryName: name,
					Iso3:        iso3,
				})
			}
		}

		if err = res.Err(); err != nil {
			return nil, err
		}

		return countries, nil
	})
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	return &countries, nil
}
