package ws

type Message struct {
	Type       string `json:"type"`
	HospitalID string `json:"hospitalId,omitempty"`
	RobotID string `json:"robotId,omitempty"`
	SDP       string `json:"sdp,omitempty"`
	Candidate string `json:"candidate,omitempty"`
	Data string `json:"data,omitempty"`
}

