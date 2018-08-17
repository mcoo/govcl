
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
    "time"
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TRegistry struct {
    IObject
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewRegistry
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewRegistry(aAccess uint32) *TRegistry {
    r := new(TRegistry)
    r.instance = Registry_Create(aAccess)
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// RegistryFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func RegistryFromInst(inst uintptr) *TRegistry {
    r := new(TRegistry)
    r.instance = inst
    r.ptr = unsafe.Pointer(inst)
    return r
}

// RegistryFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func RegistryFromObj(obj IObject) *TRegistry {
    r := new(TRegistry)
    r.instance = CheckPtr(obj)
    r.ptr = unsafe.Pointer(r.instance)
    return r
}

// RegistryFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func RegistryFromUnsafePointer(ptr unsafe.Pointer) *TRegistry {
    r := new(TRegistry)
    r.instance = uintptr(ptr)
    r.ptr = ptr
    return r
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (r *TRegistry) Free() {
    if r.instance != 0 {
        Registry_Free(r.instance)
        r.instance = 0
        r.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (r *TRegistry) Instance() uintptr {
    return r.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (r *TRegistry) UnsafeAddr() unsafe.Pointer {
    return r.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (r *TRegistry) IsValid() bool {
    return r.instance != 0
}

// TRegistryClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TRegistryClass() TClass {
    return Registry_StaticClassType()
}

// CloseKey
func (r *TRegistry) CloseKey() {
    Registry_CloseKey(r.instance)
}

// CreateKey
func (r *TRegistry) CreateKey(Key string) bool {
    return Registry_CreateKey(r.instance, Key)
}

// DeleteKey
func (r *TRegistry) DeleteKey(Key string) bool {
    return Registry_DeleteKey(r.instance, Key)
}

// DeleteValue
func (r *TRegistry) DeleteValue(Name string) bool {
    return Registry_DeleteValue(r.instance, Name)
}

// HasSubKeys
func (r *TRegistry) HasSubKeys() bool {
    return Registry_HasSubKeys(r.instance)
}

// KeyExists
func (r *TRegistry) KeyExists(Key string) bool {
    return Registry_KeyExists(r.instance, Key)
}

// LoadKey
func (r *TRegistry) LoadKey(Key string, FileName string) bool {
    return Registry_LoadKey(r.instance, Key , FileName)
}

// MoveKey
func (r *TRegistry) MoveKey(OldName string, NewName string, Delete bool) {
    Registry_MoveKey(r.instance, OldName , NewName , Delete)
}

// OpenKey
func (r *TRegistry) OpenKey(Key string, CanCreate bool) bool {
    return Registry_OpenKey(r.instance, Key , CanCreate)
}

// OpenKeyReadOnly
func (r *TRegistry) OpenKeyReadOnly(Key string) bool {
    return Registry_OpenKeyReadOnly(r.instance, Key)
}

// ReadBool
func (r *TRegistry) ReadBool(Name string) bool {
    return Registry_ReadBool(r.instance, Name)
}

// ReadDate
func (r *TRegistry) ReadDate(Name string) time.Time {
    return Registry_ReadDate(r.instance, Name)
}

// ReadDateTime
func (r *TRegistry) ReadDateTime(Name string) time.Time {
    return Registry_ReadDateTime(r.instance, Name)
}

// ReadFloat
func (r *TRegistry) ReadFloat(Name string) float64 {
    return Registry_ReadFloat(r.instance, Name)
}

// ReadInteger
func (r *TRegistry) ReadInteger(Name string) int32 {
    return Registry_ReadInteger(r.instance, Name)
}

// ReadString
func (r *TRegistry) ReadString(Name string) string {
    return Registry_ReadString(r.instance, Name)
}

// ReadTime
func (r *TRegistry) ReadTime(Name string) time.Time {
    return Registry_ReadTime(r.instance, Name)
}

// RegistryConnect
func (r *TRegistry) RegistryConnect(UNCName string) bool {
    return Registry_RegistryConnect(r.instance, UNCName)
}

// RenameValue
func (r *TRegistry) RenameValue(OldName string, NewName string) {
    Registry_RenameValue(r.instance, OldName , NewName)
}

// ReplaceKey
func (r *TRegistry) ReplaceKey(Key string, FileName string, BackUpFileName string) bool {
    return Registry_ReplaceKey(r.instance, Key , FileName , BackUpFileName)
}

// RestoreKey
func (r *TRegistry) RestoreKey(Key string, FileName string) bool {
    return Registry_RestoreKey(r.instance, Key , FileName)
}

// SaveKey
func (r *TRegistry) SaveKey(Key string, FileName string) bool {
    return Registry_SaveKey(r.instance, Key , FileName)
}

// UnLoadKey
func (r *TRegistry) UnLoadKey(Key string) bool {
    return Registry_UnLoadKey(r.instance, Key)
}

// ValueExists
func (r *TRegistry) ValueExists(Name string) bool {
    return Registry_ValueExists(r.instance, Name)
}

// WriteBool
func (r *TRegistry) WriteBool(Name string, Value bool) {
    Registry_WriteBool(r.instance, Name , Value)
}

// WriteDate
func (r *TRegistry) WriteDate(Name string, Value time.Time) {
    Registry_WriteDate(r.instance, Name , Value)
}

// WriteDateTime
func (r *TRegistry) WriteDateTime(Name string, Value time.Time) {
    Registry_WriteDateTime(r.instance, Name , Value)
}

// WriteFloat
func (r *TRegistry) WriteFloat(Name string, Value float64) {
    Registry_WriteFloat(r.instance, Name , Value)
}

// WriteInteger
func (r *TRegistry) WriteInteger(Name string, Value int32) {
    Registry_WriteInteger(r.instance, Name , Value)
}

// WriteString
func (r *TRegistry) WriteString(Name string, Value string) {
    Registry_WriteString(r.instance, Name , Value)
}

// WriteExpandString
func (r *TRegistry) WriteExpandString(Name string, Value string) {
    Registry_WriteExpandString(r.instance, Name , Value)
}

// WriteTime
func (r *TRegistry) WriteTime(Name string, Value time.Time) {
    Registry_WriteTime(r.instance, Name , Value)
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (r *TRegistry) DisposeOf() {
    Registry_DisposeOf(r.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (r *TRegistry) ClassType() TClass {
    return Registry_ClassType(r.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (r *TRegistry) ClassName() string {
    return Registry_ClassName(r.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (r *TRegistry) InstanceSize() int32 {
    return Registry_InstanceSize(r.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (r *TRegistry) InheritsFrom(AClass TClass) bool {
    return Registry_InheritsFrom(r.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (r *TRegistry) Equals(Obj IObject) bool {
    return Registry_Equals(r.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (r *TRegistry) GetHashCode() int32 {
    return Registry_GetHashCode(r.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (r *TRegistry) ToString() string {
    return Registry_ToString(r.instance)
}

// CurrentKey
func (r *TRegistry) CurrentKey() HKEY {
    return Registry_GetCurrentKey(r.instance)
}

// CurrentPath
func (r *TRegistry) CurrentPath() string {
    return Registry_GetCurrentPath(r.instance)
}

// LazyWrite
func (r *TRegistry) LazyWrite() bool {
    return Registry_GetLazyWrite(r.instance)
}

// SetLazyWrite
func (r *TRegistry) SetLazyWrite(value bool) {
    Registry_SetLazyWrite(r.instance, value)
}

// LastError
func (r *TRegistry) LastError() int32 {
    return Registry_GetLastError(r.instance)
}

// LastErrorMsg
func (r *TRegistry) LastErrorMsg() string {
    return Registry_GetLastErrorMsg(r.instance)
}

// RootKey
func (r *TRegistry) RootKey() HKEY {
    return Registry_GetRootKey(r.instance)
}

// SetRootKey
func (r *TRegistry) SetRootKey(value HKEY) {
    Registry_SetRootKey(r.instance, value)
}

// RootKeyName
func (r *TRegistry) RootKeyName() string {
    return Registry_GetRootKeyName(r.instance)
}

// Access
func (r *TRegistry) Access() uint32 {
    return Registry_GetAccess(r.instance)
}

// SetAccess
func (r *TRegistry) SetAccess(value uint32) {
    Registry_SetAccess(r.instance, value)
}

