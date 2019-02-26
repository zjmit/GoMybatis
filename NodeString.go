package GoMybatis

//字符串节点
type NodeString struct {
	value string
	t     NodeType

	//args
	expressMap          map[string]int //express表 key：name
	noConvertExpressMap map[string]int

	holder *NodeConfigHolder
}

func (it *NodeString) Type() NodeType {
	return NString
}

func (it *NodeString) Eval(env map[string]interface{}) ([]byte, error) {
	var data = it.value
	var err error
	if it.expressMap != nil {
		data, err = Replace(`#{`, it.expressMap, data, it.holder.GetSqlArgTypeConvert(), env, it.holder.GetExpressionEngineProxy())
		if err != nil {
			return nil, err
		}
	}
	if it.noConvertExpressMap != nil {
		data, err = Replace(`${`, it.noConvertExpressMap, data, it.holder.GetSqlArgTypeConvert(), env, it.holder.GetExpressionEngineProxy())
		if err != nil {
			return nil, err
		}
	}
	return []byte(data), nil
}