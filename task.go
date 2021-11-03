package main

type JcmdTask struct {
	PathJcmd     string
	PathExtaArgs string // TODO
	MainClass    string
	SubSystem    string
	TimerMs      int
	TimeoutMs    int
	Metrics      *metricMap
}

func NewJcmdTask() *JcmdTask {

	return &JcmdTask{}
}
