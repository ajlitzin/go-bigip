package bigip

type PersistenceCookies struct {
	PersistenceCookie []PersistenceCookie `json:"items,omitempty"`
}

type PersistenceCookie struct {
	Name                       string `json:"name"`
	DefaultsFrom               string `json:"defaultsFrom"`
	Kind                       string `json:"kind,omitempty"`
	Mode                       string `json:"mode,omitempty"`
	Partition                  string `json:"partition,omitempty"`
	FullPath                   string `json:"fullPath,omitempty"`
	Generation                 int    `json:"generation,omitempty"`
	SelfLink                   string `json:"selfLink,omitempty"`
	AlwaysSend                 string `json:"alwaysSend,omitempty"`
	Description                string `json:"description,omitempty"`
	AppService                 string `json:"appService,omitempty"`
	CookieEncryption           string `json:"cookieEncryption,omitempty"`
	CookieEncryptionPassphrase string `json:"cookieEncryptionPassphrase,omitempty"`
	CookieName                 string `json:"cookieName,omitempty"`
	Expiration                 string `json:"expiration,omitempty"`
	HashLength                 int    `json:"hashLength,omitempty"`
	HashOffset                 int    `json:"hashOffset,omitempty"`
	MatchAcrossPools           string `json:"matchAcrossPools,omitempty"`
	MatchAcrossServices        string `json:"matchAcrossServices,omitempty"`
	MatchAcrossVirtuals        string `json:"matchAcrossVirtuals,omitempty"`
	Method                     string `json:"method,omitempty"`
	Mirror                     string `json:"mirror,omitempty"`
	Secure                     string `json:"secure,omitempty"`
	TMPartition                string `json:"tmPartition,omitempty"`
	OverrideConnectionLimit    string `json:"overrideConnectionLimit,omitempty"`
	Timeout                    string `json:"timeout,omitempty"`
}

// PersistenceCookie returns a list of oersistence profiles.
func (b *BigIP) PersistenceCookie() (*PersistenceCookie, error) {
	var persistenceProfiles PersistenceCookie
	err, _ := b.getForEntity(&persistenceProfiles, uriLtm, uriPersistence, uriCookie)
	if err != nil {
		return nil, err
	}

	return &persistenceProfiles, nil
}

// GetPersistenceCookie gets a persistence profile by name. Returns nil if the persistence profile does not exist
func (b *BigIP) GetPersistenceCookie(name string) (*PersistenceCookie, error) {
	var persistenceProfile PersistenceCookie
	err, ok := b.getForEntity(&persistenceProfile, uriLtm, uriPersistence, uriCookie, name)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	return &persistenceProfile, nil
}

// CreatePersistenceCookie creates a new persistence profile on the BIG-IP system.
func (b *BigIP) CreatePersistenceCookie(name string, parent string) error {
	config := &PersistenceCookie{
		Name:         name,
		DefaultsFrom: parent,
	}

	return b.post(config, uriLtm, uriPersistence, uriCookie)
}

// AddPersistenceCookie adds a new persistence profile on the BIG-IP system.
func (b *BigIP) AddPersistenceCookie(config *PersistenceCookie) error {
	return b.post(config, uriLtm, uriPersistence, uriCookie)
}

// DeletePersistenceCookie removes a persistence profile.
func (b *BigIP) DeletePersistenceCookie(name string) error {
	return b.delete(uriLtm, uriPersistence, uriCookie, name)
}

// ModifyPersistenceCookie allows you to change any attribute of a persistence profile.
// Fields that can be modified are referenced in the PersistenceCookie struct.
func (b *BigIP) ModifyPersistenceCookie(name string, config *PersistenceCookie) error {
	return b.put(config, uriLtm, uriPersistence, uriCookie, name)
}

type PersistenceHashes struct {
	PersistenceHash []PersistenceHash `json:"items,omitempty"`
}

