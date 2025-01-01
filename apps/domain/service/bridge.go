package service

import "sync"

func OrDone(done <-chan interface{}, c <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case val, ok := <-c:
				if !ok {
					return
				}
				select {
				case valStream <- val:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

func Bridge(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
	valStream := make(chan interface{})
	var wg sync.WaitGroup
	go func() {
		defer close(valStream)
		for {
			var stream <-chan interface{}
			select {
			case maybeStream, ok := <-chanStream:
				if !ok {
					return
				}
				stream = maybeStream
			case <-done:
				return
			}

			wg.Add(1)
			go func(s <-chan interface{}) {
				defer wg.Done()
				for val := range OrDone(done, s) {
					select {
					case valStream <- val:
					case <-done:
						return
					}
				}
			}(stream)
		}
	}()

	go func() {
		wg.Wait()
		close(valStream)
	}()

	return valStream
}

func GenVals() <-chan <-chan interface{} {
	chanStream := make(chan (<-chan interface{}))
  go func() {
    defer close(chanStream)
  }
}
