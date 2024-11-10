package data

import (
	"ToDoList/internal/models"
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"sync"

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
	SaveTheUserTodos(uuid string, todos []*models.Todo) error
	AddTodo(uuid string, todo *models.Todo) error
	DeleteTodo(uuid string, todoID string) error
	RandomlySelectTodo(uuid string) (models.Todo, error)
	UpdateTodo(userUUID string, todoID string, todo *models.Todo) error
}

// NewTodo 建立一个新的待办事项
func NewTodo(id string, Event string, isCycle bool, description string, importanceLevel int) *models.Todo {
	return &models.Todo{
		ID:              id,
		Event:           Event,
		Description:     description,
		Completed:       false,
		IsCycle:         isCycle,
		ImportanceLevel: importanceLevel,
	}
}

// NewTodoManager 建立一个新的用户待办事项管理
func NewTodoManager(dir string) *TodoManager {
	return &TodoManager{dir: dir}
}

// NewTodoGormManager 返回一个新的 TodoGormManager
func NewTodoGormManager(db *gorm.DB) *TodoGormManager {
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
func (m *TodoManager) DeleteTodo(uuid string, todoID string) error {
	todos, err := m.ReadUserTodos(uuid)
	if err != nil {
		return err
	}
	newTodos := make([]models.Todo, 0, len(todos))
	for _, i := range todos {
		if i.ID != todoID {
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
	var todos []models.Todo
	result := m.db.Where("user_uuid = ?", uuid).Find(&todos)
	if result.Error != nil {
		return []models.Todo{}, result.Error
	}
	return todos, nil
}

// SaveTheUserTodos 保存所有 todos (会删除之前的所有 todos)
func (m *TodoGormManager) SaveTheUserTodos(uuid string, todos []*models.Todo) error {
	// 使用事务，确保操作原子性
	return m.db.Transaction(func(tx *gorm.DB) error {
		if result := tx.Where("user_uuid = ?", uuid).Delete(&models.Todo{}); result.Error != nil {
			return result.Error
		}
		// 如果提交结果为空，直接返回
		if len(todos) == 0 {
			return nil
		}
		// 给 todos 的 uuid 字段赋值
		for i := range todos {
			todos[i].UserUuid = uuid
		}
		if result := m.db.Create(&todos); result.Error != nil {
			return result.Error
		}
		return nil
	})
}

// AddTodo 添加一个 todo
func (m *TodoGormManager) AddTodo(uuid string, todo *models.Todo) error {
	todo.UserUuid = uuid
	if result := m.db.Create(&todo); result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteTodo 删除一个 todo
func (m *TodoGormManager) DeleteTodo(uuid string, todoID string) error {
	if result := m.db.Where("user_uuid = ? AND id = ?", uuid, todoID).Delete(&models.Todo{}); result.Error != nil {
		return result.Error
	}
	return nil
}

// RandomlySelectTodo 获取一个随机 todo
func (m *TodoGormManager) RandomlySelectTodo(uuid string) (models.Todo, error) {
	var todo models.Todo
	if result := m.db.Where("user_uuid = ? is_wish = ?", uuid, true).Order("RANDOM()").First(&todo); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Todo{}, result.Error
		}
	}
	return todo, nil
}

// UpdateTodo 更新一个 todo
func (m *TodoGormManager) UpdateTodo(userUUID string, todoID string, todo *models.Todo) error {
	todo.UserUuid = userUUID
	if result := m.db.Where("user_uuid = ? AND id = ?", userUUID, todoID).Updates(todo); result.Error != nil {
		return result.Error
	}
	return nil
}
