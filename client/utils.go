package client

type ResourcePostResponseData struct {
	ResourceData *ResourceData `json:"data,omitempty"`
}

type ResourceData struct {
	ID string `json:"id,omitempty"`
}

func GetCloudTypes() []string {
	cloudTypes := make([]string, 4)
	cloudTypes[0] = AWSCloudType
	cloudTypes[1] = AzureCloudType
	cloudTypes[2] = GCPCloudType
	// cloudTypes[3] = PrivateCloudType
	return cloudTypes
}

func GetPVTConnectionTypes() []string {
	PvtConnectionType := make([]string, 2)
	PvtConnectionType[0] = PrivateConnection_type
	PvtConnectionType[1] = PublicConnection_type
	return PvtConnectionType
}

func GetAWSAuthTypes() []string {
	awsAuthTypes := make([]string, 2)
	awsAuthTypes[0] = AWSIAMRoleAuth
	awsAuthTypes[1] = AWSAccessKeyAuth
	return awsAuthTypes
}

func GetIDPAccountTypes() []string {
	idpAccountTypes := make([]string, 7)
	idpAccountTypes[0] = IDPOktaAccount
	idpAccountTypes[1] = IDPAzureADAccount
	idpAccountTypes[2] = IDPOneLoginAccount
	idpAccountTypes[3] = IDPOtherAccount
	idpAccountTypes[4] = IDPPingOneAccount
	idpAccountTypes[5] = IDPPingFederateAccount
	idpAccountTypes[6] = IDPGoogleWSAccount
	return idpAccountTypes
}

func GetIDPAuthTypes() []string {
	idpAuthTypes := make([]string, 2)
	idpAuthTypes[0] = IDPOIDCAuth
	idpAuthTypes[1] = IDPSAMLAuth
	return idpAuthTypes
}

func GetIDPTypes() []string {
	idpTypes := make([]string, 2)
	idpTypes[0] = IDPPrimaryType
	idpTypes[1] = IDPPartnerType
	return idpTypes
}

func GetAPICredProvided() []string {
	apiCredProvided := make([]string, 2)
	apiCredProvided[0] = APICredYes
	apiCredProvided[1] = APICredNo
	return apiCredProvided
}

func GetIDPPRegionTypes() []string {
	regionTypes := make([]string, 4)
	regionTypes[0] = IDPEURegion
	regionTypes[1] = IDPAsiaRegion
	regionTypes[2] = IDPDefaultRegion
	regionTypes[3] = IDPUSRegion
	return regionTypes
}

func GetAppAccessTypes() []string {
	appAccessTypes := make([]string, 2)
	appAccessTypes[0] = AgentAppAccessType
	appAccessTypes[1] = AgentlessAppAccessType
	return appAccessTypes
}

func GetAppAccessTypesU2AALWeb() []string {
	appAccessTypes := make([]string, 1)
	appAccessTypes[0] = AgentlessAppAccessType
	return appAccessTypes
}

func GetAppTypes() []string {
	appTypes := make([]string, 5)
	appTypes[0] = DefaultAppType
	appTypes[1] = CitrixAppType
	appTypes[2] = TCPAppType
	appTypes[3] = CloudPlatformType
	appTypes[4] = JumpBoxType

	return appTypes
}

func GetAppTypesU2AALWeb() []string {
	appTypes := make([]string, 1)
	appTypes[0] = DefaultAppType
	return appTypes
}

func GetAppDomainTypes() []string {
	appDomainTypes := make([]string, 2)
	appDomainTypes[0] = CustomAppDomain
	appDomainTypes[1] = ProsimoAppDomain
	return appDomainTypes
}

func GetSelectedAuthTypes() []string {
	selectedAuthTypes := make([]string, 3)
	selectedAuthTypes[0] = SamlAuth
	selectedAuthTypes[1] = OidcAuth
	selectedAuthTypes[2] = OtherAuth
	return selectedAuthTypes
}

