## 100 Advanced Go Channel Problems (Medium–Hard) with Solutions

Notes
- Focus: channels-centric problems. Use goroutines/select/timeouts/cancellation where helpful.
- Solutions are concise idioms/patterns. Adapt into full programs as needed.

### 1) Bounded buffer with backpressure
Problem: Build a bounded queue using only channels that blocks producers when full.
Solution:
```go
func NewBounded[T any](cap int) (chan<- T, <-chan T) {
    buf := make(chan T, cap)
    return buf, buf
}
// Using the same channel blocks senders when full and receivers when empty.
```

### 2) Fan-in with cancellation
Problem: Merge many `<-chan T` into one `out` that stops when `done` is closed.
Solution:
```go
func FanIn[T any](done <-chan struct{}, inputs ...<-chan T) <-chan T {
    out := make(chan T)
    var wg sync.WaitGroup
    wg.Add(len(inputs))
    for _, ch := range inputs {
        ch := ch
        go func() {
            defer wg.Done()
            for v := range ch {
                select { case out <- v: case <-done: return }
            }
        }()
    }
    go func(){ wg.Wait(); close(out) }()
    return out
}
```

### 3) Fan-out workers with backpressure
Problem: Start N workers reading from one jobs channel, pushing results to one results channel.
Solution:
```go
func FanOut[I any, O any](n int, jobs <-chan I, work func(I) O) <-chan O {
    out := make(chan O)
    var wg sync.WaitGroup
    wg.Add(n)
    for i := 0; i < n; i++ {
        go func(){
            defer wg.Done()
            for j := range jobs {
                out <- work(j)
            }
        }()
    }
    go func(){ wg.Wait(); close(out) }()
    return out
}
```

### 4) Tee channel
Problem: Split one input channel to two outputs duplicating each value.
Solution:
```go
func Tee[T any](in <-chan T) (<-chan T, <-chan T) {
    a, b := make(chan T), make(chan T)
    go func(){
        defer close(a); defer close(b)
        for v := range in {
            va, vb := v, v
            // ensure both sends complete per value
            for sentA, sentB := false, false; !(sentA && sentB); {
                select {
                case a <- va: sentA = true
                case b <- vb: sentB = true
                }
            }
        }
    }()
    return a, b
}
```

### 5) Or-channel (cancellation races)
Problem: Combine multiple done channels so closing any closes the result.
Solution:
```go
func Or(chs ...<-chan struct{}) <-chan struct{} {
    switch len(chs) {
    case 0: return nil
    case 1: return chs[0]
    }
    or := make(chan struct{})
    go func(){
        defer close(or)
        select {
        case <-chs[0]:
        case <-Or(chs[1:]...):
        }
    }()
    return or
}
```

### 6) Timeout on receive
Problem: Receive from `in` but abort if no value in d.
Solution:
```go
funcRecvWithTimeout[T any](in <-chan T, d time.Duration) (T, bool) {
    var zero T
    select {
    case v, ok := <-in: return v, ok
    case <-time.After(d): return zero, false
    }
}
```

### 7) Timeout on send
Problem: Send to `out` but give up after d.
Solution:
```go
func SendWithTimeout[T any](out chan<- T, v T, d time.Duration) bool {
    select {
    case out <- v: return true
    case <-time.After(d): return false
    }
}
```

### 8) Drop-oldest buffer
Problem: Bounded channel that drops oldest element when full.
Solution:
```go
type DropOldest[T any] struct{ c chan T }

func NewDropOldest[T any](cap int) *DropOldest[T] { return &DropOldest[T]{c: make(chan T, cap)} }
func (d *DropOldest[T]) In() chan<- T { return d.c }
func (d *DropOldest[T]) Out() <-chan T { return d.c }
func (d *DropOldest[T]) TrySend(v T) {
    select { case d.c <- v: return default: }
    // full, drop one
    select { case <-d.c: default: }
    d.c <- v
}
```

### 9) Bridge channel-of-channels
Problem: Flatten a stream of channels into a single output channel.
Solution:
```go
func Bridge[T any](done <-chan struct{}, chanStream <-chan <-chan T) <-chan T {
    out := make(chan T)
    go func(){
        defer close(out)
        for ch := range chanStream {
            for v := range ch {
                select { case out <- v: case <-done: return }
            }
        }
    }()
    return out
}
```

### 10) Ordered results with unordered workers
Problem: Preserve input order when workers finish out of order.
Solution:
```go
type job[I any] struct{ idx int; val I }
type res[O any] struct{ idx int; val O }

func Ordered[I any, O any](in []I, n int, f func(I) O) []O {
    jobs := make(chan job[I])
    outs := make(chan res[O])
    go func(){
        for i, v := range in { jobs <- job[I]{i, v} }
        close(jobs)
    }()
    var wg sync.WaitGroup
    wg.Add(n)
    for i := 0; i < n; i++ { go func(){ defer wg.Done(); for j := range jobs { outs <- res[O]{j.idx, f(j.val)} } }() }
    go func(){ wg.Wait(); close(outs) }()
    out := make([]O, len(in))
    for r := range outs { out[r.idx] = r.val }
    return out
}
```

### 11) Single-flight send
Problem: Ensure only one goroutine performs a send action; others wait.
Solution:
```go
func SingleFlightSend[T any](once *sync.Once, out chan<- T, v T) {
    once.Do(func(){ out <- v })
}
```

