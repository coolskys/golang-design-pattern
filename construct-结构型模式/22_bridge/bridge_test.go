package bridge

import "testing"

func TestExampleCommonSMS(t *testing.T) {
	m := NewCommonMessage(NewSMSSender())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via SMS
}

func TestExampleCommonEmail(t *testing.T) {
	m := NewCommonMessage(NewEmailSender())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via Email
}

func TestExampleUrgencySMS(t *testing.T) {
	m := NewUrgencyMessage(NewSMSSender())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via SMS
}

func TestExampleUrgencyEmail(t *testing.T) {
	m := NewUrgencyMessage(NewEmailSender())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via Email
}
