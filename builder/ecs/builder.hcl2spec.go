// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package ecs

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType         *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion         *string           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug               *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce               *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError             *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars            map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars       []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	AccessKey                 *string           `mapstructure:"access_key" required:"true" cty:"access_key" hcl:"access_key"`
	SecretKey                 *string           `mapstructure:"secret_key" required:"true" cty:"secret_key" hcl:"secret_key"`
	Region                    *string           `mapstructure:"region" required:"true" cty:"region" hcl:"region"`
	ProjectName               *string           `mapstructure:"project_name" required:"false" cty:"project_name" hcl:"project_name"`
	ProjectID                 *string           `mapstructure:"project_id" required:"false" cty:"project_id" hcl:"project_id"`
	SecurityToken             *string           `mapstructure:"security_token" required:"false" cty:"security_token" hcl:"security_token"`
	IdentityEndpoint          *string           `mapstructure:"auth_url" required:"false" cty:"auth_url" hcl:"auth_url"`
	Insecure                  *bool             `mapstructure:"insecure" required:"false" cty:"insecure" hcl:"insecure"`
	ImageName                 *string           `mapstructure:"image_name" required:"true" cty:"image_name" hcl:"image_name"`
	ImageDescription          *string           `mapstructure:"image_description" required:"false" cty:"image_description" hcl:"image_description"`
	ImageType                 *string           `mapstructure:"image_type" required:"false" cty:"image_type" hcl:"image_type"`
	ImageTags                 map[string]string `mapstructure:"image_tags" required:"false" cty:"image_tags" hcl:"image_tags"`
	ImageMembers              []string          `mapstructure:"image_members" required:"false" cty:"image_members" hcl:"image_members"`
	ImageAutoAcceptMembers    *bool             `mapstructure:"image_auto_accept_members" required:"false" cty:"image_auto_accept_members" hcl:"image_auto_accept_members"`
	WaitImageReadyTimeout     *string           `mapstructure:"wait_image_ready_timeout" required:"false" cty:"wait_image_ready_timeout" hcl:"wait_image_ready_timeout"`
	Type                      *string           `mapstructure:"communicator" cty:"communicator" hcl:"communicator"`
	PauseBeforeConnect        *string           `mapstructure:"pause_before_connecting" cty:"pause_before_connecting" hcl:"pause_before_connecting"`
	SSHHost                   *string           `mapstructure:"ssh_host" cty:"ssh_host" hcl:"ssh_host"`
	SSHPort                   *int              `mapstructure:"ssh_port" cty:"ssh_port" hcl:"ssh_port"`
	SSHUsername               *string           `mapstructure:"ssh_username" cty:"ssh_username" hcl:"ssh_username"`
	SSHPassword               *string           `mapstructure:"ssh_password" cty:"ssh_password" hcl:"ssh_password"`
	SSHKeyPairName            *string           `mapstructure:"ssh_keypair_name" undocumented:"true" cty:"ssh_keypair_name" hcl:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string           `mapstructure:"temporary_key_pair_name" undocumented:"true" cty:"temporary_key_pair_name" hcl:"temporary_key_pair_name"`
	SSHTemporaryKeyPairType   *string           `mapstructure:"temporary_key_pair_type" cty:"temporary_key_pair_type" hcl:"temporary_key_pair_type"`
	SSHTemporaryKeyPairBits   *int              `mapstructure:"temporary_key_pair_bits" cty:"temporary_key_pair_bits" hcl:"temporary_key_pair_bits"`
	SSHCiphers                []string          `mapstructure:"ssh_ciphers" cty:"ssh_ciphers" hcl:"ssh_ciphers"`
	SSHClearAuthorizedKeys    *bool             `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys" hcl:"ssh_clear_authorized_keys"`
	SSHKEXAlgos               []string          `mapstructure:"ssh_key_exchange_algorithms" cty:"ssh_key_exchange_algorithms" hcl:"ssh_key_exchange_algorithms"`
	SSHPrivateKeyFile         *string           `mapstructure:"ssh_private_key_file" undocumented:"true" cty:"ssh_private_key_file" hcl:"ssh_private_key_file"`
	SSHCertificateFile        *string           `mapstructure:"ssh_certificate_file" cty:"ssh_certificate_file" hcl:"ssh_certificate_file"`
	SSHPty                    *bool             `mapstructure:"ssh_pty" cty:"ssh_pty" hcl:"ssh_pty"`
	SSHTimeout                *string           `mapstructure:"ssh_timeout" cty:"ssh_timeout" hcl:"ssh_timeout"`
	SSHWaitTimeout            *string           `mapstructure:"ssh_wait_timeout" undocumented:"true" cty:"ssh_wait_timeout" hcl:"ssh_wait_timeout"`
	SSHAgentAuth              *bool             `mapstructure:"ssh_agent_auth" undocumented:"true" cty:"ssh_agent_auth" hcl:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool             `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding" hcl:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int              `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts" hcl:"ssh_handshake_attempts"`
	SSHBastionHost            *string           `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host" hcl:"ssh_bastion_host"`
	SSHBastionPort            *int              `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port" hcl:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool             `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth" hcl:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string           `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username" hcl:"ssh_bastion_username"`
	SSHBastionPassword        *string           `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password" hcl:"ssh_bastion_password"`
	SSHBastionInteractive     *bool             `mapstructure:"ssh_bastion_interactive" cty:"ssh_bastion_interactive" hcl:"ssh_bastion_interactive"`
	SSHBastionPrivateKeyFile  *string           `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file" hcl:"ssh_bastion_private_key_file"`
	SSHBastionCertificateFile *string           `mapstructure:"ssh_bastion_certificate_file" cty:"ssh_bastion_certificate_file" hcl:"ssh_bastion_certificate_file"`
	SSHFileTransferMethod     *string           `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method" hcl:"ssh_file_transfer_method"`
	SSHProxyHost              *string           `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host" hcl:"ssh_proxy_host"`
	SSHProxyPort              *int              `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port" hcl:"ssh_proxy_port"`
	SSHProxyUsername          *string           `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username" hcl:"ssh_proxy_username"`
	SSHProxyPassword          *string           `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password" hcl:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string           `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval" hcl:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string           `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout" hcl:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string          `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels" hcl:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string          `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels" hcl:"ssh_local_tunnels"`
	SSHPublicKey              []byte            `mapstructure:"ssh_public_key" undocumented:"true" cty:"ssh_public_key" hcl:"ssh_public_key"`
	SSHPrivateKey             []byte            `mapstructure:"ssh_private_key" undocumented:"true" cty:"ssh_private_key" hcl:"ssh_private_key"`
	WinRMUser                 *string           `mapstructure:"winrm_username" cty:"winrm_username" hcl:"winrm_username"`
	WinRMPassword             *string           `mapstructure:"winrm_password" cty:"winrm_password" hcl:"winrm_password"`
	WinRMHost                 *string           `mapstructure:"winrm_host" cty:"winrm_host" hcl:"winrm_host"`
	WinRMNoProxy              *bool             `mapstructure:"winrm_no_proxy" cty:"winrm_no_proxy" hcl:"winrm_no_proxy"`
	WinRMPort                 *int              `mapstructure:"winrm_port" cty:"winrm_port" hcl:"winrm_port"`
	WinRMTimeout              *string           `mapstructure:"winrm_timeout" cty:"winrm_timeout" hcl:"winrm_timeout"`
	WinRMUseSSL               *bool             `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl" hcl:"winrm_use_ssl"`
	WinRMInsecure             *bool             `mapstructure:"winrm_insecure" cty:"winrm_insecure" hcl:"winrm_insecure"`
	WinRMUseNTLM              *bool             `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm" hcl:"winrm_use_ntlm"`
	Flavor                    *string           `mapstructure:"flavor" required:"true" cty:"flavor" hcl:"flavor"`
	EnterpriseProjectId       *string           `mapstructure:"enterprise_project_id" required:"false" cty:"enterprise_project_id" hcl:"enterprise_project_id"`
	AvailabilityZone          *string           `mapstructure:"availability_zone" required:"false" cty:"availability_zone" hcl:"availability_zone"`
	SourceImage               *string           `mapstructure:"source_image" required:"false" cty:"source_image" hcl:"source_image"`
	SourceImageName           *string           `mapstructure:"source_image_name" required:"false" cty:"source_image_name" hcl:"source_image_name"`
	SourceImageFilters        *FlatImageFilter  `mapstructure:"source_image_filter" required:"false" cty:"source_image_filter" hcl:"source_image_filter"`
	FloatingIP                *string           `mapstructure:"floating_ip" required:"false" cty:"floating_ip" hcl:"floating_ip"`
	ReuseIPs                  *bool             `mapstructure:"reuse_ips" required:"false" cty:"reuse_ips" hcl:"reuse_ips"`
	EIPType                   *string           `mapstructure:"eip_type" required:"false" cty:"eip_type" hcl:"eip_type"`
	EIPBandwidthSize          *int              `mapstructure:"eip_bandwidth_size" required:"false" cty:"eip_bandwidth_size" hcl:"eip_bandwidth_size"`
	SSHIPVersion              *string           `mapstructure:"ssh_ip_version" required:"false" cty:"ssh_ip_version" hcl:"ssh_ip_version"`
	VpcID                     *string           `mapstructure:"vpc_id" required:"false" cty:"vpc_id" hcl:"vpc_id"`
	Subnets                   []string          `mapstructure:"subnets" required:"false" cty:"subnets" hcl:"subnets"`
	SecurityGroups            []string          `mapstructure:"security_groups" required:"false" cty:"security_groups" hcl:"security_groups"`
	UserData                  *string           `mapstructure:"user_data" required:"false" cty:"user_data" hcl:"user_data"`
	UserDataFile              *string           `mapstructure:"user_data_file" required:"false" cty:"user_data_file" hcl:"user_data_file"`
	InstanceName              *string           `mapstructure:"instance_name" required:"false" cty:"instance_name" hcl:"instance_name"`
	InstanceMetadata          map[string]string `mapstructure:"instance_metadata" required:"false" cty:"instance_metadata" hcl:"instance_metadata"`
	ServerTags                map[string]string `mapstructure:"server_tags" required:"false" cty:"server_tags" hcl:"server_tags"`
	SpotPricing               *bool             `mapstructure:"spot_pricing" required:"false" cty:"spot_pricing" hcl:"spot_pricing"`
	SpotMaximumPrice          *string           `mapstructure:"spot_maximum_price" required:"false" cty:"spot_maximum_price" hcl:"spot_maximum_price"`
	VolumeType                *string           `mapstructure:"volume_type" required:"false" cty:"volume_type" hcl:"volume_type"`
	VolumeSize                *int              `mapstructure:"volume_size" required:"false" cty:"volume_size" hcl:"volume_size"`
	KmsKeyID                  *string           `mapstructure:"kms_key_id" required:"false" cty:"kms_key_id" hcl:"kms_key_id"`
	DataVolumes               []FlatDataVolume  `mapstructure:"data_disks" required:"false" cty:"data_disks" hcl:"data_disks"`
	Vault                     *string           `mapstructure:"vault_id" required:"false" cty:"vault_id" hcl:"vault_id"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":            &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":          &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":          &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"access_key":                   &hcldec.AttrSpec{Name: "access_key", Type: cty.String, Required: false},
		"secret_key":                   &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"region":                       &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"project_name":                 &hcldec.AttrSpec{Name: "project_name", Type: cty.String, Required: false},
		"project_id":                   &hcldec.AttrSpec{Name: "project_id", Type: cty.String, Required: false},
		"security_token":               &hcldec.AttrSpec{Name: "security_token", Type: cty.String, Required: false},
		"auth_url":                     &hcldec.AttrSpec{Name: "auth_url", Type: cty.String, Required: false},
		"insecure":                     &hcldec.AttrSpec{Name: "insecure", Type: cty.Bool, Required: false},
		"image_name":                   &hcldec.AttrSpec{Name: "image_name", Type: cty.String, Required: false},
		"image_description":            &hcldec.AttrSpec{Name: "image_description", Type: cty.String, Required: false},
		"image_type":                   &hcldec.AttrSpec{Name: "image_type", Type: cty.String, Required: false},
		"image_tags":                   &hcldec.AttrSpec{Name: "image_tags", Type: cty.Map(cty.String), Required: false},
		"image_members":                &hcldec.AttrSpec{Name: "image_members", Type: cty.List(cty.String), Required: false},
		"image_auto_accept_members":    &hcldec.AttrSpec{Name: "image_auto_accept_members", Type: cty.Bool, Required: false},
		"wait_image_ready_timeout":     &hcldec.AttrSpec{Name: "wait_image_ready_timeout", Type: cty.String, Required: false},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"temporary_key_pair_type":      &hcldec.AttrSpec{Name: "temporary_key_pair_type", Type: cty.String, Required: false},
		"temporary_key_pair_bits":      &hcldec.AttrSpec{Name: "temporary_key_pair_bits", Type: cty.Number, Required: false},
		"ssh_ciphers":                  &hcldec.AttrSpec{Name: "ssh_ciphers", Type: cty.List(cty.String), Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_key_exchange_algorithms":  &hcldec.AttrSpec{Name: "ssh_key_exchange_algorithms", Type: cty.List(cty.String), Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_certificate_file":         &hcldec.AttrSpec{Name: "ssh_certificate_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_wait_timeout":             &hcldec.AttrSpec{Name: "ssh_wait_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":               &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":       &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":             &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":             &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":         &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":         &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_interactive":      &hcldec.AttrSpec{Name: "ssh_bastion_interactive", Type: cty.Bool, Required: false},
		"ssh_bastion_private_key_file": &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_bastion_certificate_file": &hcldec.AttrSpec{Name: "ssh_bastion_certificate_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":               &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":               &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":           &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":           &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":           &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":            &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":               &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":              &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":               &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":               &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                   &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_no_proxy":               &hcldec.AttrSpec{Name: "winrm_no_proxy", Type: cty.Bool, Required: false},
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"flavor":                       &hcldec.AttrSpec{Name: "flavor", Type: cty.String, Required: false},
		"enterprise_project_id":        &hcldec.AttrSpec{Name: "enterprise_project_id", Type: cty.String, Required: false},
		"availability_zone":            &hcldec.AttrSpec{Name: "availability_zone", Type: cty.String, Required: false},
		"source_image":                 &hcldec.AttrSpec{Name: "source_image", Type: cty.String, Required: false},
		"source_image_name":            &hcldec.AttrSpec{Name: "source_image_name", Type: cty.String, Required: false},
		"source_image_filter":          &hcldec.BlockSpec{TypeName: "source_image_filter", Nested: hcldec.ObjectSpec((*FlatImageFilter)(nil).HCL2Spec())},
		"floating_ip":                  &hcldec.AttrSpec{Name: "floating_ip", Type: cty.String, Required: false},
		"reuse_ips":                    &hcldec.AttrSpec{Name: "reuse_ips", Type: cty.Bool, Required: false},
		"eip_type":                     &hcldec.AttrSpec{Name: "eip_type", Type: cty.String, Required: false},
		"eip_bandwidth_size":           &hcldec.AttrSpec{Name: "eip_bandwidth_size", Type: cty.Number, Required: false},
		"ssh_ip_version":               &hcldec.AttrSpec{Name: "ssh_ip_version", Type: cty.String, Required: false},
		"vpc_id":                       &hcldec.AttrSpec{Name: "vpc_id", Type: cty.String, Required: false},
		"subnets":                      &hcldec.AttrSpec{Name: "subnets", Type: cty.List(cty.String), Required: false},
		"security_groups":              &hcldec.AttrSpec{Name: "security_groups", Type: cty.List(cty.String), Required: false},
		"user_data":                    &hcldec.AttrSpec{Name: "user_data", Type: cty.String, Required: false},
		"user_data_file":               &hcldec.AttrSpec{Name: "user_data_file", Type: cty.String, Required: false},
		"instance_name":                &hcldec.AttrSpec{Name: "instance_name", Type: cty.String, Required: false},
		"instance_metadata":            &hcldec.AttrSpec{Name: "instance_metadata", Type: cty.Map(cty.String), Required: false},
		"server_tags":                  &hcldec.AttrSpec{Name: "server_tags", Type: cty.Map(cty.String), Required: false},
		"spot_pricing":                 &hcldec.AttrSpec{Name: "spot_pricing", Type: cty.Bool, Required: false},
		"spot_maximum_price":           &hcldec.AttrSpec{Name: "spot_maximum_price", Type: cty.String, Required: false},
		"volume_type":                  &hcldec.AttrSpec{Name: "volume_type", Type: cty.String, Required: false},
		"volume_size":                  &hcldec.AttrSpec{Name: "volume_size", Type: cty.Number, Required: false},
		"kms_key_id":                   &hcldec.AttrSpec{Name: "kms_key_id", Type: cty.String, Required: false},
		"data_disks":                   &hcldec.BlockListSpec{TypeName: "data_disks", Nested: hcldec.ObjectSpec((*FlatDataVolume)(nil).HCL2Spec())},
		"vault_id":                     &hcldec.AttrSpec{Name: "vault_id", Type: cty.String, Required: false},
	}
	return s
}