### 12) Multiplex to N consumers fairly
Problem: Distribute messages from one input evenly to N outputs.
Solution:
```go
func Distribute[T any](in <-chan T, outs []chan<- T) {
    go func(){
        defer func(){ for _, o := range outs { close(o) } }()
        i := 0
        for v := range in {
            outs[i%len(outs)] <- v
            i++
        }
    }()
}
```

### 13) Non-blocking try-receive
Problem: Poll a channel without blocking.
Solution:
```go
func TryRecv[T any](ch <-chan T) (T, bool) {
    var zero T
    select { case v := <-ch: return v, true; default: return zero, false }
}
```

### 14) Non-blocking try-send
Problem: Attempt send that fails fast if not ready.
Solution:
```go
func TrySend[T any](ch chan<- T, v T) bool {
    select { case ch <- v: return true; default: return false }
}
```

### 15) Debounce events
Problem: Coalesce bursts so you emit only once after quiet period.
Solution:
```go
func Debounce[T any](in <-chan T, d time.Duration) <-chan T {
    out := make(chan T)
    go func(){
        defer close(out)
        var t *time.Timer
        var last T
        pending := false
        reset := func(){ if t != nil { t.Stop() }; t = time.NewTimer(d) }
        for {
            if !pending { select { case v, ok := <-in:
                    if !ok { return }
                    last = v; pending = true; reset()
                }
            } else {
                select {
                case v, ok := <-in: if !ok { out <- last; return }; last = v; reset()
                case <-t.C: out <- last; pending = false
                }
            }
        }
    }()
    return out
}
```

### 16) Throttle (rate limit) via ticker
Problem: Emit at most one message per interval.
Solution:
```go
func Throttle[T any](in <-chan T, every time.Duration) <-chan T {
    out := make(chan T)
    tick := time.NewTicker(every)
    go func(){
        defer close(out); defer tick.Stop()
        var buf []T
        for {
            var sendC chan<- T
            var sendV T
            if len(buf) > 0 { sendC = out; sendV = buf[0] }
            select {
            case v, ok := <-in:
                if !ok { for _, x := range buf { out <- x }; return }
                buf = append(buf, v)
            case <-tick.C:
                if len(buf) > 0 { out <- sendV; buf = buf[1:] }
            }
        }
    }()
    return out
}
```

### 17) Token bucket using channels
Problem: Implement token bucket with capacity C and refill R/sec.
Solution:
```go
type Bucket struct{ tokens chan struct{} }

func NewBucket(capacity int, refill time.Duration) *Bucket {
    b := &Bucket{tokens: make(chan struct{}, capacity)}
    for i := 0; i < capacity; i++ { b.tokens <- struct{}{} }
    go func(){ t := time.NewTicker(refill); defer t.Stop(); for range t.C { select { case b.tokens <- struct{}{}: default: } } }()
    return b
}
func (b *Bucket) Take(ctx context.Context) error {
    select { case <-ctx.Done(): return ctx.Err(); case <-b.tokens: return nil }
}
```

### 18) Gate: allow at most K concurrent operations
Problem: Use a counting semaphore via channels.
Solution:
```go
type Gate struct{ sem chan struct{} }
func NewGate(k int) *Gate { return &Gate{sem: make(chan struct{}, k)} }
func (g *Gate) Enter(){ g.sem <- struct{}{} }
func (g *Gate) Leave(){ <-g.sem }
```

### 19) Wait for N signals without sync.WaitGroup
Problem: Use channels to wait for N goroutines.
Solution:
```go
func WaitN(n int, fn func(int)) {
    done := make(chan struct{})
    for i := 0; i < n; i++ { go func(i int){ fn(i); done <- struct{}{} }(i) }
    for i := 0; i < n; i++ { <-done }
}
```

### 20) Sliding window over stream
Problem: Maintain last W items and emit window on each new item.
Solution:
```go
func Sliding[T any](in <-chan T, w int) <-chan []T {
    out := make(chan []T)
    go func(){
        defer close(out)
        var buf []T
        for v := range in {
            buf = append(buf, v)
            if len(buf) > w { buf = buf[1:] }
            if len(buf) == w { cp := append([]T(nil), buf...); out <- cp }
        }
    }()
    return out
}
```

### 21) Dynamic fan-in/fan-out with add/remove
Problem: Allow adding/removing inputs at runtime.
Solution:
```go
type addRem[T any] struct{ add <-chan T; rem <-chan struct{} }
// Simplify: manage a slice of chans and rebuild select by ranging.
func DynamicFanIn[T any](done <-chan struct{}, control <-chan addRem[T]) <-chan T {
    out := make(chan T)
    var inputs []<-chan T
    go func(){
        defer close(out)
        for {
            select {
            case cmd, ok := <-control:
                if !ok { return }
                if cmd.add != nil { inputs = append(inputs, cmd.add) }
                // rem omitted for brevity
            default:
                for _, ch := range inputs {
                    select { case v := <-ch: out <- v; default: }
                }
                select { case <-done: return; default: time.Sleep(time.Millisecond) }
            }
        }
    }()
    return out
}
```

### 22) Drain channel safely
Problem: Empty a channel until closed without blocking forever.
Solution:
```go
func Drain[T any](ch <-chan T) []T {
    var out []T
    for v := range ch { out = append(out, v) }
    return out
}
```

### 23) Signal once then broadcast many
Problem: Convert a one-shot signal into a broadcast to many listeners.
Solution:
```go
func BroadcastOnce() (trigger func(), listeners func() <-chan struct{}) {
    var once sync.Once
    subs := make([]chan struct{}, 0)
    listeners = func() <-chan struct{} { c := make(chan struct{}); subs = append(subs, c); return c }
    trigger = func(){ once.Do(func(){ for _, c := range subs { close(c) } }) }
    return
}
```

