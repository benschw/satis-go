package db

type SatisDb struct {
	Abandoned              map[string]string `json:"abandoned,omitempty"`
	Archive                SatisArchive      `json:"archive,omitempty"`
	Comment                string            `json:"_comment,omitempty"`
	Config                 SatisConfig       `json:"config,omitempty"`
	Description            string            `json:"description,omitempty"`
	Homepage               string            `json:"homepage"`
	IncludeFilename        string            `json:"include-filename,omitempty"`
	MinimumStability       string            `json:"minimum-stability,omitempty"`
	Name                   string            `json:"name"`
	NotifyBatch            string            `json:"notify-batch,omitempty"`
	OutputDir              string            `json:"output-dir,omitempty"`
	OutputHTML             bool              `json:"output-html,omitempty"`
	Providers              bool              `json:"providers,omitempty"`
	Repositories           []SatisRepository `json:"repositories,omitempty"`
	Require                map[string]string `json:"require,omitempty"`
	RequireAll             bool              `json:"require-all,omitempty"`
	RequireDependencies    bool              `json:"require-dependencies,omitempty"`
	RequireDevDependencies bool              `json:"require-dev-dependencies,omitempty"`
	TwigTemplate           string            `json:"twig-template,omitempty"`
}

type SatisRepository struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type SatisArchive struct {
	AbsoluteDirectory string   `json:"absolute-directory,omitempty"`
	Blacklist         []string `json:"blacklist,omitempty"`
	Checksum          bool     `json:"checksum,omitempty"`
	Directory         string   `json:"directory"`
	Format            string   `json:"format,omitempty"`
	IgnoreFilters     bool     `json:"ignore-filters,omitempty"`
	PrefixURL         string   `json:"prefix-url,omitempty"`
	SkipDev           bool     `json:"skip-dev,omitempty"`
	Whitelist         []string `json:"whitelist,omitempty"`
}

type SatisConfig struct {
	AutoloaderSuffix      string                 `json:"autoloader-suffix,omitempty"`
	BinCompat             string                 `json:"bin-compat,omitempty"`
	BinDir                string                 `json:"bin-dir,omitempty"`
	CacheDir              string                 `json:"cache-dir,omitempty"`
	CacheFilesDir         string                 `json:"cache-files-dir,omitempty"`
	CacheFilesMaxsize     string                 `json:"cache-files-maxsize,omitempty"`
	CacheFilesTTL         int                    `json:"cache-files-ttl,omitempty"`
	CacheRepoDir          string                 `json:"cache-repo-dir,omitempty"`
	CacheTTL              int                    `json:"cache-ttl,omitempty"`
	CacheVcsDir           string                 `json:"cache-vcs-dir,omitempty"`
	ClassmapAuthoritative bool                   `json:"classmap-authoritative,omitempty"`
	DiscardChanges        bool                   `json:"discard-changes,omitempty"`
	GithubDomains         []string               `json:"github-domains,omitempty"`
	GithubExposeHostname  bool                   `json:"github-expose-hostname,omitempty"`
	GithubOauth           map[string]interface{} `json:"github-oauth,omitempty"`
	GithubProtocols       []string               `json:"github-protocols,omitempty"`
	HTTPBasic             map[string]interface{} `json:"http-basic,omitempty"`
	NotifyOnInstall       bool                   `json:"notify-on-install,omitempty"`
	OptimizeAutoloader    bool                   `json:"optimize-autoloader,omitempty"`
	Platform              map[string]interface{} `json:"platform,omitempty"`
	PreferredInstall      string                 `json:"preferred-install,omitempty"`
	PrependAutoloader     bool                   `json:"prepend-autoloader,omitempty"`
	ProcessTimeout        int                    `json:"process-timeout,omitempty"`
	StoreAuths            bool                   `json:"store-auths,omitempty"`
	UseIncludePath        bool                   `json:"use-include-path,omitempty"`
	VendorDir             string                 `json:"vendor-dir,omitempty"`
}
