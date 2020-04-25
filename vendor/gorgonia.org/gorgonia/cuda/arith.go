package cuda

import (
	"unsafe"

	"github.com/pkg/errors"
	"gorgonia.org/cu"
	"gorgonia.org/tensor"
)

// Code generated by gencudaengine, which is a API generation tool for Gorgonia. DO NOT EDIT.

// Add implements tensor.Adder. It does not support safe or increment operation options and will return an error if those options are passed in
func (e *Engine) Add(a tensor.Tensor, b tensor.Tensor, opts ...tensor.FuncOpt) (retVal tensor.Tensor, err error) {
	name := constructName2(a, b, "add")

	if !e.HasFunc(name) {
		return nil, errors.Errorf("Unable to perform Add(). The tensor engine does not have the function %q", name)
	}

	if err = binaryCheck(a, b); err != nil {
		return nil, errors.Wrap(err, "Basic checks failed for Add")
	}

	var reuse tensor.DenseTensor
	var safe, toReuse bool
	if reuse, safe, toReuse, _, _, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), true, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}

	var mem, memB cu.DevicePtr
	var size int64

	switch {
	case toReuse:
		mem = cu.DevicePtr(reuse.Uintptr())
		memA := cu.DevicePtr(a.Uintptr())
		memSize := int64(a.MemSize())
		e.memcpy(mem, memA, memSize)

		size = int64(logicalSize(reuse.Shape()))
		retVal = reuse
	case !safe:
		mem = cu.DevicePtr(a.Uintptr())
		retVal = a
		size = int64(logicalSize(a.Shape()))
	default:
		return nil, errors.New("Impossible state: A reuse tensor must be passed in, or the operation must be unsafe. Incr and safe operations are not supported")
	}

	memB = cu.DevicePtr(b.Uintptr())
	fn := e.f[name]
	gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ := e.ElemGridSize(int(size))
	args := []unsafe.Pointer{
		unsafe.Pointer(&mem),
		unsafe.Pointer(&memB),
		unsafe.Pointer(&size),
	}
	logf("gx %d, gy %d, gz %d | bx %d by %d, bz %d", gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ)
	logf("CUDADO %q, Mem: %v MemB: %v size %v, args %v", name, mem, memB, size, args)
	logf("LaunchKernel Params. mem: %v. Size %v", mem, size)
	e.c.LaunchAndSync(fn, gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ, 0, cu.NoStream, args)
	return
}

// AddScalar implements tensor.Adder. It does not support safe or increment operation options and will return an error if those options are passed in
func (e *Engine) AddScalar(a tensor.Tensor, b interface{}, leftTensor bool, opts ...tensor.FuncOpt) (retVal tensor.Tensor, err error) {
	name := constructName1(a, leftTensor, "add")
	if !e.HasFunc(name) {
		return nil, errors.Errorf("Unable to perform AddScalar(). The tensor engine does not have the function %q", name)
	}

	var bMem tensor.Memory
	var ok bool
	if bMem, ok = b.(tensor.Memory); !ok {
		return nil, errors.Errorf("b has to be a tensor.Memory. Got %T instead", b)
	}

	if err = unaryCheck(a); err != nil {
		return nil, errors.Wrap(err, "Basic checks failed for AddScalar")
	}

	var reuse tensor.DenseTensor
	var safe, toReuse bool
	if reuse, safe, toReuse, _, _, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), true, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}

	var mem, memB cu.DevicePtr
	var size int64

	switch {
	case toReuse:
		mem = cu.DevicePtr(reuse.Uintptr())
		memA := cu.DevicePtr(a.Uintptr())
		memSize := int64(a.MemSize())
		e.memcpy(mem, memA, memSize)

		size = int64(logicalSize(reuse.Shape()))
		retVal = reuse
	case !safe:
		mem = cu.DevicePtr(a.Uintptr())
		retVal = a
		size = int64(logicalSize(a.Shape()))
	default:
		return nil, errors.New("Impossible state: A reuse tensor must be passed in, or the operation must be unsafe. Incr and safe operations are not supported")
	}

	memB = cu.DevicePtr(bMem.Uintptr())
	if !leftTensor {
		mem, memB = memB, mem
	}

	fn := e.f[name]
	gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ := e.ElemGridSize(int(size))
	args := []unsafe.Pointer{
		unsafe.Pointer(&mem),
		unsafe.Pointer(&memB),
		unsafe.Pointer(&size),
	}
	logf("gx %d, gy %d, gz %d | bx %d by %d, bz %d", gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ)
	logf("CUDADO %q, Mem: %v size %v, args %v", name, mem, size, args)
	logf("LaunchKernel Params. mem: %v. Size %v", mem, size)
	e.c.LaunchAndSync(fn, gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ, 0, cu.NoStream, args)
	return
}