### 24) Round-robin multiplexer with fairness under slow consumers
Problem: Ensure no consumer starves when some are slow.
Solution:
```go
func RoundRobin[T any](in <-chan T, outs []chan<- T) {
    go func(){
        defer func(){ for _, o := range outs { close(o) } }()
        i := 0
        for v := range in {
            for tries := 0; tries < len(outs); tries++ {
                j := (i + tries) % len(outs)
                select { case outs[j] <- v: i = j + 1; goto next
                default: }
            }
            // if all full, block on next
            outs[i%len(outs)] <- v
            i++
        next: }
    }()
}
```

### 25) Barrier: wait until M of N have arrived
Problem: Release when M signals received.
Solution:
```go
func Barrier(M int) (arrive chan<- struct{}, released <-chan struct{}) {
    a := make(chan struct{})
    r := make(chan struct{})
    go func(){
        count := 0
        for range a {
            count++
            if count == M { close(r); return }
        }
    }()
    return a, r
}
```

### 26) Timeout-aware worker pool
Problem: Each job must complete within d or be dropped.
Solution:
```go
func PoolWithTimeout[I any, O any](n int, jobs <-chan I, d time.Duration, f func(I) O) <-chan O {
    out := make(chan O)
    var wg sync.WaitGroup
    wg.Add(n)
    for i := 0; i < n; i++ {
        go func(){
            defer wg.Done()
            for j := range jobs {
                done := make(chan O, 1)
                go func(j I){ done <- f(j) }(j)
                select { case v := <-done: out <- v; case <-time.After(d): }
            }
        }()
    }
    go func(){ wg.Wait(); close(out) }()
    return out
}
```

### 27) Retry with backoff via channels
Problem: Retry worker on failure with backoff; cancelable.
Solution:
```go
func Retry[T any](ctx context.Context, attempts int, base time.Duration, op func() (T, error)) (T, error) {
    var zero T
    for i := 0; i < attempts; i++ {
        v, err := op(); if err == nil { return v, nil }
        t := time.NewTimer(base << i)
        select { case <-ctx.Done(): t.Stop(); return zero, ctx.Err(); case <-t.C: }
    }
    return zero, errors.New("exhausted")
}
```

### 28) Pipeline cancellation propagation
Problem: Ensure closing `done` stops all stages and closes outputs.
Solution:
```go
func Stage[T any](done <-chan struct{}, in <-chan T, f func(T) T) <-chan T {
    out := make(chan T)
    go func(){
        defer close(out)
        for v := range in {
            select { case out <- f(v): case <-done: return }
        }
    }()
    return out
}
```

### 29) Priority queue using two channels
Problem: Prefer reads from `high` over `low` when both ready.
Solution:
```go
func Priority[ T any ](high, low <-chan T) <-chan T {
    out := make(chan T)
    go func(){
        defer close(out)
        for {
            select {
            case v, ok := <-high: if !ok { high = nil } else { out <- v }
            default:
                select {
                case v, ok := <-high: if !ok { high = nil } else { out <- v }
                case v, ok := <-low: if !ok { low = nil } else { out <- v }
                }
            }
            if high == nil && low == nil { return }
        }
    }()
    return out
}
```

### 30) Fan-in with per-source tagging
Problem: Preserve source id when merging channels.
Solution:
```go
type Tagged[T any] struct{ Src int; Val T }
func FanInTagged[T any](chs ...<-chan T) <-chan Tagged[T] {
    out := make(chan Tagged[T])
    var wg sync.WaitGroup
    wg.Add(len(chs))
    for i, ch := range chs { i, ch := i, ch; go func(){ defer wg.Done(); for v := range ch { out <- Tagged[T]{i, v} } }() }
    go func(){ wg.Wait(); close(out) }()
    return out
}
```

### 31) Merge sorted streams preserving order
Problem: Merge two sorted `<-chan int` into sorted output.
Solution:
```go
func MergeSorted(a, b <-chan int) <-chan int {
    out := make(chan int)
    go func(){
        defer close(out)
        var va, vb int; oka, okb := false, false
        for {
            if !oka && a != nil { if v, ok := <-a; ok { va, oka = v, true } else { a = nil } }
            if !okb && b != nil { if v, ok := <-b; ok { vb, okb = v, true } else { b = nil } }
            if !oka && !okb { return }
            if oka && (!okb || va <= vb) { out <- va; oka = false } else { out <- vb; okb = false }
        }
    }()
    return out
}
```

### 32) Duplicate suppression (dedupe)
Problem: Suppress duplicate consecutive values.
Solution:
```go
func Dedupe[T comparable](in <-chan T) <-chan T {
    out := make(chan T)
    go func(){
        defer close(out)
        var last T; have := false
        for v := range in {
            if !have || v != last { out <- v; last, have = v, true }
        }
    }()
    return out
}
```

### 33) Window by time (tumbling)
Problem: Emit items collected in fixed intervals.
Solution:
```go
func Tumbling[T any](in <-chan T, d time.Duration) <-chan []T {
    out := make(chan []T)
    go func(){
        defer close(out)
        t := time.NewTicker(d); defer t.Stop()
        var buf []T
        for {
            select {
            case v, ok := <-in:
                if !ok { if len(buf) > 0 { out <- append([]T(nil), buf...) }; return }
                buf = append(buf, v)
            case <-t.C:
                if len(buf) > 0 { out <- append([]T(nil), buf...); buf = buf[:0] }
            }
        }
    }()
    return out
}
```

