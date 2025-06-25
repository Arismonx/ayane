package main

func main() {
	Display()
	ayane := Ayane{}
	storage := NewStroage[Ayane]("todos.json")
	cmd := NewCmdFlags()
	storage.Load(&ayane)
	cmd.Execute(&ayane)
	storage.Save(&ayane)
}
