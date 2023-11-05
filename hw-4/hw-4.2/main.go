package main

import (
    "fmt"
    "sync"
)
//Cтруктурa для хранения данных и мьютекса
type MySyncMap struct {
    data map[string]interface{}
    mutex   sync.Mutex
}
//Конструктор
func NewMySyncMap() *MySyncMap {
    return &MySyncMap{
        data: make(map[string]interface{}),
    }
}
// Load возвращает значение по ключу и флаг, указывающий, было ли значение найдено
func (m *MySyncMap) Load(key string) (interface{}, bool) {
    m.mutex.Lock() // Захватываю мьютекс для чтения
    defer m.mutex.Unlock()
    value, ok := m.data[key]
    return value, ok
}
// Store сохраняет значение по ключу
func (m *MySyncMap) Store(key string, value interface{}) {
    m.mutex.Lock() // Захватываю мьютекс для чтения
    defer m.mutex.Unlock()
    m.data[key] = value
}
// Delete удаляет значение по ключу
func (m *MySyncMap) Delete(key string) {
    m.mutex.Lock() // Захватываю мьютекс для чтения
    defer m.mutex.Unlock()
    delete(m.data, key)
}

func main() {
    myMap := NewMySyncMap()

    myMap.Store("Merey", "Zhumagul")
    myMap.Store("Alisher", "Salimgaliev")

    val, ok := myMap.Load("Merey")
    if ok {
        fmt.Println("Value for Merey:", val)
    } else {
        fmt.Println("Merey not found")
    }

    myMap.Delete("Alisher")

    val, ok = myMap.Load("Alisher")
    if ok {
        fmt.Println("Value for Alisher:", val)
    } else {
        fmt.Println("Alisher not found")
    }
}
