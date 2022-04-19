package main

import "fmt"

// ПОЛИМОРФИЗМ в Go осуществляется путем использования интерфейсов по принципу
// утиной типизации - для компилятора все, у чего есть метод Name, будет человеком
//
// Human - это интерфейс описывающий Человека
// Любой тип, у которого есть метод Name с такой сигнатурой, будет реализовывать интерфейс Human
type Human interface {
	// У человека должно быть метод Name, возвращающий его имя
	Name() string
}

// Функция sayHelloToHuman принимает интерфейс Human как аргумент и здоровается с ним по имени
//
// Эта функция нужна для демонстрации полиморфизма
// Она должна будет одинаково успешно работать с разными типами
func sayHelloToHuman(human Human) {
	fmt.Println("Привет,", human.Name())
}

// Student - это тип, описывающий студента
type Student struct {
	// ИНКАПСУЛЯЦИЯ в Go происходит на уровне пакета
	//
	// Если поле или метод начинаются с маленькой буквы,
	// то они считаются приватными и доступны только внутри этого пакета
	// поэтому поле name является приватным, но метод Name публичным (потому что он с заглавной буквы)
	name string
}

// Name - это публичный метод, возвращающий приватное поле name
func (s Student) Name() string {
	return s.name
}

// Teacher - это тип, описывающий преподавателя
type Teacher struct {
	name string
}

// Name метод вернет имя преподователя
func (s Teacher) Name() string {
	return s.name
}

// Вместо НАСЛЕДОВАНИЯ в Go используется механизм встраивания
//
// Professor это тип, описывающий Профессора
// Профессор "наследует" базовый тип Преподователя (Teacher встроен в Professor)
// Но профессор дополнительно имеет научные достижения - поле scientificAchievements (массив строк)
type Professor struct {
	Teacher

	scientificAchievements []string
}

// ScientificAchievements вернет научные достижения профессора
func (p Professor) ScientificAchievements() []string {
	return p.scientificAchievements
}

// Точка входа в программу
func main() {
	// Создадим студента
	student := Student{
		name: "Студент",
	}

	// Создадим преподавателя
	teacher := Teacher{
		name: "Преподаватель",
	}

	// Поздороваемся со студентом
	sayHelloToHuman(student)

	// Поздороваемся с преподователем
	sayHelloToHuman(teacher)

	// Как видим, функция sayHelloToHuman успешно поздоровается как со студентом, так и с преподователем
	// Потому что оба эти типа имеют метод Name, то есть реализуют интерфейс Human

	// Теперь создадим профессора
	professor := Professor{
		Teacher: Teacher{
			name: "Профессор",
		},

		scientificAchievements: []string{
			"Достижение 1",
			"Достижение 2",
		},
	}

	// Как видим, функция sayHelloToHuman также успешно поздоровается с профессором, ведь
	// В тип Профессор встроен базовый тип Преподователя
	sayHelloToHuman(professor)
}

// Программа выведет
//
// Привет, Студент
// Привет, Преподаватель
// Привет, Профессор
