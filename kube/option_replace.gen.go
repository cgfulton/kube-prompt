// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var replaceOptions = []prompt.Suggest{
	prompt.Suggest{Text: "--cascade", Description: "Only relevant during a force replace. If true, cascade the deletion of the resources managed by this resource (e.g. Pods created by a ReplicationController)."},
	prompt.Suggest{Text: "-f", Description: "Filename, directory, or URL to files to use to replace the resource."},
	prompt.Suggest{Text: "--filename", Description: "Filename, directory, or URL to files to use to replace the resource."},
	prompt.Suggest{Text: "--force", Description: "Delete and re-create the specified resource"},
	prompt.Suggest{Text: "--grace-period", Description: "Only relevant during a force replace. Period of time in seconds given to the old resource to terminate gracefully. Ignored if negative."},
	prompt.Suggest{Text: "--include-extended-apis", Description: "If true, include definitions of new APIs via calls to the API server. [default true]"},
	prompt.Suggest{Text: "-o", Description: "Output mode. Use \"-o name\" for shorter output (resource/name)."},
	prompt.Suggest{Text: "--output", Description: "Output mode. Use \"-o name\" for shorter output (resource/name)."},
	prompt.Suggest{Text: "--record", Description: "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	prompt.Suggest{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	prompt.Suggest{Text: "--save-config", Description: "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future."},
	prompt.Suggest{Text: "--timeout", Description: "Only relevant during a force replace. The length of time to wait before giving up on a delete of the old resource, zero means determine a timeout from the size of the object. Any other values should contain a corresponding time unit (e.g. 1s, 2m, 3h)."},
	prompt.Suggest{Text: "--validate", Description: "If true, use a schema to validate the input before sending it"},
}
