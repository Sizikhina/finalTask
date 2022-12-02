package service

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"zlata/circular_rotation"
	"zlata/client"
	"zlata/entry_into_the_array"
	"zlata/find_missing_element"
	"zlata/sequence_check"
)

const rotation = "Циклическая ротация"
const entryIntoArray = "Чудные вхождения в массив"
const findElem = "Поиск отсутствующего элемента"
const check = "Проверка последовательности"

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	r.Get("/task/{taskName}", solveTask)
	r.Get("/tasks", solveAllTasks)

	return r
}

func solveTask(w http.ResponseWriter, r *http.Request) {
	taskName := chi.URLParam(r, "taskName")
	data := client.GetSolutionData(taskName)

	solutionCheck := client.SolutionCheck{
		UserName: "zlata_siz",
		Task:     taskName,
	}

	solutionCheck.Results.Payload = data
	var solutions [10]interface{}

	switch taskName {
	case entryIntoArray:
		parsedData := ParseIntSlices(data)
		for i, dataset := range parsedData {
			solutions[i] = entry_into_the_array.Solution(dataset)
		}

	case check:
		parsedData := ParseIntSlices(data)
		for i, dataset := range parsedData {
			solutions[i] = sequence_check.Solution(dataset)
		}
	case findElem:
		parsedData := ParseIntSlices(data)
		for i, dataset := range parsedData {
			solutions[i] = find_missing_element.Solution(dataset)
		}
	case rotation:
		parsedSlices, parsedNumbers := ParseIntSlicesAndInts(data)
		for i, dataset := range parsedSlices {
			solutions[i] = circular_rotation.Solution(dataset, parsedNumbers[i])
		}
	default:
		panic("unknown task")
	}
	solutionCheck.Results.Results = solutions

	result := client.CheckSolution(solutionCheck)
	b, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	w.Write(b)
}

type Task struct {
	Name          string      `json:"name"`
	SolutionCheck interface{} `json:"solution_check"`
}

func solveAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	var solutions1 [10]interface{}
	var solutions2 [10]interface{}
	var solutions3 [10]interface{}
	var solutions4 [10]interface{}

	//entryIntoArray
	taskName := entryIntoArray
	data := client.GetSolutionData(taskName)

	solutionCheck := client.SolutionCheck{
		UserName: "zlata_siz",
		Task:     taskName,
	}

	solutionCheck.Results.Payload = data
	parsedData := ParseIntSlices(data)

	for i, dataset := range parsedData {
		solutions1[i] = entry_into_the_array.Solution(dataset)
	}

	solutionCheck.Results.Results = solutions1
	result := client.CheckSolution(solutionCheck)

	tasks = append(tasks, Task{
		Name:          taskName,
		SolutionCheck: result,
	})

	//check
	taskName = check
	data = client.GetSolutionData(taskName)

	solutionCheck = client.SolutionCheck{
		UserName: "zlata_siz",
		Task:     taskName,
	}

	solutionCheck.Results.Payload = data
	parsedData = ParseIntSlices(data)

	for i, dataset := range parsedData {
		solutions2[i] = sequence_check.Solution(dataset)
	}

	solutionCheck.Results.Results = solutions2
	result = client.CheckSolution(solutionCheck)

	tasks = append(tasks, Task{
		Name:          taskName,
		SolutionCheck: result,
	})

	//findElem
	taskName = findElem
	data = client.GetSolutionData(taskName)

	solutionCheck = client.SolutionCheck{
		UserName: "zlata_siz",
		Task:     taskName,
	}

	solutionCheck.Results.Payload = data
	parsedData = ParseIntSlices(data)

	for i, dataset := range parsedData {
		solutions3[i] = find_missing_element.Solution(dataset)
	}

	solutionCheck.Results.Results = solutions3
	result = client.CheckSolution(solutionCheck)

	tasks = append(tasks, Task{
		Name:          taskName,
		SolutionCheck: result,
	})

	//rotation
	taskName = rotation
	data = client.GetSolutionData(taskName)

	solutionCheck = client.SolutionCheck{
		UserName: "zlata_siz",
		Task:     taskName,
	}
	solutionCheck.Results.Payload = data
	parsedSlices, parsedNumbers := ParseIntSlicesAndInts(data)

	for i, dataset := range parsedSlices {
		solutions4[i] = circular_rotation.Solution(dataset, parsedNumbers[i])
	}

	solutionCheck.Results.Results = solutions4
	result = client.CheckSolution(solutionCheck)

	tasks = append(tasks, Task{
		Name:          taskName,
		SolutionCheck: result,
	})

	b, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}
	w.Write(b)

	//в дальнейшем можно улучшить, обернув все это в цикл, так как код повторяется 4 раза
}

func ParseIntSlices(data [10][]json.RawMessage) [10][]int {
	var result [10][]int

	for i, d := range data {
		//мы знаем, что d -массив из одного элемента, который является массивом int
		if len(d) != 1 {
			panic("invalid length")
		}
		var s []int
		err := json.Unmarshal(d[0], &s)
		if err != nil {
			panic(err)
		}
		result[i] = s
	}
	return result
}

func ParseIntSlicesAndInts(data [10][]json.RawMessage) ([10][]int, [10]int) {
	var a [10][]int
	var k [10]int

	for i, d := range data {
		//мы знаем, что d -массив из двух элементов, один из которых - массив int, второй - int
		if len(d) != 2 {
			panic("invalid length")
		}
		var s []int
		err := json.Unmarshal(d[0], &s)
		if err != nil {
			panic(err)
		}

		var n int
		err = json.Unmarshal(d[1], &n)
		if err != nil {
			panic(err)
		}
		a[i] = s
		k[i] = n
	}
	return a, k
}
