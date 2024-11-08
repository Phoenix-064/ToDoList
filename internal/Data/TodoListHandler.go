package data

import (
	"ToDoList/internal/models"
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TodoManager 管理单个用户的Data
type TodoManager struct {
	dir   string
	mutex sync.RWMutex
}

type TodoGormManager struct {
	db *gorm.DB
}

type HandleTodo interface {
	ReadUserTodos(uuid string) ([]models.Todo, error)
	SaveTheUserTodos(uuid string, todos []models.Todo) error
	AddTodo(uuid string, todo models.Todo) error
	DeleteTodo(uuid string, todo models.Todo) error
	RandomlySelectTodo(uuid string) (models.Todo, error)
}

// NewTodo 建立一个新的待办事项
func NewTodo(id string, Event string) *models.Todo {
	return &models.Todo{ID: id, Event: Event, Completed: false}
}

// NewTodoManager 建立一个新的用户待办事项管理
func NewTodoManager(dir string) *TodoManager {
	return &TodoManager{dir: dir}
}

// NewTodoGormManager 返回一个新的 TodoGormManager
func NewTodoGormManager() *TodoGormManager {
	dsn := "root:123@tcp(127.0.0.1:3306)/todoList?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("无法连接数据库")
	}
	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		logrus.Error("无法连接数据库")
	}
	return &TodoGormManager{db: db}
}

// getTodoPath 获取用户Todo文件的路径
func (m *TodoManager) getTodoPath(uuid string) string {
	cleanId := strings.ReplaceAll(uuid, "-", "")
	return filepath.Join(m.dir, "user", cleanId, "todo.json")
}

// ensurePathExistence 确保路径存在，如果不存在则创建路径
func (m *TodoManager) ensurePathExistence(uuid string) error {
	dir := filepath.Dir(m.getTodoPath(uuid))
	return os.MkdirAll(dir, 0755)
}

// ReadUserTodos 读取用户todo文件
func (m *TodoManager) ReadUserTodos(uuid string) ([]models.Todo, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	path := m.getTodoPath(uuid)
	// 如果不存在文件返回空列表
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return []models.Todo{}, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return []models.Todo{}, err
	}
	var todos []models.Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return []models.Todo{}, err
	}
	return todos, nil
}

// SaveTheUserTodos 保存用户的所有todos（会删除开始的todo）
func (m *TodoManager) SaveTheUserTodos(uuid string, todos []models.Todo) error {
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
func (m *TodoManager) AddTodo(uuid string, todo models.Todo) error {
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
func (m *TodoManager) DeleteTodo(uuid string, todo models.Todo) error {
	todos, err := m.ReadUserTodos(uuid)
	if err != nil {
		return err
	}
	newTodos := make([]models.Todo, 0, len(todos))
	for _, i := range todos {
		if i != todo {
			newTodos = append(newTodos, i)
		}
	}
	return m.SaveTheUserTodos(uuid, newTodos)
}

// RandomlySelectTodo 随机读取一个todo
func (m *TodoManager) RandomlySelectTodo(uuid string) (models.Todo, error) {
	todos, err := m.ReadUserTodos(uuid)
	if err != nil {
		return models.Todo{}, err
	}
	if len(todos) == 0 {
		return models.Todo{}, nil
	}
	return todos[rand.Intn(len(todos))], nil
}

// ReadUserTodo 读取用户todos
func (m *TodoGormManager) ReadUserTodos(uuid string) ([]models.Todo, error) {
	var user models.User
	result := m.db.Preload("Todos").First(&user, "uuid = ?", uuid)
	if result.Error != nil {
		return []models.Todo{}, result.Error
	}
	todos := user.Todo
	return todos, nil
}

// SaveTheUserTodos 保存所有todos
func (m *TodoGormManager) SaveTheUserTodos(uuid string, todos []models.Todo) error {
	return nil
}