// Sub implements tensor.Suber. It does not support safe or increment operation options and will return an error if those options are passed in
func (e *Engine) Sub(a tensor.Tensor, b tensor.Tensor, opts ...tensor.FuncOpt) (retVal tensor.Tensor, err error) {
	name := constructName2(a, b, "sub")

	if !e.HasFunc(name) {
		return nil, errors.Errorf("Unable to perform Sub(). The tensor engine does not have the function %q", name)
	}

	if err = binaryCheck(a, b); err != nil {
		return nil, errors.Wrap(err, "Basic checks failed for Sub")
	}

	var reuse tensor.DenseTensor
	var safe, toReuse bool
	if reuse, safe, toReuse, _, _, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), true, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}

	var mem, memB cu.DevicePtr
	var size int64

	switch {
	case toReuse:
		mem = cu.DevicePtr(reuse.Uintptr())
		memA := cu.DevicePtr(a.Uintptr())
		memSize := int64(a.MemSize())
		e.memcpy(mem, memA, memSize)

		size = int64(logicalSize(reuse.Shape()))
		retVal = reuse
	case !safe:
		mem = cu.DevicePtr(a.Uintptr())
		retVal = a
		size = int64(logicalSize(a.Shape()))
	default:
		return nil, errors.New("Impossible state: A reuse tensor must be passed in, or the operation must be unsafe. Incr and safe operations are not supported")
	}

	memB = cu.DevicePtr(b.Uintptr())
	fn := e.f[name]
	gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ := e.ElemGridSize(int(size))
	args := []unsafe.Pointer{
		unsafe.Pointer(&mem),
		unsafe.Pointer(&memB),
		unsafe.Pointer(&size),
	}
	logf("gx %d, gy %d, gz %d | bx %d by %d, bz %d", gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ)
	logf("CUDADO %q, Mem: %v MemB: %v size %v, args %v", name, mem, memB, size, args)
	logf("LaunchKernel Params. mem: %v. Size %v", mem, size)
	e.c.LaunchAndSync(fn, gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ, 0, cu.NoStream, args)
	return
}

