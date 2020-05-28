package apis

import (
	imagev1 "github.com/adibrastegarnia/sample-k8s-controller/pkg/apis/v1"
)

func init() {
	// register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, imagev1.SchemeBuilder.AddToScheme)

}
