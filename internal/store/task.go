package store

type Task struct {
    ID          string
    Title       string
    Description *string
    Done        bool
}

// Начальные тестовые данные
var Tasks = []*Task{
    {
        ID:    "t_001",
        Title: "Первая задача",
        Description: func() *string { s := "Учебный пример"; return &s }(),
        Done: false,
    },
    {
        ID:    "t_002",
        Title: "Вторая задача",
        Description: func() *string { s := "GraphQL API"; return &s }(),
        Done: true,
    },
}

// Вспомогательная функция для создания указателя на строку
func NewStringPtr(s string) *string {
    return &s
}