### 34) Zip channels
Problem: Pair items from two channels stepwise until one closes.
Solution:
```go
func Zip[A any, B any](a <-chan A, b <-chan B) <-chan struct{A A; B B} {
    out := make(chan struct{A A; B B})
    go func(){
        defer close(out)
        for {
            va, oka := <-a
            vb, okb := <-b
            if !oka || !okb { return }
            out <- struct{A A; B B}{va, vb}
        }
    }()
    return out
}
```

### 35) Unzip channel of pairs
Problem: Split pair stream into two outputs.
Solution:
```go
func Unzip[A any, B any](in <-chan struct{A A; B B}) (<-chan A, <-chan B) {
    a, b := make(chan A), make(chan B)
    go func(){ defer close(a); defer close(b); for v := range in { a <- v.A; b <- v.B } }()
    return a, b
}
```

### 36) Sticky sessions (send per key to same worker)
Problem: Route items with same key to same worker.
Solution:
```go
func Sticky[K comparable, V any](in <-chan struct{K K; V V}, workers []chan<- V) {
    go func(){
        defer func(){ for _, w := range workers { close(w) } }()
        for kv := range in {
            idx := int(fnv.New32a().Sum32()) // placeholder, use a real hash of kv.K
            w := workers[idx%len(workers)]
            w <- kv.V
        }
    }()
}
```

### 37) Split by predicate
Problem: Route items to `yes` or `no` channels by predicate.
Solution:
```go
func Split[T any](in <-chan T, pred func(T) bool) (<-chan T, <-chan T) {
    yes, no := make(chan T), make(chan T)
    go func(){ defer close(yes); defer close(no); for v := range in { if pred(v) { yes <- v } else { no <- v } } }()
    return yes, no
}
```

### 38) First-N items then stop
Problem: Take first N values then close out.
Solution:
```go
func Take[T any](in <-chan T, n int) <-chan T {
    out := make(chan T)
    go func(){ defer close(out); for i := 0; i < n; i++ { v, ok := <-in; if !ok { return }; out <- v } }()
    return out
}
```

### 39) Skip-N items then forward
Problem: Drop first N items then forward all.
Solution:
```go
func Skip[T any](in <-chan T, n int) <-chan T {
    out := make(chan T)
    go func(){ defer close(out); for i := 0; i < n; i++ { if _, ok := <-in; !ok { return } } for v := range in { out <- v } }()
    return out
}
```

### 40) Latest-value channel (conflation)
Problem: Keep only the latest value; slow consumers get the newest.
Solution:
```go
type Latest[T any] struct{ in chan T; out chan T }
func NewLatest[T any]() *Latest[T] {
    l := &Latest[T]{in: make(chan T), out: make(chan T)}
    go func(){
        defer close(l.out)
        var cur T
        have := false
        for {
            var outC chan<- T
            if have { outC = l.out }
            select {
            case v, ok := <-l.in: if !ok { return }; cur = v; have = true
            case outC <- cur: have = false
            }
        }
    }()
    return l
}
```

### 41) Broadcast each value to N listeners with per-listener buffers
Problem: Avoid slow listener blocking others.
Solution:
```go
func PubSub[T any](in <-chan T, n, buf int) []<-chan T {
    outs := make([]chan T, n)
    for i := range outs { outs[i] = make(chan T, buf) }
    go func(){
        defer func(){ for _, o := range outs { close(o) } }()
        for v := range in {
            for _, o := range outs {
                select { case o <- v: default: /* drop for this listener */ }
            }
        }
    }()
    rs := make([]<-chan T, n)
    for i := range outs { rs[i] = outs[i] }
    return rs
}
```

### 42) Timeout for whole pipeline
Problem: Cancel pipeline if total runtime exceeds D.
Solution:
```go
func WithDeadline[T any](in <-chan T, d time.Duration) <-chan T {
    out := make(chan T)
    go func(){
        defer close(out)
        timer := time.NewTimer(d); defer timer.Stop()
        for {
            select { case v, ok := <-in: if !ok { return }; out <- v
            case <-timer.C: return }
        }
    }()
    return out
}
```

### 43) Race two producers, take first response
Problem: Query two backends; take the fastest.
Solution:
```go
func Race[T any](a, b func() T) T {
    ch := make(chan T, 2)
    go func(){ ch <- a() }(); go func(){ ch <- b() }()
    return <-ch
}
```

### 44) Context-aware receive
Problem: Receive unless context done.
Solution:
```go
func RecvCtx[T any](ctx context.Context, ch <-chan T) (T, bool) {
    var zero T
    select { case v, ok := <-ch: return v, ok; case <-ctx.Done(): return zero, false }
}
```

### 45) Context-aware send
Problem: Send unless context done.
Solution:
```go
func SendCtx[T any](ctx context.Context, ch chan<- T, v T) bool {
    select { case ch <- v: return true; case <-ctx.Done(): return false }
}
```

### 46) Work stealing among worker channels
Problem: Idle workers can pull from others' queues.
Solution:
```go
func WorkSteal[T any](queues []<-chan T) <-chan T {
    out := make(chan T)
    go func(){
        defer close(out)
        for {
            progressed := false
            for _, q := range queues {
                select { case v := <-q: out <- v; progressed = true; default: }
            }
            if !progressed { time.Sleep(time.Millisecond) }
        }
    }()
    return out
}
```

