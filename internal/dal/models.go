// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package dal

import (
	"database/sql"
	"time"
)

type Circuit struct {
	ID             string
	Name           string
	FullName       string
	PreviousNames  sql.NullString
	Type           string
	PlaceName      string
	CountryID      string
	Latitude       float64
	Longitude      float64
	TotalRacesHeld int64
}

type Constructor struct {
	ID                       string
	Name                     string
	FullName                 string
	CountryID                string
	BestChampionshipPosition sql.NullInt64
	BestStartingGridPosition sql.NullInt64
	BestRaceResult           sql.NullInt64
	TotalChampionshipWins    int64
	TotalRaceEntries         int64
	TotalRaceStarts          int64
	TotalRaceWins            int64
	Total1And2Finishes       int64
	TotalRaceLaps            int64
	TotalPodiums             int64
	TotalPodiumRaces         int64
	TotalChampionshipPoints  float64
	TotalPolePositions       int64
	TotalFastestLaps         int64
}

type ConstructorPreviousNextConstructor struct {
	ConstructorID             string
	PreviousNextConstructorID string
	YearFrom                  int64
	YearTo                    sql.NullInt64
}

type Continent struct {
	ID      string
	Code    string
	Name    string
	Demonym string
}

type Country struct {
	ID          string
	Alpha2Code  string
	Alpha3Code  string
	Name        string
	Demonym     sql.NullString
	ContinentID string
}

type Driver struct {
	ID                         string
	Name                       string
	FirstName                  string
	LastName                   string
	FullName                   string
	Abbreviation               string
	PermanentNumber            sql.NullString
	Gender                     string
	DateOfBirth                time.Time
	DateOfDeath                sql.NullTime
	PlaceOfBirth               string
	CountryOfBirthCountryID    string
	NationalityCountryID       string
	SecondNationalityCountryID sql.NullString
	BestChampionshipPosition   sql.NullInt64
	BestStartingGridPosition   sql.NullInt64
	BestRaceResult             sql.NullInt64
	TotalChampionshipWins      int64
	TotalRaceEntries           int64
	TotalRaceStarts            int64
	TotalRaceWins              int64
	TotalRaceLaps              int64
	TotalPodiums               int64
	TotalPoints                float64
	TotalChampionshipPoints    float64
	TotalPolePositions         int64
	TotalFastestLaps           int64
	TotalDriverOfTheDay        int64
	TotalGrandSlams            int64
}

type DriverFamilyRelationship struct {
	DriverID      string
	OtherDriverID string
	Type          string
}

type DriverOfTheDayResult struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	Percentage           sql.NullFloat64
}

type EngineManufacturer struct {
	ID                       string
	Name                     string
	CountryID                string
	BestChampionshipPosition sql.NullInt64
	BestStartingGridPosition sql.NullInt64
	BestRaceResult           sql.NullInt64
	TotalChampionshipWins    int64
	TotalRaceEntries         int64
	TotalRaceStarts          int64
	TotalRaceWins            int64
	TotalRaceLaps            int64
	TotalPodiums             int64
	TotalPodiumRaces         int64
	TotalChampionshipPoints  float64
	TotalPolePositions       int64
	TotalFastestLaps         int64
}

type Entrant struct {
	ID   string
	Name string
}

type FastestLap struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	Lap                  sql.NullInt64
	Time                 sql.NullString
	TimeMillis           sql.NullInt64
	Gap                  sql.NullString
	GapMillis            sql.NullInt64
	Interval             sql.NullString
	IntervalMillis       sql.NullInt64
}

type FreePractice1Result struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	Time                 sql.NullString
	TimeMillis           sql.NullInt64
	Gap                  sql.NullString
	GapMillis            sql.NullInt64
	Interval             sql.NullString
	IntervalMillis       sql.NullInt64
	Laps                 sql.NullInt64
}

