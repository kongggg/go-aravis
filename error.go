package aravis

// #cgo pkg-config: aravis-0.8
// #include <arv.h>
// #include <stdlib.h>
import "C"
import "fmt"

func handleGError(gErr *C.GError) error {
	if gErr == nil {
		return nil
	}
	defer C.g_error_free(gErr)
	err := fmt.Errorf("%s", C.GoString(gErr.message))
	fmt.Println("GError:", err)
	return err
}
