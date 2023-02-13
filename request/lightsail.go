package request

type GetRegions struct {
	SecretName string `form:"secretName" binding:"required"`
}

type CreateLightsail struct {
	Name             string `form:"name" binding:"required"`
	Zone             string `form:"zone" binding:"required"`
	AvailabilityZone string `form:"availabilityZone" binding:"required"`
	BlueprintId      string `form:"blueprintId" binding:"required"`
	BundleId         string `form:"bundleId" binding:"required"`
	SecretName       string `form:"secretName" binding:"required"`
}

type OpenLightsailPorts struct {
	Zone       string `form:"zone" binding:"required"`
	SecretName string `form:"secretName" binding:"required"`
	Name       string `form:"name" binding:"required"`
}

type LightsailAction struct {
	Zone       string `form:"zone" binding:"required"`
	SecretName string `form:"secretName" binding:"required"`
	Name       string `form:"name" binding:"required"`
}

type DeleteLightsail struct {
	Zone         string `form:"zone" binding:"required"`
	SecretName   string `form:"secretName" binding:"required"`
	Name         string `form:"name" binding:"required"`
	ResourceName string `form:"resourceName" binding:"required"`
}
