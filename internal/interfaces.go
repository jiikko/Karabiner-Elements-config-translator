package internal

type ManipulatorSerializer interface {
	FromSerialize(ConfigRuleManipulator) (map[string]interface{}, error)
}
