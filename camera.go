package aravis

// #cgo pkg-config: aravis-0.8
// #include <arv.h>
// #include <stdlib.h>
import "C"
import (
	"errors"
	"unsafe"
)

type Camera struct {
	camera *C.struct__ArvCamera
}

const (
	ACQUISITION_MODE_CONTINUOUS   = C.ARV_ACQUISITION_MODE_CONTINUOUS
	ACQUISITION_MODE_SINGLE_FRAME = C.ARV_ACQUISITION_MODE_SINGLE_FRAME
)

const (
	AUTO_OFF        = C.ARV_AUTO_OFF
	AUTO_ONCE       = C.ARV_AUTO_ONCE
	AUTO_CONTINUOUS = C.ARV_AUTO_CONTINUOUS
)

func NewCamera(name string) (Camera, error) {
	var c Camera
	var err error
	var gErr *C.GError

	cs := C.CString(name)
	c.camera, err = C.arv_camera_new(cs, &gErr)
	C.free(unsafe.Pointer(cs))

	handleGError(gErr)
	return c, err
}

func (c *Camera) CreateStream() (Stream, error) {
	var s Stream
	var err error
	var gErr *C.GError

	s.stream, err = C.arv_camera_create_stream(
		c.camera,
		nil,
		nil,
		&gErr,
	)

	handleGError(gErr)

	if s.stream == nil {
		return Stream{}, errors.New("Failed to create stream")
	}

	return s, err
}

func (c *Camera) GetDevice() (Device, error) {
	var d Device
	var err error

	d.device, err = C.arv_camera_get_device(c.camera)

	return d, err
}

func (c *Camera) GetVendorName() (string, error) {
	var gErr *C.GError
	name, err := C.arv_camera_get_vendor_name(c.camera, &gErr)

	handleGError(gErr)

	return C.GoString(name), err
}

func (c *Camera) GetModelName() (string, error) {
	var gErr *C.GError

	name, err := C.arv_camera_get_model_name(c.camera, &gErr)

	handleGError(gErr)

	return C.GoString(name), err
}

func (c *Camera) GetDeviceId() (string, error) {
	var gErr *C.GError
	id, err := C.arv_camera_get_device_id(c.camera, &gErr)
	handleGError(gErr)
	return C.GoString(id), err
}

func (c *Camera) GetSensorSize() (int, int, error) {
	var width, height int
	var gErr *C.GError
	_, err := C.arv_camera_get_sensor_size(
		c.camera,
		(*C.gint)(unsafe.Pointer(&width)),
		(*C.gint)(unsafe.Pointer(&height)),
		&gErr,
	)

	handleGError(gErr)

	return int(width), int(height), err
}

func (c *Camera) SetRegion(x, y, width, height int) {
	var gErr *C.GError
	C.arv_camera_set_region(c.camera,
		C.gint(x),
		C.gint(y),
		C.gint(width),
		C.gint(height),
		&gErr,
	)
	handleGError(gErr)
}

func (c *Camera) GetRegion() (int, int, int, int, error) {
	var x, y, width, height int
	var gErr *C.GError
	_, err := C.arv_camera_get_region(
		c.camera,
		(*C.gint)(unsafe.Pointer(&x)),
		(*C.gint)(unsafe.Pointer(&y)),
		(*C.gint)(unsafe.Pointer(&width)),
		(*C.gint)(unsafe.Pointer(&height)),
		&gErr,
	)
	handleGError(gErr)
	return int(x), int(y), int(width), int(height), err
}

func (c *Camera) GetHeightBounds() (int, int, error) {
	var min, max int
	var gErr *C.GError
	_, err := C.arv_camera_get_height_bounds(
		c.camera,
		(*C.gint)(unsafe.Pointer(&min)),
		(*C.gint)(unsafe.Pointer(&max)),
		&gErr,
	)
	handleGError(gErr)
	return int(min), int(max), err
}

func (c *Camera) GetWidthBounds() (int, int, error) {
	var min, max int
	var gErr *C.GError
	_, err := C.arv_camera_get_width_bounds(
		c.camera,
		(*C.gint)(unsafe.Pointer(&min)),
		(*C.gint)(unsafe.Pointer(&max)),
		&gErr,
	)
	handleGError(gErr)
	return int(min), int(max), err
}