// SubScalar implements tensor.Suber. It does not support safe or increment operation options and will return an error if those options are passed in
func (e *Engine) SubScalar(a tensor.Tensor, b interface{}, leftTensor bool, opts ...tensor.FuncOpt) (retVal tensor.Tensor, err error) {
	name := constructName1(a, leftTensor, "sub")
	if !e.HasFunc(name) {
		return nil, errors.Errorf("Unable to perform SubScalar(). The tensor engine does not have the function %q", name)
	}

	var bMem tensor.Memory
	var ok bool
	if bMem, ok = b.(tensor.Memory); !ok {
		return nil, errors.Errorf("b has to be a tensor.Memory. Got %T instead", b)
	}

	if err = unaryCheck(a); err != nil {
		return nil, errors.Wrap(err, "Basic checks failed for SubScalar")
	}

	var reuse tensor.DenseTensor
	var safe, toReuse bool
	if reuse, safe, toReuse, _, _, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), true, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}

	var mem, memB cu.DevicePtr
	var size int64

	switch {
	case toReuse:
		mem = cu.DevicePtr(reuse.Uintptr())
		memA := cu.DevicePtr(a.Uintptr())
		memSize := int64(a.MemSize())
		e.memcpy(mem, memA, memSize)

		size = int64(logicalSize(reuse.Shape()))
		retVal = reuse
	case !safe:
		mem = cu.DevicePtr(a.Uintptr())
		retVal = a
		size = int64(logicalSize(a.Shape()))
	default:
		return nil, errors.New("Impossible state: A reuse tensor must be passed in, or the operation must be unsafe. Incr and safe operations are not supported")
	}

	memB = cu.DevicePtr(bMem.Uintptr())
	if !leftTensor {
		mem, memB = memB, mem
	}

	fn := e.f[name]
	gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ := e.ElemGridSize(int(size))
	args := []unsafe.Pointer{
		unsafe.Pointer(&mem),
		unsafe.Pointer(&memB),
		unsafe.Pointer(&size),
	}
	logf("gx %d, gy %d, gz %d | bx %d by %d, bz %d", gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ)
	logf("CUDADO %q, Mem: %v size %v, args %v", name, mem, size, args)
	logf("LaunchKernel Params. mem: %v. Size %v", mem, size)
	e.c.LaunchAndSync(fn, gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ, 0, cu.NoStream, args)
	return
}

// Mul implements tensor.Muler. It does not support safe or increment operation options and will return an error if those options are passed in
func (e *Engine) Mul(a tensor.Tensor, b tensor.Tensor, opts ...tensor.FuncOpt) (retVal tensor.Tensor, err error) {
	name := constructName2(a, b, "mul")

	if !e.HasFunc(name) {
		return nil, errors.Errorf("Unable to perform Mul(). The tensor engine does not have the function %q", name)
	}

	if err = binaryCheck(a, b); err != nil {
		return nil, errors.Wrap(err, "Basic checks failed for Mul")
	}

	var reuse tensor.DenseTensor
	var safe, toReuse bool
	if reuse, safe, toReuse, _, _, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), true, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}

	var mem, memB cu.DevicePtr
	var size int64

	switch {
	case toReuse:
		mem = cu.DevicePtr(reuse.Uintptr())
		memA := cu.DevicePtr(a.Uintptr())
		memSize := int64(a.MemSize())
		e.memcpy(mem, memA, memSize)

		size = int64(logicalSize(reuse.Shape()))
		retVal = reuse
	case !safe:
		mem = cu.DevicePtr(a.Uintptr())
		retVal = a
		size = int64(logicalSize(a.Shape()))
	default:
		return nil, errors.New("Impossible state: A reuse tensor must be passed in, or the operation must be unsafe. Incr and safe operations are not supported")
	}

	memB = cu.DevicePtr(b.Uintptr())
	fn := e.f[name]
	gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ := e.ElemGridSize(int(size))
	args := []unsafe.Pointer{
		unsafe.Pointer(&mem),
		unsafe.Pointer(&memB),
		unsafe.Pointer(&size),
	}
	logf("gx %d, gy %d, gz %d | bx %d by %d, bz %d", gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ)
	logf("CUDADO %q, Mem: %v MemB: %v size %v, args %v", name, mem, memB, size, args)
	logf("LaunchKernel Params. mem: %v. Size %v", mem, size)
	e.c.LaunchAndSync(fn, gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ, 0, cu.NoStream, args)
	return
}

