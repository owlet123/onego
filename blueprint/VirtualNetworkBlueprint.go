package blueprint

// VirtualNetworkBlueprint to set VirtualNetwork elements
type VirtualNetworkBlueprint struct {
	Blueprint
}

// CreateAllocateVirtualNetworkBlueprint creates empty VirtualNetworkBlueprint
func CreateAllocateVirtualNetworkBlueprint() *VirtualNetworkBlueprint {
	return &VirtualNetworkBlueprint{Blueprint: *CreateBlueprint("VNET")}
}

// CreateUpdateVirtualNetworkBlueprint creates empty VirtualNetworkBlueprint
func CreateUpdateVirtualNetworkBlueprint() *VirtualNetworkBlueprint {
	return &VirtualNetworkBlueprint{Blueprint: *CreateBlueprint("TEMPLATE")}
}

// SetBridge sets Bridge of the given Virtual Network
func (vnb *VirtualNetworkBlueprint) SetBridge(bridge string) {
	vnb.SetElement("BRIDGE", bridge)
}

// SetFilterIPSpoofing sets filter ip spoofing of the given Virtual Network
func (vnb *VirtualNetworkBlueprint) SetFilterIPSpoofing(filterIPSpoofing string) {
	vnb.SetElement("FILTER_IP_SPOOFING", filterIPSpoofing)
}

// SetFilterMacSpoofing sets filter mac spoofing of the given Virtual Network
func (vnb *VirtualNetworkBlueprint) SetFilterMacSpoofing(filterMacSpoofing string) {
	vnb.SetElement("FILTER_MAC_SPOOFING", filterMacSpoofing)
}

// SetGateway sets Gateway of the given Virtual Network
func (vnb *VirtualNetworkBlueprint) SetGateway(gateway string) {
	vnb.SetElement("GATEWAY", gateway)
}

// SetMTU sets MTU of the given Virtual Network
func (vnb *VirtualNetworkBlueprint) SetMTU(mtu string) {
	vnb.SetElement("MTU", mtu)
}

// SetNetworkAddress sets network address of the given Virtual Network
func (vnb *VirtualNetworkBlueprint) SetNetworkAddress(networkAddress string) {
	vnb.SetElement("NETWORK_ADDRESS", networkAddress)
}

// SetNetworkMask sets network mask of the given Virtual Network
func (vnb *VirtualNetworkBlueprint) SetNetworkMask(networkMask string) {
	vnb.SetElement("NETWORK_MASK", networkMask)
}

// SetPhysicalDevice sets physical device of the given Virtual Network
func (vnb *VirtualNetworkBlueprint) SetPhysicalDevice(physicalDevice string) {
	vnb.SetElement("PHYDEV", physicalDevice)
}

// SetSecurityGroups sets security groups of the given Virtual Network
func (vnb *VirtualNetworkBlueprint) SetSecurityGroups(securityGroups string) {
	vnb.SetElement("SECURITY_GROUPS", securityGroups)
}

// SetVirtualLanID sets virtual lan ID of the given Virtual Network
func (vnb *VirtualNetworkBlueprint) SetVirtualLanID(virtualLanID string) {
	vnb.SetElement("VLAN_ID", virtualLanID)
}

// SetVnMad sets VN_MAD of the given Virtual Network
func (vnb *VirtualNetworkBlueprint) SetVnMad(vnMad string) {
	vnb.SetElement("VN_MAD", vnMad)
}