### 47) Split load by weighted probability
Problem: Send 70% to A and 30% to B.
Solution:
```go
func Weighted[T any](in <-chan T, a, b chan<- T, p float64) {
    go func(){ defer close(a); defer close(b); for v := range in { if rand.Float64() < p { a <- v } else { b <- v } } }()
}
```

### 48) Barrier per key (sharded barrier)
Problem: Wait per key until M arrivals per key.
Solution:
```go
func BarrierPerKey[K comparable](in <-chan K, M int) <-chan K {
    out := make(chan K)
    go func(){
        defer close(out)
        counts := map[K]int{}
        for k := range in {
            counts[k]++
            if counts[k] == M { out <- k; delete(counts, k) }
        }
    }()
    return out
}
```

### 49) Close channel from multiple goroutines safely
Problem: Ensure close happens only once.
Solution:
```go
type SafeCloser[T any] struct{ C chan T; once sync.Once }
func (s *SafeCloser[T]) Close(){ s.once.Do(func(){ close(s.C) }) }
```

### 50) Map-reduce with channels
Problem: Map items concurrently then reduce sequentially.
Solution:
```go
func MapReduce[I any, O any](in []I, mapN int, mapFn func(I) O, redFn func(O, O) O) O {
    jobs := make(chan I)
    outs := FanOut[I,O](mapN, jobs, mapFn)
    go func(){ for _, v := range in { jobs <- v }; close(jobs) }()
    var acc O; first := true
    for v := range outs { if first { acc, first = v, false } else { acc = redFn(acc, v) } }
    return acc
}
```

### 51) Cancelable generator
Problem: Generate sequence until done.
Solution:
```go
func Generator(done <-chan struct{}, start, step int) <-chan int {
    out := make(chan int)
    go func(){ defer close(out); v := start; for { select { case out <- v: v += step; case <-done: return } } }()
    return out
}
```

### 52) Backpressure across pipeline boundaries
Problem: Ensure upstream slows when downstream is slow.
Solution:
```go
// Use unbuffered channels (or small buffers) between stages to propagate backpressure.
```

### 53) Swap channel (rendezvous exchange)
Problem: Pair up goroutines to exchange values.
Solution:
```go
type exchanger[T any] struct{ ch chan T }
func NewExchanger[T any]() *exchanger[T] { return &exchanger[T]{make(chan T)} }
func (e *exchanger[T]) Exchange(v T) T { e.ch <- v; return <-e.ch }
```

### 54) Safe multiplex close when unknown producers
Problem: Many producers send to single `out` then indicate completion.
Solution:
```go
type MultiOut[T any] struct{ out chan T; done chan struct{}; wg sync.WaitGroup }
func NewMultiOut[T any]() *MultiOut[T] { return &MultiOut[T]{out: make(chan T), done: make(chan struct{})} }
func (m *MultiOut[T]) AddProducer() chan<- T { m.wg.Add(1); ch := make(chan T); go func(){ defer m.wg.Done(); for v := range ch { m.out <- v } }(); return ch }
func (m *MultiOut[T]) Close(){ m.wg.Wait(); close(m.out) }
```

### 55) Fan-in with per-source rate limit
Problem: Limit each source to R items/sec.
Solution:
```go
func RateLimit[T any](in <-chan T, d time.Duration) <-chan T {
    out := make(chan T)
    go func(){ defer close(out); t := time.NewTicker(d); defer t.Stop(); for v := range in { <-t.C; out <- v } }()
    return out
}
```

### 56) Duplicate detector over moving window
Problem: Drop items seen in last W.
Solution:
```go
func DedupeWindow[T comparable](in <-chan T, w int) <-chan T {
    out := make(chan T)
    go func(){
        defer close(out)
        q := make([]T, 0, w)
        set := map[T]int{}
        for v := range in {
            if set[v] == 0 { out <- v; q = append(q, v); set[v]++;
                if len(q) > w { old := q[0]; q = q[1:]; set[old]--; if set[old] == 0 { delete(set, old) } }
            }
        }
    }()
    return out
}
```

### 57) Merge with backoff on empty sources
Problem: If all sources empty, sleep briefly to avoid spin.
Solution:
```go
func MergeNonBusy[T any](chs ...<-chan T) <-chan T {
    out := make(chan T)
    go func(){
        defer close(out)
        for {
            progressed := false
            for i, ch := range chs {
                if ch == nil { continue }
                select { case v, ok := <-ch: if !ok { chs[i] = nil } else { out <- v; progressed = true } default: }
            }
            if !progressed { if allNil(chs) { return }; time.Sleep(time.Millisecond) }
        }
    }()
    return out
}
func allNil[T any](xs []<-chan T) bool { for _, x := range xs { if x != nil { return false } } return true }
```

### 58) Retry send with jitter
Problem: On failed try-send, wait jitter then retry.
Solution:
```go
func RetrySend[T any](ch chan<- T, v T, max time.Duration) {
    for {
        select { case ch <- v: return default: }
        time.Sleep(time.Duration(rand.Int63n(int64(max))))
    }
}
```

### 59) Semaphore with context cancel
Problem: Acquire slot or fail on ctx.
Solution:
```go
type CSem struct{ c chan struct{} }
func NewCSem(n int) *CSem { return &CSem{make(chan struct{}, n)} }
func (s *CSem) Acquire(ctx context.Context) error { select { case s.c <- struct{}{}: return nil; case <-ctx.Done(): return ctx.Err() } }
func (s *CSem) Release(){ <-s.c }
```

### 60) Producer-consumer with bounded latency
Problem: Ensure max waiting time for any item.
Solution:
```go
// Use per-item timeout when sending to next stage, drop if expired.
```

