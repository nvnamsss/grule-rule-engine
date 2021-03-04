package benchmark

import (
	"io/ioutil"
	"testing"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

func Benchmark_Grule_KnowledgeBase_Clone(b *testing.B) {
	input, _ := ioutil.ReadFile("1000_rules.grl")
	rules := string(input)
	fact := &RideFact{
		Distance: 6000,
		Duration: 121,
	}
	dctx := ast.NewDataContext()
	_ = dctx.Add("Fact", fact)

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	_ = rb.BuildRuleFromResource("load_rules_test", "0.1.1", pkg.NewBytesResource([]byte(rules)))
	_ = lib.NewKnowledgeBaseInstance("load_rules_test", "0.1.1")
	for k := 0; k < b.N; k++ {
		_ = lib.NewKnowledgeBaseInstance("load_rules_test", "0.1.1")
	}

}
