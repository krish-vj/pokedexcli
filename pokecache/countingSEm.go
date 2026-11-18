package pokecache

import "sync"

type CntSeamaphore struct{
	cnt int 
	mtx sync.Mutex
}
func (sem *CntSeamaphore) signal(){
	sem.mtx.Lock()
	defer sem.mtx.Unlock()
	sem.cnt++
}
func (sem *CntSeamaphore) wait(){
	//-1 
	sem.mtx.Lock()
	defer sem.mtx.Unlock()
	if (sem.cnt==0){
		// we need to wait for it to be non zero; how do we wait?
		// i know how we can do busy waiting but i don't wanna do busy waiting
		// do we implement queues and channel; all guys who want to trigger wait; enter the queue;
		// each signal triggers the queue to do one of the waits??
	}
	sem.cnt--

}