### 61) Multi-way select via reflect.Select
Problem: Select over dynamic list of channels.
Solution:
```go
func SelectMany[T any](chs []<-chan T) (int, T, bool) {
    cases := make([]reflect.SelectCase, len(chs))
    for i, ch := range chs { cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)} }
    i, v, ok := reflect.Select(cases)
    var zero T
    if !ok { return i, zero, false }
    return i, v.Interface().(T), true
}
```

### 62) Pipeline stage with error side channel
Problem: Emit errors on separate channel.
Solution:
```go
func StageWithErr[T any](in <-chan T, f func(T) (T, error)) (<-chan T, <-chan error) {
    out := make(chan T)
    errc := make(chan error)
    go func(){ defer close(out); defer close(errc); for v := range in { if x, err := f(v); err != nil { errc <- err } else { out <- x } } }()
    return out, errc
}
```

### 63) Heartbeat channel
Problem: Emit heartbeat ticks while work proceeds.
Solution:
```go
func Heartbeat(interval time.Duration, done <-chan struct{}) <-chan struct{} {
    hb := make(chan struct{})
    go func(){ defer close(hb); t := time.NewTicker(interval); defer t.Stop(); for { select { case <-t.C: select { case hb <- struct{}{}: default: } case <-done: return } } }()
    return hb
}
```

### 64) Timeout if no heartbeat
Problem: Fail if no heartbeat within D.
Solution:
```go
func WaitHeartbeat(hb <-chan struct{}, d time.Duration) bool {
    t := time.NewTimer(d); defer t.Stop()
    for { select { case _, ok := <-hb: if !ok { return false }; if !t.Stop() { <-t.C }; t.Reset(d) case <-t.C: return false } }
}
```

### 65) Batch by size or time whichever first
Problem: Emit batch when size=B or time=T.
Solution:
```go
func Batch[T any](in <-chan T, size int, d time.Duration) <-chan []T {
    out := make(chan []T)
    go func(){
        defer close(out)
        t := time.NewTimer(d)
        defer t.Stop()
        buf := make([]T, 0, size)
        for {
            var expire <-chan time.Time
            if len(buf) > 0 { expire = t.C }
            if len(buf) == 0 { if !t.Stop() { select { case <-t.C: default: } } }
            select {
            case v, ok := <-in:
                if !ok { if len(buf) > 0 { out <- append([]T(nil), buf...) }; return }
                buf = append(buf, v)
                if len(buf) == 1 { t.Reset(d) }
                if len(buf) >= size { out <- append([]T(nil), buf...); buf = buf[:0] }
            case <-expire:
                if len(buf) > 0 { out <- append([]T(nil), buf...); buf = buf[:0] }
            }
        }
    }()
    return out
}
```

### 66) K-way merge with heap + channels
Problem: Merge K sorted channels into one sorted.
Solution:
```go
type item struct{ val int; idx int }
type minheap []item
func (h minheap) Len() int { return len(h) }
func (h minheap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h minheap) Swap(i, j int){ h[i], h[j] = h[j], h[i] }
func (h *minheap) Push(x any){ *h = append(*h, x.(item)) }
func (h *minheap) Pop() any { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }
func MergeK(sorted []<-chan int) <-chan int {
    out := make(chan int)
    go func(){
        defer close(out)
        h := &minheap{}
        heap.Init(h)
        vals := make([]int, len(sorted))
        ok := make([]bool, len(sorted))
        for i := range sorted { if v, o := <-sorted[i]; o { vals[i], ok[i] = v, true; heap.Push(h, item{v, i}) } }
        for h.Len() > 0 {
            it := heap.Pop(h).(item)
            out <- it.val
            if v, o := <-sorted[it.idx]; o { heap.Push(h, item{v, it.idx}) }
        }
    }()
    return out
}
```

### 67) Event bus with topic channels
Problem: Publish by topic; subscribe returns a channel per topic.
Solution:
```go
type Bus struct{ mu sync.RWMutex; subs map[string][]chan any }
func NewBus() *Bus { return &Bus{subs: map[string][]chan any{}} }
func (b *Bus) Subscribe(topic string, buf int) <-chan any { b.mu.Lock(); defer b.mu.Unlock(); ch := make(chan any, buf); b.subs[topic] = append(b.subs[topic], ch); return ch }
func (b *Bus) Publish(topic string, v any) { b.mu.RLock(); defer b.mu.RUnlock(); for _, ch := range b.subs[topic] { select { case ch <- v: default: } } }
```

### 68) Duplicate tee with context cancel
Problem: Tee until ctx done.
Solution:
```go
func TeeCtx[T any](ctx context.Context, in <-chan T) (<-chan T, <-chan T) {
    a, b := make(chan T), make(chan T)
    go func(){ defer close(a); defer close(b); for { select { case v, ok := <-in: if !ok { return }; select { case a <- v: case <-ctx.Done(): return }; select { case b <- v: case <-ctx.Done(): return } case <-ctx.Done(): return } } }()
    return a, b
}
```

### 69) Fan-in preserving bursts (group by source)
Problem: Forward burst from one source atomically.
Solution:
```go
// Use a token per source to drain it fully before switching. Advanced scheduling omitted for brevity.
```

### 70) Channel barrier flush
Problem: Wait until channel queue becomes empty.
Solution:
```go
func WaitEmpty[T any](ch <-chan T) {
    for {
        select { case <-ch: default: return }
    }
}
```

