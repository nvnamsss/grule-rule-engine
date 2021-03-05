package examples

import (
	"io/ioutil"
	"testing"

	"github.com/nvnamsss/grule-rule-engine/ast"
	"github.com/nvnamsss/grule-rule-engine/builder"
	"github.com/nvnamsss/grule-rule-engine/engine"
	"github.com/nvnamsss/grule-rule-engine/pkg"
)

const ruleString = `
rule  DuplicateRule1  "Duplicate Rule 1"  salience 5 {
when
(Fact.Distance > 5000  &&   Fact.Duration > 120) && (Fact.Result == false)
Then
   Fact.Howl("1");
   Fact.NetAmount=143.320007;
   Fact.Result=true;
}

rule  DuplicateRule2  "Duplicate Rule 2"  salience 6 {
when
(Fact.Distance > 5000  &&   Fact.Duration > 120) && (Fact.Result == false)
Then
   Fact.NetAmount=143.320007;
   Fact.Result=true;
}


rule  DuplicateRule3  "Duplicate Rule 3"  salience 7 {
when
(Fact.Distance > 5000  &&   Fact.Duration > 120) && (Fact.Result == false)
Then
   Fact.Howl("3");
   Fact.NetAmount=143.320007;
   Fact.Result=true;
}


rule  DuplicateRule4  "Duplicate Rule 4"  salience 8 {
when
(Fact.Distance > 5000  &&   Fact.Duration > 120) && (Fact.Result == false)
Then
   Fact.NetAmount=143.320007;
   Fact.Result=true;
}


rule  DuplicateRule5  "Duplicate Rule 5"  salience 9 {
when
(Fact.A > 5000  &&   Fact.B == 120) && (Fact.Result == false)
Then
   Output.NetAmount=143.320007;
   Fact.Result=true;
}`

func Test_ParitialClone(t *testing.T) {
	library := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(library)
	resource := pkg.NewBytesResource([]byte(ruleString))
	rb.BuildRuleFromResource("n", "v", resource)
	kb1 := library.NewKnowledgeBaseInstance("n", "v")
	skb := library.NewPartialKnowledgeBase("n", "v", []string{"DuplicateRule1", "DuplicateRule3"})
	_ = kb1
	_ = skb

	ngin := engine.NewGruleEngine()
	dataCtx := ast.NewDataContext()
	f := &Fact{
		Distance: 10000,
		Duration: 121,
		Result:   false,
	}
	dataCtx.Add("Fact", f)
	err := ngin.Execute(dataCtx, skb)
	if err != nil {
		t.Errorf("execute error: %v", err)
	}
}

func Test_ParitialClone1000(t *testing.T) {
	input, _ := ioutil.ReadFile("benchmark/1000_rules.grl")
	library := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(library)
	resource := pkg.NewBytesResource([]byte(input))
	rb.BuildRuleFromResource("n", "v", resource)

	kb1 := library.NewKnowledgeBaseInstance("n", "v")
	skb := library.NewPartialKnowledgeBase("n", "v", []string{"DuplicateRule9"})
	_ = kb1
	_ = skb

	ngin := engine.NewGruleEngine()
	dataCtx := ast.NewDataContext()
	f := &Fact{
		Distance: 10000,
		Duration: 121,
		Result:   false,
	}
	dataCtx.Add("Fact", f)
	err := ngin.Execute(dataCtx, skb)
	if err != nil {
		t.Errorf("execute error: %v", err)
	}
}
