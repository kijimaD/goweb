Var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=db, sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}
