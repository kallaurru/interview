package talk

type Selectable interface {
	Up()
	Down()
	PgUp()
	PgDown()
	Selected() int
}
