package services

type Hba1cBuilder struct{}

func (b *Hba1cBuilder) Build(params map[string]string) map[string]interface{} {
	return map[string]interface{}{
		"uid":    "web-client",
		"age":    params["age"],
		"gender": params["gender"],
		"rdw":    params["rdw"],
		"wbc":    params["wbc"],
		"rbc":    params["rbc"],
		"hgb":    params["hgb"],
		"hct":    params["hct"],
		"mcv":    params["mcv"],
		"mch":    params["mch"],
		"mchc":   params["mchc"],
		"plt":    params["plt"],
		"neu":    params["neu"],
		"eos":    params["eos"],
		"bas":    params["bas"],
		"lym":    params["lym"],
		"mon":    params["mon"],
		"soe":    params["soe"],
		"chol":   params["chol"],
		"glu":    params["glu"],
	}
}

type LdllBuilder struct{}

func (b *LdllBuilder) Build(params map[string]string) map[string]interface{} {
	return map[string]interface{}{
		"uid":    "web-client",
		"age":    params["age"],
		"gender": params["gender"],
		"chol":   params["chol"],
		"hdl":    params["hdl"],
		"tg":     params["tg"],
	}
}

type FerrBuilder struct{}

func (b *FerrBuilder) Build(params map[string]string) map[string]interface{} {
	return map[string]interface{}{
		"uid":    "web-client",
		"age":    params["age"],
		"gender": params["gender"],
		"rdw":    params["rdw"],
		"wbc":    params["wbc"],
		"rbc":    params["rbc"],
		"hgb":    params["hgb"],
		"hct":    params["hct"],
		"mcv":    params["mcv"],
		"mch":    params["mch"],
		"mchc":   params["mchc"],
		"plt":    params["plt"],
		"neu":    params["neu"],
		"eos":    params["eos"],
		"bas":    params["bas"],
		"lym":    params["lym"],
		"mon":    params["mon"],
		"soe":    params["soe"],
		"crp":    params["crp"],
	}
}

type LdlBuilder struct{}

func (b *LdlBuilder) Build(params map[string]string) map[string]interface{} {
	return map[string]interface{}{
		"uid":    "web-client",
		"age":    params["age"],
		"gender": params["gender"],
		"rdw":    params["rdw"],
		"wbc":    params["wbc"],
		"rbc":    params["rbc"],
		"hgb":    params["hgb"],
		"hct":    params["hct"],
		"mcv":    params["mcv"],
		"mch":    params["mch"],
		"mchc":   params["mchc"],
		"plt":    params["plt"],
		"neu":    params["neu"],
		"eos":    params["eos"],
		"bas":    params["bas"],
		"lym":    params["lym"],
		"mon":    params["mon"],
		"soe":    params["soe"],
		"chol":   params["chol"],
		"glu":    params["glu"],
	}
}

type TgBuilder struct{}

func (b *TgBuilder) Build(params map[string]string) map[string]interface{} {
	return map[string]interface{}{
		"uid":    "web-client",
		"age":    params["age"],
		"gender": params["gender"],
		"rdw":    params["rdw"],
		"wbc":    params["wbc"],
		"rbc":    params["rbc"],
		"hgb":    params["hgb"],
		"hct":    params["hct"],
		"mcv":    params["mcv"],
		"mch":    params["mch"],
		"mchc":   params["mchc"],
		"plt":    params["plt"],
		"neu":    params["neu"],
		"eos":    params["eos"],
		"bas":    params["bas"],
		"lym":    params["lym"],
		"mon":    params["mon"],
		"soe":    params["soe"],
		"chol":   params["chol"],
		"glu":    params["glu"],
	}
}

type HdlBuilder struct{}

func (b *HdlBuilder) Build(params map[string]string) map[string]interface{} {
	return map[string]interface{}{
		"uid":    "web-client",
		"age":    params["age"],
		"gender": params["gender"],
		"rdw":    params["rdw"],
		"wbc":    params["wbc"],
		"rbc":    params["rbc"],
		"hgb":    params["hgb"],
		"hct":    params["hct"],
		"mcv":    params["mcv"],
		"mch":    params["mch"],
		"mchc":   params["mchc"],
		"plt":    params["plt"],
		"neu":    params["neu"],
		"eos":    params["eos"],
		"bas":    params["bas"],
		"lym":    params["lym"],
		"mon":    params["mon"],
		"soe":    params["soe"],
		"chol":   params["chol"],
		"glu":    params["glu"],
	}
}