func (c *Camera) SetBinning() {
	// TODO
}

func (c *Camera) GetBinning() (int, int, error) {
	var min, max int
	var gErr *C.GError
	_, err := C.arv_camera_get_binning(
		c.camera,
		(*C.gint)(unsafe.Pointer(&min)),
		(*C.gint)(unsafe.Pointer(&max)),
		&gErr,
	)
	handleGError(gErr)
	return int(min), int(max), err
}

func (c *Camera) SetPixelFormat() {
	// TODO
}

func (c *Camera) GetPixelFormat() {
	// TODO
}

func (c *Camera) GetPixelFormatAsString() {
	// TODO
}

func (c *Camera) SetPixelFormatFromString() {
	// TODO
}

func (c *Camera) GetAvailablePixelFormats() {
	// TODO
}

func (c *Camera) GetAvailablePixelFormatsAsDisplayNames() {
	// TODO
}

func (c *Camera) GetAvailablePixelFormatsAsStrings() {
	// TODO
}

func (c *Camera) StartAcquisition() {
	var gErr *C.GError
	C.arv_camera_start_acquisition(c.camera, &gErr)
	handleGError(gErr)
}

func (c *Camera) StopAcquisition() {
	var gErr *C.GError
	C.arv_camera_stop_acquisition(c.camera, &gErr)
	handleGError(gErr)
}

func (c *Camera) AbortAcquisition() {
	var gErr *C.GError
	C.arv_camera_abort_acquisition(c.camera, &gErr)
	handleGError(gErr)
}

func (c *Camera) SetAcquisitionMode(mode int) {
	var gErr *C.GError
	C.arv_camera_set_acquisition_mode(c.camera, C.ArvAcquisitionMode(mode), &gErr)
	handleGError(gErr)
}

func (c *Camera) SetFrameRate(frameRate float64) {
	var gErr *C.GError
	C.arv_camera_set_frame_rate(c.camera, C.double(frameRate), &gErr)
	handleGError(gErr)
}

func (c *Camera) GetFrameRate() (float64, error) {
	var gErr *C.GError
	fr, err := C.arv_camera_get_frame_rate(c.camera, &gErr)
	handleGError(gErr)
	return float64(fr), err
}

func (c *Camera) GetFrameRateBounds() (float64, float64, error) {
	var min, max float64
	var gErr *C.GError
	_, err := C.arv_camera_get_frame_rate_bounds(
		c.camera,
		(*C.double)(unsafe.Pointer(&min)),
		(*C.double)(unsafe.Pointer(&max)),
		&gErr,
	)
	handleGError(gErr)
	return float64(min), float64(max), err
}

func (c *Camera) SetTrigger(source string) {
	var gErr *C.GError
	csource := C.CString(source)
	C.arv_camera_set_trigger(c.camera, csource, &gErr)
	C.free(unsafe.Pointer(csource))
	handleGError(gErr)
}

func (c *Camera) SetTriggerSource(source string) {
	var gErr *C.GError
	csource := C.CString(source)
	C.arv_camera_set_trigger_source(c.camera, csource, &gErr)
	C.free(unsafe.Pointer(csource))
	handleGError(gErr)
}

func (c *Camera) GetTriggerSource() (string, error) {
	var gErr *C.GError
	csource, err := C.arv_camera_get_trigger_source(c.camera, &gErr)
	handleGError(gErr)
	return C.GoString(csource), err
}

func (c *Camera) SoftwareTrigger() {
	var gErr *C.GError
	C.arv_camera_software_trigger(c.camera, &gErr)
	handleGError(gErr)
}

func (c *Camera) IsExposureTimeAvailable() (bool, error) {
	var gErr *C.GError
	gboolean, err := C.arv_camera_is_exposure_time_available(c.camera, &gErr)
	handleGError(gErr)
	return toBool(gboolean), err
}

func (c *Camera) IsExposureAutoAvailable() (bool, error) {
	var gErr *C.GError
	gboolean, err := C.arv_camera_is_exposure_auto_available(c.camera, &gErr)
	handleGError(gErr)
	return toBool(gboolean), err
}

