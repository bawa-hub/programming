package validation


type ValidationChainBuilder struct {
    rules []ValidationRule
}

func NewValidationChainBuilder() *ValidationChainBuilder {
    return &ValidationChainBuilder{
        rules: make([]ValidationRule, 0),
    }
}

func (b *ValidationChainBuilder) Add(rule ValidationRule) *ValidationChainBuilder {
    b.rules = append(b.rules, rule)
    return b
}

func (b *ValidationChainBuilder) Build() ValidationRule {
    if len(b.rules) == 0 {
        return nil
    }
    for i := 0; i < len(b.rules)-1; i++ {
        b.rules[i].SetNext(b.rules[i+1])
    }
    return b.rules[0]
}


