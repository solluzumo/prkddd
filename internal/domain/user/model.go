package user

type User struct {
	ID   string
	Name string
}

// permissions значения:
// 1 - чтение собственных записей
// 2 - добавление записей себе
// 3 - чтение записей любого пользователя
// 4 - добавление записей любому пользователю
// 5 - редактирование всех таблиц
type Role struct {
	ID          string
	Name        string
	Permissions []int
}