### 71) Non-blocking drain up to N
Problem: Drain up to N items without blocking.
Solution:
```go
func DrainN[T any](ch <-chan T, n int) []T { out := make([]T, 0, n); for i := 0; i < n; i++ { select { case v := <-ch: out = append(out, v) default: return out } }; return out }
```

### 72) Interrupter: preempt a long send
Problem: Abort a blocking send when interrupted.
Solution:
```go
func SendInterruptible[T any](ch chan<- T, v T, stop <-chan struct{}) bool { select { case ch <- v: return true; case <-stop: return false } }
```

### 73) Signal edge detector
Problem: Turn a boolean stream into rising/falling edge events.
Solution:
```go
func Edges(in <-chan bool) <-chan string { out := make(chan string); go func(){ defer close(out); last := false; init := false; for v := range in { if !init { last, init = v, true; continue }; if v && !last { out <- "rise" } else if !v && last { out <- "fall" }; last = v }; }(); return out }
```

### 74) Coalesce identical adjacent bursts
Problem: Replace runs of same value by single value and count.
Solution:
```go
func Runs[T comparable](in <-chan T) <-chan struct{V T; N int} { out := make(chan struct{V T; N int}); go func(){ defer close(out); var cur T; n := 0; init := false; for v := range in { if !init { cur, n, init = v, 1, true; continue }; if v == cur { n++ } else { out <- struct{V T; N int}{cur, n}; cur, n = v, 1 } }; if init { out <- struct{V T; N int}{cur, n} } }(); return out }
```

### 75) Channel-based pool of reusable resources
Problem: Borrow/return resources via channel.
Solution:
```go
type Pool[T any] struct{ c chan T }
func NewPool[T any](makeFn func() T, n int) *Pool[T] { p := &Pool[T]{c: make(chan T, n)}; for i := 0; i < n; i++ { p.c <- makeFn() }; return p }
func (p *Pool[T]) Get() T { return <-p.c }
func (p *Pool[T]) Put(v T){ p.c <- v }
```

### 76) Mailbox with overflow counter
Problem: When full, count dropped messages instead of blocking.
Solution:
```go
type Mailbox[T any] struct{ c chan T; dropped atomic.Int64 }
func NewMailbox[T any](n int) *Mailbox[T] { return &Mailbox[T]{c: make(chan T, n)} }
func (m *Mailbox[T]) Send(v T){ select { case m.c <- v: default: m.dropped.Add(1) } }
func (m *Mailbox[T]) Recv() <-chan T { return m.c }
```

### 77) Backoff reader on empty channel
Problem: If no data, sleep with increasing backoff.
Solution:
```go
func BackoffRead[T any](ch <-chan T, max time.Duration) <-chan T { out := make(chan T); go func(){ defer close(out); d := time.Millisecond; for { select { case v, ok := <-ch: if !ok { return }; out <- v; d = time.Millisecond case <-time.After(d): if d < max { d *= 2 } } } }(); return out }
```

### 78) Ensure only latest request is processed
Problem: Cancel previous work when a new request arrives.
Solution:
```go
func LatestOnly[T any](in <-chan T, work func(context.Context, T)) {
    go func(){
        var cancel context.CancelFunc
        for v := range in {
            if cancel != nil { cancel() }
            var ctx context.Context
            ctx, cancel = context.WithCancel(context.Background())
            v := v
            go work(ctx, v)
        }
        if cancel != nil { cancel() }
    }()
}
```

### 79) Rendezvous of N goroutines
Problem: All N must arrive before any proceeds.
Solution:
```go
func Rendezvous(n int) (arrive chan<- struct{}, proceed <-chan struct{}) {
    a := make(chan struct{})
    p := make(chan struct{})
    go func(){ count := 0; for range a { count++; if count == n { close(p); return } } }()
    return a, p
}
```

### 80) Asynchronous logger with lossless shutdown
Problem: Buffer logs; on shutdown flush and close.
Solution:
```go
type Logger struct{ c chan string; done chan struct{} }
func NewLogger(n int) *Logger { l := &Logger{c: make(chan string, n), done: make(chan struct{})}; go l.loop(); return l }
func (l *Logger) loop(){ for msg := range l.c { _ = msg /* write */ }; close(l.done) }
func (l *Logger) Log(s string){ l.c <- s }
func (l *Logger) Close(){ close(l.c); <-l.done }
```

### 81) Pipeline fan-in with bounded memory
Problem: Prevent unbounded buffering between stages.
Solution:
```go
// Use small fixed-size buffers between stages and block sends when full.
```

### 82) Switchable consumer (hot-swap output)
Problem: Switch output channel at runtime without losing messages.
Solution:
```go
type Switch[T any] struct{ in chan T; mu sync.Mutex; out chan T }
func NewSwitch[T any]() *Switch[T]{ s := &Switch[T]{in: make(chan T)}; go s.loop(); return s }
func (s *Switch[T]) loop(){ var out chan T; var buf []T; for v := range s.in { if out == nil { buf = append(buf, v) } else { out <- v } } }
func (s *Switch[T]) SetOut(o chan T){ s.mu.Lock(); s.out = o; s.mu.Unlock() }
```

### 83) Duplicate suppression across time window
Problem: Drop duplicates seen within last D time.
Solution:
```go
func DedupeTime[T comparable](in <-chan T, d time.Duration) <-chan T {
    out := make(chan T)
    go func(){ defer close(out); seen := map[T]time.Time{}; t := time.NewTicker(d); defer t.Stop(); for { select {
        case v, ok := <-in: if !ok { return }; if time.Since(seen[v]) >= d { out <- v; seen[v] = time.Now() }
        case <-t.C: for k, ts := range seen { if time.Since(ts) >= d { delete(seen, k) } }
    } } }()
    return out
}
```