// MulScalar implements tensor.Muler. It does not support safe or increment operation options and will return an error if those options are passed in
func (e *Engine) MulScalar(a tensor.Tensor, b interface{}, leftTensor bool, opts ...tensor.FuncOpt) (retVal tensor.Tensor, err error) {
	name := constructName1(a, leftTensor, "mul")
	if !e.HasFunc(name) {
		return nil, errors.Errorf("Unable to perform MulScalar(). The tensor engine does not have the function %q", name)
	}

	var bMem tensor.Memory
	var ok bool
	if bMem, ok = b.(tensor.Memory); !ok {
		return nil, errors.Errorf("b has to be a tensor.Memory. Got %T instead", b)
	}

	if err = unaryCheck(a); err != nil {
		return nil, errors.Wrap(err, "Basic checks failed for MulScalar")
	}

	var reuse tensor.DenseTensor
	var safe, toReuse bool
	if reuse, safe, toReuse, _, _, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), true, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}

	var mem, memB cu.DevicePtr
	var size int64

	switch {
	case toReuse:
		mem = cu.DevicePtr(reuse.Uintptr())
		memA := cu.DevicePtr(a.Uintptr())
		memSize := int64(a.MemSize())
		e.memcpy(mem, memA, memSize)

		size = int64(logicalSize(reuse.Shape()))
		retVal = reuse
	case !safe:
		mem = cu.DevicePtr(a.Uintptr())
		retVal = a
		size = int64(logicalSize(a.Shape()))
	default:
		return nil, errors.New("Impossible state: A reuse tensor must be passed in, or the operation must be unsafe. Incr and safe operations are not supported")
	}

	memB = cu.DevicePtr(bMem.Uintptr())
	if !leftTensor {
		mem, memB = memB, mem
	}

	fn := e.f[name]
	gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ := e.ElemGridSize(int(size))
	args := []unsafe.Pointer{
		unsafe.Pointer(&mem),
		unsafe.Pointer(&memB),
		unsafe.Pointer(&size),
	}
	logf("gx %d, gy %d, gz %d | bx %d by %d, bz %d", gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ)
	logf("CUDADO %q, Mem: %v size %v, args %v", name, mem, size, args)
	logf("LaunchKernel Params. mem: %v. Size %v", mem, size)
	e.c.LaunchAndSync(fn, gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ, 0, cu.NoStream, args)
	return
}

// Div implements tensor.Diver. It does not support safe or increment operation options and will return an error if those options are passed in
func (e *Engine) Div(a tensor.Tensor, b tensor.Tensor, opts ...tensor.FuncOpt) (retVal tensor.Tensor, err error) {
	name := constructName2(a, b, "div")

	if !e.HasFunc(name) {
		return nil, errors.Errorf("Unable to perform Div(). The tensor engine does not have the function %q", name)
	}

	if err = binaryCheck(a, b); err != nil {
		return nil, errors.Wrap(err, "Basic checks failed for Div")
	}

	var reuse tensor.DenseTensor
	var safe, toReuse bool
	if reuse, safe, toReuse, _, _, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), true, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}

	var mem, memB cu.DevicePtr
	var size int64

	switch {
	case toReuse:
		mem = cu.DevicePtr(reuse.Uintptr())
		memA := cu.DevicePtr(a.Uintptr())
		memSize := int64(a.MemSize())
		e.memcpy(mem, memA, memSize)

		size = int64(logicalSize(reuse.Shape()))
		retVal = reuse
	case !safe:
		mem = cu.DevicePtr(a.Uintptr())
		retVal = a
		size = int64(logicalSize(a.Shape()))
	default:
		return nil, errors.New("Impossible state: A reuse tensor must be passed in, or the operation must be unsafe. Incr and safe operations are not supported")
	}

	memB = cu.DevicePtr(b.Uintptr())
	fn := e.f[name]
	gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ := e.ElemGridSize(int(size))
	args := []unsafe.Pointer{
		unsafe.Pointer(&mem),
		unsafe.Pointer(&memB),
		unsafe.Pointer(&size),
	}
	logf("gx %d, gy %d, gz %d | bx %d by %d, bz %d", gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ)
	logf("CUDADO %q, Mem: %v MemB: %v size %v, args %v", name, mem, memB, size, args)
	logf("LaunchKernel Params. mem: %v. Size %v", mem, size)
	e.c.LaunchAndSync(fn, gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ, 0, cu.NoStream, args)
	return
}

