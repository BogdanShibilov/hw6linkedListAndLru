package main

func main() {
	c := NewCache(-1)

	_ = c.Set("aaa", 100)
	_ = c.Set("bbb", 200)
	_ = c.Set("ccc", 300)
	_ = c.Set("ddd", 400)
	_, _ = c.Get("bbb")

	c.PrintQueue()
}