type FreePractice2Result struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	Time                 sql.NullString
	TimeMillis           sql.NullInt64
	Gap                  sql.NullString
	GapMillis            sql.NullInt64
	Interval             sql.NullString
	IntervalMillis       sql.NullInt64
	Laps                 sql.NullInt64
}

type FreePractice3Result struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	Time                 sql.NullString
	TimeMillis           sql.NullInt64
	Gap                  sql.NullString
	GapMillis            sql.NullInt64
	Interval             sql.NullString
	IntervalMillis       sql.NullInt64
	Laps                 sql.NullInt64
}

type FreePractice4Result struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	Time                 sql.NullString
	TimeMillis           sql.NullInt64
	Gap                  sql.NullString
	GapMillis            sql.NullInt64
	Interval             sql.NullString
	IntervalMillis       sql.NullInt64
	Laps                 sql.NullInt64
}

type GrandPrix struct {
	ID             string
	Name           string
	FullName       string
	ShortName      string
	Abbreviation   string
	CountryID      sql.NullString
	TotalRacesHeld int64
}

type PitStop struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	Stop                 sql.NullInt64
	Lap                  sql.NullInt64
	Time                 sql.NullString
	TimeMillis           sql.NullInt64
}

type PreQualifyingResult struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	Time                 sql.NullString
	TimeMillis           sql.NullInt64
	Gap                  sql.NullString
	GapMillis            sql.NullInt64
	Interval             sql.NullString
	IntervalMillis       sql.NullInt64
	Laps                 sql.NullInt64
}

type Qualifying1Result struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	Time                 sql.NullString
	TimeMillis           sql.NullInt64
	Gap                  sql.NullString
	GapMillis            sql.NullInt64
	Interval             sql.NullString
	IntervalMillis       sql.NullInt64
	Laps                 sql.NullInt64
}

type Qualifying2Result struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	Time                 sql.NullString
	TimeMillis           sql.NullInt64
	Gap                  sql.NullString
	GapMillis            sql.NullInt64
	Interval             sql.NullString
	IntervalMillis       sql.NullInt64
	Laps                 sql.NullInt64
}

type QualifyingResult struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	Time                 sql.NullString
	TimeMillis           sql.NullInt64
	Q1                   sql.NullString
	Q1Millis             sql.NullInt64
	Q2                   sql.NullString
	Q2Millis             sql.NullInt64
	Q3                   sql.NullString
	Q3Millis             sql.NullInt64
	Gap                  sql.NullString
	GapMillis            sql.NullInt64
	Interval             sql.NullString
	IntervalMillis       sql.NullInt64
	Laps                 sql.NullInt64
}

type Race struct {
	ID                     int64
	Year                   int64
	Round                  int64
	Date                   time.Time
	Time                   sql.NullString
	GrandPrixID            string
	OfficialName           string
	QualifyingFormat       string
	SprintQualifyingFormat sql.NullString
	CircuitID              string
	CircuitType            string
	CourseLength           float64
	Laps                   int64
	Distance               float64
	ScheduledLaps          sql.NullInt64
	ScheduledDistance      sql.NullFloat64
	PreQualifyingDate      sql.NullTime
	PreQualifyingTime      sql.NullString
	FreePractice1Date      sql.NullTime
	FreePractice1Time      sql.NullString
	FreePractice2Date      sql.NullTime
	FreePractice2Time      sql.NullString
	FreePractice3Date      sql.NullTime
	FreePractice3Time      sql.NullString
	FreePractice4Date      sql.NullTime
	FreePractice4Time      sql.NullString
	Qualifying1Date        sql.NullTime
	Qualifying1Time        sql.NullString
	Qualifying2Date        sql.NullTime
	Qualifying2Time        sql.NullString
	QualifyingDate         sql.NullTime
	QualifyingTime         sql.NullString
	SprintQualifyingDate   sql.NullTime
	SprintQualifyingTime   sql.NullString
	SprintRaceDate         sql.NullTime
	SprintRaceTime         sql.NullString
	WarmingUpDate          sql.NullTime
	WarmingUpTime          sql.NullString
}

