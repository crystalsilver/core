// Code generated by protoc-gen-go. DO NOT EDIT.
// source: capabilities.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Capabilities struct {
	Cpu []*CPUDevice `protobuf:"bytes,1,rep,name=cpu" json:"cpu,omitempty"`
	Mem *RAMDevice   `protobuf:"bytes,2,opt,name=mem" json:"mem,omitempty"`
	Gpu []*GPUDevice `protobuf:"bytes,3,rep,name=gpu" json:"gpu,omitempty"`
}

func (m *Capabilities) Reset()                    { *m = Capabilities{} }
func (m *Capabilities) String() string            { return proto.CompactTextString(m) }
func (*Capabilities) ProtoMessage()               {}
func (*Capabilities) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Capabilities) GetCpu() []*CPUDevice {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Capabilities) GetMem() *RAMDevice {
	if m != nil {
		return m.Mem
	}
	return nil
}

func (m *Capabilities) GetGpu() []*GPUDevice {
	if m != nil {
		return m.Gpu
	}
	return nil
}

type CPUDevice struct {
	// Num describes the CPU number on a board.
	Num int32 `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
	// VendorId describes vendor id as a string, for example "GenuineIntel".
	VendorId string `protobuf:"bytes,2,opt,name=vendorId" json:"vendorId,omitempty"`
	// Model describes model.
	Model string `protobuf:"bytes,3,opt,name=model" json:"model,omitempty"`
	// ModelName describes full model name.
	// For example "Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz".
	ModelName string `protobuf:"bytes,4,opt,name=modelName" json:"modelName,omitempty"`
	// Cores describes number of cores on a CPU device.
	Cores int32 `protobuf:"varint,5,opt,name=cores" json:"cores,omitempty"`
	// Maximum configured clock frequency of the device in MHz.
	ClockFrequency float64 `protobuf:"fixed64,6,opt,name=clockFrequency" json:"clockFrequency,omitempty"`
	// CacheSize shows CPU cache size.
	CacheSize int32 `protobuf:"varint,7,opt,name=cacheSize" json:"cacheSize,omitempty"`
	// Stepping describes CPU stepping level.
	Stepping int32 `protobuf:"varint,8,opt,name=stepping" json:"stepping,omitempty"`
	// Flags describes supported extensions on a CPU.
	Flags []string `protobuf:"bytes,9,rep,name=flags" json:"flags,omitempty"`
}

func (m *CPUDevice) Reset()                    { *m = CPUDevice{} }
func (m *CPUDevice) String() string            { return proto.CompactTextString(m) }
func (*CPUDevice) ProtoMessage()               {}
func (*CPUDevice) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CPUDevice) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *CPUDevice) GetVendorId() string {
	if m != nil {
		return m.VendorId
	}
	return ""
}

func (m *CPUDevice) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *CPUDevice) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *CPUDevice) GetCores() int32 {
	if m != nil {
		return m.Cores
	}
	return 0
}

func (m *CPUDevice) GetClockFrequency() float64 {
	if m != nil {
		return m.ClockFrequency
	}
	return 0
}

func (m *CPUDevice) GetCacheSize() int32 {
	if m != nil {
		return m.CacheSize
	}
	return 0
}

func (m *CPUDevice) GetStepping() int32 {
	if m != nil {
		return m.Stepping
	}
	return 0
}

func (m *CPUDevice) GetFlags() []string {
	if m != nil {
		return m.Flags
	}
	return nil
}

type RAMDevice struct {
	Total uint64 `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
	Used  uint64 `protobuf:"varint,2,opt,name=used" json:"used,omitempty"`
}

func (m *RAMDevice) Reset()                    { *m = RAMDevice{} }
func (m *RAMDevice) String() string            { return proto.CompactTextString(m) }
func (*RAMDevice) ProtoMessage()               {}
func (*RAMDevice) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *RAMDevice) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *RAMDevice) GetUsed() uint64 {
	if m != nil {
		return m.Used
	}
	return 0
}

