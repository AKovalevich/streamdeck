package StreamDeck

import (
	"errors"
	"sync"

	"github.com/google/gousb"
)

type USBDevice struct {
	sync.Mutex
	context     *gousb.Context
	device      *gousb.Device
	intf        *gousb.Interface
	inEndpoint  *gousb.InEndpoint
	outEndpoint *gousb.OutEndpoint
	connected   bool
	log         Logger
	productID   uint16
	vendorID    uint16
}

func (usbDevice *USBDevice) IsConnected() bool {
	usbDevice.Lock()
	defer usbDevice.Unlock()
	return usbDevice.connected
}

func (usbDevice *USBDevice) GetVendorID() uint16 {
	usbDevice.Lock()
	defer usbDevice.Unlock()
	return usbDevice.vendorID
}

func (usbDevice *USBDevice) GetProductID() uint16 {
	usbDevice.Lock()
	defer usbDevice.Unlock()
	return usbDevice.productID
}

func (usbDevice *USBDevice) GetSerialNumber() (string, error) {
	usbDevice.Lock()
	defer usbDevice.Unlock()
	return usbDevice.device.SerialNumber()
}

func (usbDevice *USBDevice) SetConnected(connected bool) {
	usbDevice.Lock()
	usbDevice.connected = connected
	usbDevice.Unlock()
}

func (usbDevice *USBDevice) Close() error {
	var err error

	if usbDevice.IsConnected() {
		usbDevice.intf.Close()
		err = usbDevice.context.Close()
		if err == nil {
			err = usbDevice.device.Close()
		}
	}

	return err
}

func (usbDevice *USBDevice) write(data []byte) (int, error) {
	return usbDevice.outEndpoint.Write(data)
}

func (usbDevice *USBDevice) read(data []byte) (int, error) {
	count, err := usbDevice.inEndpoint.Read(data)
	if err != nil {
		usbDevice.SetConnected(false)
	}

	return count, err
}

func NewUSBDevice(productID, vendorID uint16) *USBDevice {
	return &USBDevice{
		productID: productID,
		vendorID:  vendorID,
		connected: false,
	}
}

func (usbDevice *USBDevice) Connect() error {
	ctx := gousb.NewContext()
	devices, err := ctx.OpenDevices(findUSBDevice(usbDevice.productID, usbDevice.vendorID))
	if err != nil {
		return err
	}

	if len(devices) <= 0 {
		return errors.New("no one devices")
	}

	usbDevice.device = devices[0]
	usbDevice.context = ctx

	// Detach the device from whichever process already
	// has it.
	err = usbDevice.device.SetAutoDetach(true)
	if err != nil {
		return err
	}

	for num := range usbDevice.device.Desc.Configs {
		config, _ := usbDevice.device.Config(num)
		defer config.Close()

		// Iterate through available interfaces for this configuration
		for _, desc := range config.Desc.Interfaces {
			intf, _ := config.Interface(desc.Number, 0)

			// Iterate through endpoints available for this interface.
			for _, endpointDesc := range intf.Setting.Endpoints {
				// We only want to read, so we're looking for IN endpoints.
				if endpointDesc.Direction == gousb.EndpointDirectionIn {
					endpoint, err := intf.InEndpoint(endpointDesc.Number)
					if err != nil {
						return err
					}

					usbDevice.intf = intf
					usbDevice.inEndpoint = endpoint
				}

				if endpointDesc.Direction == gousb.EndpointDirectionOut {
					endpoint, err := intf.OutEndpoint(endpointDesc.Number)
					if err != nil {
						return err
					}

					usbDevice.outEndpoint = endpoint
				}
			}
		}
	}

	usbDevice.SetConnected(true)

	return nil
}

func findUSBDevice(product, vendor uint16) func(desc *gousb.DeviceDesc) bool {
	return func(desc *gousb.DeviceDesc) bool {
		return desc.Product == gousb.ID(product) && desc.Vendor == gousb.ID(vendor)
	}
}