type PersistenceHash struct {
	Kind                    string `json:"kind,omitempty"`
	DefaultsFrom            string `json:"defaultsFrom"`
	Name                    string `json:"name"`
	Partition               string `json:"partition,omitempty"`
	FullPath                string `json:"fullPath,omitempty"`
	Generation              int    `json:"generation,omitempty"`
	SelfLink                string `json:"selfLink,omitempty"`
	HashAlgorithm           string `json:"hashAlgorithm,omitempty"`
	HashBufferLimit         int    `json:"hashBufferLimit,omitempty"`
	HashLength              int    `json:"hashLength,omitempty"`
	HashOffset              int    `json:"hashOffset,omitempty"`
	MatchAcrossPools        string `json:"matchAcrossPools,omitempty"`
	MatchAcrossServices     string `json:"matchAcrossServices,omitempty"`
	MatchAcrossVirtuals     string `json:"matchAcrossVirtuals,omitempty"`
	AppService              string `json:"appService,omitempty"`
	Description             string `json:"description,omitempty"`
	HashEndPattern          string `json:"hashEndPattern,omitempty"`
	HashStartPattern        string `json:"hashStartPattern,omitempty"`
	Mode                    string `json:"mode,omitempty"`
	Rule                    string `json:"rule,omitempty"`
	TMPartition             string `json:"tmPartition,omitempty"`
	Mirror                  string `json:"mirror,omitempty"`
	OverrideConnectionLimit string `json:"overrideConnectionLimit,omitempty"`
	Timeout                 string `json:"timeout,omitempty"`
}

// PersistenceHash returns a list of persistence profiles.
func (b *BigIP) PersistenceHash() (*PersistenceHash, error) {
	var persistenceProfiles PersistenceHash
	err, _ := b.getForEntity(&persistenceProfiles, uriLtm, uriPersistence, uriHash)
	if err != nil {
		return nil, err
	}

	return &persistenceProfiles, nil
}

// GetPersistenceHash gets a persistence profile by name. Returns nil if the persistence profile does not exist
func (b *BigIP) GetPersistenceHash(name string) (*PersistenceHash, error) {
	var persistenceProfile PersistenceHash
	err, ok := b.getForEntity(&persistenceProfile, uriLtm, uriPersistence, uriHash, name)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	return &persistenceProfile, nil
}

// CreatePersistenceHash creates a new persistence profile on the BIG-IP system.
func (b *BigIP) CreatePersistenceHash(name string, parent string) error {
	config := &PersistenceHash{
		Name:         name,
		DefaultsFrom: parent,
	}

	return b.post(config, uriLtm, uriPersistence, uriHash)
}

// AddPersistenceHash adds a new persistence profile on the BIG-IP system.
func (b *BigIP) AddPersistenceHash(config *PersistenceHash) error {
	return b.post(config, uriLtm, uriPersistence, uriHash)
}

// DeletePersistenceHash removes a persistence profile.
func (b *BigIP) DeletePersistenceHash(name string) error {
	return b.delete(uriLtm, uriPersistence, uriHash, name)
}

// ModifyPersistenceHash allows you to change any attribute of a persistence profile.
// Fields that can be modified are referenced in the PersistenceHash struct.
func (b *BigIP) ModifyPersistenceHash(name string, config *PersistenceHash) error {
	return b.put(config, uriLtm, uriPersistence, uriHash, name)
}

type PersistenceSourceAddres struct {
	PersistenceSourceAddr []PersistenceSourceAddr `json:"items,omitempty"`
}

type PersistenceSourceAddr struct {
	Name                    string `json:"name"`
	DefaultsFrom            string `json:"defaultsFrom"`
	Kind                    string `json:"kind,omitempty"`
	Partition               string `json:"partition,omitempty"`
	FullPath                string `json:"fullPath,omitempty"`
	Generation              int    `json:"generation,omitempty"`
	SelfLink                string `json:"selfLink,omitempty"`
	HashAlgorithm           string `json:"hashAlgorithm,omitempty"`
	MapProxies              string `json:"mapProxies,omitempty"`
	MatchAcrossPools        string `json:"matchAcrossPools,omitempty"`
	MatchAcrossServices     string `json:"matchAcrossServices,omitempty"`
	MatchAcrossVirtuals     string `json:"matchAcrossVirtuals,omitempty"`
	Mirror                  string `json:"mirror,omitempty"`
	OverrideConnectionLimit string `json:"overrideConnectionLimit,omitempty"`
	Timeout                 string `json:"timeout,omitempty"`
}

// PersistenceSourceAddr returns a list of persistence profiles.
func (b *BigIP) PersistenceSourceAddr() (*PersistenceSourceAddr, error) {
	var persistenceProfiles PersistenceSourceAddr
	err, _ := b.getForEntity(&persistenceProfiles, uriLtm, uriPersistence, uriSourceAddr)
	if err != nil {
		return nil, err
	}

	return &persistenceProfiles, nil
}

// GetPersistenceSourceAddr gets a persistence profile by name. Returns nil if the persistence profile does not exist
func (b *BigIP) GetPersistenceSourceAddr(name string) (*PersistenceSourceAddr, error) {
	var persistenceProfile PersistenceSourceAddr
	err, ok := b.getForEntity(&persistenceProfile, uriLtm, uriPersistence, uriSourceAddr, name)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	return &persistenceProfile, nil
}