func GetAppProtocols() []string {
	appProtocols := make([]string, 7)
	appProtocols[0] = HTTPAppProtocol
	appProtocols[1] = HTTPSAppProtocol
	appProtocols[2] = SSHAppProtocol
	appProtocols[3] = RDPAppProtocol
	appProtocols[4] = VNCAppProtocol
	appProtocols[5] = TCPAppProtocol
	appProtocols[6] = UDPAppProtocol
	return appProtocols
}

func GetAppProtocolsLWeb() []string {
	U2AALWebappProtocols := make([]string, 5)
	U2AALWebappProtocols[0] = HTTPU2AALWebAppProtocol
	U2AALWebappProtocols[1] = HTTPSU2AALWebAppProtocol
	U2AALWebappProtocols[2] = SSHU2AALWebAppProtocol
	U2AALWebappProtocols[3] = RDPU2AALWebAppProtocol
	U2AALWebappProtocols[4] = VNCU2AALWebAppProtocol
	return U2AALWebappProtocols
}

func GetAppProtocolsLFQDN() []string {
	fqdnappProtocols := make([]string, 3)
	fqdnappProtocols[0] = TCPfqdnAppProtocol
	fqdnappProtocols[1] = UDPfqdnAppProtocol
	fqdnappProtocols[2] = DNSfqdnAppProtocol
	return fqdnappProtocols
}

func GetCloudConnectionOptions() []string {
	connectionOptions := make([]string, 2)
	connectionOptions[0] = PublicCloudConnectionOption
	connectionOptions[1] = PrivateCloudConnectionOption
	return connectionOptions
}

func GetConnectorPlacementOptions() []string {
	ConnectorPlacementOptions := make([]string, 3)
	ConnectorPlacementOptions[0] = WorkloadVpcConnectorPlacementOptions
	ConnectorPlacementOptions[1] = NoneConnectorPlacementOptions
	ConnectorPlacementOptions[2] = InfraVPCConnectorPlacementOptions
	return ConnectorPlacementOptions
}

func GetCloudTypeOptions() []string {
	CloudtypeOptions := make([]string, 2)
	CloudtypeOptions[0] = PublicCloudConnectionOption
	CloudtypeOptions[1] = PrivateCloudConnectionOption
	return CloudtypeOptions
}
func AppOnboardConnOptn() []string {
	ConnOptn := make([]string, 9)
	ConnOptn[0] = Optnpeering
	ConnOptn[1] = OptnawsPrivateLink
	ConnOptn[2] = OptntransitGateway
	ConnOptn[3] = OptnawsVpnGateway
	ConnOptn[4] = OptnazurePrivateLink
	ConnOptn[5] = OptnazureTransitVnet
	ConnOptn[6] = OptnvwanHub
	ConnOptn[7] = Optnpublic
	ConnOptn[8] = Optnprivate
	return ConnOptn
}

// To be reviewed if it's required anymore.
func AppHostedOptn() []string {
	HostedOptn := make([]string, 1)
	HostedOptn[0] = HostedPrivate
	return HostedOptn
}

func GetDNSServiceTypes() []string {
	dnsServiceTypes := make([]string, 3)
	dnsServiceTypes[0] = ManualDNSServiceType
	dnsServiceTypes[1] = AWSRoute53DNSServiceType
	dnsServiceTypes[2] = ProsimoDNSServiceType
	return dnsServiceTypes
}

func GetAppOnboardingOptimization() []string {
	optOtions := make([]string, 3)
	optOtions[0] = PerformanceEnhancedOptOption
	optOtions[1] = CostSavingOptOption
	optOtions[2] = FastLaneOptOption
	return optOtions
}

func GetAppOnboardingPolicy() []string {
	policyOtions := make([]string, 3)
	policyOtions[0] = AllowAllPolicyOption
	policyOtions[1] = DenyAllPolicyOption
	policyOtions[2] = CustomizePolicyOption
	return policyOtions
}

