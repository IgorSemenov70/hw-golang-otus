package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func pipelineWorker(in In, done In, bi chan interface{}) {
	// Читает данные из входного канала, и кладёт в промежуточный
	defer close(bi)
	for {
		select {
		case <-done:
			return
		case v, ok := <-in:
			if !ok {
				return
			}
			bi <- v
		}
	}
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if len(stages) == 0 {
		return nil
	}
	out := in
	for _, stage := range stages {
		bi := make(Bi)

		go pipelineWorker(out, done, bi)

		out = stage(bi)
	}
	return out
}
