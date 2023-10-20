package main

import (
	"fmt"
	"github.com/kongggg/go-aravis"
	"image"
	"image/jpeg"
	"os"
	"time"
)

var exposureTime float64 = 1000
var gain float64 = 4

func main() {
	var (
		n   uint = 0
		err error
	)

	for {
		fmt.Println("Refreshing")
		// 1. inet update device list
		aravis.UpdateDeviceList()

		// 2. get devices list
		n, err = aravis.GetNumDevices()
		if err != nil {
			fmt.Println("aravis.GetNumDevices()", err)
			return
		}

		if n < 1 {
			continue
		}
		fmt.Println("Devices:", n)
		break
	}

	/*	// 1. inet update device list
		aravis.UpdateDeviceList()

		// 2. get devices list
		n, err := aravis.GetNumDevices()
		if err != nil {
			fmt.Println("aravis.GetNumDevices()", err)
			return
		}
		fmt.Println("Devices:", n)*/

	// 3. get device
	name, err := aravis.GetDeviceId(n - 1)
	if err != nil {
		fmt.Println("aravis.GetDeviceId(n)", err)
		return
	}
	fmt.Println("Device: ", name)

	// 3.1. get device ip
	address, err := aravis.GetDeviceAddress(n - 1)
	if err != nil {
		fmt.Println("aravis.GetDeviceId(n)", err)
		return
	}
	fmt.Println("Device Address: ", address)

	// 4. get camera
	camera, err := aravis.NewCamera(name)
	if err != nil {
		fmt.Println("aravis.NewCamera(name)", err)
		return
	}
	defer camera.Close()

	// 5. get camera sensor size
	maxWidth, maxHeight, err := camera.GetSensorSize()
	if err != nil {
		fmt.Println("camera.GetSensorSize()", err)
		return
	}
	fmt.Println("Sensor:", maxWidth, maxHeight)

	// 6. set camera param
	camera.SetRegion(0, 0, maxWidth, maxHeight)
	camera.SetExposureTimeAuto(aravis.AUTO_OFF)
	camera.SetExposureTime(exposureTime)
	camera.SetGain(gain)
	camera.SetFrameRate(3.75)
	camera.SetAcquisitionMode(aravis.ACQUISITION_MODE_SINGLE_FRAME)

	// 7. get camera payload size
	size, err := camera.GetPayloadSize()
	if err != nil {
		fmt.Println("camera.GetPayloadSize()", err)
		return
	}
	fmt.Println("Payload:", size)

	// 8. get camera region
	_, _, width, height, err := camera.GetRegion()
	if err != nil {
		fmt.Println("camera.GetRegion()", err)
		return
	}
	fmt.Println("Region:", width, height)

	// 9. create stream
	stream, err := camera.CreateStream()
	if err != nil {
		fmt.Println("camera.CreateStream()", err)
		return
	}
	defer stream.Close()

	// 10. create buffer to hold stream
	buffer, err := aravis.NewBuffer(size)
	if err != nil {
		fmt.Println("aravis.NewBuffer(size)", err)
		return
	}

	// 11. push stream 2 buffer
	stream.PushBuffer(buffer)

	// 12. start acquisition
	camera.StartAcquisition()
	defer camera.StopAcquisition()

	buffer, err = stream.TimeoutPopBuffer(time.Second)
	if err != nil {
		fmt.Println("stream.TimeoutPopBuffer(time.Second)", err)
		return
	}

	data, err := buffer.GetData()
	if err != nil {
		fmt.Println("buffer.GetData()", err)
		return
	}
	fmt.Println("Data:", len(data))

	// Image is in red-green bayer format
	img := aravis.NewBayerRG(
		image.Rectangle{Max: image.Point{X: width, Y: height}},
	)
	img.Pix = data

	f, err := os.Create("xxxx.jpg")
	if err != nil {
		fmt.Println("os.Create(\"xxxx.jpg\")", err)
		return
	}
	defer f.Close()

	err = jpeg.Encode(f, img, nil)
	if err != nil {
		fmt.Println("jpeg.Encode(f, img, nil)", err)
		return
	}

	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println("f.Stat()", err)
		return
	}
	fmt.Println("Img:", fileInfo)
}