// DivScalar implements tensor.Diver. It does not support safe or increment operation options and will return an error if those options are passed in
func (e *Engine) DivScalar(a tensor.Tensor, b interface{}, leftTensor bool, opts ...tensor.FuncOpt) (retVal tensor.Tensor, err error) {
	name := constructName1(a, leftTensor, "div")
	if !e.HasFunc(name) {
		return nil, errors.Errorf("Unable to perform DivScalar(). The tensor engine does not have the function %q", name)
	}

	var bMem tensor.Memory
	var ok bool
	if bMem, ok = b.(tensor.Memory); !ok {
		return nil, errors.Errorf("b has to be a tensor.Memory. Got %T instead", b)
	}

	if err = unaryCheck(a); err != nil {
		return nil, errors.Wrap(err, "Basic checks failed for DivScalar")
	}

	var reuse tensor.DenseTensor
	var safe, toReuse bool
	if reuse, safe, toReuse, _, _, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), true, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}

	var mem, memB cu.DevicePtr
	var size int64

	switch {
	case toReuse:
		mem = cu.DevicePtr(reuse.Uintptr())
		memA := cu.DevicePtr(a.Uintptr())
		memSize := int64(a.MemSize())
		e.memcpy(mem, memA, memSize)

		size = int64(logicalSize(reuse.Shape()))
		retVal = reuse
	case !safe:
		mem = cu.DevicePtr(a.Uintptr())
		retVal = a
		size = int64(logicalSize(a.Shape()))
	default:
		return nil, errors.New("Impossible state: A reuse tensor must be passed in, or the operation must be unsafe. Incr and safe operations are not supported")
	}

	memB = cu.DevicePtr(bMem.Uintptr())
	if !leftTensor {
		mem, memB = memB, mem
	}

	fn := e.f[name]
	gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ := e.ElemGridSize(int(size))
	args := []unsafe.Pointer{
		unsafe.Pointer(&mem),
		unsafe.Pointer(&memB),
		unsafe.Pointer(&size),
	}
	logf("gx %d, gy %d, gz %d | bx %d by %d, bz %d", gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ)
	logf("CUDADO %q, Mem: %v size %v, args %v", name, mem, size, args)
	logf("LaunchKernel Params. mem: %v. Size %v", mem, size)
	e.c.LaunchAndSync(fn, gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ, 0, cu.NoStream, args)
	return
}

// Pow implements tensor.Power. It does not support safe or increment operation options and will return an error if those options are passed in
func (e *Engine) Pow(a tensor.Tensor, b tensor.Tensor, opts ...tensor.FuncOpt) (retVal tensor.Tensor, err error) {
	name := constructName2(a, b, "pow")

	if !e.HasFunc(name) {
		return nil, errors.Errorf("Unable to perform Pow(). The tensor engine does not have the function %q", name)
	}

	if err = binaryCheck(a, b); err != nil {
		return nil, errors.Wrap(err, "Basic checks failed for Pow")
	}

	var reuse tensor.DenseTensor
	var safe, toReuse bool
	if reuse, safe, toReuse, _, _, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), true, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}

	var mem, memB cu.DevicePtr
	var size int64

	switch {
	case toReuse:
		mem = cu.DevicePtr(reuse.Uintptr())
		memA := cu.DevicePtr(a.Uintptr())
		memSize := int64(a.MemSize())
		e.memcpy(mem, memA, memSize)

		size = int64(logicalSize(reuse.Shape()))
		retVal = reuse
	case !safe:
		mem = cu.DevicePtr(a.Uintptr())
		retVal = a
		size = int64(logicalSize(a.Shape()))
	default:
		return nil, errors.New("Impossible state: A reuse tensor must be passed in, or the operation must be unsafe. Incr and safe operations are not supported")
	}

	memB = cu.DevicePtr(b.Uintptr())
	fn := e.f[name]
	gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ := e.ElemGridSize(int(size))
	args := []unsafe.Pointer{
		unsafe.Pointer(&mem),
		unsafe.Pointer(&memB),
		unsafe.Pointer(&size),
	}
	logf("gx %d, gy %d, gz %d | bx %d by %d, bz %d", gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ)
	logf("CUDADO %q, Mem: %v MemB: %v size %v, args %v", name, mem, memB, size, args)
	logf("LaunchKernel Params. mem: %v. Size %v", mem, size)
	e.c.LaunchAndSync(fn, gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ, 0, cu.NoStream, args)
	return
}