type RaceConstructorStanding struct {
	RaceID               int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	ConstructorID        string
	EngineManufacturerID string
	Points               float64
	PositionsGained      sql.NullInt64
}

type RaceDatum struct {
	RaceID                                   int64
	Type                                     string
	PositionDisplayOrder                     int64
	PositionNumber                           sql.NullInt64
	PositionText                             string
	DriverNumber                             string
	DriverID                                 string
	ConstructorID                            string
	EngineManufacturerID                     string
	TyreManufacturerID                       string
	PracticeTime                             sql.NullString
	PracticeTimeMillis                       sql.NullInt64
	PracticeGap                              sql.NullString
	PracticeGapMillis                        sql.NullInt64
	PracticeInterval                         sql.NullString
	PracticeIntervalMillis                   sql.NullInt64
	PracticeLaps                             sql.NullInt64
	QualifyingTime                           sql.NullString
	QualifyingTimeMillis                     sql.NullInt64
	QualifyingQ1                             sql.NullString
	QualifyingQ1Millis                       sql.NullInt64
	QualifyingQ2                             sql.NullString
	QualifyingQ2Millis                       sql.NullInt64
	QualifyingQ3                             sql.NullString
	QualifyingQ3Millis                       sql.NullInt64
	QualifyingGap                            sql.NullString
	QualifyingGapMillis                      sql.NullInt64
	QualifyingInterval                       sql.NullString
	QualifyingIntervalMillis                 sql.NullInt64
	QualifyingLaps                           sql.NullInt64
	StartingGridPositionGridPenalty          sql.NullString
	StartingGridPositionGridPenaltyPositions sql.NullInt64
	StartingGridPositionTime                 sql.NullString
	StartingGridPositionTimeMillis           sql.NullInt64
	RaceSharedCar                            sql.NullBool
	RaceLaps                                 sql.NullInt64
	RaceTime                                 sql.NullString
	RaceTimeMillis                           sql.NullInt64
	RaceTimePenalty                          sql.NullString
	RaceTimePenaltyMillis                    sql.NullInt64
	RaceGap                                  sql.NullString
	RaceGapMillis                            sql.NullInt64
	RaceGapLaps                              sql.NullInt64
	RaceInterval                             sql.NullString
	RaceIntervalMillis                       sql.NullInt64
	RaceReasonRetired                        sql.NullString
	RacePoints                               sql.NullFloat64
	RaceGridPositionNumber                   sql.NullInt64
	RaceGridPositionText                     sql.NullString
	RacePositionsGained                      sql.NullInt64
	RacePitStops                             sql.NullInt64
	RaceFastestLap                           sql.NullBool
	RaceDriverOfTheDay                       sql.NullBool
	RaceGrandSlam                            sql.NullBool
	FastestLapLap                            sql.NullInt64
	FastestLapTime                           sql.NullString
	FastestLapTimeMillis                     sql.NullInt64
	FastestLapGap                            sql.NullString
	FastestLapGapMillis                      sql.NullInt64
	FastestLapInterval                       sql.NullString
	FastestLapIntervalMillis                 sql.NullInt64
	PitStopStop                              sql.NullInt64
	PitStopLap                               sql.NullInt64
	PitStopTime                              sql.NullString
	PitStopTimeMillis                        sql.NullInt64
	DriverOfTheDayPercentage                 sql.NullFloat64
}

type RaceDriverStanding struct {
	RaceID               int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverID             string
	Points               float64
	PositionsGained      sql.NullInt64
}

