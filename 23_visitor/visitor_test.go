package visitor

import (
	"testing"
)

func TestExampleRequestVisitor(t *testing.T) {
	c := &CustomerCol{}
	// 添加对应的访问方式
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Add(NewIndividualCustomer("bob"))
	// 遍历访问
	c.Accept(&ServiceRequestVisitor{})
	// Output:
	// serving enterprise customer A company
	// serving enterprise customer B company
	// serving individual customer bob
}

func TestExampleAnalysis(t *testing.T) {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewIndividualCustomer("bob"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Accept(&AnalysisVisitor{})
	// Output:
	// analysis enterprise customer A company
	// analysis enterprise customer B company
}
