package main

import ()

func main() {
	todos := Todos{}
	todos.add("Buy milk")
	todos.add("Running")
	todos.toggle(0)
	todos.print()
}