type RaceResult struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	SharedCar            sql.NullBool
	Laps                 sql.NullInt64
	Time                 sql.NullString
	TimeMillis           sql.NullInt64
	TimePenalty          sql.NullString
	TimePenaltyMillis    sql.NullInt64
	Gap                  sql.NullString
	GapMillis            sql.NullInt64
	GapLaps              sql.NullInt64
	Interval             sql.NullString
	IntervalMillis       sql.NullInt64
	ReasonRetired        sql.NullString
	Points               sql.NullFloat64
	GridPositionNumber   sql.NullInt64
	GridPositionText     sql.NullString
	PositionsGained      sql.NullInt64
	PitStops             sql.NullInt64
	FastestLap           sql.NullBool
	DriverOfTheDay       sql.NullBool
	GrandSlam            sql.NullBool
}

type Season struct {
	Year int64
}

type SeasonConstructorStanding struct {
	Year                 int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	ConstructorID        string
	EngineManufacturerID string
	Points               float64
}

type SeasonDriverStanding struct {
	Year                 int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverID             string
	Points               float64
}

type SeasonEntrant struct {
	Year      int64
	EntrantID string
	CountryID string
}

type SeasonEntrantConstructor struct {
	Year                 int64
	EntrantID            string
	ConstructorID        string
	EngineManufacturerID string
}

type SeasonEntrantDriver struct {
	Year                 int64
	EntrantID            string
	ConstructorID        string
	EngineManufacturerID string
	DriverID             string
	Rounds               sql.NullString
	RoundsText           sql.NullString
	TestDriver           bool
}

type SeasonEntrantTyreManufacturer struct {
	Year                 int64
	EntrantID            string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
}

type SprintQualifyingResult struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	Time                 sql.NullString
	TimeMillis           sql.NullInt64
	Q1                   sql.NullString
	Q1Millis             sql.NullInt64
	Q2                   sql.NullString
	Q2Millis             sql.NullInt64
	Q3                   sql.NullString
	Q3Millis             sql.NullInt64
	Gap                  sql.NullString
	GapMillis            sql.NullInt64
	Interval             sql.NullString
	IntervalMillis       sql.NullInt64
	Laps                 sql.NullInt64
}

type SprintRaceResult struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	Laps                 sql.NullInt64
	Time                 sql.NullString
	TimeMillis           sql.NullInt64
	TimePenalty          sql.NullString
	TimePenaltyMillis    sql.NullInt64
	Gap                  sql.NullString
	GapMillis            sql.NullInt64
	GapLaps              sql.NullInt64
	Interval             sql.NullString
	IntervalMillis       sql.NullInt64
	ReasonRetired        sql.NullString
	Points               sql.NullFloat64
	GridPositionNumber   sql.NullInt64
	GridPositionText     sql.NullString
	PositionsGained      sql.NullInt64
}

type SprintStartingGridPosition struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	GridPenalty          sql.NullString
	GridPenaltyPositions sql.NullInt64
	Time                 sql.NullString
	TimeMillis           sql.NullInt64
}

type StartingGridPosition struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	GridPenalty          sql.NullString
	GridPenaltyPositions sql.NullInt64
	Time                 sql.NullString
	TimeMillis           sql.NullInt64
}

type TyreManufacturer struct {
	ID                       string
	Name                     string
	CountryID                string
	BestStartingGridPosition sql.NullInt64
	BestRaceResult           sql.NullInt64
	TotalRaceEntries         int64
	TotalRaceStarts          int64
	TotalRaceWins            int64
	TotalRaceLaps            int64
	TotalPodiums             int64
	TotalPodiumRaces         int64
	TotalPolePositions       int64
	TotalFastestLaps         int64
}

type WarmingUpResult struct {
	RaceID               int64
	Year                 int64
	Round                int64
	PositionDisplayOrder int64
	PositionNumber       sql.NullInt64
	PositionText         string
	DriverNumber         string
	DriverID             string
	ConstructorID        string
	EngineManufacturerID string
	TyreManufacturerID   string
	Time                 sql.NullString
	TimeMillis           sql.NullInt64
	Gap                  sql.NullString
	GapMillis            sql.NullInt64
	Interval             sql.NullString
	IntervalMillis       sql.NullInt64
	Laps                 sql.NullInt64
}
