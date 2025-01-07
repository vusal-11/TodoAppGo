package todoapp

type User struct {
    Id       int    `json:"-" db:"id"`                        // Это поле не будет сериализоваться в JSON
    Name     string `json:"name" binding:"required"`  // Поле будет обязательным
    Username string `json:"username" binding:"required"` // Исправлено закрытие кавычки
    Password string `json:"password" binding:"required"` // Исправлено тег json
}