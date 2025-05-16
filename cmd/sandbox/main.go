package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/fengdotdev/golibs-nativedrive/interfaces"
	"github.com/fengdotdev/golibs-nativedrive/v1/ndrive"
	"github.com/fengdotdev/golibs-traits/trait"
)

func SomeStream(ctx context.Context) chan []byte {

	// Create a channel to send data
	dataChan := make(chan []byte)

	// Start a goroutine to send data
	go func() {
		for i := 0; i < 10; i++ {
			dataChan <- []byte(fmt.Sprintf("Data chunk %d", i))
		}
		close(dataChan)
	}()

	return dataChan
}

type FutureStream struct {
	Channel   chan []byte
	ErrorChan chan error
}

func NewFutureStream() FutureStream {
	return FutureStream{
		Channel:   make(chan []byte),
		ErrorChan: make(chan error),
	}
}
func FIleStream2(ctx context.Context, filepath string) FutureStream {
	future := FutureStream{
		Channel:   make(chan []byte, 10),
		ErrorChan: make(chan error, 1),
	}

	go func() {
		defer close(future.Channel)
		defer close(future.ErrorChan)

		file, err := os.Open(filepath)
		if err != nil {
			future.ErrorChan <- err
			return
		}
		defer file.Close()

		buf := make([]byte, 1024)
		for {
			n, err := file.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				future.ErrorChan <- err
				return
			}
			if n == 0 {
				break
			}

			select {
			case future.Channel <- buf[:n]:
			case <-ctx.Done():
				future.ErrorChan <- ctx.Err()
				return
			}
		}

		future.ErrorChan <- nil
	}()

	return future
}

func FIleStream(ctx context.Context, filepath string) FutureStream {

	future := NewFutureStream()

	// Start a goroutine to send data
	go func() {

		file, err := os.Open(filepath)
		if err != nil {
			future.ErrorChan <- err
			return
		}
		defer file.Close()
		// Read the file in chunks
		buf := make([]byte, 1024)

		for {
			n, err := file.Read(buf)
			if err != nil {
				if err.Error() == "EOF" {
					break
				}
				future.ErrorChan <- err
				return
			}
			if n == 0 {
				break
			}
			// Send the data to the channel
			future.Channel <- buf[:n]
		}

		close(future.Channel)
		close(future.ErrorChan)
	}()

	return future
}

func StreamReader(ctx context.Context, dataChan chan []byte) {
	for data := range dataChan {
		fmt.Println("Received:", string(data))
	}
}

func StreamFileReader(ctx context.Context, future FutureStream) {

	for data := range future.Channel {
		fmt.Println("Received:", string(data))
	}

	if err := <-future.ErrorChan; err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Stream completed successfully")
	}
}

func main() {

	ctx := context.Background()
	pathfile := "/home/feng/Descargas/Presentación libro Nervios Craneales editorial -002.mp4"

	future := FIleStream(ctx, pathfile)

	StreamFileReader(ctx, future)

}

// / place holder
func Foo() {
	workingDirRaw, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	workingDir := workingDirRaw + "/testfolder"

	var drive trait.CRUDWithCTX[string, interfaces.Element] = ndrive.NewNDrive(workingDir)

	fmt.Println(drive)

	fmt.Println("Drive created successfully")
	ctx := context.Background()

	fmt.Println(ctx)

}