// CreatePersistenceSourceAddr creates a new persistence profile on the BIG-IP system.
func (b *BigIP) CreatePersistenceSourceAddr(name string, parent string) error {
	config := &PersistenceSourceAddr{
		Name:         name,
		DefaultsFrom: parent,
	}

	return b.post(config, uriLtm, uriPersistence, uriSourceAddr)
}

// AddPersistenceSourceAddr adds a new persistence profile on the BIG-IP system.
func (b *BigIP) AddPersistenceSourceAddr(config *PersistenceSourceAddr) error {
	return b.post(config, uriLtm, uriPersistence, uriSourceAddr)
}

// DeletePersistenceSourceAddr removes a persistence profile.
func (b *BigIP) DeletePersistenceSourceAddr(name string) error {
	return b.delete(uriLtm, uriPersistence, uriSourceAddr, name)
}

// ModifyPersistenceSourceAddr allows you to change any attribute of a persistence profile.
// Fields that can be modified are referenced in the PersistenceSourceAddr struct.
func (b *BigIP) ModifyPersistenceSourceAddr(name string, config *PersistenceSourceAddr) error {
	return b.put(config, uriLtm, uriPersistence, uriSourceAddr, name)
}

type SSLCertificateInfo struct {
	FileName    string `json:"file_name"`
	IsBundled   bool   `json:"is_bundled"`
	Certificate struct {
		CertInfo struct {
			ID    string      `json:"id"`
			Email interface{} `json:"email"`
		} `json:"cert_info"`
		ExpirationString string      `json:"expiration_string"`
		CertType         string      `json:"cert_type"`
		KeyType          string      `json:"key_type"`
		Version          int         `json:"version"`
		ExpirationDate   int         `json:"expiration_date"`
		SerialNumber     interface{} `json:"serial_number"`
		BitLength        int         `json:"bit_length"`
		Issuer           struct {
			DivisionName     string `json:"division_name"`
			StateName        string `json:"state_name"`
			LocalityName     string `json:"locality_name"`
			OrganizationName string `json:"organization_name"`
			CountryName      string `json:"country_name"`
			CommonName       string `json:"common_name"`
		} `json:"issuer"`
		Subject struct {
			DivisionName     string `json:"division_name"`
			StateName        string `json:"state_name"`
			LocalityName     string `json:"locality_name"`
			OrganizationName string `json:"organization_name"`
			CountryName      string `json:"country_name"`
			CommonName       string `json:"common_name"`
		} `json:"subject"`
	} `json:"certificate"`
}

// FTPProfile represents a FTP Profile configuration
type FTPProfile struct {
	Kind                 string `json:"kind"`
	Name                 string `json:"name"`
	Partition            string `json:"partition,omitempty"`
	FullPath             string `json:"fullPath,omitempty"`
	Generation           int    `json:"generation,omitempty"`
	SelfLink             string `json:"selfLink,omitempty"`
	InheritParentProfile string `json:"inheritParentProfile,omitempty"`
	Port                 int    `json:"port,omitempty"`
	Security             string `json:"security,omitempty"`
	TranslateExtended    string `json:"translateExtended,omitempty"`
	DefaultsFrom         string `json:defaultsFrom`
}

// FTPProfiles is an array of FTPProfile structs
type FTPProfiles []FTPProfile

// FTPProfiles returns a list of FTP profiles.
func (b *BigIP) FTPProfiles() (*FTPProfiles, error) {
	var serverFTPProfiles FTPProfiles
	err, _ := b.getForEntity(&serverFTPProfiles, uriLtm, uriProfile, uriFtp)
	if err != nil {
		return nil, err
	}

	return &serverFTPProfiles, nil
}

// GetFTPProfile gets a FTP profile by name. Returns nil if the FTP profile does not exist
func (b *BigIP) GetFTPProfile(name string) (*FTPProfile, error) {
	var serverFTPProfile FTPProfile
	err, ok := b.getForEntity(&serverFTPProfile, uriLtm, uriProfile, uriFtp, name)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	return &serverFTPProfile, nil
}

// CreateFTPProfile creates a new FTP profile on the BIG-IP system.
func (b *BigIP) CreateFTPProfile(name string) error {
	config := &FTPProfile{
		Name: name,
	}

	return b.post(config, uriLtm, uriProfile, uriFtp)
}

