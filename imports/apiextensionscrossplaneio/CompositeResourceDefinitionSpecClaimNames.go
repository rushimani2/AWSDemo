package apiextensionscrossplaneio


// ClaimNames specifies the names of an optional composite resource claim.
//
// When claim names are specified Crossplane will create a namespaced 'composite resource claim' CRD that corresponds to the defined composite resource. This composite resource claim acts as a namespaced proxy for the composite resource; creating, updating, or deleting the claim will create, update, or delete a corresponding composite resource. You may add claim names to an existing CompositeResourceDefinition, but they cannot be changed or removed once they have been set.
type CompositeResourceDefinitionSpecClaimNames struct {
	// kind is the serialized kind of the resource.
	//
	// It is normally CamelCase and singular. Custom resource instances will use this value as the `kind` attribute in API calls.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// plural is the plural name of the resource to serve.
	//
	// The custom resources are served under `/apis/<group>/<version>/.../<plural>`. Must match the name of the CustomResourceDefinition (in the form `<names.plural>.<group>`). Must be all lowercase.
	Plural *string `field:"required" json:"plural" yaml:"plural"`
	// categories is a list of grouped resources this custom resource belongs to (e.g. 'all'). This is published in API discovery documents, and used by clients to support invocations like `kubectl get all`.
	Categories *[]*string `field:"optional" json:"categories" yaml:"categories"`
	// listKind is the serialized kind of the list for this resource.
	//
	// Defaults to "`kind`List".
	// Default: kind`List".
	//
	ListKind *string `field:"optional" json:"listKind" yaml:"listKind"`
	// shortNames are short names for the resource, exposed in API discovery documents, and used by clients to support invocations like `kubectl get <shortname>`.
	//
	// It must be all lowercase.
	ShortNames *[]*string `field:"optional" json:"shortNames" yaml:"shortNames"`
	// singular is the singular name of the resource.
	//
	// It must be all lowercase. Defaults to lowercased `kind`.
	// Default: lowercased `kind`.
	//
	Singular *string `field:"optional" json:"singular" yaml:"singular"`
}