func (c *Camera) SetExposureTime(time float64) {
	var gErr *C.GError
	C.arv_camera_set_exposure_time(c.camera, C.double(time), &gErr)
	handleGError(gErr)
}

func (c *Camera) GetExposureTime() (float64, error) {
	var gErr *C.GError
	cdouble, err := C.arv_camera_get_exposure_time(c.camera, &gErr)
	handleGError(gErr)
	return float64(cdouble), err
}

func (c *Camera) GetExposureTimeBounds() {
	// TODO
}

func (c *Camera) SetExposureTimeAuto(mode int) {
	var gErr *C.GError
	C.arv_camera_set_exposure_time_auto(c.camera, C.ArvAuto(mode), &gErr)
	handleGError(gErr)
}

func (c *Camera) GetExposureTimeAuto() {
	// TODO
}

func (c *Camera) SetGain(gain float64) {
	var gErr *C.GError
	C.arv_camera_set_gain(c.camera, C.double(gain), &gErr)
	handleGError(gErr)
}

func (c *Camera) GetGain() (float64, error) {
	var gErr *C.GError
	cgain, err := C.arv_camera_get_gain(c.camera, &gErr)
	handleGError(gErr)
	return float64(cgain), err
}

func (c *Camera) GetGainBounds() (float64, float64, error) {
	var min, max float64
	var gErr *C.GError
	_, err := C.arv_camera_get_gain_bounds(
		c.camera,
		(*C.double)(unsafe.Pointer(&min)),
		(*C.double)(unsafe.Pointer(&max)),
		&gErr,
	)
	handleGError(gErr)
	return float64(min), float64(max), err
}

func (c *Camera) SetGainAuto() {
	// TODO
}

func (c *Camera) GetPayloadSize() (uint, error) {
	var gErr *C.GError
	csize, err := C.arv_camera_get_payload(c.camera, &gErr)
	handleGError(gErr)
	return uint(csize), err
}

func (c *Camera) IsGVDevice() (bool, error) {
	cbool, err := C.arv_camera_is_gv_device(c.camera)
	return toBool(cbool), err
}

func (c *Camera) GVGetNumStreamChannels() (int, error) {
	var gErr *C.GError
	cint, err := C.arv_camera_gv_get_n_stream_channels(c.camera, &gErr)
	handleGError(gErr)
	return int(cint), err
}

func (c *Camera) GVSelectStreamChannels(id int) {
	var gErr *C.GError
	C.arv_camera_gv_select_stream_channel(c.camera, C.gint(id), &gErr)
	handleGError(gErr)
}

func (c *Camera) GVGetCurrentStreamChannel() (int, error) {
	var gErr *C.GError
	cint, err := C.arv_camera_gv_get_current_stream_channel(c.camera, &gErr)
	handleGError(gErr)
	return int(cint), err
}

func (c *Camera) GVGetPacketDelay() (int64, error) {
	var gErr *C.GError
	cint64, err := C.arv_camera_gv_get_packet_delay(c.camera, &gErr)
	handleGError(gErr)
	return int64(cint64), err
}

func (c *Camera) GVSetPacketDelay(delay int64) {
	var gErr *C.GError
	C.arv_camera_gv_set_packet_delay(c.camera, C.gint64(delay), &gErr)
	handleGError(gErr)
}

func (c *Camera) GVGetPacketSize() (int, error) {
	var gErr *C.GError
	csize, err := C.arv_camera_gv_get_packet_size(c.camera, &gErr)
	handleGError(gErr)
	return int(csize), err
}

func (c *Camera) GVSetPacketSize(size int) {
	var gErr *C.GError
	C.arv_camera_gv_set_packet_size(c.camera, C.gint(size), &gErr)
	handleGError(gErr)
}

func (c *Camera) GetChunkMode() (bool, error) {
	var gErr *C.GError
	mode, err := C.arv_camera_get_chunk_mode(c.camera, &gErr)
	handleGError(gErr)
	return toBool(mode), err
}

func (c *Camera) Close() {
	C.g_object_unref(C.gpointer(c.camera))
}
