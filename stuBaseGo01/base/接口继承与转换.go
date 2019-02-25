package base

import "fmt"

type humaner interface {
	say()
}
type boyer interface {
	game()
}
type maner interface {
	humaner
	boyer
}
type student11 struct {
	id   int
	name string
	age  int
}
type teacher struct {
	id   int
	name string
	age  int
}
func (st *student11) say(){
	fmt.Printf("hello,我是学生,我的名字叫%s \n",st.name)
}
func (st *student11) game(){
	fmt.Println("hello,我在玩游戏,我在偷偷玩游戏")
}
func (st *teacher) say(){
	fmt.Println("hello,我是老师,我的名字叫%s",st.name)
}
func (st *teacher) game(){
	fmt.Println("hello,我在玩游戏,我在光明正大的玩游戏")
}
func main() {
	st:=student11{1,"aa",11}
	var man maner
	var boy boyer
	man=&st
	man.say()
	man.game()
	boy=&st
	boy.game()
	boy=man//将超集转为子集
	boy.game()

}