package main

var currentId int

var messages Messages

// initial fake data
func init() {
	RepoCreateMessage(Message{Author: "Homer", Content: "I don't believe in communication"})
	RepoCreateMessage(Message{Author: "Lisa", Content: "Give it a chance"})
}

/*
func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}
*/

func RepoCreateMessage(m Message) Message {
	currentId += 1
	m.Id = currentId
	messages = append(messages, m)
	return m
}

/*
func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(messages[:i], messages[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
*/