func GetNetworkOnboardingPolicy() []string {
	policyOtions := make([]string, 3)
	policyOtions[0] = AllowAllPolicyOptionNetwork
	policyOtions[1] = DenyAllPolicyOptionNetwork
	policyOtions[2] = CustomizePolicyOption
	return policyOtions
}

func GetAppOnboardTypes() []string {
	appOnboardTypes := make([]string, 3)
	appOnboardTypes[0] = BehindFabricAppOnboardType
	appOnboardTypes[1] = AccessingFabricAppOnboardType
	appOnboardTypes[2] = BothAppOnboardType
	return appOnboardTypes
}
func GetAppOnboardTypesU2AALWeb() []string {
	appOnboardTypes := make([]string, 1)
	appOnboardTypes[0] = BehindFabricAppOnboardType
	return appOnboardTypes
}

func AppOnboardRegionType() []string {
	appOnboardRegionTypes := make([]string, 2)
	appOnboardRegionTypes[0] = ActiveRegionType
	appOnboardRegionTypes[1] = BackUpRegionType
	return appOnboardRegionTypes
}

func GetInteractionTypes() []string {
	interactionTypes := make([]string, 2)
	interactionTypes[0] = UserToAppInteractionType
	interactionTypes[1] = AppToAppInteractionType
	return interactionTypes
}

func GetAddressTypes() []string {
	addressTypes := make([]string, 2)
	addressTypes[0] = FQDNAddessType
	addressTypes[1] = IPAddressType
	return addressTypes
}

func GetTgwAppRoutetableType() []string {
	TgwAppRoutetableType := make([]string, 2)
	TgwAppRoutetableType[0] = ModifyTgwAppRoutetableType
	TgwAppRoutetableType[1] = MaintainTgwAppRoutetableType
	return TgwAppRoutetableType
}
func GetAddressTypesU2AALWeb() []string {
	addressTypes := make([]string, 1)
	addressTypes[0] = FQDNAddessType
	return addressTypes
}

func GetWafModeTypes() []string {
	wafModeTypes := make([]string, 3)
	wafModeTypes[0] = WafModeEnforce
	wafModeTypes[1] = WafModeDetect
	return wafModeTypes
}
func GetDyRiskNameTypes() []string {
	DyRiskNameTypes := make([]string, 3)
	DyRiskNameTypes[0] = DRNameAlert
	DyRiskNameTypes[1] = DRNameMfa
	DyRiskNameTypes[2] = DRNameLockUser
	return DyRiskNameTypes
}

func GetPolicyResourceTypes() []string {
	PolicyResourceTypes := make([]string, 11)
	PolicyResourceTypes[0] = PolicyMatchUser
	PolicyResourceTypes[1] = PolicyMatchNetwork
	PolicyResourceTypes[2] = PolicyMatchDevice
	PolicyResourceTypes[3] = PolicyMatchTime
	PolicyResourceTypes[4] = PolicyMatchUrl
	PolicyResourceTypes[5] = PolicyMatchApplication
	PolicyResourceTypes[6] = PolicyMatchAppAdvanced
	PolicyResourceTypes[7] = PolicyMatchAppIDP
	PolicyResourceTypes[8] = PolicyMatchAppDP
	PolicyResourceTypes[9] = PolicyMatchPosition
	PolicyResourceTypes[10] = PolicyMatchNetworkACL

	return PolicyResourceTypes
}

func GetPolicyResourceOperation() []string {
	PolicyResourceOperation := make([]string, 10)
	PolicyResourceOperation[0] = PolicyMatchOperationIs
	PolicyResourceOperation[1] = PolicyMatchOperationIsNot
	PolicyResourceOperation[2] = PolicyMatchOperationContains
	PolicyResourceOperation[3] = PolicyMatchOperationDoesNotContain
	PolicyResourceOperation[4] = PolicyMatchOperationStartsWith
	PolicyResourceOperation[5] = PolicyMatchOperationEndsWith
	PolicyResourceOperation[6] = PolicyMatchOperationIn
	PolicyResourceOperation[7] = PolicyMatchOperationNotIn
	PolicyResourceOperation[8] = PolicyMatchOperationIsAtLeast
	PolicyResourceOperation[9] = PolicyMatchOperationBetween

	return PolicyResourceOperation
}

