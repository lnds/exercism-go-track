package react

type ValueCell struct {
	id      int
	reactor *CellsReactor
}

func (cell ValueCell) Value() int {
	return cell.reactor.GetValue(cell.id)
}

func (cell ValueCell) SetValue(val int) {
	if cell.Value() == val {
		return
	}
	cell.reactor.SetValue(cell.id, val)
}

type ComputeCell1 struct {
	id         int
	reactor    *CellsReactor
	fn         func(int) int
	dependency Cell
	callbacks  map[int]func(int)
}

func (cell ComputeCell1) Value() int {
	val := cell.dependency.Value()
	oldValue := cell.reactor.GetCompute(cell.id)
	result := cell.fn(val)
	if result == oldValue {
		return oldValue
	}
	cell.reactor.SetCompute(cell.id, result)
	return result
}

func (cell ComputeCell1) AddCallback(fn func(int)) Canceler {
	return cell.reactor.AddCallback(cell.id, fn)
}

type ComputeCell2 struct {
	id          int
	reactor     *CellsReactor
	fn          func(int, int) int
	dependency1 Cell
	dependency2 Cell
}

type CallbackCanceler struct {
	cellId     int
	callbackId int
	reactor    *CellsReactor
}

func (canceler CallbackCanceler) Cancel() {
	// TODO
}

func (cell ComputeCell2) Value() int {
	val1 := cell.dependency1.Value()
	val2 := cell.dependency2.Value()
	result := cell.fn(val1, val2)
	return result
}

func (cell ComputeCell2) AddCallback(fn func(int)) Canceler {
	return cell.reactor.AddCallback(cell.id, fn)
}

type callback func(int)

type CellsReactor struct {
	values       []int
	computes     []int
	inputCells   []InputCell
	computeCells []ComputeCell
	callbacks    map[int]map[int]callback
}

func New() *CellsReactor {
	return &CellsReactor{callbacks: make(map[int]map[int]callback)}
}

func (r *CellsReactor) AddCallback(cellId int, fn callback) Canceler {
	m, ok := r.callbacks[cellId]
	if !ok {
		m = make(map[int]callback)
		r.callbacks[cellId] = m
	}
	callbackId := len(m)
	m[callbackId] = fn
	return CallbackCanceler{cellId: cellId, callbackId: callbackId, reactor: r}
}

func (r *CellsReactor) GetValue(id int) int {
	return r.values[id]
}

func (r *CellsReactor) SetValue(id, val int) {
	r.values[id] = val
	for cellId := range r.callbacks {
		oldValue := r.computes[cellId]
		newValue := r.computeCells[cellId].Value()
		if oldValue != newValue {
			for _, cb := range r.callbacks[cellId] {
				cb(newValue)
			}
		}
	}
}

func (r *CellsReactor) GetCompute(id int) int {
	return r.computes[id]
}

func (r *CellsReactor) SetCompute(id, val int) {
	r.computes[id] = val
}

func (r *CellsReactor) CreateInput(val int) InputCell {
	id := len(r.values)
	r.values = append(r.values, val)
	cell := ValueCell{id: id, reactor: r}
	r.inputCells = append(r.inputCells, cell)
	return cell
}

func (r *CellsReactor) CreateCompute1(dependency Cell, fn func(int) int) ComputeCell {
	cell := ComputeCell1{
		id:         len(r.computes),
		reactor:    r,
		fn:         fn,
		dependency: dependency,
		callbacks:  map[int]func(int){},
	}
	r.computes = append(r.computes, fn(dependency.Value()))
	r.computeCells = append(r.computeCells, cell)
	return cell
}

func (r *CellsReactor) CreateCompute2(cell1, cell2 Cell, fn func(int, int) int) ComputeCell {
	cell := ComputeCell2{
		id:          len(r.computes),
		reactor:     r,
		fn:          fn,
		dependency1: cell1,
		dependency2: cell2,
	}
	r.computes = append(r.computes, fn(cell1.Value(), cell2.Value()))
	r.computeCells = append(r.computeCells, cell)
	return cell
}