### 84) Sample every Nth item
Problem: Keep every Nth, drop others.
Solution:
```go
func SampleN[T any](in <-chan T, n int) <-chan T { out := make(chan T); go func(){ defer close(out); i := 0; for v := range in { if i%n == 0 { out <- v }; i++ } }(); return out }
```

### 85) Shuffle using buffered channels
Problem: Randomize item order.
Solution:
```go
func Shuffle[T any](in <-chan T, buf int) <-chan T { out := make(chan T); go func(){ defer close(out); arr := make([]T, 0, buf); for v := range in { arr = append(arr, v); if len(arr) == buf { rand.Shuffle(len(arr), func(i,j int){ arr[i], arr[j] = arr[j], arr[i] }); for _, x := range arr { out <- x }; arr = arr[:0] } }; for _, x := range arr { out <- x } }(); return out }
```

### 86) Gate by time (allow bursts then close gate)
Problem: Allow sends only during window D after open.
Solution:
```go
func TimeGate[T any](in <-chan T, d time.Duration) <-chan T { out := make(chan T); go func(){ defer close(out); deadline := time.After(d); for { select { case v, ok := <-in: if !ok { return }; select { case out <- v: case <-deadline: return } case <-deadline: return } } }(); return out }
```

### 87) Parallel map with max in-flight per key
Problem: Limit in-flight work per key.
Solution:
```go
// Use a map[K]*Gate and acquire before sending to worker.
```

### 88) Split into K shards
Problem: Deterministically route to K shard channels.
Solution:
```go
func Shard[T any](in <-chan T, k int, hash func(T) uint64) []<-chan T { outs := make([]chan T, k); for i := range outs { outs[i] = make(chan T) }; go func(){ defer func(){ for _, o := range outs { close(o) } }(); for v := range in { idx := int(hash(v) % uint64(k)); outs[idx] <- v } }(); res := make([]<-chan T, k); for i := range outs { res[i] = outs[i] }; return res }
```

### 89) Merge with per-source fairness
Problem: Avoid starving slow producers.
Solution:
```go
// Round-robin select over sources using indices and try-recv; block if none ready.
```

### 90) Circuit breaker using channels
Problem: Trip open after E errors, half-open after cool-down.
Solution:
```go
type Breaker struct{ trips int; max int; cool time.Duration; ch chan struct{} }
func NewBreaker(max int, cool time.Duration) *Breaker { return &Breaker{max: max, cool: cool, ch: make(chan struct{}, 1)} }
func (b *Breaker) Allow() bool { if b.trips >= b.max { return false }; b.ch <- struct{}{}; return true }
func (b *Breaker) Fail(){ b.trips++; if b.trips >= b.max { go func(){ time.Sleep(b.cool); b.trips = 0 }() } }
```

### 91) Async request/response with correlation id
Problem: Match responses to requests using channels.
Solution:
```go
type Req struct{ ID int; Reply chan string }
func Client(out chan<- Req, id int) string { r := Req{ID: id, Reply: make(chan string, 1)}; out <- r; return <-r.Reply }
func Server(in <-chan Req){ for r := range in { r.Reply <- fmt.Sprintf("ok:%d", r.ID) } }
```

### 92) Pipeline that never deadlocks on close
Problem: Ensure every stage closes its output when input is closed.
Solution:
```go
// Pattern: range over input, send to output, and defer close(output).
```

### 93) Safe broadcast close with late subscribers
Problem: New listeners get immediate closed signal.
Solution:
```go
type OnceClosed struct{ once sync.Once; c chan struct{} }
func NewOnceClosed() *OnceClosed { return &OnceClosed{c: make(chan struct{})} }
func (o *OnceClosed) Close(){ o.once.Do(func(){ close(o.c) }) }
func (o *OnceClosed) Done() <-chan struct{} { return o.c }
```

### 94) Backpressure-aware HTTP worker skeleton
Problem: Queue requests with bounded channel; return 503 if full.
Solution:
```go
// if !TrySend(jobs, req) { w.WriteHeader(503); return }
```

### 95) Turn callback API into channel
Problem: Wrap event callback into receive-only channel.
Solution:
```go
func FromCallback[T any](register func(func(T))) <-chan T { ch := make(chan T, 1); register(func(v T){ select { case ch <- v: default: } }); return ch }
```

### 96) Turn channel into iterator function
Problem: Return `next()` that yields values until closed.
Solution:
```go
func Iterator[T any](ch <-chan T) func() (T, bool) { return func() (T, bool) { v, ok := <-ch; return v, ok } }
```

### 97) Backpressure signal to producer
Problem: Producer pauses when consumer indicates slow.
Solution:
```go
// Use an additional feedback channel consumer->producer to signal pause/resume.
```

### 98) Funnel many cancel signals to one worker
Problem: Worker listens to many stop signals efficiently.
Solution:
```go
func FunnelStops(stops ...<-chan struct{}) <-chan struct{} { return Or(stops...) }
```

### 99) Ensure graceful goroutine leak avoidance
Problem: Prevent goroutine leaks on blocked sends.
Solution:
```go
// Always include a done/cancel/select case to unblock or drop.
```

### 100) Deterministic tests for channel code
Problem: Test pipelines without flaky sleeps.
Solution:
```go
// Use small, buffered channels; inject fake timers/tickers; use select with timeouts in tests.
```