// PowScalar implements tensor.Power. It does not support safe or increment operation options and will return an error if those options are passed in
func (e *Engine) PowScalar(a tensor.Tensor, b interface{}, leftTensor bool, opts ...tensor.FuncOpt) (retVal tensor.Tensor, err error) {
	name := constructName1(a, leftTensor, "pow")
	if !e.HasFunc(name) {
		return nil, errors.Errorf("Unable to perform PowScalar(). The tensor engine does not have the function %q", name)
	}

	var bMem tensor.Memory
	var ok bool
	if bMem, ok = b.(tensor.Memory); !ok {
		return nil, errors.Errorf("b has to be a tensor.Memory. Got %T instead", b)
	}

	if err = unaryCheck(a); err != nil {
		return nil, errors.Wrap(err, "Basic checks failed for PowScalar")
	}

	var reuse tensor.DenseTensor
	var safe, toReuse bool
	if reuse, safe, toReuse, _, _, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), true, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}

	var mem, memB cu.DevicePtr
	var size int64

	switch {
	case toReuse:
		mem = cu.DevicePtr(reuse.Uintptr())
		memA := cu.DevicePtr(a.Uintptr())
		memSize := int64(a.MemSize())
		e.memcpy(mem, memA, memSize)

		size = int64(logicalSize(reuse.Shape()))
		retVal = reuse
	case !safe:
		mem = cu.DevicePtr(a.Uintptr())
		retVal = a
		size = int64(logicalSize(a.Shape()))
	default:
		return nil, errors.New("Impossible state: A reuse tensor must be passed in, or the operation must be unsafe. Incr and safe operations are not supported")
	}

	memB = cu.DevicePtr(bMem.Uintptr())
	if !leftTensor {
		mem, memB = memB, mem
	}

	fn := e.f[name]
	gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ := e.ElemGridSize(int(size))
	args := []unsafe.Pointer{
		unsafe.Pointer(&mem),
		unsafe.Pointer(&memB),
		unsafe.Pointer(&size),
	}
	logf("gx %d, gy %d, gz %d | bx %d by %d, bz %d", gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ)
	logf("CUDADO %q, Mem: %v size %v, args %v", name, mem, size, args)
	logf("LaunchKernel Params. mem: %v. Size %v", mem, size)
	e.c.LaunchAndSync(fn, gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ, 0, cu.NoStream, args)
	return
}

