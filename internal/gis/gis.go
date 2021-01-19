package gis

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	lg "github.com/sirupsen/logrus"
)

// DataBase info
type ParamDB struct {
	Base *sql.DB
	Log  lg.FieldLogger
}

// Data for configuration of DB connecttion
type Config struct {
	User string
	Pass string
	Db   string
	Host string
	Port string
}

// Book info
type City struct {
	Id     int64
	Title  string
	Coords string
}

// Create connection to db
func ConnectToDB(conf Config, log lg.FieldLogger) (*ParamDB, error) {
	var err error = nil
	var par ParamDB
	par.Base = nil
	par.Log = log

	par.Log.Println("Start connection to database.")

	if conf.User == "" || conf.Pass == "" || conf.Db == "" ||
		conf.Host == "" || conf.Port == "" {
		err = errors.New("Bad connection parameters")
		return &par, err
	}

	// Connect to user database
	connstr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		conf.User, conf.Pass, conf.Host, conf.Port, conf.Db)
	par.Log.Printf("Connection string: %s.", connstr)

	db, err := openDB(connstr)
	if err != nil {
		log.Fatalf("GIS: open DB error: %s.", err.Error())
		err = errors.New("Bad connection to user db")
		return &par, err
	}
	par.Base = db

	par.Log.Println("Set connection to database.")
	return &par, nil
}

// Close user database
func (par *ParamDB) Close() (err error) {
	if par.Base == nil {
		return nil
	}
	par.Log.Printf("Close database: %v.", par)

	if err = par.Base.Close(); err != nil {
		return err
	}

	par.Base = nil
	return nil
}

// Check connect to DB
func openDB(confstr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", confstr)

	if err != nil {
		err = errors.New("Bad connection to db")
		return nil, err
	}

	// Ping the database
	if err = db.Ping(); err != nil {
		err = errors.New("Couldn't ping the database")
		db.Close()
		return nil, err
	}
	return db, nil
}

// Insert city info into database
func (b *City) InsertCity(par *ParamDB) (int64, error) {
	if par.Base == nil {
		return 0, errors.New("DB not opened")
	}

	query := fmt.Sprintf("INSERT INTO city_records (title, coords) VALUES('%s', '%s') RETURNING id;", b.Title, b.Coords)
	par.Log.Printf("DB query: %s.", query)
	row := par.Base.QueryRow(query)

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, errors.New("Error insert city - bad id")
	}
	return id, nil
}

// Delete city from database
func (b *City) DeleteCity(par *ParamDB) error {
	if par.Base == nil {
		return errors.New("DB not opened")
	}

	if b.Id <= 0 {
		return errors.New("Id not set")
	}

	query := fmt.Sprintf("DELETE FROM city_records where id = %v;", b.Id)
	_, err := par.Base.Exec(query)

	if err != nil {
		return errors.New("Error delete city")
	}
	return nil
}

// Update city in the database
func (b *City) UpdateCity(par *ParamDB) error {
	if par.Base == nil {
		return errors.New("DB not opened")
	}

	if b.Id <= 0 {
		return errors.New("Id not set")
	}

	var query string
	switch {
	case b.Coords != "" && b.Title != "":
		query = fmt.Sprintf("UPDATE city_records SET coords = '%s', title = '%s' WHERE id = %v;", b.Coords, b.Title, b.Id)
	case b.Coords != "":
		query = fmt.Sprintf("UPDATE city_records SET author = '%s' WHERE id = %v;", b.Coords, b.Id)
	case b.Title != "":
		query = fmt.Sprintf("UPDATE city_records SET title = '%s' WHERE id = %v;", b.Title, b.Id)
	default:
		return nil
	}
	fmt.Println(query)
	_, err := par.Base.Exec(query)

	if err != nil {
		return errors.New("Error update city")
	}
	return nil
}

// Select city from database
func (b *City) SelectCity(par *ParamDB) (*[]City, error) {
	bb := make([]City, 0)
	if par.Base == nil {
		return &bb, errors.New("DB not opened")
	}

	var query string
	if b.Id > 0 {
		query = fmt.Sprintf("SELECT id, title, coords FROM city_records WHERE id = %v ORDER BY id;", b.Id)
	} else {
		query = fmt.Sprintf("SELECT id, title, coords FROM city_records ORDER BY id;")
	}

	rows, err := par.Base.Query(query)
	if err != nil {
		return &bb, errors.New("Select error")
	}
	defer rows.Close()

	for rows.Next() {
		var city City
		if err := rows.Scan(&city.Id, &city.Title, &city.Coords); err != nil {
			return &bb, errors.New("Select scan error")
		}
		bb = append(bb, city)
	}

	if err != nil {
		return &bb, errors.New("Error update city")
	}
	return &bb, nil
}
