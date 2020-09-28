//T *w //pointer point to first  unrefresh element
//T *r //pointer point to frist prefetch element
//T *f //pointer point to next be refresh element
// atomic_ptr <T> c //c  share atomic pointer

//constructor
func LockFreeQueue() {
	//push terminator element into queue
	//Initilize r,w,f and c to point the terminator
	//if in the queue is invaild the c to NULL
}
func Write(T value) {
	//Append the value to the queue
	//add new terminator element
	//update the postion of pointer f
}
func Flush() bool {
	//if there are no-flushed items,do nothing
	//try to compare and set c to f
	// if (above cas unsuc){
	// 	set c to f
	// 	set w to f/*  */
	// 	return false
	// }
	// set w to f
	// return true
}
func read() {
	// try to prefetch a value
	// if above prefetch faild{
	// 	return false
	// }
	// store the value of the front of the queue
	// pop tht front of the queue
	// return the value
}