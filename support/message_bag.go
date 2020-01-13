package support

type MessageBag struct {
	messages map[string][]string
	keys     []string
}

func (m *MessageBag) Add(key string, msg string) {
	if m.messages == nil {
		m.messages = make(map[string][]string)
	}
	if _, ok := m.messages[key]; ok {
		m.messages[key] = append(m.messages[key], msg)
	} else {
		m.keys = append(m.keys, key)
		m.messages[key] = []string{msg}
	}
}

func (m *MessageBag) Messages() map[string][]string {
	return m.messages
}

func (m *MessageBag) First() string {
	if len(m.messages) <= 0 {
		return ""
	}
	return m.messages[m.keys[0]][0]
}