// AddFTPProfile adds a new FTP profile on the BIG-IP system.
func (b *BigIP) AddFTPProfile(config *FTPProfile) error {
	return b.post(config, uriLtm, uriProfile, uriFtp)
}

// DeleteFTPProfile removes a FTP profile.
func (b *BigIP) DeleteFTPProfile(name string) error {
	return b.delete(uriLtm, uriProfile, uriFtp, name)
}

// ModifyFTPProfile allows you to change any attribute of a FTP profile.
// Fields that can be modified are referenced in the VirtualServer struct.
func (b *BigIP) ModifyFTPProfile(name string, config *FTPProfile) error {
	return b.put(config, uriLtm, uriProfile, uriFtp, name)
}

// FastL4Profile is a representation of a fastL4Profile configuration
type FastL4Profile struct {
	Kind                    string `json:"kind,omitempty"`
	Name                    string `json:"name"`
	Partition               string `json:"partition,omitempty"`
	FullPath                string `json:"fullPath,omitempty"`
	Generation              int    `json:"generation,omitempty"`
	SelfLink                string `json:"selfLink,omitempty"`
	HardwareSynCookie       string `json:"hardwareSynCookie,omitempty"`
	IdleTimeout             string `json:"idleTimeout,omitempty"`
	IPTosToClient           string `json:"ipTosToClient,omitempty"`
	IPTosToServer           string `json:"ipTosToServer,omitempty"`
	KeepAliveInterval       string `json:"keepAliveInterval,omitempty"`
	LinkQosToClient         string `json:"linkQosToClient,omitempty"`
	LinkQosToServer         string `json:"linkQosToServer,omitempty"`
	LooseClose              string `json:"looseClose,omitempty"`
	LooseInitialization     string `json:"looseInitialization,omitempty"`
	MssOverride             int    `json:"mssOverride,omitempty"`
	PriorityToClient        string `json:"priorityToClient,omitempty"`
	PriorityToServer        string `json:"priorityToServer,omitempty"`
	PvaAcceleration         string `json:"pvaAcceleration,omitempty"`
	PvaDynamicClientPackets int    `json:"pvaDynamicClientPackets,omitempty"`
	PvaDynamicServerPackets int    `json:"pvaDynamicServerPackets,omitempty"`
	PvaFlowAging            string `json:"pvaFlowAging,omitempty"`
	PvaFlowEvict            string `json:"pvaFlowEvict,omitempty"`
	PvaOffloadDynamic       string `json:"pvaOffloadDynamic,omitempty"`
	PvaOffloadState         string `json:"pvaOffloadState,omitempty"`
	ReassembleFragments     string `json:"reassembleFragments,omitempty"`
	ReceiveWindowSize       int    `json:"receiveWindowSize,omitempty"`
	ResetOnTimeout          string `json:"resetOnTimeout,omitempty"`
	RttFromClient           string `json:"rttFromClient,omitempty"`
	RttFromServer           string `json:"rttFromServer,omitempty"`
	ServerSack              string `json:"serverSack,omitempty"`
	ServerTimestamp         string `json:"serverTimestamp,omitempty"`
	SoftwareSynCookie       string `json:"softwareSynCookie,omitempty"`
	TCPCloseTimeout         string `json:"tcpCloseTimeout,omitempty"`
	TCPGenerateIsn          string `json:"tcpGenerateIsn,omitempty"`
	TCPHandshakeTimeout     string `json:"tcpHandshakeTimeout,omitempty"`
	TCPStripSack            string `json:"tcpStripSack,omitempty"`
	TCPTimestampMode        string `json:"tcpTimestampMode,omitempty"`
	TCPWscaleMode           string `json:"tcpWscaleMode,omitempty"`
	DefaultsFrom            string `json:"defaultsFrom"`
}

// FastL4Profiles is an array of FastL4Profile structs
type FastL4Profiles []FastL4Profile

// FastL4Profiles returns a list of fastL4 profiles.
func (b *BigIP) FastL4Profiles() (*FastL4Profiles, error) {
	var serverFastL4Profiles FastL4Profiles
	err, _ := b.getForEntity(&serverFastL4Profiles, uriLtm, uriProfile, uriFastL4)
	if err != nil {
		return nil, err
	}

	return &serverFastL4Profiles, nil
}

// GetFastL4Profile gets a fastL4 profile by name. Returns nil if the fastL4 profile does not exist
func (b *BigIP) GetFastL4Profile(name string) (*FastL4Profile, error) {
	var serverFastL4Profile FastL4Profile
	err, ok := b.getForEntity(&serverFastL4Profile, uriLtm, uriProfile, uriFastL4, name)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	return &serverFastL4Profile, nil
}

