package kubernetes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/portainer/agent"
	"github.com/portainer/agent/http/security"
	httperror "github.com/portainer/libhttp/error"
)

// Handler is the HTTP handler used to handle volume browsing operations.
type Handler struct {
	*mux.Router
	kubernetesDeployer agent.KubernetesDeployer
}

// NewHandler returns a pointer to an Handler
func NewHandler(notaryService *security.NotaryService, kubernetesDeployer agent.KubernetesDeployer) *Handler {
	h := &Handler{
		Router:             mux.NewRouter(),
		kubernetesDeployer: kubernetesDeployer,
	}

	h.Handle("/kubernetes/stack",
		notaryService.DigitalSignatureVerification(httperror.LoggerHandler(h.kubernetesDeploy))).Methods(http.MethodPost)

	return h
}