// Mod implements tensor.Moder. It does not support safe or increment operation options and will return an error if those options are passed in
func (e *Engine) Mod(a tensor.Tensor, b tensor.Tensor, opts ...tensor.FuncOpt) (retVal tensor.Tensor, err error) {
	name := constructName2(a, b, "mod")

	if !e.HasFunc(name) {
		return nil, errors.Errorf("Unable to perform Mod(). The tensor engine does not have the function %q", name)
	}

	if err = binaryCheck(a, b); err != nil {
		return nil, errors.Wrap(err, "Basic checks failed for Mod")
	}

	var reuse tensor.DenseTensor
	var safe, toReuse bool
	if reuse, safe, toReuse, _, _, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), true, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}

	var mem, memB cu.DevicePtr
	var size int64

	switch {
	case toReuse:
		mem = cu.DevicePtr(reuse.Uintptr())
		memA := cu.DevicePtr(a.Uintptr())
		memSize := int64(a.MemSize())
		e.memcpy(mem, memA, memSize)

		size = int64(logicalSize(reuse.Shape()))
		retVal = reuse
	case !safe:
		mem = cu.DevicePtr(a.Uintptr())
		retVal = a
		size = int64(logicalSize(a.Shape()))
	default:
		return nil, errors.New("Impossible state: A reuse tensor must be passed in, or the operation must be unsafe. Incr and safe operations are not supported")
	}

	memB = cu.DevicePtr(b.Uintptr())
	fn := e.f[name]
	gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ := e.ElemGridSize(int(size))
	args := []unsafe.Pointer{
		unsafe.Pointer(&mem),
		unsafe.Pointer(&memB),
		unsafe.Pointer(&size),
	}
	logf("gx %d, gy %d, gz %d | bx %d by %d, bz %d", gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ)
	logf("CUDADO %q, Mem: %v MemB: %v size %v, args %v", name, mem, memB, size, args)
	logf("LaunchKernel Params. mem: %v. Size %v", mem, size)
	e.c.LaunchAndSync(fn, gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ, 0, cu.NoStream, args)
	return
}

// ModScalar implements tensor.Moder. It does not support safe or increment operation options and will return an error if those options are passed in
func (e *Engine) ModScalar(a tensor.Tensor, b interface{}, leftTensor bool, opts ...tensor.FuncOpt) (retVal tensor.Tensor, err error) {
	name := constructName1(a, leftTensor, "mod")
	if !e.HasFunc(name) {
		return nil, errors.Errorf("Unable to perform ModScalar(). The tensor engine does not have the function %q", name)
	}

	var bMem tensor.Memory
	var ok bool
	if bMem, ok = b.(tensor.Memory); !ok {
		return nil, errors.Errorf("b has to be a tensor.Memory. Got %T instead", b)
	}

	if err = unaryCheck(a); err != nil {
		return nil, errors.Wrap(err, "Basic checks failed for ModScalar")
	}

	var reuse tensor.DenseTensor
	var safe, toReuse bool
	if reuse, safe, toReuse, _, _, err = handleFuncOpts(a.Shape(), a.Dtype(), a.DataOrder(), true, opts...); err != nil {
		return nil, errors.Wrap(err, "Unable to handle funcOpts")
	}

	var mem, memB cu.DevicePtr
	var size int64

	switch {
	case toReuse:
		mem = cu.DevicePtr(reuse.Uintptr())
		memA := cu.DevicePtr(a.Uintptr())
		memSize := int64(a.MemSize())
		e.memcpy(mem, memA, memSize)

		size = int64(logicalSize(reuse.Shape()))
		retVal = reuse
	case !safe:
		mem = cu.DevicePtr(a.Uintptr())
		retVal = a
		size = int64(logicalSize(a.Shape()))
	default:
		return nil, errors.New("Impossible state: A reuse tensor must be passed in, or the operation must be unsafe. Incr and safe operations are not supported")
	}

	memB = cu.DevicePtr(bMem.Uintptr())
	if !leftTensor {
		mem, memB = memB, mem
	}

	fn := e.f[name]
	gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ := e.ElemGridSize(int(size))
	args := []unsafe.Pointer{
		unsafe.Pointer(&mem),
		unsafe.Pointer(&memB),
		unsafe.Pointer(&size),
	}
	logf("gx %d, gy %d, gz %d | bx %d by %d, bz %d", gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ)
	logf("CUDADO %q, Mem: %v size %v, args %v", name, mem, size, args)
	logf("LaunchKernel Params. mem: %v. Size %v", mem, size)
	e.c.LaunchAndSync(fn, gridDimX, gridDimY, gridDimZ, blockDimX, blockDimY, blockDimZ, 0, cu.NoStream, args)
	return
}