// CreateFastL4Profile creates a new fastL4 profile on the BIG-IP system.
func (b *BigIP) CreateFastL4Profile(name string, parent string) error {
	config := &FastL4Profile{
		Name:         name,
		DefaultsFrom: parent,
	}

	return b.post(config, uriLtm, uriProfile, uriFastL4)
}

// AddFastL4Profile adds a new fastL4 profile on the BIG-IP system.
func (b *BigIP) AddFastL4Profile(config *FastL4Profile) error {
	return b.post(config, uriLtm, uriProfile, uriFastL4)
}

// DeleteFastL4Profile removes a fastL4 profile.
func (b *BigIP) DeleteFastL4Profile(name string) error {
	return b.delete(uriLtm, uriProfile, uriFastL4, name)
}

// ModifyFastL4Profile allows you to change any attribute of a fastL4 profile.
// Fields that can be modified are referenced in the VirtualServer struct.
func (b *BigIP) ModifyFastL4Profile(name string, config *FastL4Profile) error {
	return b.put(config, uriLtm, uriProfile, uriFastL4, name)
}

// OneConnectProfiles
// Documentation: https://devcentral.f5.com/wiki/iControlREST.APIRef_tm_ltm_profile_oneConnect.ashx

// OneConnectProfiles contains a list of every oneConnect profile on the BIG-IP system.
type OneConnectProfiles struct {
	OneConnectProfiles []OneConnectProfile `json:"items"`
}

// OneConnectProfile contains information about each oneConnect profile. You can use all
// of these fields when modifying a oneConnect profile.
type OneConnectProfile struct {
	Kind                string `json:"kind,omitempty"`
	Name                string `json:"name"`
	Partition           string `json:"partition,omitempty"`
	FullPath            string `json:"fullPath,omitempty"`
	Generation          int    `json:"generation,omitempty"`
	SelfLink            string `json:"selfLink,omitempty"`
	IdleTimeoutOverride string `json:"idleTimeoutOverride,omitempty"`
	MaxAge              int    `json:"maxAge,omitempty"`
	MaxReuse            int    `json:"maxReuse,omitempty"`
	MaxSize             int    `json:"maxSize,omitempty"`
	SharePools          string `json:"sharePools,omitempty"`
	SourceMask          string `json:"sourceMask,omitempty"`
	DefaultsFrom        string `json:"defaultsFrom"`
}

// OneConnectProfiles returns a list of oneConnect profiles.
func (b *BigIP) OneConnectProfiles() (*OneConnectProfiles, error) {
	var oneConnectProfiles OneConnectProfiles
	err, _ := b.getForEntity(&oneConnectProfiles, uriLtm, uriProfile, uriOneConnect)
	if err != nil {
		return nil, err
	}

	return &oneConnectProfiles, nil
}

// GetOneConnectProfile gets a oneConnect profile by name. Returns nil if the oneConnect profile does not exist
func (b *BigIP) GetOneConnectProfile(name string) (*OneConnectProfile, error) {
	var oneConnectProfile OneConnectProfile
	err, ok := b.getForEntity(&oneConnectProfile, uriLtm, uriProfile, uriOneConnect, name)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}

	return &oneConnectProfile, nil
}

// CreateOneConnectProfile creates a new oneConnect profile on the BIG-IP system.
func (b *BigIP) CreateOneConnectProfile(name string, parent string) error {
	config := &OneConnectProfile{
		Name:         name,
		DefaultsFrom: parent,
	}

	return b.post(config, uriLtm, uriProfile, uriOneConnect)
}

// AddOneConnectProfile adds a new oneConnect profile on the BIG-IP system.
func (b *BigIP) AddOneConnectProfile(config *OneConnectProfile) error {
	return b.post(config, uriLtm, uriProfile, uriOneConnect)
}

// DeleteOneConnectProfile removes a oneConnect profile.
func (b *BigIP) DeleteOneConnectProfile(name string) error {
	return b.delete(uriLtm, uriProfile, uriOneConnect, name)
}

// ModifyOneConnectProfile allows you to change any attribute of a sever-oneConnect profile.
// Fields that can be modified are referenced in the VirtualClient struct.
func (b *BigIP) ModifyOneConnectProfile(name string, config *OneConnectProfile) error {
	return b.put(config, uriLtm, uriProfile, uriOneConnect, name)
}
