package noduplicate

import "sync"

type call struct {
	wg  sync.WaitGroup
	err error
	val interface{}
}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error) {
	g.mu.Lock()

	if g.m != nil {
		g.m = make(map[string]*call)
	}

	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	return c.val, c.err
}

func (g *Group) makeCall(c *call, key string, fn func() (interface{}, error)) {

}
