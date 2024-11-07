package data

import (
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
)

type Todo struct {
	ID              string
	Event           string
	ImportanceLevel int
	Completed       bool
}

// TodoManager 管理单个用户的Data
type TodoManager struct {
	dir   string
	mutex sync.RWMutex
}

type HandleTodo interface {
	findTodo(string) []Todo
	deleteTodo(string) error
	addTodo(uuid string, todo Todo) error
	updateTodo(former Todo, Later Todo) error
}

// newTodo 建立一个新的待办事项
func newTodo(Event string) *Todo {
	return &Todo{Event: Event, Completed: false}
}

// NewTodoManager 建立一个新的用户待办事项管理
func NewTodoManager(dir string) *TodoManager {
	return &TodoManager{dir: dir}
}

// getTodoPath 获取用户Todo文件的路径
func (m *TodoManager) getTodoPath(uuid string) string {
	return filepath.Join(m.dir, "user", uuid, "todo.json")
}

// ensurePathExistence 确保路径存在，如果不存在则创建路径
func (m *TodoManager) ensurePathExistence(uuid string) error {
	dir := filepath.Dir(m.getTodoPath(uuid))
	return os.MkdirAll(dir, 0755)
}

// ReadUserTodos 读取用户todo文件
func (m *TodoManager) ReadUserTodos(uuid string) ([]Todo, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	path := m.getTodoPath(uuid)
	// 如果不存在文件返回空列表
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []Todo{}, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return []Todo{}, err
	}
	var todos []Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return []Todo{}, err
	}
	return todos, nil
}

// SaveTheUserTodos 保存用户的所有todos（会删除开始的todo）
func (m *TodoManager) SaveTheUserTodos(uuid string, todos []Todo) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if err := m.ensurePathExistence(uuid); err != nil { // 注意，这里的错误是路径不存在，创建路径时的错误
		return err
	}
	// 创建临时文件，确保操作的原子性
	path := m.getTodoPath(uuid)
	tempPath := path + ".temp"
	data, err := json.MarshalIndent(todos, "", "    ")
	if err != nil {
		return err
	}
	if err = os.WriteFile(tempPath, data, 0644); err != nil {
		return err
	}
	return os.Rename(tempPath, path)
}

// AddTodo 添加单个待办事项
func (m *TodoManager) AddTodo(uuid string, todo Todo) error {
	todos, err := m.ReadUserTodos(uuid)
	if err != nil {
		return err
	}
	todos = append(todos, todo)
	err = m.SaveTheUserTodos(uuid, todos)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTodo 删除一个todo
func (m *TodoManager) DeleteTodo(uuid string, todo Todo) error {
	todos, err := m.ReadUserTodos(uuid)
	if err != nil {
		return err
	}
	newTodos := make([]Todo, 0, len(todos))
	for _, i := range todos {
		if i != todo {
			newTodos = append(newTodos, i)
		}
	}
	return m.SaveTheUserTodos(uuid, newTodos)
}

// RandomlySelectTodo 随机读取一个todo
func (m *TodoManager) RandomlySelectTodo(uuid string) (Todo, error) {
	todos, err := m.ReadUserTodos(uuid)
	if err != nil {
		return Todo{}, err
	}
	return todos[rand.Intn(len(todos)-1)], nil
}
