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
	err, _ := b.getForEntity(&persistenceProfiles, uriLtm, uriPersistences)
	if err != nil {
		return nil, err
	}

	return &persistenceProfiles, nil
}

// GetPersistenceCookie gets a persistence profile by name. Returns nil if the persistence profile does not exist
func (b *BigIP) GetPersistenceCookie(name string) (*PersistenceCookie, error) {
	var persistenceProfile PersistenceCookie
	err, ok := b.getForEntity(&persistenceProfile, uriLtm, uriProfile, uriPersistenceCookie, name)
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

	return b.post(config, uriLtm, uriProfile, uriPersistenceCookie)
}

// AddPersistenceCookie adds a new persistence profile on the BIG-IP system.
func (b *BigIP) AddPersistenceCookie(config *PersistenceCookie) error {
	return b.post(config, uriLtm, uriProfile, uriPersistenceCookie)
}

// DeletePersistenceCookie removes a persistence profile.
func (b *BigIP) DeletePersistenceCookie(name string) error {
	return b.delete(uriLtm, uriProfile, uriPersistenceCookie, name)
}

// ModifyPersistenceCookie allows you to change any attribute of a persistence profile.
// Fields that can be modified are referenced in the PersistenceCookie struct.
func (b *BigIP) ModifyPersistenceCookie(name string, config *PersistenceCookie) error {
	return b.put(config, uriLtm, uriProfile, uriPersistenceCookie, name)
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

// PersistenceHash returns a list of oersistence profiles.
func (b *BigIP) PersistenceHash() (*PersistenceHash, error) {
	var persistenceProfiles PersistenceHash
	err, _ := b.getForEntity(&persistenceProfiles, uriLtm, uriPersistences)
	if err != nil {
		return nil, err
	}

	return &persistenceProfiles, nil
}

// GetPersistenceHash gets a persistence profile by name. Returns nil if the persistence profile does not exist
func (b *BigIP) GetPersistenceHash(name string) (*PersistenceHash, error) {
	var persistenceProfile PersistenceHash
	err, ok := b.getForEntity(&persistenceProfile, uriLtm, uriProfile, uriPersistenceHash, name)
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

	return b.post(config, uriLtm, uriProfile, uriPersistenceHash)
}

// AddPersistenceHash adds a new persistence profile on the BIG-IP system.
func (b *BigIP) AddPersistenceHash(config *PersistenceHash) error {
	return b.post(config, uriLtm, uriProfile, uriPersistenceHash)
}

// DeletePersistenceHash removes a persistence profile.
func (b *BigIP) DeletePersistenceHash(name string) error {
	return b.delete(uriLtm, uriProfile, uriPersistenceHash, name)
}

// ModifyPersistenceHash allows you to change any attribute of a persistence profile.
// Fields that can be modified are referenced in the PersistenceHash struct.
func (b *BigIP) ModifyPersistenceHash(name string, config *PersistenceHash) error {
	return b.put(config, uriLtm, uriProfile, uriPersistenceHash, name)
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

// PersistenceSourceAddr returns a list of oersistence profiles.
func (b *BigIP) PersistenceSourceAddr() (*PersistenceSourceAddr, error) {
	var persistenceProfiles PersistenceSourceAddr
	err, _ := b.getForEntity(&persistenceProfiles, uriLtm, uriPersistences)
	if err != nil {
		return nil, err
	}

	return &persistenceProfiles, nil
}

// GetPersistenceSourceAddr gets a persistence profile by name. Returns nil if the persistence profile does not exist
func (b *BigIP) GetPersistenceSourceAddr(name string) (*PersistenceSourceAddr, error) {
	var persistenceProfile PersistenceSourceAddr
	err, ok := b.getForEntity(&persistenceProfile, uriLtm, uriProfile, uriPersistenceSourceAddr, name)
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

	return b.post(config, uriLtm, uriProfile, uriPersistenceSourceAddr)
}

// AddPersistenceSourceAddr adds a new persistence profile on the BIG-IP system.
func (b *BigIP) AddPersistenceSourceAddr(config *PersistenceSourceAddr) error {
	return b.post(config, uriLtm, uriProfile, uriPersistenceSourceAddr)
}

// DeletePersistenceSourceAddr removes a persistence profile.
func (b *BigIP) DeletePersistenceSourceAddr(name string) error {
	return b.delete(uriLtm, uriProfile, uriPersistenceSourceAddr, name)
}

// ModifyPersistenceSourceAddr allows you to change any attribute of a persistence profile.
// Fields that can be modified are referenced in the PersistenceSourceAddr struct.
func (b *BigIP) ModifyPersistenceSourceAddr(name string, config *PersistenceSourceAddr) error {
	return b.put(config, uriLtm, uriProfile, uriPersistenceSourceAddr, name)
}