func GetPolicyappAccessType() []string {
	PolicyappAccessType := make([]string, 2)
	PolicyappAccessType[0] = PolicyTransit
	PolicyappAccessType[1] = PolicyAccess
	return PolicyappAccessType
}

func GetPolicyActionTypes() []string {
	PolicyActionTypes := make([]string, 6)
	PolicyActionTypes[0] = PolicyActionAllow
	PolicyActionTypes[1] = PolicyActionDeny
	// PolicyActionTypes[2] = PolicyActionBypass
	// PolicyActionTypes[3] = PolicyActionSkipWAF
	// PolicyActionTypes[4] = PolicyActionAlert
	// PolicyActionTypes[5] = PolicyActionMFA
	return PolicyActionTypes
}

func GetCacheTimeUnit() []string {
	CacheTimeUnit := make([]string, 4)
	CacheTimeUnit[0] = CacheTimeUnitDays
	CacheTimeUnit[1] = CacheTimeUnitHours
	CacheTimeUnit[2] = CacheTimeUnitMinutes
	CacheTimeUnit[3] = CacheTimeUnitSeconds
	return CacheTimeUnit
}

func GetCacheType() []string {
	CacheType := make([]string, 3)
	CacheType[0] = CacheTypeDynamic
	CacheType[1] = CacheTypeStaticLongLived
	CacheType[2] = CacheTypeStaticShortLived
	return CacheType
}

func GetEDRvendorTypes() []string {
	EDRvendorTypes := make([]string, 1)
	EDRvendorTypes[0] = TypeCrowdStrike
	return EDRvendorTypes
}

func GetEDRbaseurlTypes() []string {
	EDRbaseurlTypes := make([]string, 4)
	EDRbaseurlTypes[0] = ApiCrowdStrike
	EDRbaseurlTypes[1] = ApiUSCrowdStrike
	EDRbaseurlTypes[2] = ApiEUCrowdStrike
	EDRbaseurlTypes[3] = ApiLaggerGCWCrowdStrike
	return EDRbaseurlTypes
}

func GetEDRProfileInputTypes() []string {
	EDRProfileInputTypes := make([]string, 3)
	EDRProfileInputTypes[0] = TypeEnabled
	EDRProfileInputTypes[1] = TypeDisabled
	EDRProfileInputTypes[2] = TypeNa
	return EDRProfileInputTypes
}

func GetDPRiskLevelTypes() []string {
	DPRiskLevelTypes := make([]string, 3)
	DPRiskLevelTypes[0] = RiskLevelHigh
	DPRiskLevelTypes[1] = RiskLevelMedium
	DPRiskLevelTypes[2] = RiskLevelLow
	return DPRiskLevelTypes
}

func GetOsTypes() []string {
	DPOsTypes := make([]string, 2)
	DPOsTypes[0] = OsWindows
	DPOsTypes[1] = OsMac
	return DPOsTypes
}

func GetOsOperatorTypes() []string {
	DPOsOperatorTypes := make([]string, 3)
	DPOsOperatorTypes[0] = TypeIS
	DPOsOperatorTypes[1] = TypeISNOT
	DPOsOperatorTypes[2] = TypeISATLEAST
	return DPOsOperatorTypes
}

func GetGroupingTypes() []string {
	GroupingTypes := make([]string, 6)
	GroupingTypes[0] = TypeUSER
	GroupingTypes[1] = TypeAPP
	GroupingTypes[2] = TypeDEVICE
	GroupingTypes[3] = TypeTIME
	GroupingTypes[4] = TypeIP_RANGE
	GroupingTypes[5] = TypeGEO
	return GroupingTypes
}
