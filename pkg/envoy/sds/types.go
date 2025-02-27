// Package sds implements Envoy's Secret Discovery Service (SDS).
package sds

import (
	"github.com/openservicemesh/osm/pkg/catalog"
	"github.com/openservicemesh/osm/pkg/certificate"
	"github.com/openservicemesh/osm/pkg/configurator"

	"github.com/openservicemesh/osm/pkg/identity"
	"github.com/openservicemesh/osm/pkg/logger"
)

var (
	log = logger.New("envoy/sds")
)

// sdsImpl is the type that implements the internal functionality of SDS
type sdsImpl struct {
	svcAccount  identity.K8sServiceAccount
	meshCatalog catalog.MeshCataloger
	cfg         configurator.Configurator
	certManager certificate.Manager
}
