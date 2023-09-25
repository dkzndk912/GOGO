package main

func main() {
	HotSpurs := "hm ? j madi?"
	replacePlayer := string.NewReplacer("?", "son")
	player := replacePlayer.Replace(HotSpurs)
	fmt.println(player)
}
