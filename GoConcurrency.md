## What is concurrent programming?
- Concurrent programming is a way of programming in which a program can be made to utilize multiple cores and processors so that parts of the program can run in parallel.
- ***But it is'nt always effective due to the overhead of threading itself and because it's easier to make mistakes in threaded programs.***

## Benefits of concurrent programming in Go
- Go provides a high-level support for concurrent programming that makes it easier to do it correctly.
- ***Concurrent processing is done in goroutines that are much more lightweight than threads***.
- Automatic gerbage collector releives the burden of memory management which can be complicated in concurrent programs.

## CSP
- Go's high-level API for writing concurrent programs is based on CSP -***Communicating Sequential Processes***.
- *This means that synchronization is acheived by sending and receiving data via thread-safe channels.*
- This removes the need of explicit locking and all the care needed to lock and unlock at the right times.
- The CSP approach greatly simplifies writing concurrent programs.
- **Several thousands or even tens of thousands of goroutines can be created without any problem.**

## Approach to writing concurrent programs
- You split up the processing over multiple goroutines (in addition to the main orchestrating goroutine) and either output results as soon as they are computed or gather the results for outputting at the end.
- **PITFALLS**:
  - ***Premature termination***
    - The main goroutine terminates without waiting for the other goroutines to finish.
    - Solved by keeping the main goroutine alive untill all the other goroutines have finished working.
  - ***Non-termination***
    - The main goroutine and all the processing goroutines are alive though all the work has been done.
    - Happens because of a failure to report the completion.
- **Solution** - A simple solution to the above problems is to make the main goroutine wait for a "done" channel to report that the work is finished.

## Channels
- ***Channels provide a lock-free means of communicating between concurrently running goroutines***.
- By default, channels are bi-directional. You can send values into them and recieve values from them.
- **Unbuffered channels** - A goroutine is will be blocked on a send operation to a unbuffered channel, and will remain blocked untill another goroutine tries to receive from that channel.
- **Closing a channel** 
  - A channel can be closed using **`close()`** function by a sender goroutine, so that the receiver goroutine will know that sender has stopped sending.
  - Using a **`for .. range, select or checked reveive using <- operator`** a gorutine can know if the channel is closed.
  - It is perfectly sensible not to close channels that are never checked for being closed. Channels are lightweight and do not hold up resources like say an open file.
  
## Sample program
- ```go
    func main() {
      jobs := make(chan Job)
      done := make(chan bool, len(jobList))
      go func() {
        for _, job := range jobList {
          jobs <- job // Blocks waiting for a receive
          close(jobs) 
        }
      }()

      go func() {
        for job := range jobs { // Blocks waiting for a send
          fmt.Println(job) // Do one job
          done <- true 
        }
      }()

      for i := 0; i < len(jobList); i++ {
        <-done // Blocks waiting for a receive
      }
    }
    ```