type GPUDevice struct {
	// Name describes full GPU device name.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// VendorId describes vendor id.
	VendorId uint64 `protobuf:"varint,2,opt,name=vendorId" json:"vendorId,omitempty"`
	// VendorName describes vendor name, for example "NVIDIA" or "AMD".
	VendorName string `protobuf:"bytes,3,opt,name=vendorName" json:"vendorName,omitempty"`
	// Total maximum memory size the device can hold.
	MaxMemorySize uint64 `protobuf:"varint,4,opt,name=maxMemorySize" json:"maxMemorySize,omitempty"`
	// Maximum configured clock frequency of the device in MHz.
	MaxClockFrequency uint64 `protobuf:"varint,5,opt,name=maxClockFrequency" json:"maxClockFrequency,omitempty"`
	// OpenCL major version.
	OpenCLDeviceVersionMajor int32 `protobuf:"varint,6,opt,name=openCLDeviceVersionMajor" json:"openCLDeviceVersionMajor,omitempty"`
	// OpenCL minor version.
	OpenCLDeviceVersionMinor int32 `protobuf:"varint,7,opt,name=openCLDeviceVersionMinor" json:"openCLDeviceVersionMinor,omitempty"`
}

func (m *GPUDevice) Reset()                    { *m = GPUDevice{} }
func (m *GPUDevice) String() string            { return proto.CompactTextString(m) }
func (*GPUDevice) ProtoMessage()               {}
func (*GPUDevice) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GPUDevice) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GPUDevice) GetVendorId() uint64 {
	if m != nil {
		return m.VendorId
	}
	return 0
}

func (m *GPUDevice) GetVendorName() string {
	if m != nil {
		return m.VendorName
	}
	return ""
}

func (m *GPUDevice) GetMaxMemorySize() uint64 {
	if m != nil {
		return m.MaxMemorySize
	}
	return 0
}

func (m *GPUDevice) GetMaxClockFrequency() uint64 {
	if m != nil {
		return m.MaxClockFrequency
	}
	return 0
}

func (m *GPUDevice) GetOpenCLDeviceVersionMajor() int32 {
	if m != nil {
		return m.OpenCLDeviceVersionMajor
	}
	return 0
}

func (m *GPUDevice) GetOpenCLDeviceVersionMinor() int32 {
	if m != nil {
		return m.OpenCLDeviceVersionMinor
	}
	return 0
}

func init() {
	proto.RegisterType((*Capabilities)(nil), "sonm.Capabilities")
	proto.RegisterType((*CPUDevice)(nil), "sonm.CPUDevice")
	proto.RegisterType((*RAMDevice)(nil), "sonm.RAMDevice")
	proto.RegisterType((*GPUDevice)(nil), "sonm.GPUDevice")
}

