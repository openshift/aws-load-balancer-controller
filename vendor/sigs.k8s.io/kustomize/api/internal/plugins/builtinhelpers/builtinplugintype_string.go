// Code generated by "stringer -type=BuiltinPluginType"; DO NOT EDIT.

package builtinhelpers

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[AnnotationsTransformer-1]
	_ = x[ConfigMapGenerator-2]
	_ = x[IAMPolicyGenerator-3]
	_ = x[HashTransformer-4]
	_ = x[ImageTagTransformer-5]
	_ = x[LabelTransformer-6]
	_ = x[NamespaceTransformer-7]
	_ = x[PatchJson6902Transformer-8]
	_ = x[PatchStrategicMergeTransformer-9]
	_ = x[PatchTransformer-10]
	_ = x[PrefixSuffixTransformer-11]
	_ = x[PrefixTransformer-12]
	_ = x[SuffixTransformer-13]
	_ = x[ReplicaCountTransformer-14]
	_ = x[SecretGenerator-15]
	_ = x[ValueAddTransformer-16]
	_ = x[HelmChartInflationGenerator-17]
	_ = x[ReplacementTransformer-18]
}

const _BuiltinPluginType_name = "UnknownAnnotationsTransformerConfigMapGeneratorIAMPolicyGeneratorHashTransformerImageTagTransformerLabelTransformerNamespaceTransformerPatchJson6902TransformerPatchStrategicMergeTransformerPatchTransformerPrefixSuffixTransformerPrefixTransformerSuffixTransformerReplicaCountTransformerSecretGeneratorValueAddTransformerHelmChartInflationGeneratorReplacementTransformer"

var _BuiltinPluginType_index = [...]uint16{0, 7, 29, 47, 65, 80, 99, 115, 135, 159, 189, 205, 228, 245, 262, 285, 300, 319, 346, 368}

func (i BuiltinPluginType) String() string {
	if i < 0 || i >= BuiltinPluginType(len(_BuiltinPluginType_index)-1) {
		return "BuiltinPluginType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BuiltinPluginType_name[_BuiltinPluginType_index[i]:_BuiltinPluginType_index[i+1]]
}