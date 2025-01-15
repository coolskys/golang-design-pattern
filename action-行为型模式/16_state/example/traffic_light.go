package example

type LightController interface {
	State() string             // 展示当前状态
	Next(*TrafficLight) string // 行为
}

type TrafficLight struct {
	State LightController
}

func NewTrafficLight(state LightController) *TrafficLight {
	if state == nil {
		state = NewGreenLightState()
	}
	return &TrafficLight{
		State: state,
	}
}

// 改变状态的方法
func (t *TrafficLight) changeState(state LightController) {
	t.State = state
}

type DefaultLightState struct {
	stateName string
}

type redLightState struct {
	DefaultLightState
}

func NewRedLightState() *redLightState {
	return &redLightState{
		DefaultLightState: DefaultLightState{
			stateName: "红灯",
		},
	}
}

func (r *redLightState) State() string {
	return r.stateName
}

func (r *redLightState) Next(light *TrafficLight) string {
	y := NewYellowLightState()
	light.changeState(y)
	return y.stateName
}

type yellowLightState struct {
	DefaultLightState
}

func NewYellowLightState() *yellowLightState {
	return &yellowLightState{
		DefaultLightState: DefaultLightState{
			stateName: "黄灯",
		},
	}
}

func (y *yellowLightState) State() string {
	return y.stateName
}

func (y *yellowLightState) Next(light *TrafficLight) string {
	g := NewGreenLightState()
	light.changeState(g)
	return g.stateName
}

type greenLightState struct {
	DefaultLightState
}

func NewGreenLightState() *greenLightState {
	return &greenLightState{
		DefaultLightState: DefaultLightState{
			stateName: "绿灯",
		},
	}
}

func (g *greenLightState) State() string {
	return g.stateName
}

func (g *greenLightState) Next(light *TrafficLight) string {
	r := NewRedLightState()
	light.changeState(r)
	return r.stateName
}