func init() { proto.RegisterFile("capabilities.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0xa5, 0x37, 0xc9, 0xba, 0x29, 0xbf, 0x1b, 0x0f, 0x8d, 0x88, 0xc4, 0x41, 0x64, 0x0e, 0x32,
	0x07, 0xc5, 0x8b, 0x37, 0x89, 0x28, 0x82, 0x23, 0xd2, 0xa2, 0xf7, 0xde, 0x9e, 0x32, 0xb6, 0xa6,
	0x3f, 0xec, 0x24, 0xeb, 0xae, 0x7f, 0xc8, 0x1f, 0xe9, 0x45, 0xba, 0xa2, 0x49, 0xd6, 0x61, 0x6e,
	0xf5, 0xea, 0xbd, 0x7e, 0x55, 0xf5, 0x12, 0xe0, 0x5a, 0x05, 0x75, 0x6a, 0x5a, 0xd3, 0x1b, 0xec,
	0x36, 0x21, 0xfa, 0xde, 0xf3, 0xbc, 0xf3, 0xce, 0xae, 0x7e, 0xc0, 0xb5, 0x7a, 0xc1, 0xf1, 0x07,
	0x90, 0xe9, 0x30, 0x08, 0x56, 0x65, 0xeb, 0xab, 0x4f, 0x6e, 0x6e, 0x92, 0x66, 0x53, 0xbf, 0xff,
	0xf8, 0x12, 0xcf, 0x8c, 0x46, 0x99, 0xb8, 0x24, 0xb1, 0x68, 0xc5, 0x51, 0xc5, 0x66, 0x89, 0x7c,
	0xb1, 0xfd, 0x27, 0xb1, 0x68, 0x93, 0xa4, 0x09, 0x83, 0xc8, 0x96, 0x2e, 0xaf, 0x67, 0x97, 0x26,
	0x0c, 0xab, 0xdf, 0x0c, 0xca, 0xc9, 0x98, 0xdf, 0x82, 0xcc, 0x0d, 0x56, 0xb0, 0x8a, 0xad, 0x0b,
	0x99, 0x4a, 0x7e, 0x17, 0x4e, 0xce, 0xd0, 0xed, 0x7c, 0x7c, 0xb3, 0xa3, 0x51, 0xa5, 0x9c, 0x30,
	0xbf, 0x03, 0x85, 0xf5, 0x3b, 0x6c, 0x45, 0x46, 0xc4, 0x08, 0xf8, 0x3d, 0x28, 0xa9, 0x78, 0xa7,
	0x2c, 0x8a, 0x9c, 0x98, 0xb9, 0x91, 0xde, 0x68, 0x1f, 0xb1, 0x13, 0x05, 0xcd, 0x18, 0x01, 0x7f,
	0x04, 0x37, 0x74, 0xeb, 0xf5, 0xb7, 0x57, 0x11, 0xbf, 0x0f, 0xe8, 0xf4, 0x85, 0x38, 0xae, 0xd8,
	0x9a, 0xc9, 0xff, 0xba, 0xc9, 0x5b, 0x2b, 0xfd, 0x05, 0x3f, 0x98, 0x9f, 0x28, 0xae, 0x90, 0xc3,
	0xdc, 0x48, 0xbb, 0x76, 0x3d, 0x86, 0x60, 0x5c, 0x23, 0x4e, 0x88, 0x9c, 0x70, 0x9a, 0xfb, 0xb9,
	0x55, 0x4d, 0x27, 0xca, 0x2a, 0x4b, 0xbb, 0x12, 0x58, 0x3d, 0x83, 0x72, 0x8a, 0x2c, 0x49, 0x7a,
	0xdf, 0xab, 0x96, 0xce, 0xcf, 0xe5, 0x08, 0x38, 0x87, 0x7c, 0xe8, 0x70, 0x3c, 0x3e, 0x97, 0x54,
	0xaf, 0x7e, 0x1d, 0x41, 0x39, 0xe5, 0x98, 0x14, 0x2e, 0xdd, 0xca, 0xe8, 0x56, 0xaa, 0xf7, 0x62,
	0xcb, 0x17, 0xb1, 0xdd, 0x07, 0x18, 0x6b, 0x4a, 0x68, 0xcc, 0x6e, 0xd1, 0xe1, 0x0f, 0xe1, 0xba,
	0x55, 0xe7, 0x5b, 0xb4, 0x3e, 0x5e, 0xd0, 0xa1, 0x39, 0x19, 0x5c, 0x6e, 0xf2, 0xc7, 0x70, 0xdb,
	0xaa, 0xf3, 0xfa, 0x72, 0x6a, 0x05, 0x29, 0xf7, 0x09, 0xfe, 0x1c, 0x84, 0x0f, 0xe8, 0xea, 0xb7,
	0xe3, 0xce, 0x9f, 0x30, 0x76, 0xc6, 0xbb, 0xad, 0xfa, 0xea, 0x23, 0x45, 0x5d, 0xc8, 0x83, 0xfc,
	0xa1, 0xb7, 0xc6, 0xf9, 0xf8, 0xf7, 0x1b, 0x1c, 0xe4, 0x4f, 0x8f, 0xe9, 0x27, 0x7f, 0xfa, 0x27,
	0x00, 0x00, 0xff, 0xff, 0xba, 0x2c, 0xe9, 0x3a, 0xfa, 0x02, 0x00, 0x00,
}
